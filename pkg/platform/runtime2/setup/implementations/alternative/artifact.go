package alternative

import "github.com/ActiveState/cli/pkg/platform/runtime2/model"

type ArtifactSetup struct {
}

func NewArtifactSetup(artifactID model.ArtifactID) *ArtifactSetup {
	return &ArtifactSetup{}
}

func (as *ArtifactSetup) NeedsSetup() bool {
	panic("implement me")
}

func (as *ArtifactSetup) Move(from string) error {
	panic("implement me")
}

func (as *ArtifactSetup) MetaDataCollection() error {
	panic("implement me")
}

func (as *ArtifactSetup) Relocate() error {
	panic("implement me")
}