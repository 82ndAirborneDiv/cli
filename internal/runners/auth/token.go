package auth

import (
	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
	"github.com/ActiveState/cli/pkg/platform/authentication"
)

func tokenAuth(token string) error {
	auth := authentication.Get()
	return auth.AuthenticateWithModel(&mono_models.Credentials{
		Token: token,
	})
}
