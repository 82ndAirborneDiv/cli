package auth

import (
	"github.com/ActiveState/cli/internal/failures"
	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
	"github.com/ActiveState/cli/pkg/platform/authentication"
)

func tokenAuth() *failures.Failure {
	auth := authentication.Get()
	return auth.AuthenticateWithModel(&mono_models.Credentials{
		Token: Flags.Token,
	})
}
