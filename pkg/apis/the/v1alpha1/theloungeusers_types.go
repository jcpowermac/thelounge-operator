package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TheloungeUsersSpec defines the desired state of TheloungeUsers
// +k8s:openapi-gen=true
type TheloungeUsersSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Password  string `json:"password,omitempty"`
}

// TheloungeUsersStatus defines the observed state of TheloungeUsers
// +k8s:openapi-gen=true
type TheloungeUsersStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TheloungeUsers is the Schema for the theloungeusers API
// +k8s:openapi-gen=true
type TheloungeUsers struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TheloungeUsersSpec   `json:"spec,omitempty"`
	Status TheloungeUsersStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TheloungeUsersList contains a list of TheloungeUsers
type TheloungeUsersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TheloungeUsers `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TheloungeUsers{}, &TheloungeUsersList{})
}
