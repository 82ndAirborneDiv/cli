package integration

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/osutils"
	"github.com/ActiveState/cli/internal/testhelpers/e2e"
	"github.com/ActiveState/cli/internal/testhelpers/tagsuite"
	"github.com/ActiveState/cli/pkg/projectfile"
	"github.com/ActiveState/termtest"
	"github.com/stretchr/testify/suite"
)

type openKeyFn func(path string) (osutils.RegistryKey, error)

type cmdEnv struct {
	openKeyFn openKeyFn
	// whether this updates the system environment
	userScope bool
}

func newCmdEnv(userScope bool) *cmdEnv {
	openKeyFn := osutils.OpenSystemKey
	if userScope {
		openKeyFn = osutils.OpenUserKey
	}
	return &cmdEnv{
		openKeyFn: openKeyFn,
		userScope: userScope,
	}
}

func (c *cmdEnv) set(name, newValue string) error {
	key, err := c.openKeyFn(getEnvironmentPath(c.userScope))
	if err != nil {
		return err
	}
	defer key.Close()

	_, valType, err := key.GetStringValue(name)
	if err != nil {
		return err
	}
	return osutils.SetStringValue(key, name, valType, newValue)
}

func (c *cmdEnv) get(name string) (string, error) {
	key, err := c.openKeyFn(getEnvironmentPath(c.userScope))
	if err != nil {
		return "", err
	}
	defer key.Close()

	v, _, err := key.GetStringValue(name)
	return v, err
}

func getEnvironmentPath(userScope bool) string {
	if userScope {
		return "Environment"
	}
	return `SYSTEM\ControlSet001\Control\Session Manager\Environment`
}

func scriptPath(t *testing.T, legacy bool) string {
	name := "install.ps1"
	if runtime.GOOS != "windows" {
		name = "install.sh"
	}
	if legacy {
		name = "legacy-" + name
	}
	root := environment.GetRootPathUnsafe()
	subdir := "installers"

	exec := filepath.Join(root, subdir, name)
	if !fileutils.FileExists(exec) {
		t.Fatalf("Could not find install script %s", exec)
	}

	return exec
}

type InstallScriptsIntegrationTestSuite struct {
	tagsuite.Suite
	cfg    projectfile.ConfigGetter
	server *http.Server
}

func expectLegacyStateToolInstallation(cp *termtest.ConsoleProcess, addToPathAnswer string) {
	cp.Expect("Installing to")
	cp.Expect("Continue?")
	cp.SendLine("y")
	cp.Expect("Fetching the latest version")
	cp.Expect("Allow $PATH to be appended in your")
	cp.SendLine(addToPathAnswer)
	cp.Expect("State Tool installation complete")
}

func expectStateToolInstallation(cp *termtest.ConsoleProcess) {
	cp.Expect("Installing to")
	cp.Expect("Continue?")
	cp.SendLine("y")
	cp.Expect("Fetching the latest version")
	cp.Expect("State Tool installation complete")
}

func expectStateToolInstallationWindows(cp *termtest.ConsoleProcess) {
	cp.Expect("Installing to")
	cp.Expect("Continue?")
	cp.SendLine("y")
	cp.Expect("Fetching the latest version")
	cp.Expect("State Tool successfully installed to")
}

func expectDefaultActivation(cp *termtest.ConsoleProcess) {
	cp.Expect("Activating Virtual Environment")
	cp.Expect("Choose Destination")
	cp.Send("")
	cp.Expect("Cloning Repository")
	cp.Expect("Installing")
	cp.ExpectLongString("Successfully configured ActiveState/Perl-5.32 as the global default project")
	cp.Expect("activated state")
	cp.SendLine("exit")
}

func env(useTestBranch bool) []string {
	var res []string
	if useTestBranch {
		res = append(res, "_TEST_UPDATE_URL=http://localhost:"+testPort)
	}
	return res
}

func (suite *InstallScriptsIntegrationTestSuite) TestLegacyInstallSh() {
	if runtime.GOOS == "windows" {
		suite.T().SkipNow()
	}
	suite.OnlyRunForTags(tagsuite.InstallScripts, tagsuite.Critical)

	ts := e2e.New(suite.T(), false, env(true)...)
	defer ts.Close()

	script := scriptPath(suite.T(), true)

	cp := ts.SpawnCmdWithOpts("bash", e2e.WithArgs(script, "-t", ts.Dirs.Work))
	expectLegacyStateToolInstallation(cp, "n")
	cp.Expect("State Tool Installed")
	cp.ExpectExitCode(0)
}

func (suite *InstallScriptsIntegrationTestSuite) TestInstallSh() {
	if runtime.GOOS == "windows" {
		suite.T().SkipNow()
	}
	suite.OnlyRunForTags(tagsuite.InstallScripts, tagsuite.Critical)

	tests := []struct {
		Name        string
		TestInstall bool
		Channel     string
	}{
		{"install-local-test-update", true, constants.BranchName},
		// Todo Replace the target branch for this test to release, as soon as we have a working deployment there.
		{"install-release", false, "beta"},
	}

	for _, tt := range tests {
		if !tt.TestInstall {
			// Todo update on beta
			suite.T().Skipf("Non-local State Tool installations will only work once we have a new State Tool installer deployed.")
		}
		suite.Run(tt.Name, func() {
			ts := e2e.New(suite.T(), false, env(tt.TestInstall)...)
			defer ts.Close()

			script := scriptPath(suite.T(), false)

			cp := ts.SpawnCmdWithOpts("bash", e2e.WithArgs(script, "-t", ts.Dirs.Work, "-b", tt.Channel))
			expectStateToolInstallation(cp)
			cp.Expect("State Tool Installed")
			cp.ExpectExitCode(0)
		})
	}
}

func (suite *InstallScriptsIntegrationTestSuite) TestInstallPerl5_32Default() {
	suite.OnlyRunForTags(tagsuite.InstallScripts, tagsuite.Critical)
	suite.runInstallTest("-c", "state activate ActiveState/Perl-5.32 --default")
}

func (suite *InstallScriptsIntegrationTestSuite) TestInstallPerl5_32ActivateDefault() {
	suite.OnlyRunForTags(tagsuite.InstallScripts, tagsuite.Critical)
	suite.runInstallTest("--activate-default", "ActiveState/Perl-5.32")
}

func (suite *InstallScriptsIntegrationTestSuite) TestInstallPs1() {
	if runtime.GOOS != "windows" {
		suite.T().SkipNow()
	}
	suite.OnlyRunForTags(tagsuite.InstallScripts, tagsuite.Critical)

	tests := []struct {
		Name        string
		TestInstall bool
		Channel     string
	}{
		{"install-local-test-update", true, constants.BranchName},
		// Todo Replace the target branch for this test to release, as soon as we have a working deployment there.
		{"install-release", false, "beta"},
	}

	for _, tt := range tests {
		if !tt.TestInstall {
			// Todo update on beta
			suite.T().Skipf("Non-local State Tool installations will only work once we have a new State Tool installer deployed.")
		}
		suite.Run(tt.Name, func() {
			ts := e2e.New(suite.T(), false, env(tt.TestInstall)...)
			defer ts.Close()

			script := scriptPath(suite.T(), false)

			isAdmin, err := osutils.IsWindowsAdmin()
			suite.Require().NoError(err, "Could not determine if running as administrator")

			cmdEnv := newCmdEnv(!isAdmin)
			oldPathEnv, err := cmdEnv.get("PATH")
			suite.Require().NoError(err, "could not get PATH")

			defer func() {
				err := cmdEnv.set("PATH", oldPathEnv)
				suite.Assert().NoError(err, "Unexpected error re-setting paths")
			}()

			cp := ts.SpawnCmdWithOpts("powershell.exe", e2e.WithArgs(script, "-t", ts.Dirs.Work, "-b", tt.Channel))
			expectStateToolInstallationWindows(cp)
			cp.ExpectExitCode(0)

			pathEnv, err := cmdEnv.get("PATH")
			suite.Require().NoError(err, "could not get PATH")
			paths := strings.Split(pathEnv, string(os.PathListSeparator))
			suite.Assert().Contains(paths, ts.Dirs.Work, "Could not find installation path in PATH")
		})
	}
}

func (suite *InstallScriptsIntegrationTestSuite) TestInstallPerl5_32DefaultWindows() {
	suite.OnlyRunForTags(tagsuite.InstallScripts, tagsuite.Critical)
	suite.runInstallTestWindows("-c", "\"state activate ActiveState/Perl-5.32 --default\"")
}

func (suite *InstallScriptsIntegrationTestSuite) TestInstallPerl5_32_ActivateDefaultWindows() {
	suite.OnlyRunForTags(tagsuite.InstallScripts, tagsuite.Critical)
	suite.runInstallTestWindows("-activate-default", "ActiveState/Perl-5.32")
}

func (suite *InstallScriptsIntegrationTestSuite) runInstallTest(installScriptArgs ...string) {
	if runtime.GOOS != "linux" {
		suite.T().SkipNow()
	}

	ts := e2e.New(suite.T(), false, env(true)...)
	defer ts.Close()

	script := scriptPath(suite.T(), false)

	cp := ts.SpawnCmdWithOpts(
		"bash",
		e2e.WithArgs("-c", fmt.Sprintf("cp $HOME/.bashrc %s/bashrc.bak", ts.Dirs.Work)),
	)
	cp.ExpectExitCode(0)

	defer func() {
		cp = ts.SpawnCmdWithOpts(
			"bash",
			e2e.WithArgs("-c", fmt.Sprintf("cp %s/.bashrc.bak $HOME/.bashrc", ts.Dirs.Work)),
		)
	}()

	computedCommand := append([]string{script, "-t", ts.Dirs.Work, "-b", constants.BranchName}, installScriptArgs...)

	cp.ExpectExitCode(0)
	cp = ts.SpawnCmdWithOpts(
		"bash",
		e2e.WithArgs(computedCommand...),
		e2e.AppendEnv("ACTIVESTATE_CLI_DISABLE_RUNTIME=false", "SHELL=bash"),
	)
	expectStateToolInstallation(cp)

	expectDefaultActivation(cp)
	cp.ExpectExitCode(0)

	// we need to run an interactive bash session to ensure that the modified ~/.bashrc is being parsed
	cp = ts.SpawnCmd("bash")
	cp.SendLine("echo $PATH; exit")
	// Expect Global Binary directory on PATH
	globalBinDir := filepath.Join(ts.Dirs.Cache, "bin")
	cp.ExpectLongString(globalBinDir, 1*time.Second)
	// expect State Tool Installation directory
	cp.ExpectLongString(ts.Dirs.Work, 1*time.Second)
	cp.ExpectExitCode(0)
}

func (suite *InstallScriptsIntegrationTestSuite) runInstallTestWindows(installScriptArgs ...string) {
	if runtime.GOOS != "windows" {
		suite.T().SkipNow()
	}

	ts := e2e.New(suite.T(), false, env(true)...)
	defer ts.Close()

	script := scriptPath(suite.T(), false)

	isAdmin, err := osutils.IsWindowsAdmin()
	suite.Require().NoError(err, "Could not determine if running as administrator")

	cmdEnv := newCmdEnv(!isAdmin)
	oldPathEnv, err := cmdEnv.get("PATH")
	suite.Require().NoError(err, "could not get PATH")

	defer func() {
		err := cmdEnv.set("PATH", oldPathEnv)
		suite.Assert().NoError(err, "Unexpected error re-setting paths")
	}()

	computedCommand := append([]string{script, "-t", ts.Dirs.Work, "-b", constants.BranchName}, installScriptArgs...)

	cp := ts.SpawnCmdWithOpts(
		"powershell.exe",
		e2e.WithArgs(computedCommand...),
		e2e.AppendEnv("ACTIVESTATE_CLI_DISABLE_RUNTIME=false"))
	expectStateToolInstallationWindows(cp)
	expectDefaultActivation(cp)
	cp.ExpectExitCode(0)

	pathEnv, err := cmdEnv.get("PATH")
	suite.Require().NoError(err, "could not get PATH")
	paths := strings.Split(pathEnv, string(os.PathListSeparator))
	// The global binary directory is only added to the PATH for non-Administrator users
	if !isAdmin {
		suite.Assert().Contains(paths, filepath.Join(ts.Dirs.Cache, "bin"), "Could not find global binary directory on PATH")
	}
	suite.Assert().Contains(paths, ts.Dirs.Work, "Could not find installation path in PATH")
}

func (suite *InstallScriptsIntegrationTestSuite) BeforeTest(suiteName, testName string) {
	var err error
	root, err := environment.GetRootPath()
	suite.Require().NoError(err)
	testUpdateDir := filepath.Join(root, "build", "update")
	suite.Require().DirExists(testUpdateDir, "You need to run `state run generate-updates` for this test to work.")
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(testUpdateDir)))
	suite.server = &http.Server{Addr: "localhost:" + testPort, Handler: mux}
	go func() {
		_ = suite.server.ListenAndServe()
	}()

	suite.cfg, err = config.Get()
	suite.Require().NoError(err)
}

func (suite *InstallScriptsIntegrationTestSuite) AfterTest(suiteName, testName string) {
	err := suite.server.Shutdown(context.Background())
	suite.Require().NoError(err)
}

func TestInstallScriptsIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(InstallScriptsIntegrationTestSuite))
}
