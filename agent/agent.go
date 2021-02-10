/**
 * @author: xiandan
 * @createTime: 2024/06/13 下午5:36
 * @description:
 */

package agent

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"log"
	"os"
)

var (
	tools       = DefinitionTools()
	maxCounts   = 5 // 最大次数
	client      *openai.Client
	ChatHistory *History
	maxMemory   = 5
	//model       = os.Getenv("OPENAI_MODEL")
	model = "gpt-4o"
)

/**
 * @description:
 * @param userInput
 * @return 是否继续
 */

func Talk2AI(userInput string, toolCall bool, currentCount int) {
	currentCount++
	if currentCount > maxCounts && toolCall {
		fmt.Println("单次AI迭代超出次数限制,请重新开始对话")
		ChatHistory.Clear()
		return
	}
	ctx := context.Background()
	if !toolCall {
		ChatHistory.Add(
			openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: userInput,
			},
		)
	}
	toolCall = false
	fmt.Printf("AI>: %v", "生成中请稍等...\r")
	resp, err := client.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: ChatHistory.Messages,
			Tools:    []openai.Tool{tools},
		},
	)
	// 第一次调用
	if err != nil || len(resp.Choices) != 1 {
		log.Printf("Completion error: err:%v len(choices):%v\n", err, len(resp.Choices))
		return
	}
	msg := resp.Choices[0].Message
	fmt.Printf("AI>: %v\n", msg.Content)
	ChatHistory.Add(msg)
	// 未使用工具则直接返回
	if len(msg.ToolCalls) == 0 {
		return
	}

	// 调用工具
	toolResList := ExecTool(msg)

	ChatHistory.AddToolMessage(toolResList)
	// 迭代调用AI
	Talk2AI("", true, currentCount)
}

/*
history相关`
*/
type History struct {
	Messages []openai.ChatCompletionMessage
}

var prompt = `
# 角色
你是一个Kubernetes优化器。你的技术娴熟，可以生成、运行以及优化kubectl命令。你的任务是提供尽可能简洁、可靠的方案，通过减少输出结果以便降低上下文复杂度。

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

func (h *History) Get() *History {
	return h
}

// 压缩历史消息
func (h *History) Add(message openai.ChatCompletionMessage) {
	ChatHistory.Messages = append(ChatHistory.Messages, message)
}

// 压缩历史消息
func (h *History) AddToolMessage(toolCallResultList []ToolCallResult) {

	for _, toolCallResult := range toolCallResultList {
		msg := ""
		if toolCallResult.err != nil {
			msg = fmt.Sprintf("执行失败,%v", toolCallResult.err)
		} else {
			//if strings.Split( toolCallResult.Output,"\n")
			msg = toolCallResult.Output
		}
		h.Messages = append(h.Messages, openai.ChatCompletionMessage{
			Role:       openai.ChatMessageRoleTool,
			Content:    msg,
			ToolCallID: toolCallResult.ID,
			Name:       toolCallResult.Name,
		})
	}
}

func (h *History) Clear() {
	ChatHistory.Messages = ChatHistory.Messages[:1]
}

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	apiBase := os.Getenv("OPENAI_API_BASE")
	if apiKey == "" || apiBase == "" {
		log.Fatal("OPENAI_API_KEY or OPENAI_API_BASE is not set")
	}
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = apiBase
	client = openai.NewClientWithConfig(config)
	ChatHistory = &History{
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompt,
			},
		},
	}

}
