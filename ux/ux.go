package ux

import (
	"fmt"
	"log"
	"os"
	"time"

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

func (c *Configuration) Run(conf *kubernetes.Configuration, poll time.Duration) {

	//preview------------------------
	p1 := termui.NewPar("Warming up the processing core...\n")
	p1.Border = false
	p1.Width = 40
	p1.Height = 10
	p1.X = 0
	p1.Y = 0

	//--------------------------------
	bigview := termui.NewTable()
	bigview.FgColor = termui.ColorWhite
	bigview.BgColor = termui.ColorDefault
	bigview.Rows = [][]string{[]string{"Namespace", "Deployments", "Type", "Replicas", "Status"}}
	bigview.Width = termui.TermWidth()
	bigview.Height = termui.TermHeight()
	//--------------------------------
	//Namespaces forms the first loop for recursing deployments within

	drawBigView := func() {
		dataSet := [][]string{[]string{"Namespace", "Deployments", "Type", "Replicas", "Status"}}
		namespacelist, err := conf.GetNamespaces()
		if err != nil {
			log.Println(fmt.Sprintf("namespaces: %s", err.Error()))

		}

		for _, namespace := range namespacelist.Items {
			deploymentlist, err := conf.GetDeployments(namespace.Name)
			if err != nil {
				log.Println(fmt.Sprintf("deployment: %s", err.Error()))

			}
			for _, deployment := range deploymentlist.Items {
				row := []string{namespace.Name, deployment.Name,
					"Deployment", fmt.Sprintf("%d/%d", int(*deployment.Spec.Replicas), int(deployment.Status.AvailableReplicas)),
					deployment.Status.Conditions[len(deployment.Status.Conditions)-1].Message}
				dataSet = append(dataSet, row)
			}

			stslist, err := conf.GetStatefulSets(namespace.Name)
			if err != nil {
				log.Println(fmt.Sprintf("statefulset: %s", err.Error()))

			}
			for _, sts := range stslist.Items {

				status := ""
				if len(sts.Status.Conditions) > 0 {
					status = sts.Status.Conditions[len(sts.Status.Conditions)-1].Message
				}
				row := []string{namespace.Name, sts.Name, "StatefulSet", fmt.Sprintf("%d/%d", int(*sts.Spec.Replicas), int(sts.Status.CurrentReplicas)),
					status}
				dataSet = append(dataSet, row)
			}
		}
		bigview.Rows = dataSet
		termui.Render(bigview)
	}

	termui.Render(p1)

	drawBigView()

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		bigview.Width = termui.TermWidth()
		termui.Clear()
		termui.Render(bigview)
	})

	termui.Handle("/timer/5s", func(e termui.Event) {

		drawBigView()
	})

	termui.Handle("/sys/kbd/q", func(e termui.Event) {

		termui.StopLoop()
		termui.Close()
		os.Exit(0)
	})

	termui.Loop()
}
