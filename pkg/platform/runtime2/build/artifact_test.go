package build

import (
	"testing"

	"github.com/ActiveState/cli/pkg/platform/runtime2/testhelper"
	"github.com/stretchr/testify/assert"
)

// TestArtifactsFromRecipe ensures that we are able to parse a recipe correctly
// This is probably good to do, as it is more complicated
func TestArtifactsFromRecipe(t *testing.T) {
	t.Skip("not implemented yet")
	tests := []struct {
		Name                  string
		recipeName            string
		expectedArtifactNames []string // TODO: expect full artifact structure maybe
	}{
		{
			"camel recipe",
			"camel",
			[]string{},
		},
		{
			"alternative recipe",
			"alternative",
			[]string{},
		},
		{
			"alternative with bundles",
			"alternative-with-bundles",
			[]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			recipe := testhelper.LoadRecipe(t, tt.recipeName)
			/*res := */ ArtifactsFromRecipe(recipe)

			// TODO: ensure some expectations on result
		})
	}
}

func TestRequestedArtifactChanges(t *testing.T) {
	t.Skip("not implemented yet")
	tests := []struct {
		Name            string
		baseRecipeName  string
		newRecipeName   string
		expectedChanges ArtifactChanges
	}{
		{
			"no changes",
			"base",
			"base",
			ArtifactChanges{},
		},
		{
			"one package added",
			"base",
			"one-requirement-added",
			ArtifactChanges{},
		},
		{
			"one package removed",
			"base",
			"one-requirement-removed",
			ArtifactChanges{},
		},
		{
			"one package updated",
			"base",
			"one-requirement-updated",
			ArtifactChanges{},
		},
		{
			"complex changes",
			"base",
			"complex-changes",
			ArtifactChanges{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			old := testhelper.LoadRecipe(t, tt.baseRecipeName)
			new := testhelper.LoadRecipe(t, tt.newRecipeName)
			res := RequestedArtifactChanges(old, new)

			assert.Equal(t, tt.expectedChanges, res)
		})
	}
}

func TestResolvedArtifactChanges(t *testing.T) {
	t.Skip("not implemented yet")
	tests := []struct {
		Name            string
		baseRecipeName  string
		newRecipeName   string
		expectedChanges ArtifactChanges
	}{
		{
			"no changes",
			"base",
			"base",
			ArtifactChanges{},
		},
		{
			"one package added",
			"base",
			"one-requirement-added",
			ArtifactChanges{},
		},
		{
			"one package removed",
			"base",
			"one-requirement-removed",
			ArtifactChanges{},
		},
		{
			"one package updated",
			"base",
			"one-requirement-updated",
			ArtifactChanges{},
		},
		{
			"complex changes",
			"base",
			"complex-changes",
			ArtifactChanges{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			old := testhelper.LoadRecipe(t, tt.baseRecipeName)
			new := testhelper.LoadRecipe(t, tt.newRecipeName)
			res := ResolvedArtifactChanges(old, new)

			assert.Equal(t, tt.expectedChanges, res)
		})
	}
}

func TestArtifactDownloads(t *testing.T) {
	// TODO: test that we can extract download URIs from build status response
}
