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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	// +optional
	Replicas *int32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
	// +optional
	Env string `json:"env" protobuf:"bytes,2,opt,name=env"`
	// +optional
	WorkloadRef *ObjectReference `json:"workload_ref,omitempty" protobuf:"bytes,3,opt,name=workload_ref"`
	// +optional
	AppRepo *AppRepo `json:"appRepo" protobuf:"bytes,4,opt,name=appRepo"`
}

// ObjectRef holds a references to the Kubernetes object
type ObjectReference struct {
	// API Version of the referent
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,1,opt,name=apiVersion"`
	// Kind of the referent
	Kind string `json:"kind,omitempty" protobuf:"bytes,2,opt,name=kind"`
	// Name of the referent
	Name string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
}

type AppRepo struct {
	RepoPath  string `json:"repoPath,omitempty" protobuf:"bytes,1,opt,name=apiVersion"`
	RepoType  string `json:"repoType,omitempty" protobuf:"bytes,2,opt,name=repoType"`
	RepoKey   string `json:"repoKey,omitempty" protobuf:"bytes,3,opt,name=repoKey"`
	RepoValue string `json:"repoValue,omitempty" protobuf:"bytes,4,opt,name=repoValue"`
	RepoSub   string `json:"repoSub,omitempty" protobuf:"bytes,5,opt,name=repoSub"`
}

// ApplicationStatus defines the observed state of Application
type ApplicationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// Application is the Schema for the applications API
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec,omitempty"`
	Status ApplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
