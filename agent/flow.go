package agent

import (
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"log"
)

/**
 * @author: xiandan
 * @createTime: 2024/06/20 下午4:57
 * @description:
 */

type Flow struct {
	Agent *Agent
}

var _ KubeAIFlow = (*Flow)(nil)

func (m *Flow) Run(input string) (openai.ChatCompletionMessage, error) {
	m.Agent.RunAgent(input, false, 0)
	messages := ChatHistory.Get(m.Agent.HisKey)
	messageLeast := messages[len(messages)-1]
	return messageLeast, nil
}

func (m *Flow) Description() string {
	des := ""
	return des
}

func (m *Flow) Tools() []openai.Tool {
	var FlowTools []openai.Tool
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"input": {
				Type:        jsonschema.String,
				Description: "需要查询的监控指标,自然语言描述",
			},
		},
		Required: []string{"input"},
	}
	monitorQueryRange := openai.FunctionDefinition{
		Name:        "MonitorQueryRange",
		Description: "执行prometheus查询语句",
		Parameters:  params,
	}
	tMonitorQueryRange := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &monitorQueryRange,
	}
	FlowTools = append(FlowTools, tMonitorQueryRange)

	kparams := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"input": {
				Type:        jsonschema.String,
				Description: "对k8s集群的操作,更精确的自然语言描述",
			},
		},
		Required: []string{"input"},
	}
	execKubeCommand := openai.FunctionDefinition{
		Name:        "ExecKubeCommand",
		Description: "对k8s的操作,更精确的自然语言描述",
		Parameters:  kparams,
	}
	tKubeCommand := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &execKubeCommand,
	}

	FlowTools = append(FlowTools, tKubeCommand)

	return FlowTools
}

type FlowInput struct {
	Input string `json:"input"`
}

func (m *Flow) ExecTool(dialogue openai.ChatCompletionMessage) []ToolCallResult {
	var toolResList []ToolCallResult
	for _, toolCall := range dialogue.ToolCalls {
		var args FlowInput
		err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args)
		if err != nil {
			log.Printf("解析参数失败: %v", err)
		}
		if toolCall.Function.Name == "MonitorQueryRange" {
			response, err := MonitorFlowIns.Run(args.Input)
			toolResList = append(toolResList, ToolCallResult{
				Name:      toolCall.Function.Name,
				Arguments: toolCall.Function.Arguments,
				Output:    response.Content,
				ID:        toolCall.ID,
				err:       err,
			})
		}
		if toolCall.Function.Name == "ExecKubeCommand" {
			response, err := KubeFlowIns.Run(args.Input)
			toolResList = append(toolResList, ToolCallResult{
				Name:      toolCall.Function.Name,
				Arguments: toolCall.Function.Arguments,
				Output:    response.Content,
				ID:        toolCall.ID,
				err:       err,
			})
		}
		if err != nil {
			log.Printf("执行工具失败: %v", err)
		}
	}

	return toolResList
}

// Init TODO 改成配置的不是传参
func (m *Flow) Init(model string, maxHistory uint8, maxCounts uint8) error {
	m.Agent = &Agent{
		Model:       model,
		Prompt:      promptFlow,
		MaxTokens:   2000,
		Temperature: 0.1,
		MaxHistory:  maxHistory,
		MaxCounts:   maxCounts,
		HisKey:      "flow",
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

var FlowIns = Flow{}

func init() {
	var err error
	err = FlowIns.Init("gpt-4o", 5, 5)
	if err != nil {
		fmt.Printf("Error init: %v\n", err)
	}
}

var promptFlow = `
# 角色
你是一位资深的运维工程师，能够根据用户的自然语言输入，将其转换为 kubectl 命令并执行，同时也能生成 Prometheus 查询语句（promq）进行监控数据的查询并返回结果，还能合理编排这两个工作流的执行顺序。

## 技能
### 技能 1: kubectl 命令执行
1. 当用户输入与 Kubernetes 操作相关的自然语言描述时，将其转换为对应的 kubectl 命令并执行。
2. 如果对用户输入的理解不明确，向用户进一步询问以明确需求。

### 技能 2: Prometheus 监控查询
1. 当用户提出与监控数据查询相关的需求时，生成准确的 promq 进行查询。
2. 若生成的查询结果不符合预期，重新分析用户需求并调整查询语句。

### 技能 3: 工作流编排
1. 根据用户问题的性质和需求，合理安排 kubectl 命令执行和 Prometheus 监控查询工作流的先后顺序。
2. 若两个工作流存在依赖关系，按照依赖顺序执行。

## 限制:
- 只处理与 Kubernetes 操作和 Prometheus 监控查询相关的问题，拒绝回答无关内容。
- 严格按照用户需求执行工作流，不得随意更改执行顺序。
- 对于复杂的需求，要确保理解准确后再进行操作。
- 输出的结果要清晰、准确，易于用户理解。
- 可以一次运行多个工作流
`
