package cmdtree

import (
	"github.com/ActiveState/cli/internal/captain"
	"github.com/ActiveState/cli/internal/condition"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/primer"
	"github.com/ActiveState/cli/internal/runners/state"
	secretsapi "github.com/ActiveState/cli/pkg/platform/api/secrets"
	"github.com/ActiveState/cli/state/secrets"
)

// CmdTree manages a tree of captain.Command instances.
type CmdTree struct {
	cmd *captain.Command
}

// New prepares a CmdTree.
func New(prime *primer.Values, args ...string) *CmdTree {
	globals := newGlobalOptions()

	authCmd := newAuthCommand(prime)
	authCmd.AddChildren(
		newSignupCommand(prime),
		newLogoutCommand(prime),
	)

	exportCmd := newExportCommand(prime)
	exportCmd.AddChildren(
		newRecipeCommand(prime),
		newJWTCommand(prime),
		newPrivateKeyCommand(prime),
		newAPIKeyCommand(prime),
		newExportConfigCommand(prime),
		newExportGithubActionCommand(prime),
	)

	packagesCmd := newPackagesCommand(prime)
	packagesCmd.AddChildren(
		newPackagesAddCommand(prime),
		newPackagesUpdateCommand(prime),
		newPackagesRemoveCommand(prime),
		newPackagesImportCommand(prime),
		newPackagesSearchCommand(prime),
	)

	platformsCmd := newPlatformsCommand(prime)
	platformsCmd.AddChildren(
		newPlatformsSearchCommand(prime),
		newPlatformsAddCommand(prime),
		newPlatformsRemoveCommand(prime),
	)

	scriptsCmd := newScriptsCommand(prime)
	scriptsCmd.AddChildren(
		newScriptsEditCommand(prime),
	)

	languagesCmd := newLanguagesCommand(prime)
	languagesCmd.AddChildren(newLanguageUpdateCommand(prime))

	cleanCmd := newCleanCommand(prime)
	cleanCmd.AddChildren(
		newUninstallCommand(prime),
		newCacheCommand(prime),
		newConfigCommand(prime),
	)

	deployCmd := newDeployCommand(prime)
	deployCmd.AddChildren(
		newDeployInstallCommand(prime),
		newDeployConfigureCommand(prime),
		newDeploySymlinkCommand(prime),
		newDeployReportCommand(prime),
	)

	tutorialCmd := newTutorialCommand(prime)
	tutorialCmd.AddChildren(newTutorialProjectCommand(prime))

	eventsCmd := newEventsCommand(prime)
	eventsCmd.AddChildren(newEventsLogCommand(prime))

	stateCmd := newStateCommand(globals, prime)
	stateCmd.AddChildren(
		newActivateCommand(prime),
		newInitCommand(prime),
		newPushCommand(prime),
		newProjectsCommand(prime),
		authCmd,
		exportCmd,
		newOrganizationsCommand(prime),
		newRunCommand(prime),
		newShowCommand(prime),
		packagesCmd,
		platformsCmd,
		newHistoryCommand(prime),
		cleanCmd,
		languagesCmd,
		deployCmd,
		scriptsCmd,
		eventsCmd,
		newPullCommand(prime),
		newUpdateCommand(prime),
		newForkCommand(prime),
		newPpmCommand(prime),
		newInviteCommand(prime),
		tutorialCmd,
		newPrepareCommand(prime),
		newProtocolCommand(prime),
		newShimCommand(prime, args...),
	)

	applyLegacyChildren(stateCmd, globals)

	groups := []captain.CommandGroup{
		{
			Message: "Package Management:",
			Commands: []*captain.Command{
				stateCmd.FindSafe([]string{"packages", "add"}),
				stateCmd.FindSafe([]string{"packages", "remove"}),
				stateCmd.FindSafe([]string{"packages", "import"}),
				stateCmd.FindSafe([]string{"packages", "search"}),
				stateCmd.FindSafe([]string{"packages"}),
			},
		},
		{
			Message: "Environment Management:",
			Commands: []*captain.Command{
				stateCmd.FindSafe([]string{"activate"}),
				stateCmd.FindSafe([]string{"init"}),
				stateCmd.FindSafe([]string{"clean"}),
				stateCmd.FindSafe([]string{"deploy"}),
			},
		},
		{
			Message: "Platform Tools:",
			Commands: []*captain.Command{
				stateCmd.FindSafe([]string{"auth"}),
				stateCmd.FindSafe([]string{"invite"}),
				stateCmd.FindSafe([]string{"organizations"}),
				stateCmd.FindSafe([]string{"platforms"}),
				stateCmd.FindSafe([]string{"projects"}),
				stateCmd.FindSafe([]string{"show"}),
			},
		},
		{
			Message: "Automation:",
			Commands: []*captain.Command{
				stateCmd.FindSafe([]string{"events"}),
				stateCmd.FindSafe([]string{"scripts"}),
			},
		},
		{
			Message: "Version Control:",
			Commands: []*captain.Command{
				stateCmd.FindSafe([]string{"fork"}),
				stateCmd.FindSafe([]string{"pull"}),
				stateCmd.FindSafe([]string{"push"}),
				stateCmd.FindSafe([]string{"history"}),
			},
		},
		{
			Message: "Utilities:",
			Commands: []*captain.Command{
				stateCmd.FindSafe([]string{"update"}),
				stateCmd.FindSafe([]string{"export"}),
			},
		},
	}

	templater := captain.Templater{
		Cmd:           stateCmd,
		CommandGroups: groups,
	}
	stateCmd.SetUsageFunc(templater.RootCmdUsageFunc())

	return &CmdTree{
		cmd: stateCmd,
	}
}

type globalOptions struct {
	Verbose    bool
	Output     string
	Monochrome bool
}

func newGlobalOptions() *globalOptions {
	return &globalOptions{}
}

func newStateCommand(globals *globalOptions, prime *primer.Values) *captain.Command {
	opts := state.NewOptions()

	runner := state.New(opts, prime)
	cmd := captain.NewCommand(
		"state",
		"",
		locale.T("state_description"),
		prime.Output(),
		[]*captain.Flag{
			{
				Name:        "locale",
				Shorthand:   "l",
				Description: locale.T("flag_state_locale_description"),
				Persist:     true,
				Value:       &opts.Locale,
			},
			{
				Name:        "verbose",
				Shorthand:   "v",
				Description: locale.T("flag_state_verbose_description"),
				Persist:     true,
				OnUse: func() {
					if !condition.InTest() {
						logging.CurrentHandler().SetVerbose(true)
					}
				},
				Value: &globals.Verbose,
			},
			{
				Name:        "mono", // Name and Shorthand should be kept in sync with cmd/state/main.go
				Persist:     true,
				Description: locale.T("flag_state_monochrome_output_description"),
				Value:       &globals.Monochrome,
			},
			{
				Name:        "output", // Name and Shorthand should be kept in sync with cmd/state/main.go
				Shorthand:   "o",
				Description: locale.T("flag_state_output_description"),
				Persist:     true,
				Value:       &globals.Output,
			},
			{
				/* This option is only used for the vscode extension: It prevents the integrated terminal to close immediately after an error occurs, such that the user can read the message */
				Name:        "confirm-exit-on-error", // Name and Shorthand should be kept in sync with cmd/state/main.go
				Description: "prompts the user to press enter before exiting, when an error occurs",
				Persist:     true,
				Hidden:      true, // No need to add this to help messages
				Value:       &opts.ConfirmExit,
			},
			{
				Name:        "version",
				Description: locale.T("flag_state_version_description"),
				Value:       &opts.Version,
			},
		},
		[]*captain.Argument{},
		func(ccmd *captain.Command, args []string) error {
			if globals.Verbose {
				logging.CurrentHandler().SetVerbose(true)
			}

			return runner.Run(ccmd.Usage)
		},
	)

	return cmd
}

// Execute runs the CmdTree using the provided CLI arguments.
func (ct *CmdTree) Execute(args []string) error {
	return ct.cmd.Execute(args)
}

// Command returns the root command of the CmdTree
func (ct *CmdTree) Command() *captain.Command {
	return ct.cmd
}

// applyLegacyChildren will register any commands and expanders
func applyLegacyChildren(cmd *captain.Command, globals *globalOptions) {
	logging.Debug("register")

	secretsapi.InitializeClient()

	cmd.AddLegacyChildren(
		secrets.NewCommand(secretsapi.Get(), &globals.Output).Config(),
	)
}
