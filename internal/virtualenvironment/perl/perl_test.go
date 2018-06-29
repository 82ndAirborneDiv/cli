package perl

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ActiveState/cli/internal/artifact"
	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/distribution"
	"github.com/ActiveState/cli/internal/environment"
)

func setup(t *testing.T) {
	root, _ := environment.GetRootPath()
	os.Chdir(filepath.Join(root, "test"))

	datadir := config.GetDataDir()
	os.RemoveAll(filepath.Join(datadir, "virtual"))
	os.RemoveAll(filepath.Join(datadir, "packages"))
	os.RemoveAll(filepath.Join(datadir, "languages"))
	os.RemoveAll(filepath.Join(datadir, "artifacts"))
}

func TestLanguage(t *testing.T) {
	venv := &VirtualEnvironment{}
	assert.Equal(t, "perl", venv.Language(), "Should return perl")
}

func TestDataDir(t *testing.T) {
	venv := &VirtualEnvironment{}
	assert.Empty(t, venv.DataDir())

	venv.SetDataDir("/foo")
	assert.NotEmpty(t, venv.DataDir(), "Should set the datadir")
}

func TestLanguageMeta(t *testing.T) {
	setup(t)

	venv := &VirtualEnvironment{}
	assert.Nil(t, venv.Artifact(), "Should not have artifact info")

	venv.SetArtifact(&artifact.Artifact{
		Meta: &artifact.Meta{
			Name: "test",
		},
		Path: "test",
	})
	assert.NotNil(t, venv.Artifact(), "Should have artifact info")
}

func TestLoadPackageFromPath(t *testing.T) {
	setup(t)

	venv := &VirtualEnvironment{}

	datadir := filepath.Join(os.TempDir(), "as-state-test")
	os.RemoveAll(datadir)
	os.Mkdir(datadir, os.ModePerm)
	venv.SetDataDir(datadir)

	dist, fail := distribution.Obtain()
	assert.NoError(t, fail.ToError())

	var language *artifact.Artifact
	for _, lang := range dist.Languages {
		if lang.Meta.Name == venv.Language() {
			language = lang
			break
		}
	}

	assert.NotNil(t, language, "Should retrieve language from dist")

	artf := dist.Artifacts[language.Hash][0]
	fail = venv.LoadArtifact(artf)
	if runtime.GOOS != "windows" {
		assert.NoError(t, fail.ToError(), "Loads artifact without errors")
	} else {
		// Since creating symlinks on Windows requires admin privilages for now,
		// artifacts should not load correctly.
		assert.Error(t, fail, "Symlinking requires admin privilages for now")
	}

	// Todo: Test with datadir as source, not the archived version
	if runtime.GOOS != "windows" {
		assert.FileExists(t, filepath.Join(datadir, "language", "lib", "artifact.json"), "Should create a package symlink")
	} else {
		// Since creating symlinks on Windows requires admin privilages for now,
		// the symlinked file should not exist.  Check if it was created or not. Skip if not.
		_, err := os.Stat(filepath.Join(datadir, "language", "lib", "artifact.json"))
		if err == nil {
			assert.FileExists(t, filepath.Join(datadir, "language", "lib", "artifact.json"), "Should create a package symlink")
		}
	}
}

func TestActivate(t *testing.T) {
	setup(t)

	venv := &VirtualEnvironment{}

	venv.SetArtifact(&artifact.Artifact{
		Meta: &artifact.Meta{
			Name:    "perl",
			Version: "5.26.1",
		},
		Path: "test",
	})

	datadir := config.GetDataDir()
	datadir = filepath.Join(datadir, "test")
	venv.SetDataDir(datadir)

	venv.Activate()

	assert.DirExists(t, filepath.Join(venv.DataDir(), "bin"))
}
