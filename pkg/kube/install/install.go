package install

import (
	"github.com/roboll/kube-vault-controller/pkg/kube"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(kube.GroupVersion,
		&kube.SecretClaim{},
		&kube.SecretClaimList{},
		&metav1.ListOptions{},
		&metav1.DeleteOptions{},
	)
	metav1.AddToGroupVersion(scheme, kube.GroupVersion)
	return nil
}

const importPrefix = "github.com/roboll/kube-vault-controller/pkg/kube"

//var accessor = meta.NewAccessor()

// availableVersions lists all known external versions for this group from most preferred to least preferred
var availableVersions = []schema.GroupVersion{kube.GroupVersion}

func init() {
	SchemeBuilder.AddToScheme(scheme.Scheme)
}
