// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var autoscaleOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--cpu-percent", Description: "The target average CPU utilization (represented as a percent of requested CPU) over all the pods. If it's not specified or negative, a default autoscaling policy will be used."},
	prompt.Suggest{Text: "--dry-run", Description: "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource."},
	prompt.Suggest{Text: "--field-manager", Description: "Name of the manager used to track field ownership."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to autoscale."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to autoscale."},
	prompt.Suggest{Text: "-k", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--kustomize", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--max", Description: "The upper limit for the number of pods that can be set by the autoscaler. Required."},
	prompt.Suggest{Text: "--min", Description: "The lower limit for the number of pods that can be set by the autoscaler. If it's not specified or negative, the server will apply a default value."},
	prompt.Suggest{Text: "--name", Description: "The name for the newly created object. If not specified, the name of the input resource will be used."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file)."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file)."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future."},
	prompt.Suggest{Text: "--show-managed-fields", Description: "If true, keep the managedFields when printing objects in JSON or YAML format."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]. Usage:\n  kubectl autoscale (-f FILENAME | TYPE NAME | TYPE/NAME) [--min=MINPODS] --max=MAXPODS [--cpu-percent=CPU] [options] Use \"kubectl options\" for a list of global command-line options (applies to all commands)."},
}
