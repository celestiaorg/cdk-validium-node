package sequencesender

import (
	"context"

	"github.com/0xPolygon/cdk-validium-node/etherman/types"
	"github.com/0xPolygon/cdk-validium-node/log"
	"github.com/rollkit/celestia-openrpc/types/blob"
)

func (s *SequenceSender) submitCelestiaBlob(ctx context.Context, sequences []types.Sequence) error {

	log.Warn("Sequences: ", sequences)
	var blobs []*blob.Blob
	for _, seq := range sequences {
		if len(seq.BatchL2Data) > 0 {
			log.Warn("BatchL2Data", seq.BatchL2Data)
			dataBlob, err := blob.NewBlobV0(s.celestiaDA.Namespace, seq.BatchL2Data)
			blobs = append(blobs, dataBlob)
			if err != nil {
				log.Warn("Error creating blob", "err", err)
				return err
			}
		}
	}

	if len(blobs) == 0 {
		log.Warn("No BatchData to post to Celestia")
		return nil
	}

	commitments, err := blob.CreateCommitments(blobs)
	if err != nil {
		log.Warn("Error creating commitment", "err", err)
		return err
	}

	log.Info("Blobs: ", blobs)
	height, err := s.celestiaDA.Client.Blob.Submit(ctx, blobs)
	if err != nil {
		log.Warn("Blob Submission error", "err", err)
		return err
	}

	log.Info("Submmited blobs ", "height: ", height, " blobs: ", blobs)

	for index, seq := range sequences {
		s.ethTxManager.AddCommitment(ctx, seq.BatchNumber, height, commitments[index], nil)
	}

	return nil
}
