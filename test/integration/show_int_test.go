package integration

import (
	"strings"
	"testing"

	"github.com/ActiveState/cli/internal/testhelpers/e2e"
	"github.com/stretchr/testify/suite"
)

type ShowIntegrationTestSuite struct {
	suite.Suite
}

func (suite *ShowIntegrationTestSuite) TestShow() {
	ts := e2e.New(suite.T(), false)
	defer ts.Close()

	suite.PrepareActiveStateYAML(ts)

	cp := ts.Spawn("show")
	cp.Expect(`Name: project`)
	cp.Expect(`Organization: ActiveState`)
	cp.ExpectExitCode(0)
}

func (suite *ShowIntegrationTestSuite) PrepareActiveStateYAML(ts *e2e.Session) {
	asyData := strings.TrimSpace(`
project: "https://platform.activestate.com/ActiveState/project?commitID=00010001-0001-0001-0001-000100010001"
namespace: github.com/ActiveState/CodeIntel
environments: dev,qa,prod
platforms:
  - name: Linux64Label
    os: linux
    architecture: amd64
    libc: glibc-2.25
    compiler: gcc-7
  - name: Windows10Label
    os: windows
    version: 10
  - name: MacOSLabel
    os: macos
    version: 10.9
    compiler: clang-4
languages:
  - name: Go
    version: 1.10
    constraints:
        platform: Windows10Label,Linux64Label
        environment: dev,qa,prod
    packages:
      - name: golang.org/x/crypto
        version: "*"
        build:
          debug: $variable.DEBUG
      - name: gopkg.in/yaml.v2
        version: "2.*"
        build:
          override: --foo --bar --debug $variable.DEBUG --libDir $variable.PYTHONPATH --libc $platform.libc
constants:
  - name: DEBUG
    value: true
  - name: PYTHONPATH
    value: '%projectDir%/src:%projectDir%/tests'
    constraints:
        environment: dev,qa
  - name: PYTHONPATH
    value: '%projectDir%/src:%projectDir%/tests'
secrets:
  user:
    - name: user-secret
      description: user-secret-description
  project:
    - name: project-secret
      description: project-secret-description
events:
  - name: FIRST_INSTALL
    value: '%pythonExe% %projectDir%/setup.py prepare'
  - name: AFTER_UPDATE
    value: '%pythonExe% %projectDir%/setup.py prepare'
scripts:
  - name: tests
    value: pytest %projectDir%/tests
  - name: debug
    value: debug foo
`)

	ts.PrepareActiveStateYAML(asyData)
}

func TestShowIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(ShowIntegrationTestSuite))
}