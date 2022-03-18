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
	return &Bar{title: title, total: total, b: progressbar.NewOptions(total,
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
}

func (b *Bar) Add() {
	b.b.Add(1)
}
