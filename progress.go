package display

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
)

var (
	progressTheme = progressbar.Theme{
		Saucer:        "[green]▓[reset]",
		SaucerHead:    "[green]▒[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}
)

func BarWithPrefix(max int, prefix, title string) *progressbar.ProgressBar {

	desc, width := GetBarDescription(max, prefix, title)

	pb := progressbar.NewOptions(max, progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(width),
		progressbar.OptionShowCount(),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionSetDescription(desc),
		progressbar.OptionOnCompletion(finishProgressBar),
		progressbar.OptionSetTheme(progressTheme),
	)

	pb.RenderBlank()
	return pb
}

func GetBarDescription(max int, prefix, title string) (string, int) {
	strMax := fmt.Sprintf("%d", max)

	width := LINEWIDTH - len(title) - len(prefix) - (len(strMax) * 2) - 12

	desc := title
	if prefix != "" {
		width = width - 1
		desc = fmt.Sprintf("%s %s", color.CyanString(prefix), title)
	}
	return desc, width
}

func finishProgressBar() {
	fmt.Println("")
}
