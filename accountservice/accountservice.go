//go:generate mockgen -destination mock_accountservice/mock_accountservice.go github.com/altercpre/aparok-sync/accountservice Service
package accountservice

import (
	"github.com/altercpre/aparok-sync/app"
	"github.com/altercpre/aparok-sync/commonspace/object/accountdata"
)

const CName = "common.accountservice"

type Service interface {
	app.Component
	Account() *accountdata.AccountKeys
}

type Config struct {
	PeerId     string `yaml:"peerId"`
	PeerKey    string `yaml:"peerKey"`
	SigningKey string `yaml:"signingKey"`
}

type ConfigGetter interface {
	GetAccount() Config
}
