// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var runOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--annotations", Description: "Annotations to apply to the pod."},
	prompt.Suggest{Text: "--attach", Description: "If true, wait for the Pod to start running, and then attach to the Pod as if 'kubectl attach ...' were called.  Default false, unless '-i/--stdin' is set, in which case the default is true. With '--restart=Never' the exit code of the container process is returned."},
	prompt.Suggest{Text: "--cascade", Description: "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background."},
	prompt.Suggest{Text: "--command", Description: "If true and extra arguments are present, use them as the 'command' field in the container, rather than the 'args' field which is the default."},
	prompt.Suggest{Text: "--dry-run", Description: "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource."},
	prompt.Suggest{Text: "--env", Description: "Environment variables to set in the container."},
	prompt.Suggest{Text: "--expose", Description: "If true, create a ClusterIP service associated with the pod.  Requires `--port`."},
	prompt.Suggest{Text: "--field-manager", Description: "Name of the manager used to track field ownership."},
	prompt.Suggest{Text: "-f", Description: "to use to replace the resource."},
	prompt.Suggest{Text: "--filename", Description: "to use to replace the resource."},
	prompt.Suggest{Text: "--force", Description: "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation."},
	prompt.Suggest{Text: "--grace-period", Description: "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion)."},
	prompt.Suggest{Text: "--image", Description: "The image for the container to run."},
	prompt.Suggest{Text: "--image-pull-policy", Description: "The image pull policy for the container.  If left empty, this value will not be specified by the client and defaulted by the server."},
	prompt.Suggest{Text: "-k", Description: "Process a kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "--kustomize", Description: "Process a kustomization directory. This flag can't be used together with -f or -R."},
	prompt.Suggest{Text: "-l", Description: "Comma separated labels to apply to the pod. Will override previous values."},
	prompt.Suggest{Text: "--labels", Description: "Comma separated labels to apply to the pod. Will override previous values."},
	prompt.Suggest{Text: "--leave-stdin-open", Description: "If the pod is started in interactive mode or with stdin, leave stdin open after the first attach completes. By default, stdin will be closed after the first attach completes."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file)."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file)."},
	prompt.Suggest{Text: "--override-type", Description: "The method used to override the generated object: json, merge, or strategic."},
	prompt.Suggest{Text: "--overrides", Description: "An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field."},
	prompt.Suggest{Text: "--pod-running-timeout", Description: "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running"},
	prompt.Suggest{Text: "--port", Description: "The port that this container exposes."},
	prompt.Suggest{Text: "--privileged", Description: "If true, run the container in privileged mode."},
	prompt.Suggest{Text: "-q", Description: "If true, suppress prompt messages."},
	prompt.Suggest{Text: "--quiet", Description: "If true, suppress prompt messages."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--restart", Description: "The restart policy for this Pod.  Legal values [Always, OnFailure, Never]."},
	prompt.Suggest{Text: "--rm", Description: "If true, delete the pod after it exits.  Only valid when attaching to the container, e.g. with '--attach' or with '-i/--stdin'."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future."},
	prompt.Suggest{Text: "--show-managed-fields", Description: "If true, keep the managedFields when printing objects in JSON or YAML format."},
	prompt.Suggest{Text: "-i", Description: "Keep stdin open on the container in the pod, even if nothing is attached."},
	prompt.Suggest{Text: "--stdin", Description: "Keep stdin open on the container in the pod, even if nothing is attached."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--timeout", Description: "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object"},
	prompt.Suggest{Text: "-t", Description: "Allocate a TTY for the container in the pod."},
	prompt.Suggest{Text: "--tty", Description: "Allocate a TTY for the container in the pod."},
	prompt.Suggest{Text: "--wait", Description: "If true, wait for resources to be gone before returning. This waits for finalizers. Usage:\n  kubectl run NAME --image=image [--env=\"key=value\"] [--port=port] [--dry-run=server|client] [--overrides=inline-json] [--command] -- [COMMAND] [args...] [options] Use \"kubectl options\" for a list of global command-line options (applies to all commands)."},
}
