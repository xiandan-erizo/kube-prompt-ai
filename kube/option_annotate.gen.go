// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var annotateOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all", Description: "Select all resources, in the namespace of the specified resource types."},
	prompt.Suggest{Text: "-A", Description: "If true, check the specified action in all namespaces."},
	prompt.Suggest{Text: "--all-namespaces", Description: "If true, check the specified action in all namespaces."},
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--dry-run", Description: "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource."},
	prompt.Suggest{Text: "--field-manager", Description: "Name of the manager used to track field ownership."},
	prompt.Suggest{Text: "--field-selector", Description: "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to update the annotation"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to update the annotation"},
	prompt.Suggest{Text: "-k", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--kustomize", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--list", Description: "If true, display the annotations for a given resource."},
	prompt.Suggest{Text: "--local", Description: "If true, annotation will NOT contact api-server but run locally."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file)."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file)."},
	prompt.Suggest{Text: "--overwrite", Description: "If true, allow annotations to be overwritten, otherwise reject annotation updates that overwrite existing annotations."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--resource-version", Description: "If non-empty, the annotation update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints."},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints."},
	prompt.Suggest{Text: "--show-managed-fields", Description: "If true, keep the managedFields when printing objects in JSON or YAML format."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]. Usage:\n  kubectl annotate [--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version] [options] Use \"kubectl options\" for a list of global command-line options (applies to all commands)."},
}
