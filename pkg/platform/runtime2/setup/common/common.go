// Package common defines types that are shared between the Setup and the
// specialized implementations for specific build engines.
package common

import (
	"github.com/ActiveState/cli/pkg/platform/api/buildlogstream"
	runtime "github.com/ActiveState/cli/pkg/platform/runtime2"
	"github.com/ActiveState/cli/pkg/platform/runtime2/artifact"
)

// Setuper is the interface for an implementation of runtime setup functions
// These need to be specialized for each BuildEngine type
type Setuper interface {
	ArtifactSetup(artifactID runtime.ArtifactID) ArtifactSetuper
	PostInstall() error
}

// ArtifactSetuper is the interface for an implementation of artifact setup functions
// These need to be specialized for each BuildEngine type
type ArtifactSetuper interface {
	NeedsSetup() bool
	Move(tmpInstallDir string) error
	MetaDataCollection() error
	Relocate() error
}

type ArtifactChanges struct {
	Added   []runtime.ArtifactID
	Updated []runtime.ArtifactID
	Removed []runtime.ArtifactID
}
type MessageHandler interface {
	buildlogstream.MessageHandler

	// ChangeSummary summarizes the changes to the current project during the InstallRuntime() call.
	// This summary is printed as soon as possible, providing the State Tool user with an idea of the complexity of the requested build.
	// The arguments are for the changes introduced in the latest commit that this Setup is setting up.
	// TODO: Decide if we want to have a method to de-activate the change summary for activations where it does not make sense.
	ChangeSummary(artifacts map[runtime.ArtifactID]artifact.Artifact, requested ArtifactChanges, changed ArtifactChanges)
	ArtifactDownloadStarting(artifactName string)
	ArtifactDownloadCompleted(artifactName string, number, total int)
	ArtifactDownloadFailed(artifactName string, errorMsg string)
}
