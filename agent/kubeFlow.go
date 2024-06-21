package agent

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/c-bata/kube-prompt/internal/debug"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"log"
	"os"
	"os/exec"
	"strings"
)

/**
 * @author: xiandan
 * @createTime: 2024/06/20 下午5:58
 * @description:
 */

type KubeFlow struct {
	Agent *Agent
}

var _ KubeAIFlow = (*KubeFlow)(nil)

func (m *KubeFlow) Run(input string) (openai.ChatCompletionMessage, error) {
	m.Agent.RunAgent(input, false, 0)
	messages := ChatHistory.Get(m.Agent.HisKey)
	messageLeast := messages[len(messages)-1]
	return messageLeast, nil
}

func (m *KubeFlow) Description() string {
	des := "运维工程师"
	return des
}

func (m *KubeFlow) Tools() []openai.Tool {
	var KubeTools []openai.Tool
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"flag": {
				Type:        jsonschema.String,
				Description: "kubectl命令行的flags参数",
			},
			"options": {
				Type: jsonschema.String,
				Enum: []string{"kubectl命令行的options参数"},
			},
			"other": {
				Type: jsonschema.String,
				Enum: []string{"这里可以是管道符及后面的命令,例如 | wc -l 用来统计数量, 但不允许输入危险指令"},
			},
		},
		Required: []string{"flags", "options"},
	}
	execKubeReadCommandFunction := openai.FunctionDefinition{
		Name:        "ExecKubeReadCommand",
		Description: "执行kubectl命令,仅可执行get，describe，log等只读，不对集群修改的命令,注意,get describe等参数也要写不然工具不认",
		Parameters:  params,
	}
	KubeTools = append(KubeTools, openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &execKubeReadCommandFunction,
	})
	return KubeTools
}

type KubectlArgs struct {
	Options string `json:"options"`
	Flag    string `json:"flag"`
	Other   string `json:"other"`
}

func (m *KubeFlow) ExecTool(dialogue openai.ChatCompletionMessage) []ToolCallResult {
	var toolResList []ToolCallResult
	for _, toolCall := range dialogue.ToolCalls {
		if toolCall.Function.Name == "ExecKubeReadCommand" {
			var args KubectlArgs
			err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args)
			response, err := ExecKubeReadCommand(args.Flag, args.Options, args.Other)
			if err != nil {
				log.Printf("ExecKubeReadCommand error: %v", err)
			} else {
				//log.Printf("ExecKubeReadCommand response: %v", response)
			}

			toolResList = append(toolResList, ToolCallResult{
				Name:      toolCall.Function.Name,
				Arguments: toolCall.Function.Arguments,
				Output:    response,
				ID:        toolCall.ID,
				err:       err,
			})
		}
	}

	return toolResList
}

func ExecKubeReadCommand(flags string, options string, other string) (string, error) {
	_, err := confirmExecKubeCommand(flags, options, other)
	if err != nil {
		return "", err
	}
	return ExecuteAndGetResult(fmt.Sprintf("%s %s", flags, options))
}

func confirmExecKubeCommand(flags string, options string, other string) (bool, error) {
	PrintRed("\n执行命令: %s %s %s\n", flags, options, other)
	reader := bufio.NewReader(os.Stdin)
	PrintRed("请输入y确认执行命令: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时发生错误:", err)
		return false, fmt.Errorf("用户取消执行命令")
	}
	if input != "y\n" {
		return false, fmt.Errorf("用户取消执行命令")
	}
	return true, nil
}

func ExecuteAndGetResult(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		debug.Log("you need to pass the something arguments")
		return "", fmt.Errorf("you need to pass the something arguments")
	}

	out := &bytes.Buffer{}
	errSout := &bytes.Buffer{}
	cmd := exec.Command("/bin/sh", "-c", "kubectl "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	cmd.Stderr = errSout

	if err := cmd.Run(); err != nil {
		errStr := string(errSout.Bytes())
		fmt.Printf("执行命令失败%s", errStr)
		return "", fmt.Errorf(errStr)
	}
	r := string(out.Bytes())
	return r, nil
}

func PrintRed(message string, a ...any) {

	msgOut := fmt.Sprintf(message, a...)

	red := "\033[31m"
	reset := "\033[0m" // 用于之后重置颜色

	fmt.Print(red + msgOut + reset)
}

// Init TODO 改成配置的不是传参
func (m *KubeFlow) Init(model string, maxHistory uint8, maxCounts uint8) error {
	m.Agent = &Agent{
		Model:       model,
		Prompt:      KubePrompt,
		MaxTokens:   2000,
		Temperature: 0.1,
		MaxHistory:  maxHistory,
		MaxCounts:   maxCounts,
		HisKey:      "monitor",
		Tools:       m.Tools(),
		ExecTool:    m.ExecTool,
	}
	his := map[string][]*Message{}
	his[m.Agent.HisKey] = []*Message{
		{
			Message: openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleSystem,
				Content: m.Agent.Prompt,
			},
			Length: uint8(len(m.Agent.Prompt))},
	}
	ChatHistory = &History{
		his: his,
	}

	return nil
}

var KubeFlowIns = KubeFlow{}

func init() {
	var err error
	err = KubeFlowIns.Init("gpt-4o", 5, 5)
	if err != nil {
		fmt.Printf("Error init: %v\n", err)
	}
}

var KubePrompt = `
# 角色
你是一个运维工程师,你非常精通Kubernetes。你的技术娴熟，可以生成、运行以及优化kubectl命令。你的任务是提供尽可能简洁、可靠的方案，通过减少输出结果以便降低上下文复杂度。

## 技能
### 技能 1: 生成kubectl命令
- 根据用户的需求生成kubectl命令。

### 技能 2: 优化kubectl命令
- 对生成的kubectl命令进行优化，尽可能减少输出结果并保持命令的可靠性。

### 技能 3: 执行kubectl命令
- 执行生成并优化的kubectl命令，获取并返回执行结果。

### 技能 4: 判断是否需要执行命令
- 充分理解用户的问题，当已经获取到足够的信息来回答用户的问题时，无需再执行新的命令。

## 限制：
- 只讨论与kubectl命令生成和优化有关的话题。
- 根据用户的需求使用恰当的命令。
- 优化的角度是减少输出结果和提升命令的可靠性。
- 尽量减少上下文复杂度，提供清晰简洁
- 你的输出环境是shell的环境,请注意格式和缩进
- 查看日志时不要超过200行
`
