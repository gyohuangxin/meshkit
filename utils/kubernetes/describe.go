package kubernetes

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/kubectl/pkg/describe"
)

// ServiceOptions give control of which service to discover and which port to discover.
type DescriberOptions struct {
	Name       string // Name of the kubernetes obj
	Namespace  string // Namespace of the kubernetes obj
	ShowEvents bool
	ChunkSize  int64
	Type       DescribeType
}

type DescribeType int

const (
	Service = iota
	Pod
	Namespace
	Job
)

func Describe(client kubernetes.Interface, options DescriberOptions) (string, error) {
	var settings describe.DescriberSettings
	settings.ShowEvents = options.ShowEvents
	settings.ChunkSize = options.ChunkSize

	switch options.Type {
	case Pod:
		var describer *describe.PodDescriber
		describer.Interface = client
		return describer.Describe(options.Namespace, options.Name, settings)
	case Namespace:
		var describer *describe.NamespaceDescriber
		describer.Interface = client
		return describer.Describe(options.Namespace, options.Name, settings)
	case Service:
		var describer *describe.ServiceDescriber
		describer.Interface = client
		return describer.Describe(options.Namespace, options.Name, settings)
	case Job:
		var describer *describe.JobDescriber
		describer.Interface = client
		return describer.Describe(options.Namespace, options.Name, settings)
	default:
		return "", fmt.Errorf("invalid kubernetes object type")
	}
}
