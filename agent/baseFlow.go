package agent

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

/**
 * @author: xiandan
 * @createTime: 2024/06/19 上午11:55
 * @description:
 */

var (
	ChatHistory *History
	maxCounts   = 5 // 最大次数
	client      *openai.Client
	//ai_model       = os.Getenv("OPENAI_MODEL")
	ai_model = "gpt-4o"
)

type KubeAIFlow interface {
	Run(input string) (openai.ChatCompletionMessage, error)
	Init(Model string, maxCounts uint8, MaxHistory uint8) error
	Description() string
	Tools() []openai.Tool
	ExecTool(dialogue openai.ChatCompletionMessage) []ToolCallResult
}

type ToolCallResult struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
	Output    string `json:"output"`
	ID        string `json:"id"`
	err       error
}
type Agent struct {
	Model       string
	Tools       []openai.Tool // 工具
	Prompt      string        // prompt
	MaxTokens   int           // 最大tokens
	Temperature float32       // 0.0-1.0
	TopP        float32
	Stop        []string
	MaxHistory  uint8 // 最大历史记录
	MaxCounts   uint8 // 最大次数
	Stream      bool
	HisKey      string
	ExecTool    func(dialogue openai.ChatCompletionMessage) []ToolCallResult
}

// RunAgent 运行AI
func (a *Agent) RunAgent(userInput string, toolCall bool, currentCount int) {
	currentCount++
	if currentCount > maxCounts && toolCall {
		fmt.Println("单次AI迭代超出次数限制,请重新开始对话")
		ChatHistory.Clear(a.HisKey)
		return
	}
	ctx := context.Background()
	if !toolCall {
		ChatHistory.Add(a.HisKey,
			openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: userInput,
			},
		)
	}
	toolCall = false
	if a.HisKey == "flow" {
		fmt.Printf("AI>: %v", "生成中请稍等...\r")
	}
	resp, err := client.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model:    ai_model,
			Messages: ChatHistory.Get(a.HisKey),
			Tools:    a.Tools,
		},
	)
	// 第一次调用
	if err != nil || len(resp.Choices) != 1 {
		log.Printf("Completion error: err:%v len(choices):%v\n", err, len(resp.Choices))
		return
	}

	msg := resp.Choices[0].Message
	if a.HisKey == "flow" {
		fmt.Printf("\nAI>: %v\n", msg.Content)
	} else {
		PrintRed(fmt.Sprintf("\nAI>: %v\n", msg.Content))
	}
	ChatHistory.Add(a.HisKey, msg)
	// 未使用工具则直接返回
	if len(msg.ToolCalls) == 0 {
		return
	}

	// 调用工具
	toolResList := a.ExecTool(msg)

	ChatHistory.AddToolMessage(a.HisKey, toolResList)
	// 迭代调用AI
	a.RunAgent("", true, currentCount)
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

}
