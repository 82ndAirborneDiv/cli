package ppm

import (
	"fmt"
	"strings"

	"github.com/ActiveState/cli/internal/analytics"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/primer"
	"github.com/ActiveState/cli/internal/runbits"
	"github.com/ActiveState/cli/pkg/projectfile"
)

type Shim struct {
	out     output.Outputer
	project *projectfile.Project
}

func NewShim(prime *primer.Values) *Shim {
	return &Shim{
		out:     prime.Output(),
		project: prime.Projectfile(),
	}
}

func (s *Shim) RunPPM(args ...string) error {
	s.printForwardInfo("ppm", "packages", "ppm_print_forward")
	return s.shim("ppm", "packages", args...)
}

func (s *Shim) RunInstall(args ...string) error {
	s.printForwardInfo("install", "packages add", "ppm_print_forward_failure")
	return s.shim("install", "packages add", args...)
}

func (s *Shim) RunUpgrade(args ...string) error {
	s.printForwardInfo("upgrade", "packages update", "ppm_print_forward_failure")
	return s.shim("upgrade", "packages update", args...)
}

func (s *Shim) RunRemove(args ...string) error {
	s.printForwardInfo("remove", "packages remove", "ppm_print_forward_failure")
	return s.shim("remove", "packages remove", args...)
}

func (s *Shim) RunList(args ...string) error {
	s.printForwardInfo("list", "packages", "ppm_print_forward")
	return s.shim("list", "packages", args...)
}

func (s *Shim) shim(intercepted, replaced string, args ...string) error {
	err := s.executeShim(intercepted, replaced, args...)
	if err != nil {
		analytics.EventWithLabel(analytics.CatPPMShimCmd, intercepted, fmt.Sprintf("error: %v", errs.Join(err, " :: ").Error()))
	} else {
		analytics.EventWithLabel(analytics.CatPPMShimCmd, intercepted, "success")
	}
	return err
}

func (s *Shim) executeShim(intercepted, replaced string, args ...string) error {
	if s.project == nil {
		// TODO: Replace this function call when conversion flow is complete
		analytics.Event(analytics.CatPPMShimCmd, "tutorial")
		return tutorial()
	}

	commands := strings.Split(replaced, " ")
	replacedArgs := args
	if len(commands) > 1 {
		replaced = commands[0]
		replacedArgs = commands[1:]
		replacedArgs = append(replacedArgs, args...)
	}

	forwarded := []string{replaced}
	forwarded = append(forwarded, replacedArgs...)
	return runbits.Invoke(s.out, forwarded...)
}

func tutorial() error {
	// Placeholder until conversion flow is complete
	return nil
}

func (s *Shim) printForwardInfo(intercepted, replaced, localeID string) {
	forwarded := []string{"state", replaced}
	s.out.Print(locale.Tr(localeID, strings.Join(forwarded, " "), intercepted))
}

func (s *Shim) PrintSuggestion(ppmIntent, newCommand, docLink string) error {
	s.out.Print(locale.Tr("ppm_print_suggestion", ppmIntent, newCommand, docLink))
	return nil
}

func (s *Shim) PrintDefault() error {
	s.out.Print(strings.TrimSpace(locale.T("ppm_header_message")))
	return nil
}

func (s *Shim) PrintMain() error {
	s.out.Print(locale.T("ppm_print_main"))
	return nil
}
