package clean

import (
	"errors"
	"os"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/primer"
)

type confirmAble interface {
	Confirm(title, message string, defaultChoice *bool) (bool, error)
}

type Uninstall struct {
	out         output.Outputer
	confirm     confirmAble
	configPath  string
	cachePath   string
	installPath string
}

type UninstallParams struct {
	Force bool
}

type primeable interface {
	primer.Outputer
	primer.Prompter
}

func NewUninstall(prime primeable) (*Uninstall, error) {
	return newUninstall(prime.Output(), prime.Prompt())
}

func newUninstall(out output.Outputer, confirm confirmAble) (*Uninstall, error) {
	installPath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	cfg := config.Get()

	return &Uninstall{
		out:         out,
		confirm:     confirm,
		installPath: installPath,
		configPath:  cfg.ConfigPath(),
		cachePath:   cfg.CachePath(),
	}, nil
}

func (u *Uninstall) Run(params *UninstallParams) error {
	if os.Getenv(constants.ActivatedStateEnvVarName) != "" {
		return errors.New(locale.T("err_uninstall_activated"))
	}

	if !params.Force {
		ok, err := u.confirm.Confirm(locale.T("confirm"), locale.T("uninstall_confirm"), new(bool))
		if err != nil {
			return err
		}
		if !ok {
			return nil
		}
	}

	return u.runUninstall()
}
