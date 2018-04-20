package ux

import (
	"github.com/gigforks/termui"
	ui "github.com/gigforks/termui"
)

type Configuration struct {
}

func NewConfiguration() *Configuration {
	err := termui.Init()
	if err != nil {
		panic(err)
	}

	return &Configuration{}
}

func (c *Configuration) Run() {

	defer termui.Close()
	rows1 := [][]string{
		[]string{"header1", "header2", "header3"},
		[]string{"你好吗", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"2016", "10", "11"},
	}

	p := ui.NewPar(":PRESS q TO QUIT DEMO")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "Text Box"
	p.BorderFg = ui.ColorCyan
	p.Handle("/timer/1s", func(e ui.Event) {
		cnt := e.Data.(ui.EvtTimer)
		if cnt.Count%2 == 0 {
			p.TextFgColor = ui.ColorRed
		} else {
			p.TextFgColor = ui.ColorWhite
		}
	})
	table1 := termui.NewTable()
	table1.Rows = rows1
	table1.FgColor = termui.ColorWhite
	table1.BgColor = termui.ColorDefault
	table1.Y = 0
	table1.X = 0
	table1.Width = 62
	table1.Height = 7

	termui.Render(table1)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	ui.Handle("/timer/1s", func(e ui.Event) {
		termui.StopLoop()
	})
	termui.Loop()

}
