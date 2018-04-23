package ux

import (
	"github.com/AlexsJones/kubelord/kubernetes"
	"github.com/gigforks/termui"
)

type Configuration struct {
}

func (c *Configuration) Exit() {
	termui.Close()
}
func NewConfiguration() *Configuration {
	err := termui.Init()
	if err != nil {
		panic(err)
	}

	return &Configuration{}
}

func (c *Configuration) Run(curator *kubernetes.Curator) {

	table1 := termui.NewTable()
	table1.Rows = curator.Namespaces
	table1.FgColor = termui.ColorWhite
	table1.BgColor = termui.ColorDefault
	table1.Y = 0
	table1.X = 0
	table1.Width = 62
	table1.Height = 7

	termui.Render(table1)

	termui.Loop()
}
