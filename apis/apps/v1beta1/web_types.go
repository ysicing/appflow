/*
Copyright 2022 ysicing <i@ysicing.me>.

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WebSpec defines the desired state of Web
type WebSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// If unspecified, defaults to 1.
	Replicas *int32 `json:"replicas,omitempty"`
	Image    string `json:"image"`
	// +optional
	ImagePullSecrets string `json:"imagePullSecrets,omitempty"`
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
	// +optional
	Envs []corev1.EnvVar `json:"envs,omitempty"`
	// +optional
	Volumes Volumes `json:"volumes,omitempty"`
	// +optional
	Service Service `json:"service,omitempty"`
	// +optional
	Ingress Ingress `json:"ingress,omitempty"`
}

type Volumes struct {
	Name string `json:"name"`
	// +optional
	Type string `json:"type,omitempty"`
	// +optional
	Class string `json:"class,omitempty"`
	// +optional
	Mode string       `json:"mode,omitempty"`
	Path []VolumePath `json:"path"`
	// +optional
	Size string `json:"size,omitempty"`
}

type VolumePath struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Service struct {
	// +optional
	Type  string        `json:"type,omitempty"`
	Ports []ServicePort `json:"ports"`
}

type ServicePort struct {
	Port     int32           `json:"port"`
	Protocol corev1.Protocol `json:"protocol,omitempty"`
}

type Ingress struct {
	Class  string        `json:"class,omitempty"`
	Domain []IngressHost `json:"domain"`
}

type IngressHost struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
	Path string `json:"path,omitempty"`
	// +optional
	TLS string `json:"tls,omitempty"`
}

// WebStatus defines the observed state of Web
type WebStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	PVC        bool `json:"pvc"`
	Deployment bool `json:"deployment"`
	Service    bool `json:"service"`
	Ingress    bool `json:"ingress"`
	Ready      bool `json:"ready"`
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="PVC",type="boolean",JSONPath=".status.pvc",description="The already status of pvc"
//+kubebuilder:printcolumn:name="Deployment",type="boolean",JSONPath=".status.deployment",description="The already status of deployment"
//+kubebuilder:printcolumn:name="Service",type="boolean",JSONPath=".status.service",description="The already status of service"
//+kubebuilder:printcolumn:name="Ingress",type="boolean",JSONPath=".status.ingress",description="The already status of ingress"
//+kubebuilder:printcolumn:name="Ready",type="boolean",JSONPath=".status.ready",description="status of web"
//+kubebuilder:printcolumn:name="Image",type="string",JSONPath=".spec.image",description="Web Image"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC."

// Web is the Schema for the webs API
type Web struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebSpec   `json:"spec,omitempty"`
	Status WebStatus `json:"status,omitempty"`
}

//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true

// WebList contains a list of Web
type WebList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Web `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Web{}, &WebList{})
}
