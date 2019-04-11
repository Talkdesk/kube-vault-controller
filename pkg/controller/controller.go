package controller

import (
	"time"

	vaultapi "github.com/hashicorp/vault/api"
	"github.com/roboll/kube-vault-controller/pkg/kube"
	"github.com/roboll/kube-vault-controller/pkg/vault"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

type Controller struct {
	SecretClaimController *cache.Controller
}

type Config struct {
	Namespace  string
	SyncPeriod time.Duration
}

func New(config *Config, vconfig *vaultapi.Config, kconfig *rest.Config) (*Controller, error) {
	claimSource, err := newSecretClaimSource(kconfig, config.Namespace)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	vaultController, err := vault.NewController(vconfig, kconfig)
	if err != nil {
		return nil, err
	}

	watchList := cache.NewListWatchFromClient
	_, claimCtrl := cache.NewInformer(claimSource, &kube.SecretClaim{}, config.SyncPeriod, newSecretClaimHandler(vaultController))
	//_, secretCtrl := cache.NewInformer(secretSource, &v1.Secret{}, 500, newSecretHandler(vaultController, claims))

	return &Controller{
		SecretClaimController: claimCtrl,
	}, nil
}

func (ctrl *Controller) Run(stop chan struct{}) {
	claimStop := make(chan struct{})
	go ctrl.SecretClaimController.Run(claimStop)

	<-stop
	claimStop <- struct{}{}
}
