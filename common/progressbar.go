package common

import (
	"github.com/schollz/progressbar/v3"
	"os"
)

func SetupProgressbar(size uint64, description string) *progressbar.ProgressBar {
	return progressbar.NewOptions64(int64(size),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionSetDescription(description),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionShowElapsedTimeOnFinish(),
		//progressbar.OptionClearOnFinish(),
		progressbar.OptionOnCompletion(func() {
			_, _ = os.Stderr.Write([]byte("\n"))
		}),
	)
}
