package installer

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/logging"
)

type Installation struct {
	fromDir string
	toDir   string
	backups []string
}

func backupFiles(targetFiles []string) ([]string, error) {
	var renamed []string
	for _, t := range targetFiles {
		if fileutils.TargetExists(t) {
			newName := fmt.Sprintf("%s.bak", t)
			if fileutils.TargetExists(newName) {
				_ = os.Remove(newName)
			}
			var err error
			if runtime.GOOS == "windows" {
				// Note: We could also rename the files here, in which case the installation would always pass on Windows (even if the executable is still running), but we would not be able to remove the .bak file after the installation (if the executable is still running).
				// We opted for this solution, as it is more compliant with other Windows installations, where the installation requires all existing programmes to be stopped prior to updates.
				err = fileutils.CopyFile(t, newName)
			} else {
				err = os.Rename(t, newName)
			}
			if err != nil {
				// restore already renamed files and return with error
				_ = restoreFiles(renamed)
				return nil, errs.Wrap(err, "Failed to backup file %s", t)
			}
			renamed = append(renamed, newName)
		}
	}
	return renamed, nil
}

func restoreFiles(backupFiles []string) error {
	var errors []error
	for _, b := range backupFiles {
		origName := strings.TrimSuffix(b, ".bak")
		err := os.Rename(b, origName)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return errs.Wrap(errors[0], "Failed to restore some files.")
	}
	return nil
}

func New(fromDir, toDir string) *Installation {
	return &Installation{
		fromDir, toDir, nil,
	}
}

func (i *Installation) RemoveBackup() error {
	var es []error
	for _, b := range i.backups {
		err := os.Remove(b)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			es = append(es, err)
		}
	}
	if len(es) > 0 {
		return errs.Wrap(es[0], "Failed to remove some back-up files")
	}

	return nil
}

func (i *Installation) BackupFiles() error {
	// Todo: https://www.pivotaltracker.com/story/show/177600107
	// Get target file paths.
	var targetFiles []string
	for _, file := range fileutils.ListDir(i.fromDir, false) {
		targetFile := filepath.Join(i.toDir, filepath.Base(file))
		targetFiles = append(targetFiles, targetFile)
	}
	logging.Debug("Target files=%s", strings.Join(targetFiles, ","))

	backups, err := backupFiles(targetFiles)
	if err != nil {
		return errs.Wrap(err, "Backup of existing files failed.")
	}
	i.backups = backups
	return nil
}

func (i *Installation) Rollback() error {
	return restoreFiles(i.backups)
}

func (i *Installation) Install() error {
	if err := i.BackupFiles(); err != nil {
		return errs.Wrap(err, "Failed to backup original files.")
	}
	if err := fileutils.CopyAndRenameFiles(i.fromDir, i.toDir); err != nil {
		return errs.Wrap(err, "Failed to copy installation files to dir %s", i.toDir)
	}
	if err := InstallSystemFiles(i.toDir); err != nil {
		return errs.Wrap(err, "Installation of system files failed.")
	}

	return nil
}
