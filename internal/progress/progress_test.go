package progress

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vbauerster/mpb"
)

type mockTask struct {
	Error error
}

func (mt *mockTask) mockFileSizeTask(cb FileSizeCallback) error {
	cb(10000)
	cb(20000)
	cb(10000)
	return mt.Error
}

// Test
func TestReportProgressDynamically(t *testing.T) {

	cases := []struct {
		Name   string
		Error  error
		UseMpb bool
	}{

		{
			Name:   "with progressbar",
			Error:  nil,
			UseMpb: true,
		},
		{
			Name:   "with progressbar and error",
			Error:  fmt.Errorf("test error"),
			UseMpb: true,
		},
		{
			Name:   "without progressbar",
			Error:  nil,
			UseMpb: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {

			var progress *mpb.Progress
			if tt.UseMpb {
				progress = mpb.New()
			}

			mt := mockTask{Error: nil}
			err := ReportProgressDynamically(mt.mockFileSizeTask, progress, 0)

			assert.NoError(t, err, "expected no error")

			if tt.UseMpb {
				progress.Wait()
			}
		})
	}

}
