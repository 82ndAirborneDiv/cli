package platforms

import (
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/pkg/platform/model"
)

// Search manages the searching execution context.
type Search struct{}

// NewSearch prepares a search execution context for use.
func NewSearch() *Search {
	return &Search{}
}

// Run executes the search behavior.
func (s *Search) Run() (*SearchResult, error) {
	logging.Debug("Execute platforms search")

	return newSearchResult()
}

// SearchResult represents the output data of a search.
type SearchResult struct {
	Platforms []*Platform `json:"platforms"`
}

func newSearchResult() (*SearchResult, error) {
	platforms, fail := model.FetchPlatforms()
	if fail != nil {
		return nil, fail
	}

	result := SearchResult{
		Platforms: makePlatformsFromModelPlatforms(platforms),
	}

	return &result, nil
}
