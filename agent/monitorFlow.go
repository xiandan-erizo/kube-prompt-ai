package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/api"
	proV1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"log"
	"net/http"
	"os"
	"time"
)

/**
 * @author: xiandan
 * @createTime: 2024/06/19 上午10:07
 * @description:
 */

type MonitorFlow struct {
	Agent *Agent
}

var _ KubeAIFlow = (*MonitorFlow)(nil)

func (m *MonitorFlow) Run(input string) (openai.ChatCompletionMessage, error) {
	m.Agent.RunAgent(input, false, 0)
	messages := ChatHistory.Get(m.Agent.HisKey)
	messageLeast := messages[len(messages)-1]
	return messageLeast, nil
}

func (m *MonitorFlow) Description() string {
	des := "监控内容获取工作流"
	return des
}

func (m *MonitorFlow) Tools() []openai.Tool {
	var monitorTools []openai.Tool
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"query": {
				Type:        jsonschema.String,
				Description: "promq",
			},
			"step": {
				Type: jsonschema.Integer,
				Enum: []string{"时间间隔,为了保证数据量不会太大,单位是分钟"},
			},
		},
		Required: []string{"flags", "options"},
	}
	monitorQueryRange := openai.FunctionDefinition{
		Name:        "MonitorQueryRange",
		Description: "执行prometheus查询语句",
		Parameters:  params,
	}
	t := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &monitorQueryRange,
	}
	monitorTools = append(monitorTools, t)

	return monitorTools
}

func (m *MonitorFlow) ExecTool(dialogue openai.ChatCompletionMessage) []ToolCallResult {
	var toolResList []ToolCallResult
	for _, toolCall := range dialogue.ToolCalls {
		if toolCall.Function.Name == "MonitorQueryRange" {
			var args MonitorQuery
			err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args)
			if args.Step == 0 {
				args.Step = 10
			}
			response, err := MonitorQueryRange(args.Query, args.Step)
			if err != nil {
				log.Printf("ExecKubeReadCommand error: %v", err)
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

// Init TODO 改成配置的不是传参
func (m *MonitorFlow) Init(model string, maxHistory uint8, maxCounts uint8) error {
	m.Agent = &Agent{
		Model:       model,
		Prompt:      promptMonitor,
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

var Client api.Client

var MonitorFlowIns = MonitorFlow{}

func init22() {
	var err error
	Client, err = CreatClient()
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
	}

	err = MonitorFlowIns.Init("gpt-4o", 5, 5)
	if err != nil {
		fmt.Printf("Error init: %v\n", err)
	}
}

type CustomRoundTripper struct {
	Transport http.RoundTripper
}

func (c *CustomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// todo delete
	au := os.Getenv("Authorization")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", au))
	return c.Transport.RoundTrip(req)
}

func CreatClient() (api.Client, error) {
	client, err := api.NewClient(api.Config{
		Address: "",
		RoundTripper: &CustomRoundTripper{
			Transport: http.DefaultTransport,
		},
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		return nil, err
	}
	return client, nil
}

type MonitorQuery struct {
	Query string `json:"query"`
	Step  int    `json:"step"`
}

func MonitorQueryRange(query string, Step int) (string, error) {
	v1api := proV1.NewAPI(Client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stepDuration := time.Duration(Step) * time.Minute

	r := proV1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  stepDuration,
	}

	result, warnings, err := v1api.QueryRange(ctx, query, r, proV1.WithTimeout(5*time.Second))
	if err != nil {
		return "", err
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}

	return processResult(result), nil
}

func timeS(timestampMs int64) string {
	// 将Unix时间戳转换为time.Time类型
	t := time.Unix(int64(timestampMs), 0)
	// 格式化输出时间
	return t.Format("2006-01-02 15:04:05")

}

func processResult(result model.Value) string {
	var resultStr = ""
	switch v := result.(type) {
	case model.Vector:
		for _, sample := range v {
			resultStr += fmt.Sprintf("Metric: %+v, Value: %f\n", sample.Metric, sample.Value)

		}
	case model.Matrix:
		for _, series := range v {
			resultStr += fmt.Sprintf("Series for metric: %+v\n", series.Metric)
			for _, sample := range series.Values {
				timestampMs := timeS(int64(sample.Timestamp) / 1000)
				value := sample.Value
				resultStr += fmt.Sprintf("Timestamp: %s, Value: %f\n", timestampMs, value)
			}
		}
	default:
		//fmt.Printf("Unsupported result type: %T\n", v)
		resultStr += fmt.Sprintf("Unsupported result type: %T\n", v)
	}
	return resultStr
}

var promptMonitor = `
- Role: 监控分析师
- Background: 用户需要从Prometheus监控系统中检索和分析时间序列数据。
- Profile: 您是一位专业的监控分析师，熟悉PromQL语法和监控数据的分析方法。
- Skills: 掌握PromQL语法、数据分析、监控系统知识。
- Goals: 设计一个prompt，帮助用户构建和优化PromQL查询语句，以获取所需的监控数据。
- Constrains: 查询语句需要精确、高效，并且能够正确反映监控数据的趋势和模式。
- Workflow:
  1. 确定用户想要查询的数据指标和时间范围。
  2. 根据用户需求构建PromQL查询语句。
  3. 执行查询并呈现结果。
  命名空间始终为 dingding
- Examples:
  - 查询过去1小时内CPU使用率的平均值："rate(process_cpu_seconds_total{namespace="dingding", app=~"release-rest-server"}[1h])"
  - 查询jvm用率的平均值："sum by (pod) (jvm_memory_bytes_used{namespace="dingding", app="release-rest-server", area="heap"})"
- Initialization: 欢迎使用PromQL查询构建服务。请告诉我您想要查询哪些监控指标，以及您的具体需求。
`
