package cmdtree

import (
	"github.com/ActiveState/cli/internal/primer"
	"github.com/ActiveState/cli/internal/runners/update"

	"github.com/ActiveState/cli/internal/captain"
	"github.com/ActiveState/cli/internal/locale"
)

func newUpdateCommand(prime *primer.Values) *captain.Command {
	runner := update.New(prime)
	params := update.Params{}

	cmd := captain.NewCommand(
		"update",
		locale.Tl("update_title", "Updating The State Tool"),
		locale.Tl("update_description", "Updates the State Tool to the latest available version"),
		prime.Output(),
		prime.Config(),
		[]*captain.Flag{
			{
				Name:        "set-channel",
				Description: locale.Tl("update_channel", "Switches to the given update channel, eg. 'release'."),
				Value:       &params.Channel,
			},
		},
		[]*captain.Argument{},
		func(cmd *captain.Command, args []string) error {
			return runner.Run(&params)
		},
	)
	cmd.SetGroup(UtilsGroup)
	cmd.SetSkipChecks(true)
	return cmd
}

func newUpdateLockCommand(prime *primer.Values) *captain.Command {
	runner := update.NewLock(prime)
	params := update.LockParams{}

	cmd := captain.NewCommand(
		"lock",
		locale.Tl("lock_title", "Lock the State Tool version"),
		locale.Tl("lock_description", "Lock the State Tool at the current version, this disables automatic updates."),
		prime.Output(),
		prime.Config(),
		[]*captain.Flag{
			{
				Name:        "set-channel",
				Description: locale.Tl("update_channel", "Switches to the given update channel, eg. 'release'."),
				Value:       &params.Channel,
			},
			{
				Name: "force",
				Description: locale.Tl(
					"flag_update_force",
					"Automatically confirm that you would like to update the State Tool version that your project is locked to."),
				Value: &params.Force,
			},
		},
		[]*captain.Argument{},
		func(cmd *captain.Command, args []string) error {
			return runner.Run(&params)
		},
	)
	cmd.SetSkipChecks(true)
	return cmd
}
