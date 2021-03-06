package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
type TrainedModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrainedModelSpec   `json:"spec,omitempty"`
	Status            TrainedModelStatus `json:"status,omitempty"`
}

// TrainedModelList contains a list of TrainedModel
// +kubebuilder:object:root=true
type TrainedModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// +listType=set
	Items []TrainedModel `json:"items"`
}

type TrainedModelSpec struct {
	// Required field for parent inference service
	InferenceService string `json:"inferenceService"`
	// Predictor model spec
	PredictorModel ModelSpec `json:"predictorModel"`
}
type ModelSpec struct {
	// Storage URI for the model repository
	StorageURI string `json:"storageUri"`
	// Machine Learning <framework name>
	// The values could be: "tensorflow","pytorch","sklearn","onnx","xgboost", "myawesomeinternalframework" etc.
	Framework string `json:"framework"`
	// Maximum memory this model will consume, this field is used to decide if a model server has enough memory to load this model.
	Memory resource.Quantity `json:"memory,omitempty"`
}
