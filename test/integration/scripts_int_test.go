package integration

import (
	"strings"
	"testing"

	"github.com/ActiveState/cli/internal/testhelpers/e2e"
	"github.com/ActiveState/cli/internal/testhelpers/tagsuite"
	"github.com/stretchr/testify/suite"
)

type ScriptsIntegrationTestSuite struct {
	tagsuite.Suite
}

func (suite *ScriptsIntegrationTestSuite) setupConfigFile(ts *e2e.Session) {
	configFileContent := strings.TrimSpace(`
project: "https://platform.activestate.com/ScriptOrg/ScriptProject?commitID=00010001-0001-0001-0001-000100010001"
scripts:
  - name: first-script
    value: echo "first script"
    constraints:
      os: macos,linux
  - name: first-script
    value: echo first script
    constraints:
      os: windows
  - name: second-script
    value: print("second script")
    language: python3
`)

	ts.PrepareActiveStateYAML(configFileContent)
}

func TestScriptsIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(ScriptsIntegrationTestSuite))
}
