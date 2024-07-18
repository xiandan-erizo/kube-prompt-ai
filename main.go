package main

import (
	"fmt"
	"github.com/xiandan-erizo/kube-prompt-ai/internal/debug"
	"github.com/xiandan-erizo/kube-prompt-ai/kube"
	"os"

	prompt "github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

var (
	version  string
	revision string
)

func main() {
	c, err := kube.NewCompleter()
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	defer debug.Teardown()
	fmt.Printf("kube-prompt for ai %s (rev-%s)\n", version, revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
	fmt.Println("input #AI to talk with AI🤖")
	defer fmt.Println("Bye!👋🏻")
	p := prompt.New(
		kube.Executor,
		c.Complete,
		prompt.OptionTitle("kube-prompt: interactive kubernetes client"),
		prompt.OptionPrefix(fmt.Sprintf("%s>>> ", kube.InputModel)),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)
	p.Run()
}
