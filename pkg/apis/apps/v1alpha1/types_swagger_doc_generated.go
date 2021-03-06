/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_PetSet = map[string]string{
	"":       "PetSet represents a set of pods with consistent identities. Identities are defined as:\n - Network: A single stable DNS and hostname.\n - Storage: As many VolumeClaims as requested.\nThe PetSet guarantees that a given network identity will always map to the same storage identity. PetSet is currently in alpha and subject to change without notice.",
	"spec":   "Spec defines the desired identities of pets in this set.",
	"status": "Status is the current status of Pets in this PetSet. This data may be out of date by some window of time.",
}

func (PetSet) SwaggerDoc() map[string]string {
	return map_PetSet
}

var map_PetSetList = map[string]string{
	"": "PetSetList is a collection of PetSets.",
}

func (PetSetList) SwaggerDoc() map[string]string {
	return map_PetSetList
}

var map_PetSetSpec = map[string]string{
	"":                     "A PetSetSpec is the specification of a PetSet.",
	"replicas":             "Replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.",
	"selector":             "Selector is a label query over pods that should match the replica count. If empty, defaulted to labels on the pod template. More info: http://releases.k8s.io/v1.4.0-alpha.0/docs/user-guide/labels.md#label-selectors",
	"template":             "Template is the object that describes the pod that will be created if insufficient replicas are detected. Each pod stamped out by the PetSet will fulfill this Template, but have a unique identity from the rest of the PetSet.",
	"volumeClaimTemplates": "VolumeClaimTemplates is a list of claims that pets are allowed to reference. The PetSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pet. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.",
	"serviceName":          "ServiceName is the name of the service that governs this PetSet. This service must exist before the PetSet, and is responsible for the network identity of the set. Pets get DNS/hostnames that follow the pattern: pet-specific-string.serviceName.default.svc.cluster.local where \"pet-specific-string\" is managed by the PetSet controller.",
}

func (PetSetSpec) SwaggerDoc() map[string]string {
	return map_PetSetSpec
}

var map_PetSetStatus = map[string]string{
	"":                   "PetSetStatus represents the current state of a PetSet.",
	"observedGeneration": "most recent generation observed by this autoscaler.",
	"replicas":           "Replicas is the number of actual replicas.",
}

func (PetSetStatus) SwaggerDoc() map[string]string {
	return map_PetSetStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
