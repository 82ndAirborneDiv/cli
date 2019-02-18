package auth

import (
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/print"
	"github.com/ActiveState/cli/pkg/platform/api"
	"github.com/ActiveState/cli/pkg/platform/api/models"
)

func tokenAuth() {
	loginOK, err := api.Authenticate(&models.Credentials{
		Token: Args.Token,
	})

	if err != nil {
		failures.Handle(err, locale.T("err_auth_failed_unknown_cause"))
		return
	}

	print.Info(locale.T("login_success_welcome_back", map[string]string{
		"Name": loginOK.Payload.User.Username,
	}))
}
