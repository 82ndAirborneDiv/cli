package installation_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/installation"
	"github.com/phayes/permbits"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var stateToolTestFile string = "state"
var otherTestFile string = "other"
var installerTestFile string = "state-installer"

var stateToolTestFileContent []byte
var installerTestFileContent []byte
var updatedTestFileContent []byte = []byte("#!/bin/bash\necho updated;")

func init() {
	if runtime.GOOS == "windows" {
		stateToolTestFile = "state.exe"
		otherTestFile = "other.exe"
		installerTestFile = "state-installer.exe"
	}
}

func buildTestExecutable(t *testing.T, dir, exe string) {
	root, err := environment.GetRootPath()
	require.NoError(t, err)

	cmd := exec.Command(
		"go", "build", "-o", exe,
		filepath.Join(root, "internal", "testhelpers", "installation", dir),
	)
	err = cmd.Run()
	require.NoError(t, err)
}

func copyStateToolTestFile(t *testing.T, targetPath string) {
	if stateToolTestFileContent == nil {
		td, err := ioutil.TempDir("", "")
		require.NoError(t, err)
		fp := filepath.Join(td, stateToolTestFile)

		buildTestExecutable(t, "testcmd", fp)
		stateToolTestFileContent, err = ioutil.ReadFile(fp)
		require.NoError(t, err)
	}
	err := ioutil.WriteFile(targetPath, stateToolTestFileContent, 0755)
	require.NoError(t, err)
}

func copyInstallerTestFile(t *testing.T, targetPath string) {
	if installerTestFileContent == nil {
		td, err := ioutil.TempDir("", "")
		require.NoError(t, err)
		fp := filepath.Join(td, installerTestFile)

		buildTestExecutable(t, "testinst", fp)
		installerTestFileContent, err = ioutil.ReadFile(fp)
		require.NoError(t, err)
	}
	err := ioutil.WriteFile(targetPath, installerTestFileContent, 0755)
	require.NoError(t, err)
}

func initTempInstallDirs(t *testing.T, withAutoInstall bool) (string, string) {
	fromDir, err := ioutil.TempDir("", "from*")
	require.NoError(t, err)
	toDir, err := ioutil.TempDir("", "to*")
	require.NoError(t, err)
	for _, df := range []string{otherTestFile, stateToolTestFile} {
		// populate fromDir with a file that is going to be installed
		err = ioutil.WriteFile(filepath.Join(fromDir, df), updatedTestFileContent, 0775)
		require.NoError(t, err, "Failed to write test file %s", df)

	}
	// populate State Tool test file that gets replaced in installation directory
	copyStateToolTestFile(t, filepath.Join(toDir, stateToolTestFile))

	if withAutoInstall {
		copyInstallerTestFile(t, filepath.Join(fromDir, installerTestFile))
	}

	return fromDir, toDir
}

func assertPermissions(t *testing.T, fp string) {
	info, err := os.Stat(fp)
	require.NoError(t, err)
	pb := permbits.FileMode(info.Mode())
	assert.True(t, pb.UserRead(), "%s should be readable")
	if runtime.GOOS != "windows" {
		// Windows does not need an executable flag (just the correct file ending)
		assert.True(t, pb.UserExecute(), "%s should be executable")
	}
}

func assertSuccessfulInstallation(t *testing.T, toDir string) {
	for _, df := range []string{stateToolTestFile, otherTestFile} {
		fp := filepath.Join(toDir, df)
		assert.FileExists(t, fp, "Expected test file %s to exist", fp)
		b, err := ioutil.ReadFile(fp)
		require.NoError(t, err)
		if !bytes.Equal(updatedTestFileContent, b) {
			t.Errorf("Test file %s was not correctly updated", fp)
		}
		assertPermissions(t, fp)
	}
}

func assertRevertedInstallation(t *testing.T, toDir string) {
	fp := filepath.Join(toDir, stateToolTestFile)
	b, err := ioutil.ReadFile(fp)
	require.NoError(t, err)
	if !bytes.Equal(stateToolTestFileContent, b) {
		t.Error("State Tool test file was not correctly restored.")
	}
	assertPermissions(t, fp)
}

// TestInstallation tests that an installation is working if there are no obstacles like running processes
func TestInstallation(t *testing.T) {
	tests := []struct {
		Name                      string
		SimulateAdminInstallation bool
		ExpectSuccess             bool
	}{
		{"successful", false, true},
		{"update-without-permissions", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			from, to := initTempInstallDirs(t, false)

			if tt.SimulateAdminInstallation {
				// Simulate that a previous installation has been installed with administrator rights:
				// Remove the "Writable"-permission for installed files
				err := os.Chmod(to, 0550)
				require.NoError(t, err)
				err = os.Chmod(filepath.Join(to, stateToolTestFile), 0550)
				require.NoError(t, err)
			}

			inst, err := func() (*installation.Installation, error) {
				inst, err := installation.New(from, to)
				if err != nil {
					return nil, err
				}

				err = inst.Install()
				return inst, err
			}()
			if tt.ExpectSuccess {
				require.NoError(t, err)

				err = inst.Close()
				require.NoError(t, err)

				assertSuccessfulInstallation(t, to)
			} else {
				require.Error(t, err)
				assertRevertedInstallation(t, to)
			}

		})

	}
}

func TestInstallationWhileProcessesAreActive(t *testing.T) {
	from, to := initTempInstallDirs(t, false)

	// run the old command which waits for two seconds.
	cmd := exec.Command(filepath.Join(to, stateToolTestFile), "2")
	err := cmd.Start()
	require.NoError(t, err)

	inst, err := installation.New(from, to)
	require.NoError(t, err)

	errC := make(chan error)
	go func() {
		err := inst.Install()
		if err != nil {
			errC <- err
		}
		errC <- inst.Close()
	}()

	err = cmd.Wait()
	require.NoError(t, err, "%s", errs.Join(err, ": "))

	select {
	case err := <-errC:
		if runtime.GOOS == "windows" {
			assert.Error(t, err, "Installation should fail on Windows.")
			assertRevertedInstallation(t, to)
		} else {
			require.NoError(t, err)
			assertSuccessfulInstallation(t, to)
		}
	case <-time.After(time.Second * 2):
		t.Fatalf("Timeout waiting for installation to finish")
	}
}

// TestAutoUpdate tests that an executable can update itself, by spawning the
// installer process which eventually replaces the calling executable.
func TestAutoUpdate(t *testing.T) {
	tests := []struct {
		Name          string
		Timeout       string
		ExpectSuccess bool
	}{
		{
			"replaced-executable-is-running",
			"0",
			// when the replaced executable is still running, the auto-update should fail on Windows
			runtime.GOOS != "windows",
		},
		{
			"replaced-executable-shut-down",
			"2",
			// when the replaced executable is stopped, the auto-update should always pass
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			from, to := initTempInstallDirs(t, true)
			defer os.RemoveAll(from)
			defer os.RemoveAll(to)

			cfg, err := config.Get()
			require.NoError(t, err)
			configPath := cfg.ConfigPath()

			// run the executable (called stateToolTestFile) that is trying to
			// replace itself by triggering the installer
			c := exec.Command(filepath.Join(to, stateToolTestFile), from,
				filepath.Join(from, installerTestFile), tt.Timeout)
			var stdout bytes.Buffer
			var stderr bytes.Buffer
			c.Stdout = &stdout
			c.Stderr = &stderr
			// We set the CONFIGDIR to ensure that the installers log file is written there
			c.Env = os.Environ()
			c.Env = append(c.Env, "ACTIVESTATE_CLI_CONFIGDIR="+configPath)
			err = c.Run()
			require.NoError(t, err, "Error running auto-replacing test file: %v, stderr=%s", err, stderr.String())

			// regenerate the log file name for the installer from its PID
			pid, err := strconv.ParseInt(stdout.String(), 10, 32)
			require.NoError(t, err)
			assert.NotEqual(t, 0, pid)

			logFile := installation.LogfilePath(configPath, int(pid))

			// poll for successful auto-update
			for i := 0; i < 20; i++ {
				time.Sleep(time.Millisecond * 200)

				logs, err := ioutil.ReadFile(logFile)
				require.NoError(t, err)
				if strings.Contains(string(logs), "was successful") || strings.Contains(string(logs), "Installation failed") {
					break
				}
			}

			logs, err := ioutil.ReadFile(logFile)
			require.NoError(t, err)

			if tt.ExpectSuccess {
				assert.Containsf(t, string(logs), "was successful", "logs should contain 'was successful', got=%s", string(logs))
				assert.Contains(t, string(logs), "Target files=", "logs should contain 'Target files=', got=%s", logs)

				assertSuccessfulInstallation(t, to)
			} else {
				assert.Containsf(t, string(logs), "Installation failed", "logs should contains 'Installation failed', got=%s", string(logs))
				if runtime.GOOS == "windows" {
					// On Windows we expect the restore to fail on the State Tool executable.
					// It cannot be restored, because it is still running, but that is okay, because for the same reason it could not be modified beforehand.
					assert.Contains(t, string(logs), "Failed to restore some files")
				} else {
					assert.Contains(t, string(logs), "Successfully restored original files.")
				}

				assertRevertedInstallation(t, to)
			}
		})
	}
}
