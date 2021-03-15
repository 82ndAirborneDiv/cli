package main

import (
	"fmt"
	"os"

	"github.com/ActiveState/cli/cmd/state-tray/internal/open"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/getlantern/systray"
	"github.com/gobuffalo/packr"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	box := packr.NewBox("assets")
	systray.SetIcon(box.Bytes("icon.ico"))
	systray.SetTooltip("ActiveState State Tool")

	mAbout := systray.AddMenuItem("About State Tool", "Information about the State Tool")

	systray.AddSeparator()

	mDoc := systray.AddMenuItem("Documentation", "Open State Tool Docs")

	mPlatform := systray.AddMenuItem("ActiveState Platform", "")
	mLearn := mPlatform.AddSubMenuItem("Learn", "Open ActiveState blog")
	mSupport := mPlatform.AddSubMenuItem("Support", "Open support page")
	mAccount := mPlatform.AddSubMenuItem("Account", "Open your account page")

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Exit", "")

	for {
		select {
		case <-mAbout.ClickedCh:
			logging.Debug("About event")
			err := open.Prompt("state --version")
			if err != nil {
				handleError(errs.Wrap(err, "Could not open command prompt"))
			}
		case <-mDoc.ClickedCh:
			logging.Debug("Documentation event")
			// Not implemented
		case <-mLearn.ClickedCh:
			logging.Debug("Learn event")
			// Not implemented
		case <-mSupport.ClickedCh:
			logging.Debug("Support event")
			// Not implemented
		case <-mAccount.ClickedCh:
			logging.Debug("Account event")
			// Not implemented
		case <-mQuit.ClickedCh:
			logging.Debug("Quit event")
			systray.Quit()
			return
		}
	}
}

func onExit() {
	// Not implemented
}

func handleError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	logging.Error("Systray encountered an error: %v", err)
}
