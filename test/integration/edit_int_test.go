package integration

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/constraints"
	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/testhelpers/integration"
	"github.com/ActiveState/cli/pkg/projectfile"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type EditIntegrationTestSuite struct {
	integration.Suite
	cleanup func()
}

func (suite *EditIntegrationTestSuite) SetupTest() {
	suite.Suite.SetupTest()

	var tempDir string
	tempDir, suite.cleanup = suite.PrepareTemporaryWorkingDirectory("EditIntegrationTest")

	root := environment.GetRootPathUnsafe()
	editorScript := filepath.Join(root, "test", "integration", "assets", "editor", "main.go")

	fail := fileutils.CopyFile(editorScript, filepath.Join(tempDir, "editor", "main.go"))
	suite.Require().NoError(fail.ToError())

	configFileContent := strings.TrimSpace(`
project: "https://platform.activestate.com/EditOrg/EditProject?commitID=00010001-0001-0001-0001-000100010001"
scripts:
  - name: test-script
    value: echo "hello test"
    constraints:
      os: macos,linux
  - name: test-script
    value: echo hello test
    constraints:
      os: windows
`)

	projectFile := &projectfile.Project{}
	err := yaml.Unmarshal([]byte(configFileContent), projectFile)
	suite.Require().NoError(err)

	projectFile.SetPath(filepath.Join(tempDir, constants.ConfigFileName))
	fail = projectFile.Save()
	suite.Require().NoError(fail.ToError())

	editorScriptDir := filepath.Join(tempDir, "editor")
	suite.SetWd(editorScriptDir)

	var extension string
	if runtime.GOOS == "windows" {
		extension = ".exe"
	}
	suite.SpawnCustom("go", "build", "-o", "editor"+extension)
	suite.Wait()

	suite.SetWd(tempDir)
	suite.Require().FileExists(filepath.Join(editorScriptDir, "editor"+extension))
	suite.AppendEnv([]string{fmt.Sprintf("EDITOR=%s", filepath.Join(editorScriptDir, "editor"+extension))})
}

func (suite *EditIntegrationTestSuite) TearDownTest() {
	suite.Suite.TearDownTest()
	projectfile.Reset()
	suite.cleanup()
}

func (suite *EditIntegrationTestSuite) TestEdit() {
	suite.Spawn("scripts", "edit", "test-script")
	suite.Expect("Watching file changes")
	suite.Expect("Are you done editing?")
	// Can't consistently get this line detected on CI
	// suite.Expect("Script changes detected")
	suite.SendLine("Y")
	suite.Wait()
}

func (suite *EditIntegrationTestSuite) TestEdit_NonInteractive() {
	suite.AppendEnv([]string{"ACTIVESTATE_NONINTERACTIVE=true"})
	suite.Spawn("scripts", "edit", "test-script")
	suite.Expect("Watching file changes")
	// Can't consistently get this line detected on CI
	// suite.Expect("Script changes detected")
	suite.Quit()
}

func (suite *EditIntegrationTestSuite) TestEdit_UpdateCorrectPlatform() {
	suite.Spawn("scripts", "edit", "test-script")
	suite.SendLine("Y")
	suite.Wait()

	time.Sleep(time.Second * 2) // let CI env catch up

	project := projectfile.Get()
	for _, script := range project.Scripts {
		if script.Name == "test-script" {
			if !constraints.IsConstrained(script.Constraints) {
				suite.Contains(script.Value, "more info!")
			}
		}
	}
}

func TestEditIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(EditIntegrationTestSuite))
}
