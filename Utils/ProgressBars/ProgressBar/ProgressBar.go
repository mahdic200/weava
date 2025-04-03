package ProgressBar

import (
	"fmt"
	"os"
	"time"

	"github.com/k0kubun/go-ansi"
	pb "github.com/schollz/progressbar/v3"
)

/* Store the returned value in a variable which in order to be able to use it
example :
bar := ProgressBar.Default("Doing something : ", 1000)
don't forget to clear the bar in case of any errors : bar.Exit()
*/
func Default(name string, total int) *pb.ProgressBar {
	bar := pb.NewOptions(
		total,
		pb.OptionSetWriter(ansi.NewAnsiStdout()),
		pb.OptionEnableColorCodes(true),
		pb.OptionSetDescription(name),
		pb.OptionSetWriter(os.Stderr),
		pb.OptionSetWidth(14),
		pb.OptionShowTotalBytes(true),
		pb.OptionThrottle(65*time.Millisecond),
		pb.OptionShowCount(),
		pb.OptionShowIts(),
		pb.OptionSetTheme(pb.Theme{
			Saucer: "[green]=[reset]",
			SaucerHead: "[green]>[reset]",
			SaucerPadding: " ",
			BarStart: "[",
			BarEnd: "]",
		}),
		pb.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		pb.OptionSpinnerType(14),
		pb.OptionFullWidth(),
		pb.OptionSetRenderBlankState(true),
	)

	return bar
}
