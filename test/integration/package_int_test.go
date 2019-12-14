package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ActiveState/cli/internal/testhelpers/integration"
)

type PackageIntegrationTestSuite struct {
	integration.Suite
}

func (suite *PackageIntegrationTestSuite) TestPackage_listing() {
	tempDir, cleanup := suite.PrepareTemporaryWorkingDirectory("package_no_args")
	defer cleanup()

	suite.AppendEnv([]string{"ACTIVESTATE_CLI_DISABLE_RUNTIME=false"})

	asyData := `project: "https://platform.activestate.com/ActiveState-CLI/Python3"`
	suite.Require().NoError(setupASY(tempDir, asyData))

	suite.Run("simple", func() {
		suite.Spawn("package")
		suite.Expect("Name")
		suite.Wait()
	})

	suite.Run("with commit", func() {
		suite.Spawn("package", "--commit", "c780f643-724b-49bb-aca9-194e3c072f64")
		suite.Expect("Name")
		suite.Wait()
	})

	suite.Run("with commit junk val", func() {
		suite.Spawn("package", "--commit", "junk")
		suite.Expect("Cannot obtain")
		suite.Wait()
	})

	suite.Run("with commit unknown id", func() {
		suite.Spawn("package", "--commit", "00010001-0001-0001-0001-000100010001")
		suite.Expect("No data")
		suite.Wait()
	})
}

func TestPackageIntegrationTestSuite(t *testing.T) {
	_ = suite.Run

	integration.RunParallel(t, new(PackageIntegrationTestSuite))
}
