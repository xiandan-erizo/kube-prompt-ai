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
 * @createTime: 2024/03/01 16:21
 * @description:
 */

// 这里后期优化成从命令文件中获取
var ReadOlny = []string{"get", "describe", "log", "top"}

type ToolCallResult struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
	Output    string `json:"output"`
	ID        string `json:"id"`
	err       error
}

type KubectlArgs struct {
	Options string `json:"options"`
	Flag    string `json:"flag"`
	Other   string `json:"other"`
}

func ExecTool(dialogue openai.ChatCompletionMessage) []ToolCallResult {
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
	//if !In(flags, ReadOlny) && flags != "" {
	//	return "error", fmt.Errorf("only get, describe, log command can be executed")
	//}
	//if !In(options, ReadOlny) && options != "" {
	//	return "error", fmt.Errorf("only get, describe, log command can be executed")
	//}
	_, err := confirmExecKubeCommand(flags, options, other)
	if err != nil {
		return "", err
	}
	return ExecuteAndGetResult(fmt.Sprintf("%s %s", flags, options))
}

func confirmExecKubeCommand(flags string, options string, other string) (bool, error) {
	//fmt.Printf("\n执行命令: %s %s %s\n", flags, options, other)
	PrintRead("\n执行命令: %s %s %s\n", flags, options, other)
	reader := bufio.NewReader(os.Stdin)
	//fmt.Print("请输入y确认执行命令: ")
	PrintRead("请输入y确认执行命令: ")
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

func ExecKubePatchCommand(flags string, options string) (string, error) {
	for _, flag := range flags {
		fmt.Printf("flag:%s \n", flag)
	}

	for _, option := range options {
		fmt.Printf("flag:%s \n", option)
	}

	return "success", nil
}

func In(arg string, array []string) bool {
	for _, a := range array {
		if arg == a {
			return true
		}
	}
	return false
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
	//debug.Log(r)
	return r, nil
}

func DefinitionTools() openai.Tool {
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
	t := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &execKubeReadCommandFunction,
	}

	return t
}

func PrintRead(message string, a ...any) {

	msgOut := fmt.Sprintf(message, a...)

	red := "\033[31m"
	reset := "\033[0m" // 用于之后重置颜色

	fmt.Print(red + msgOut + reset)
}
