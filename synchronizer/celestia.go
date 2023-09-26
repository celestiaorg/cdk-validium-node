package synchronizer

import (
	"context"

	"github.com/0xPolygon/cdk-validium-node/log"
)

func (s *ClientSynchronizer) getBlobData(height uint64, commitment []byte) ([]byte, error) {

	log.Info("Requesting data from Celestia", "namespace", s.celestiaDA.Cfg.NamespaceId, "height", height)

	blob, err := s.celestiaDA.Client.Blob.Get(context.Background(), height, s.celestiaDA.Namespace, commitment)
	if err != nil {
		return nil, err
	}

	log.Info("Succesfully fetched data from Celestia", "namespace", s.celestiaDA.Namespace, "height", height, "commitment", blob.Commitment)

	return blob.Data, nil
}
