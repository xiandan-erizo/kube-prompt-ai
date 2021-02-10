// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var applyOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all", Description: "Select all resources in the namespace of the specified resource types."},
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--cascade", Description: "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background."},
	prompt.Suggest{Text: "--dry-run", Description: "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource."},
	prompt.Suggest{Text: "--field-manager", Description: "Name of the manager used to track field ownership."},
	prompt.Suggest{Text: "-f", Description: "The files that contain the configurations to apply."},
	prompt.Suggest{Text: "--filename", Description: "The files that contain the configurations to apply."},
	prompt.Suggest{Text: "--force", Description: "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation."},
	prompt.Suggest{Text: "--force-conflicts", Description: "If true, server-side apply will force the changes against conflicts."},
	prompt.Suggest{Text: "--grace-period", Description: "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion)."},
	prompt.Suggest{Text: "-k", Description: "Process a kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--kustomize", Description: "Process a kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--openapi-patch", Description: "If true, use openapi to calculate diff when the openapi presents and the resource can be found in the openapi spec. Otherwise, fall back to use baked-in types."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file)."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file)."},
	prompt.Suggest{Text: "--overwrite", Description: "Automatically resolve conflicts between the modified and live configuration by using values from the modified configuration"},
	prompt.Suggest{Text: "--prune", Description: "Automatically delete resource objects, that do not appear in the configs and are created by either apply or create --save-config. Should be used with either -l or --all."},
	prompt.Suggest{Text: "--prune-whitelist", Description: "Overwrite the default whitelist with <group/version/kind> for --prune"},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints."},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints."},
	prompt.Suggest{Text: "--server-side", Description: "If true, apply runs in the server instead of the client."},
	prompt.Suggest{Text: "--show-managed-fields", Description: "If true, keep the managedFields when printing objects in JSON or YAML format."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--timeout", Description: "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object"},
	prompt.Suggest{Text: "--validate", Description: "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields."},
	prompt.Suggest{Text: "--wait", Description: "If true, wait for resources to be gone before returning. This waits for finalizers. Usage:\n  kubectl apply (-f FILENAME | -k DIRECTORY) [options] Use \"kubectl <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."},
}
