package activate

import (
	"path/filepath"

	"github.com/ActiveState/cli/internal/analytics"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/defact"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/osutils"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/primer"
	"github.com/ActiveState/cli/internal/subshell"
	"github.com/ActiveState/cli/pkg/project"
	"github.com/spf13/viper"
)

type Activate struct {
	namespaceSelect  namespaceSelectAble
	activateCheckout CheckoutAble
	out              output.Outputer
	config           defact.DefaultConfigurer
	subshell         subshell.SubShell
}

type ActivateParams struct {
	Namespace     *project.Namespaced
	PreferredPath string
	Command       string
	Default       bool
}

type primeable interface {
	primer.Outputer
	primer.Subsheller
}

func NewActivate(prime primeable, namespaceSelect namespaceSelectAble, activateCheckout CheckoutAble) *Activate {
	return &Activate{
		namespaceSelect,
		activateCheckout,
		prime.Output(),
		viper.GetViper(),
		prime.Subshell(),
	}
}

func (r *Activate) Run(params *ActivateParams) error {
	return r.run(params, activationLoop)
}

func (r *Activate) run(params *ActivateParams, activatorLoop activationLoopFunc) error {
	logging.Debug("Activate with namespace=%v, path=%v", params.Namespace, params.PreferredPath)

	targetPath, err := r.setupPath(params.Namespace.String(), params.PreferredPath)
	if err != nil {
		if !params.Namespace.IsValid() {
			return failures.FailUserInput.Wrap(err)
		}
		err := r.activateCheckout.Run(params.Namespace.String(), targetPath)
		if err != nil {
			return err
		}
	}

	// Send google analytics event with label set to project namespace
	names, fail := project.ParseNamespaceOrConfigfile(params.Namespace.String(), filepath.Join(targetPath, constants.ConfigFileName))
	if fail != nil {
		names = &project.Namespaced{}
		logging.Debug("error resolving namespace: %v", fail.ToError())
	}
	analytics.EventWithLabel(analytics.CatRunCmd, "activate", names.String())

	if params.Command != "" {
		r.subshell.SetActivateCommand(params.Command)
	}

	return activatorLoop(r.out, r.config, r.subshell, targetPath, params.Default, activate)
}

func (r *Activate) setupPath(namespace string, preferredPath string) (string, error) {
	var (
		targetPath string
		err        error
	)

	switch {
	// Checkout via namespace (eg. state activate org/project) and set resulting path
	case namespace != "":
		targetPath, err = r.namespaceSelect.Run(namespace, preferredPath)
	// Use the user provided path
	case preferredPath != "":
		targetPath, err = preferredPath, nil
	// Get path from working directory
	default:
		targetPath, err = osutils.Getwd()
	}
	if err != nil {
		return "", err
	}

	proj, fail := project.FromPath(targetPath)
	if fail != nil {
		return targetPath, fail
	}

	return filepath.Dir(proj.Source().Path()), nil
}
