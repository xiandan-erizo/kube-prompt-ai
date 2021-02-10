package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/c-bata/kube-prompt/internal/optionconv"
)

var (
	output       string
	pkg          string
	variableName string
)

var in = "Display one or many resources.\n\n Prints a table of the most important information about the specified resources. You can filter the list using a label selector and the --selector flag. If the desired resource type is namespaced you will only see results in your current namespace unless you pass --all-namespaces.\n\n By specifying the output as 'template' and providing a Go template as the value of the --template flag, you can filter the attributes of the fetched resources.\n\nUse \"kubectl api-resources\" for a complete list of supported resources.\n\nExamples:\n  # List all pods in ps output format\n  kubectl get pods\n  \n  # List all pods in ps output format with more information (such as node name)\n  kubectl get pods -o wide\n  \n  # List a single replication controller with specified NAME in ps output format\n  kubectl get replicationcontroller web\n  \n  # List deployments in JSON output format, in the \"v1\" version of the \"apps\" API group\n  kubectl get deployments.v1.apps -o json\n  \n  # List a single pod in JSON output format\n  kubectl get -o json pod web-pod-13je7\n  \n  # List a pod identified by type and name specified in \"pod.yaml\" in JSON output format\n  kubectl get -f pod.yaml -o json\n  \n  # List resources from a directory with kustomization.yaml - e.g. dir/kustomization.yaml\n  kubectl get -k dir/\n  \n  # Return only the phase value of the specified pod\n  kubectl get -o template pod/web-pod-13je7 --template={{.status.phase}}\n  \n  # List resource information in custom columns\n  kubectl get pod test-pod -o custom-columns=CONTAINER:.spec.containers[0].name,IMAGE:.spec.containers[0].image\n  \n  # List all replication controllers and services together in ps output format\n  kubectl get rc,services\n  \n  # List one or more resources by their type and names\n  kubectl get rc/web service/frontend pods/web-pod-13je7\n  \n  # List status subresource for a single pod.\n  kubectl get pod web-pod-13je7 --subresource status\n\nOptions:\n    -A, --all-namespaces=false:\n        If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.\n\n    --allow-missing-template-keys=true:\n        If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.\n\n    --chunk-size=500:\n        Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future.\n\n    --field-selector='':\n        Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.\n\n    -f, --filename=[]:\n        Filename, directory, or URL to files identifying the resource to get from a server.\n\n    --ignore-not-found=false:\n        If the requested object does not exist the command will return exit code 0.\n\n    -k, --kustomize='':\n        Process the kustomization directory. This flag can't be used together with -f or -R.\n\n    -L, --label-columns=[]:\n        Accepts a comma separated list of labels that are going to be presented as columns. Names are case-sensitive. You can also use multiple flag options like -L label1 -L label2...\n\n    --no-headers=false:\n        When using the default or custom-column output format, don't print headers (default print headers).\n\n    -o, --output='':\n        Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file, custom-columns, custom-columns-file, wide). See custom columns [https://kubernetes.io/docs/reference/kubectl/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [https://kubernetes.io/docs/reference/kubectl/jsonpath/].\n\n    --output-watch-events=false:\n        Output watch event objects when --watch or --watch-only is used. Existing objects are output as initial ADDED events.\n\n    --raw='':\n        Raw URI to request from the server.  Uses the transport specified by the kubeconfig file.\n\n    -R, --recursive=false:\n        Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.\n\n    -l, --selector='':\n        Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.\n\n    --server-print=true:\n        If true, have the server return the appropriate table output. Supports extension APIs and CRDs.\n\n    --show-kind=false:\n        If present, list the resource type for the requested object(s).\n\n    --show-labels=false:\n        When printing, show all labels as the last column (default hide labels column)\n\n    --show-managed-fields=false:\n        If true, keep the managedFields when printing objects in JSON or YAML format.\n\n    --sort-by='':\n        If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string.\n\n    --subresource='':\n        If specified, gets the subresource of the requested object. Must be one of [status scale]. This flag is alpha and may change in the future.\n\n    --template='':\n        Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].\n\n    -w, --watch=false:\n        After listing/getting the requested object, watch for changes.\n\n    --watch-only=false:\n        Watch for changes to the requested object(s), without listing/getting first.\n\nUsage:\n  kubectl get [(-o|--output=)json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file|custom-columns|custom-columns-file|wide] (TYPE[.VERSION][.GROUP] [NAME | -l label] | TYPE[.VERSION][.GROUP]/NAME ...) [flags] [options]\n\nUse \"kubectl options\" for a list of global command-line options (applies to all commands).\n\n    -A, --all-namespaces=false:\n        If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.\n"

func convert() error {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	f, err := optionconv.GetOptionsFromHelpTextNew(string(bytes))
	if err != nil {
		return err
	}
	options := optionconv.SplitOptions(f)
	suggests := optionconv.ConvertToSuggestions(options)
	if output == "" {
		//_, err = pp.Fprintln(os.Stdout, suggests)
		for _, suggest := range suggests {
			_, err = fmt.Println(fmt.Sprintf("  flags: %s Description: %s\n", suggest.Text, suggest.Description))
		}

	} else {
		f, err := os.Create(output)
		if err != nil {
			return err
		}
		defer f.Close()

		fmt.Fprintf(f, "// Code generated by 'option-gen'. DO NOT EDIT.\n\n")
		fmt.Fprintf(f, "package %s\n\n", pkg)
		fmt.Fprintln(f, `import (`)
		fmt.Fprintln(f, `prompt "github.com/c-bata/go-prompt"`)
		fmt.Fprintln(f, ")")
		fmt.Fprintln(f, "")
		fmt.Fprintf(f, "var %s = []prompt.Suggest{\n", variableName)
		for _, s := range suggests {
			fmt.Fprintf(f, "%#v,\n", s)
		}
		fmt.Fprintln(f, "}")
	}
	return err
}

func main() {
	flag.StringVar(&output, "o", "", "output file. print stdout if empty")
	flag.StringVar(&pkg, "pkg", "kube", "package name")
	flag.StringVar(&variableName, "var", "flagXXX", "variable name")
	flag.Parse()

	if err := convert(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
