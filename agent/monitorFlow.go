package agent

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	proV1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"net/http"
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

func (m *MonitorFlow) Run(input string) {
	m.Agent.RunAgent(input, false, 0)
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
				Type: jsonschema.String,
				Enum: []string{"时间间隔"},
			},
		},
		Required: []string{"flags", "options"},
	}
	monitorQueryRange := openai.FunctionDefinition{
		Name:        "MonitorQueryRange",
		Description: "执行kubectl命令,仅可执行get，describe，log等只读，不对集群修改的命令,注意,get describe等参数也要写不然工具不认",
		Parameters:  params,
	}
	t := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &monitorQueryRange,
	}
	monitorTools = append(monitorTools, t)

	return monitorTools
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
	}

	return nil
}

var Client api.Client

func init() {
	var err error
	Client, err = CreatClient()
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
	}
}

type CustomRoundTripper struct {
	Transport http.RoundTripper
}

func (c *CustomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// todo delete
	//au := os.Getenv("Authorization")
	au := "c3Vua2VzaTpTdW5AdmlhMDgxNEA="
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", au))
	return c.Transport.RoundTrip(req)
}

func CreatClient() (api.Client, error) {
	client, err := api.NewClient(api.Config{
		Address: "http://prometheus-dd.ekuaibao.net/",
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
			//fmt.Printf("Metric: %+v, Value: %f\n", sample.Metric, sample.Value)
			resultStr += fmt.Sprintf("Metric: %+v, Value: %f\n", sample.Metric, sample.Value)
		}
	case model.Matrix:
		for _, series := range v {
			fmt.Printf("Series for metric: %+v\n", series.Metric)
			for _, sample := range series.Values {
				timestampMs := timeS(int64(sample.Timestamp) / 1000)
				value := sample.Value
				fmt.Printf("Timestamp: %s, Value: %f\n", timestampMs, value)
				resultStr += fmt.Sprintf("Timestamp: %s, Value: %f\n", timestampMs, value)
			}
		}
	default:
		fmt.Printf("Unsupported result type: %T\n", v)
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
- OutputFormat: 输出通常为时间序列图表或数据表格。
- Workflow:
  1. 确定用户想要查询的数据指标和时间范围。
  2. 根据用户需求构建PromQL查询语句。
  3. 执行查询并呈现结果。
- Examples:
  - 查询过去1小时内CPU使用率的平均值："avg(rate(container_cpu_usage_seconds_total{container_label_com_docker_swarm_service_name = "my_service"}[1h])) by (instance)"
  - 获取特定服务的请求总数："sum(rate(http_requests_total{job = "my_job"}[5m])) by (status)"
  - 计算特定指标在过去2小时内的增长率："increase(metric_name[2h])"
- Initialization: 欢迎使用PromQL查询构建服务。请告诉我您想要查询哪些监控指标，以及您的具体需求。
`
