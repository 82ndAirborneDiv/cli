// +build !windows

package runtime

import (
	"github.com/ActiveState/archiver"
	"github.com/ActiveState/cli/internal/progress"
)

// InstallerExtension is used to identify whether an artifact is one that we should care about
const InstallerExtension = ".tar.gz"

// Archiver returns the archiver to use
func Archiver() archiver.Archiver {
	return archiver.DefaultTarGz
}

// Unarchiver returns the unarchiver to use
func Unarchiver() archiver.Unarchiver {
	return archiver.DefaultTarGz
}

// UnarchiverWithProgress returns the ProgressUnarchiver to use
func UnarchiverWithProgress() *progress.TarGzArchiveReader {
	return &progress.TarGzArchiveReader{*archiver.DefaultTarGz}
}
