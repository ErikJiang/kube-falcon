/*


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
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DeployObjectSpec defines the desired state of DeployObject
type DeployObjectSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Minimum=0

	// Replicas is pod replica num
	Replicas *int32 `json:"replicas"`

	// +kubebuilder:validation:MinLength=0

	// Image is container image address
	Image string `json:"image"`

	// Resources describes the compute resource requirements
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// EnvVar represents an environment variable present in a Container.
	// +optional
	Env []corev1.EnvVar `json:"env,omitempty"`

	// ServicePort contains information on service's port.
	// +optional
	Ports []corev1.ServicePort `json:"ports,omitempty"`
}

// DeployObjectStatus defines the observed state of DeployObject
type DeployObjectStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// DeploymentStatus is the most recently observed status of the Deployment.
	appsv1.DeploymentStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// DeployObject is the Schema for the deployobjects API
type DeployObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeployObjectSpec   `json:"spec,omitempty"`
	Status DeployObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeployObjectList contains a list of DeployObject
type DeployObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeployObject `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeployObject{}, &DeployObjectList{})
}
