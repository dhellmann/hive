package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// ClusterProvisionSpec defines the results of provisioning a cluster.
type ClusterProvisionSpec struct {

	// ClusterDeployment references the cluster deployment provisioned.
	ClusterDeployment corev1.LocalObjectReference `json:"clusterDeployment"`

	// PodSpec is the spec to use for the installer pod.
	PodSpec corev1.PodSpec `json:"podSpec"`

	// Attempt is which attempt number of the cluster deployment that this ClusterProvision is
	Attempt int `json:"attempt"`

	// Stage is the stage of provisioning that the cluster deployment has reached.
	Stage ClusterProvisionStage `json:"stage"`

	// ClusterID is a globally unique identifier for this cluster generated during installation. Used for reporting metrics among other places.
	ClusterID *string `json:"clusterID,omitempty"`

	// InfraID is an identifier for this cluster generated during installation and used for tagging/naming resources in cloud providers.
	InfraID *string `json:"infraID,omitempty"`

	// InstallLog is the log from the installer.
	InstallLog *string `json:"installLog,omitempty"`

	// Metadata is the metadata.json generated by the installer, providing metadata information about the cluster created.
	Metadata *runtime.RawExtension `json:"metadata,omitempty"`

	// AdminKubeconfigSecret references the secret containing the admin kubeconfig for this cluster.
	AdminKubeconfigSecret *corev1.LocalObjectReference `json:"adminKubeconfigSecret,omitempty"`

	// AdminPasswordSecret references the secret containing the admin username/password which can be used to login to this cluster.
	AdminPasswordSecret *corev1.LocalObjectReference `json:"adminPasswordSecret,omitempty"`

	// PrevClusterID is the cluster ID of the previous failed provision attempt.
	PrevClusterID *string `json:"prevClusterID,omitempty"`

	// PrevInfraID is the infra ID of the previous failed provision attempt.
	PrevInfraID *string `json:"prevInfraID,omitempty"`
}

// ClusterProvisionStatus defines the observed state of ClusterProvision.
type ClusterProvisionStatus struct {
	Job *corev1.LocalObjectReference `json:"job,omitempty"`

	// Conditions includes more detailed status for the cluster provision
	// +optional
	Conditions []ClusterProvisionCondition `json:"conditions,omitempty"`
}

// ClusterProvisionStage is the stage of provisioning.
type ClusterProvisionStage string

const (
	// ClusterProvisionStageProvisioning indicates that the cluster provision is ongoing
	ClusterProvisionStageProvisioning ClusterProvisionStage = "provisioning"
	// ClusterProvisionStageComplete indicates that the cluster provision completed successfully.
	ClusterProvisionStageComplete ClusterProvisionStage = "complete"
	// ClusterProvisionStageFailed indicates that the cluster provision failed.
	ClusterProvisionStageFailed ClusterProvisionStage = "failed"
)

// ClusterProvisionCondition contains details for the current condition of a cluster provision
type ClusterProvisionCondition struct {
	// Type is the type of the condition.
	Type ClusterProvisionConditionType `json:"type"`
	// Status is the status of the condition.
	Status corev1.ConditionStatus `json:"status"`
	// LastProbeTime is the last time we probed the condition.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// LastTransitionTime is the last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// Reason is a unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Message is a human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// ClusterProvisionConditionType is a valid value for ClusterProvisionCondition.Type
type ClusterProvisionConditionType string

const (
	// ClusterProvisionCompletedCondition is set when a cluster provision completes.
	ClusterProvisionCompletedCondition ClusterProvisionConditionType = "ClusterProvisionCompleted"

	// ClusterProvisionFailedCondition is set when a cluster provision fails.
	ClusterProvisionFailedCondition ClusterProvisionConditionType = "ClusterProvisionFailed"

	// ClusterProvisionJobCreated is set when the install job is created for a cluster provision.
	ClusterProvisionJobCreated ClusterProvisionConditionType = "ClusterProvisionJobCreated"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterProvision is the Schema for the clusterprovisions API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ClusterDeployment",type="string",JSONPath=".spec.clusterDeployment.name"
// +kubebuilder:printcolumn:name="Stage",type="string",JSONPath=".spec.stage"
// +kubebuilder:printcolumn:name="InfraID",type="string",JSONPath=".spec.infraID"
// +kubebuilder:resource:path=clusterprovisions
type ClusterProvision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterProvisionSpec   `json:"spec,omitempty"`
	Status ClusterProvisionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterProvisionList contains a list of ClusterProvision
type ClusterProvisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterProvision `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterProvision{}, &ClusterProvisionList{})
}
