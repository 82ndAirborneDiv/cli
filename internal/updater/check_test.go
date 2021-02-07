package updater

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ActiveState/cli/internal/environment"
)

func setup(t *testing.T, withVersion bool) {
	cwd, err := environment.GetRootPath()
	require.NoError(t, err, "Should fetch cwd")
	path := filepath.Join(cwd, "internal", "updater", "testdata")
	if withVersion {
		path = filepath.Join(path, "withversion")
	}
	err = os.Chdir(path)
	require.NoError(t, err, "Should change dir without issue.")
}

func TestTimeout(t *testing.T) {
	info, err := timeout(func() (*Info, error) {
		return &Info{}, nil
	}, time.Second)
	assert.NoError(t, err, "no timeout")
	assert.NotNil(t, info, "no timeout")

	info, err = timeout(func() (*Info, error) {
		return nil, errors.New("some error")
	}, time.Second)
	assert.Nil(t, info, "non-timeout error")
	assert.Equal(t, err.Error(), "some error", "non-timeout error")

	info, err = timeout(func() (*Info, error) {
		time.Sleep(time.Second)
		return &Info{}, nil
	}, time.Millisecond)
	assert.Nil(t, info, "timeout")
	assert.Equal(t, err.Error(), "timeout", "timeout error")
}
