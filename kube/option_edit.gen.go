// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var editOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--field-manager", Description: "Name of the manager used to track field ownership."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files to use to edit the resource"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files to use to edit the resource"},
	prompt.Suggest{Text: "-k", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--kustomize", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file)."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file)."},
	prompt.Suggest{Text: "--output-patch", Description: "Output the patch if the resource is edited."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future."},
	prompt.Suggest{Text: "--show-managed-fields", Description: "If true, keep the managedFields when printing objects in JSON or YAML format."},
	prompt.Suggest{Text: "--subresource", Description: "If specified, edit will operate on the subresource of the requested object. Must be one of [status]. This flag is alpha and may change in the future."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--validate", Description: "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields."},
	prompt.Suggest{Text: "--windows-line-endings", Description: "Defaults to the line ending native to your platform. Usage:\n  kubectl edit (RESOURCE/NAME | -f FILENAME) [options] Use \"kubectl options\" for a list of global command-line options (applies to all commands)."},
}
