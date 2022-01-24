/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// CanaryStatus defines the observed state of Canary
type CanaryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// CanaryStrategy defines parameters for a Replica Based Canary
type CanarySpec struct {
	// CanaryService holds the name of a service which selects pods with canary version and don't select any pods with stable version.
	// +optional
	CanaryService string `json:"canaryService,omitempty" protobuf:"bytes,1,opt,name=canaryService"`
	// StableService holds the name of a service which selects pods with stable version and don't select any pods with canary version.
	// +optional
	StableService string `json:"stableService,omitempty" protobuf:"bytes,2,opt,name=stableService"`
	// Steps define the order of phases to execute the canary deployment
	// +optional
	Steps []CanaryStep `json:"steps,omitempty" protobuf:"bytes,3,rep,name=steps"`
	// TrafficRouting hosts all the supported service meshes supported to enable more fine-grained traffic routing
	TrafficRouting *TrafficRouting `json:"trafficRouting,omitempty" protobuf:"bytes,4,opt,name=trafficRouting"`
	// +optional
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty" protobuf:"bytes,5,opt,name=maxUnavailable"`
	// +optional
	MaxSurge *intstr.IntOrString `json:"maxSurge,omitempty" protobuf:"bytes,6,opt,name=maxSurge"`
	// +optional
	ScaleDownDelaySeconds *int32 `json:"scaleDownDelaySeconds,omitempty" protobuf:"varint,7,opt,name=scaleDownDelaySeconds"`
	// ScaleDownDelayRevisionLimit limits the number of old RS that can run at one time before getting scaled down
	// +optional
	ScaleDownDelayRevisionLimit *int32 `json:"scaleDownDelayRevisionLimit,omitempty" protobuf:"varint,8,opt,name=scaleDownDelayRevisionLimit"`
	// +optional
	AbortScaleDownDelaySeconds *int32 `json:"abortScaleDownDelaySeconds,omitempty" protobuf:"varint,9,opt,name=abortScaleDownDelaySeconds"`
	// +optional
	DynamicStableScale bool `json:"dynamicStableScale,omitempty" protobuf:"varint,10,opt,name=dynamicStableScale"`
}

type TrafficRouting struct {
	// Istio holds Istio specific configuration to route traffic
	Istio *IstioTrafficRouting `json:"istio,omitempty" protobuf:"bytes,1,opt,name=istio"`
	// Nginx holds Nginx Ingress specific configuration to route traffic
	Nginx *NginxTrafficRouting `json:"nginx,omitempty" protobuf:"bytes,2,opt,name=nginx"`
	// Nginx holds ALB Ingress specific configuration to route traffic
}

// NginxTrafficRouting configuration for Nginx ingress controller to control traffic routing
type NginxTrafficRouting struct {
	// AnnotationPrefix has to match the configured annotation prefix on the nginx ingress controller
	// +optional
	AnnotationPrefix string `json:"annotationPrefix,omitempty" protobuf:"bytes,1,opt,name=annotationPrefix"`
	// StableIngress refers to the name of an `Ingress` resource in the same namespace as the `Rollout`
	StableIngress string `json:"stableIngress" protobuf:"bytes,2,opt,name=stableIngress"`
	// +optional
	AdditionalIngressAnnotations map[string]string `json:"additionalIngressAnnotations,omitempty" protobuf:"bytes,3,rep,name=additionalIngressAnnotations"`
}

// IstioTrafficRouting configuration for Istio service mesh to enable fine grain configuration
type IstioTrafficRouting struct {
	// VirtualService references an Istio VirtualService to modify to shape traffic
	VirtualService *IstioVirtualService `json:"virtualService,omitempty" protobuf:"bytes,1,opt,name=virtualService"`
	// DestinationRule references an Istio DestinationRule to modify to shape traffic
	DestinationRule *IstioDestinationRule `json:"destinationRule,omitempty" protobuf:"bytes,2,opt,name=destinationRule"`
	// VirtualServices references a list of Istio VirtualService to modify to shape traffic
	VirtualServices []IstioVirtualService `json:"virtualServices,omitempty" protobuf:"bytes,3,opt,name=virtualServices"`
}

// IstioVirtualService holds information on the virtual service the rollout needs to modify
type IstioVirtualService struct {
	// Name holds the name of the VirtualService
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// A list of HTTP routes within VirtualService to edit. If omitted, VirtualService must have a single route of this type.
	Routes []string `json:"routes,omitempty" protobuf:"bytes,2,rep,name=routes"`
	// A list of TLS/HTTPS routes within VirtualService to edit. If omitted, VirtualService must have a single route of this type.
	TLSRoutes []TLSRoute `json:"tlsRoutes,omitempty" protobuf:"bytes,3,rep,name=tlsRoutes"`
}

// TLSRoute holds the information on the virtual service's TLS/HTTPS routes that are desired to be matched for changing weights.
type TLSRoute struct {
	// Port number of the TLS Route desired to be matched in the given Istio VirtualService.
	Port int64 `json:"port,omitempty" protobuf:"bytes,1,opt,name=port"`
	// A list of all the SNI Hosts of the TLS Route desired to be matched in the given Istio VirtualService.
	SNIHosts []string `json:"sniHosts,omitempty" protobuf:"bytes,2,rep,name=sniHosts"`
}

// IstioDestinationRule is a reference to an Istio DestinationRule to modify and shape traffic
type IstioDestinationRule struct {
	// Name holds the name of the DestinationRule
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// CanarySubsetName is the subset name to modify labels with canary ReplicaSet pod template hash value
	CanarySubsetName string `json:"canarySubsetName" protobuf:"bytes,2,opt,name=canarySubsetName"`
	// StableSubsetName is the subset name to modify labels with stable ReplicaSet pod template hash value
	StableSubsetName string `json:"stableSubsetName" protobuf:"bytes,3,opt,name=stableSubsetName"`
}

// CanaryStep defines a step of a canary deployment.
type CanaryStep struct {
	// SetWeight sets what percentage of the newRS should receive
	SetWeight *int32 `json:"setWeight,omitempty" protobuf:"varint,1,opt,name=setWeight"`
	// +optional
	Pause *Pause `json:"pause,omitempty" protobuf:"bytes,2,opt,name=pause"`
	// +optional
	SetCanaryScale *SetCanaryScale `json:"setCanaryScale,omitempty" protobuf:"bytes,3,opt,name=setCanaryScale"`
}

type SetCanaryScale struct {
	// Weight sets the percentage of replicas the newRS should have
	// +optional
	Weight *int32 `json:"weight,omitempty" protobuf:"varint,1,opt,name=weight"`
	// Replicas sets the number of replicas the newRS should have
	// +optional
	Replicas *int32 `json:"replicas,omitempty" protobuf:"varint,2,opt,name=replicas"`
	// MatchTrafficWeight cancels out previously set Replicas or Weight, effectively activating SetWeight
	// +optional
	MatchTrafficWeight bool `json:"matchTrafficWeight,omitempty" protobuf:"varint,3,opt,name=matchTrafficWeight"`
}

type Pause struct {
	// Duration the amount of time to wait before moving to the next step.
	// +optional
	Duration *intstr.IntOrString `json:"duration,omitempty" protobuf:"bytes,1,opt,name=duration"`
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// Canary is the Schema for the canaries API
type Canary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CanarySpec   `json:"spec,omitempty"`
	Status CanaryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CanaryList contains a list of Canary
type CanaryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Canary `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Canary{}, &CanaryList{})
}
