package ux

import (
	"os"
	"time"

	"github.com/AlexsJones/kubelord/kubernetes"
	"github.com/gigforks/termui"
	spin "github.com/tj/go-spin"
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

	s := spin.New()

	//Async fetch handler...
	go func() {
		//initial synchronization
		c.dataBuffer <- dataFetch(conf, "")

		for {
			elapsed := time.Since(c.lastFetch)
			if elapsed > time.Second*5 {
				c.dataBuffer <- dataFetch(conf, "")
				c.lastFetch = time.Now()
			}
		}
	}()

	termui.Handle("/timer/1s", func(e termui.Event) {
		//UI Run
		func() {
			select {
			case data := <-c.dataBuffer:
				bigview.Rows = data
				termui.Body.Align()
				bigview.Rows[0][0] = s.Next()
				termui.Render(termui.Body)
				//TODO: It is much faster to do the data filtering here
			default:
				bigview.Rows[0][0] = s.Next()
				termui.Render(termui.Body)
			}

		}()
	})

	termui.Handle("/sys/kbd/q", func(e termui.Event) {
		termui.StopLoop()
		termui.Close()
		os.Exit(0)
	})

	termui.Loop()
}
