package alternative

import "github.com/ActiveState/cli/pkg/platform/runtime2/build"

// var _ runtime.Setuper = &Setup{}
// var _ runtime.ArtifactSetuper = &ArtifactSetup{}

type Setup struct {
}

type ArtifactSetup struct {
}

func NewSetup() *Setup {
	return &Setup{}
}

func NewArtifactSetup(artifactID build.ArtifactID) *ArtifactSetup {
	return &ArtifactSetup{}
}

func (s *Setup) PostInstall() error {
	panic("implement me")
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
