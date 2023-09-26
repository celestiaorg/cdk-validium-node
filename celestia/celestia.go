package celestia

import (
	"context"
	"encoding/hex"
	"errors"

	openrpc "github.com/rollkit/celestia-openrpc"
	"github.com/rollkit/celestia-openrpc/types/share"
)

type CelestiaConfig struct {
	Enable      bool   `mapstructure:"Enable"`
	Rpc         string `mapstructure:"Rpc"`
	NamespaceId string `mapstructure:"NamespaceId"`
	AuthToken   string `mapstructure:"AuthToken"`
}

type CelestiaDA struct {
	Cfg       CelestiaConfig
	Client    *openrpc.Client
	Namespace share.Namespace
}

func NewCelestiaDA(cfg CelestiaConfig) (*CelestiaDA, error) {
	daClient, err := openrpc.NewClient(context.Background(), cfg.Rpc, cfg.AuthToken)
	if err != nil {
		return nil, err
	}

	if cfg.NamespaceId == "" {
		return nil, errors.New("namespace id cannot be blank")
	}
	nsBytes, err := hex.DecodeString(cfg.NamespaceId)
	if err != nil {
		return nil, err
	}

	namespace, err := share.NewBlobNamespaceV0(nsBytes)
	if err != nil {
		return nil, err
	}

	return &CelestiaDA{
		Cfg:       cfg,
		Client:    daClient,
		Namespace: namespace,
	}, nil
}
