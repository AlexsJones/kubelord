package ux

import (
	"os"
	"time"

	"github.com/AlexsJones/kubelord/kubernetes"
	"github.com/gigforks/termui"
)

//Configuration ...
type Configuration struct {
	lastFetch  time.Time
	dataBuffer chan [][]string
}

//NewConfiguration generation ...
func NewConfiguration() *Configuration {
	err := termui.Init()
	if err != nil {
		panic(err)
	}

	return &Configuration{dataBuffer: make(chan [][]string), lastFetch: time.Now()}
}

//Run the primary uiloop and datafetch cycle...
func (c *Configuration) Run(conf *kubernetes.Configuration, poll time.Duration) {

	//preview------------------------
	help := termui.NewPar("Filter out with (d)eployments (s)tatefulsets (c)ronjobs (r)eset/(q)uit")
	help.Border = false
	help.Width = termui.TermWidth()
	help.Height = 10
	help.PaddingTop = 5
	help.X = 0
	help.Y = termui.TermHeight()
	//--------------------------------
	bigview := termui.NewTable()
	bigview.FgColor = termui.ColorWhite
	bigview.Border = false
	bigview.BgColor = termui.ColorDefault
	bigview.Rows = [][]string{[]string{"Namespace", "Deployments", "Type", "Replicas", "Status"}}
	bigview.Width = termui.TermWidth()
	bigview.Height = termui.TermHeight()
	//--------------------------------
	termui.Body.AddRows(termui.NewRow(termui.NewCol(termui.TermWidth(), 0, bigview)), termui.NewCol(termui.TermWidth(), 0, help))
	termui.Body.Align()

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		bigview.Width = termui.TermWidth()
		termui.Clear()
		termui.Body.Align()
		termui.Render(termui.Body)
	})

	//Async fetch handler...
	go func() {
		//initial synchronization
		c.dataBuffer <- dataFetch(conf, 0)

		for {
			elapsed := time.Since(c.lastFetch)
			if elapsed > time.Second*5 {
				c.dataBuffer <- dataFetch(conf, int(elapsed.Seconds()))
				c.lastFetch = time.Now()
			}
		}
	}()

	termui.Handle("/timer/1s", func(e termui.Event) {
		//UI Run
		t := e.Data.(termui.EvtTimer)
		func(i int) {
			select {
			case data := <-c.dataBuffer:
				bigview.Rows = data
			default:
			}
			termui.Render(termui.Body)
		}(int(t.Count))
	})

	termui.Handle("/sys/kbd/q", func(e termui.Event) {
		termui.StopLoop()
		termui.Close()
		os.Exit(0)
	})

	termui.Loop()
}
