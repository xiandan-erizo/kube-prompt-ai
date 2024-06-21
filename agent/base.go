package agent

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
)

/**
 * @author: xiandan
 * @createTime: 2024/06/19 上午10:20
 * @description:
 */

/*
history相关`
*/

type Message struct {
	Message openai.ChatCompletionMessage
	Length  uint8
}

type History struct {
	his map[string][]*Message
}

func (h *History) Get(key string) []openai.ChatCompletionMessage {
	var msgs []openai.ChatCompletionMessage
	if _, ok := ChatHistory.his[key]; !ok {
		ChatHistory.his[key] = make([]*Message, 0)
	}
	for _, msg := range ChatHistory.his[key] {
		msgs = append(msgs, msg.Message)
	}
	return msgs
}

// 添加消息
func (h *History) Add(key string, message openai.ChatCompletionMessage) {
	// TODO 这里补充压缩历史消息的逻辑
	msg := Message{
		Message: message,
		Length:  uint8(len(message.Content)),
	}
	ChatHistory.his[key] = append(ChatHistory.his[key], &msg)
}

// 压缩历史消息
func (h *History) AddToolMessage(key string, toolCallResultList []ToolCallResult) {
	for _, toolCallResult := range toolCallResultList {
		msg := ""
		if toolCallResult.err != nil {
			msg = fmt.Sprintf("执行失败,%v", toolCallResult.err)
		} else {
			//if strings.Split( toolCallResult.Output,"\n")
			msg = toolCallResult.Output
		}

		msgTool := &Message{
			Message: openai.ChatCompletionMessage{
				Role:       openai.ChatMessageRoleTool,
				Content:    msg,
				ToolCallID: toolCallResult.ID,
				Name:       toolCallResult.Name,
			},
			Length: uint8(len(msg)),
		}

		h.his[key] = append(h.his[key], msgTool)
	}
}

func (h *History) Clear(key string) {
	if len(ChatHistory.Get(key)) > 2 {
		log.Printf("%s 对话上下文已清除", key)
		ChatHistory.his[key] = ChatHistory.his[key][:1]
	}
}
