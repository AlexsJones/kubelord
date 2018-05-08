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
	help := termui.NewPar("(q)uit")
	help.Border = true
	help.Width = termui.TermWidth()
	help.Height = 3
	help.X = 0
	help.Y = 0
	//--------------------------------
	bigview := termui.NewTable()
	bigview.FgColor = termui.ColorWhite
	bigview.Border = true
	bigview.BgColor = termui.ColorDefault
	bigview.Rows = [][]string{[]string{"Namespace", "Deployments", "Type", "Replicas", "Status"}}
	bigview.Width = termui.TermWidth()
	bigview.Height = termui.TermHeight()
	//--------------------------------
	showDeployments := true
	showStatefulSets := true
	showCronJobs := true
	//--------------------------------
	termui.Body.AddRows(termui.NewCol(3, 0, help, bigview))
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
			if elapsed > poll {
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
				bigview.Rows = func() [][]string {
					return data
				}()
				termui.Body.Align()
				bigview.Rows[0][0] = s.Next()
				termui.Render(termui.Body)

			default:
				bigview.Rows[0][0] = s.Next()
				termui.Render(termui.Body)
			}

		}()
	})
	termui.Handle("/sys/kbd/d", func(e termui.Event) {
		showDeployments = !showDeployments
	})
	termui.Handle("/sys/kbd/s", func(e termui.Event) {
		showStatefulSets = !showStatefulSets
	})
	termui.Handle("/sys/kbd/c", func(e termui.Event) {
		showCronJobs = !showCronJobs
	})
	termui.Handle("/sys/kbd/r", func(e termui.Event) {
		showCronJobs = false
		showDeployments = false
		showStatefulSets = false
	})
	termui.Handle("/sys/kbd/q", func(e termui.Event) {
		termui.StopLoop()
		termui.Close()
		os.Exit(0)
	})

	termui.Loop()
}
