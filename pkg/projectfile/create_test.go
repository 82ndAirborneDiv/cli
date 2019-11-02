package projectfile

import (
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/ActiveState/cli/internal/constants"

	"github.com/ActiveState/cli/internal/fileutils"

	"github.com/ActiveState/cli/internal/failures"
	"github.com/go-openapi/strfmt"
)

func Test_Create(t *testing.T) {
	var tempDir = fileutils.TempDirUnsafe()
	var uuid = strfmt.UUID("00010001-0001-0001-0001-000100010001")
	type args struct {
		org       string
		project   string
		directory string
		commitID  *strfmt.UUID
	}
	tests := []struct {
		name         string
		args         args
		want         *failures.Failure
		wantCreated  bool
		wantContents string
	}{
		{
			"orgName/projName",
			args{"orgName", "projName", tempDir, &uuid},
			nil,
			true,
			"orgName/projName?commitID=" + uuid.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args.org, tt.args.project, tt.args.commitID, tt.args.directory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createProjectFile() = %v, want %v", got, tt.want)
			}
			configFile := filepath.Join(tempDir, constants.ConfigFileName)
			if created := fileutils.FileExists(configFile); created != tt.wantCreated {
				t.Errorf("%s created: %v, but wanted: %v", constants.ConfigFileName, created, tt.wantCreated)
			}
			fileContents := fileutils.ReadFileUnsafe(configFile)
			if !strings.Contains(string(fileContents), tt.wantContents) {
				t.Errorf("%s does not contain: '%s', actual contents: '%v'", constants.ConfigFileName, tt.wantContents, string(fileContents))
			}
		})
	}
}
