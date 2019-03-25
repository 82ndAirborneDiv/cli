package secrets

import (
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/keypairs"
	secretsModels "github.com/ActiveState/cli/pkg/platform/api/secrets/secrets_models"
)

// ShareFromDiff decrypts a source user's secrets that they are sharing and re-encrypts those secrets using
// the public key of a target user provided in the UserSecretDiff struct. This is effectively "copying" a set
// of secrets for use by another user.
func ShareFromDiff(sourceKeypair keypairs.Keypair, diff *secretsModels.UserSecretDiff) ([]*secretsModels.UserSecretShare, *failures.Failure) {
	targetPubKey, failure := keypairs.ParseRSAPublicKey(*diff.PublicKey)
	if failure != nil {
		return nil, failure
	}

	targetShares := make([]*secretsModels.UserSecretShare, len(diff.Shares))
	for idx, sourceShare := range diff.Shares {
		decrVal, failure := sourceKeypair.DecodeAndDecrypt(*sourceShare.Value)
		if failure != nil {
			return nil, failure
		}

		targetSecret, failure := targetPubKey.EncryptAndEncode(decrVal)
		if failure != nil {
			return nil, failure
		}

		targetShares[idx] = &secretsModels.UserSecretShare{
			ProjectID: sourceShare.ProjectID,
			Name:      sourceShare.Name,
			Value:     &targetSecret,
		}
	}
	return targetShares, nil
}
