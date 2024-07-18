package kube

import (
	"bytes"
	"fmt"
	"github.com/xiandan-erizo/kube-prompt-ai/agent"
	"github.com/xiandan-erizo/kube-prompt-ai/internal/debug"
	"os"
	"os/exec"
	"strings"
)

var aiModel = false
var InputModel = "user"

func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if strings.ToLower(s) == "#ai" {
		fmt.Println("Enter AI Conversation Mode\r\n")
		InputModel = "User"
		agent.ChatHistory.Clear("ops")
		aiModel = true
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	} else if s == "/clear" {
		agent.ChatHistory.Clear("ops")
		//agent.ChatHeistory.Clear("flow")
		agent.ChatHistory.Clear("monitor")
		return
	} else if s == "exai" && aiModel {
		agent.ChatHistory.Clear("ops")
		fmt.Println("Exit AI Conversation Mode\n")
		aiModel = false
		return
	} else if aiModel && s != "" {
		_, err := agent.FlowIns.Run(s)
		if err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
		}
		return
	}

	cmd := exec.Command("/bin/sh", "-c", "kubectl "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}

func ExecuteAndGetResult(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		debug.Log("you need to pass the something arguments")
		return ""
	}

	out := &bytes.Buffer{}
	cmd := exec.Command("/bin/sh", "-c", "kubectl "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		debug.Log(err.Error())
		return ""
	}
	r := string(out.Bytes())
	return r
}
