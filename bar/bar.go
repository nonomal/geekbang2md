package bar

import (
	"github.com/schollz/progressbar/v3"
)

type Interface interface {
	Add()
}

type Bar struct {
	title string
	total int
	b     *progressbar.ProgressBar
}

func NewBar(title string, total int) *Bar {
	b := &Bar{title: title, total: total + 1, b: progressbar.NewOptions(total,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetDescription("[Download]: "+title),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)}
	b.b.Add(1)
	return b
}

func (b *Bar) Add() {
	b.b.Add(1)
}
