//go:generate codecgen -o types_codec.go types.go
package kube

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	APIGroup        = "vaultproject.io"
	APIVersion      = "v1"
	APIGroupVersion = APIGroup + "/" + APIVersion

	ResourceSecretClaims = "secretclaims"
)

var (
	GroupVersion = schema.GroupVersion{
		Group:   APIGroup,
		Version: APIVersion,
	}
)

func (in *SecretClaim) DeepCopyInto(out *SecretClaim) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = SecretSpec{
		Type:  in.Spec.Type,
		Path:  in.Spec.Path,
		Data:  in.Spec.Data,
		Renew: in.Spec.Renew,
	}
}

func (in *SecretClaim) DeepCopyObject() runtime.Object {
	out := SecretClaim{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *SecretClaimList) DeepCopyObject() runtime.Object {
	out := SecretClaimList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]SecretClaim, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}

type SecretSpec struct {
	Type  v1.SecretType          `json:"type"`
	Path  string                 `json:"path"`
	Data  map[string]interface{} `json:"data"`
	Renew int64                  `json:"renew"`
}

type SecretClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SecretSpec `json:"spec"`
}

type SecretClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []SecretClaim `json:"items"`
}

type SecretClaimManager interface {
	CreateOrUpdateSecret(claim *SecretClaim, force bool) error
	DeleteSecret(claim *SecretClaim) error
}
