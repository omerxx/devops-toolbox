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

type Reporter struct {
	Kind    string `json:"kind,omitempty"`
	Key     string `json:"key,omitempty"`
	Channel string `json:"channel,omitempty"`
}

// PodTrackerSpec defines the desired state of PodTracker
type PodTrackerSpec struct {
	Name   string   `json:"name,omitempty"`
	Report Reporter `json:"report,omitempty"`
}

// PodTrackerStatus defines the observed state of PodTracker
type PodTrackerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodTracker is the Schema for the podtrackers API
type PodTracker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodTrackerSpec   `json:"spec,omitempty"`
	Status PodTrackerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodTrackerList contains a list of PodTracker
type PodTrackerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodTracker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodTracker{}, &PodTrackerList{})
}
