package ux

import (
	"fmt"
	"log"
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

	//--------------------------------
	bigview := termui.NewTable()
	bigview.FgColor = termui.ColorWhite
	bigview.BgColor = termui.ColorDefault
	bigview.Rows = [][]string{[]string{"Namespace", "Deployments", "Type", "Replicas", "Status"}}
	bigview.Width = termui.TermWidth()
	bigview.Height = termui.TermHeight()
	//--------------------------------
	//Namespaces forms the first loop for recursing deployments within
	namespacelist, err := conf.GetNamespaces()
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, namespace := range namespacelist.Items {
		deploymentlist, err := conf.GetDeployments(namespace.Name)
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, deployment := range deploymentlist.Items {
			row := []string{namespace.Name, deployment.Name,
				"Deployment", fmt.Sprintf("%d/%d", int(*deployment.Spec.Replicas), int(deployment.Status.AvailableReplicas)),
				deployment.Status.Conditions[len(deployment.Status.Conditions)-1].Message}
			bigview.Rows = append(bigview.Rows, row)
		}

		stslist, err := conf.GetStatefulSets(namespace.Name)
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, sts := range stslist.Items {

			row := []string{namespace.Name, sts.Name, "StatefulSet", fmt.Sprintf("%d/%d", int(*sts.Spec.Replicas), int(sts.Status.CurrentReplicas)),
				sts.Status.Conditions[len(sts.Status.Conditions)-1].Message}
			bigview.Rows = append(bigview.Rows, row)
		}
	}
	//Render body -------------------

	termui.Render(bigview)

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		termui.Body.Width = termui.TermWidth()
		termui.Body.Align()
		termui.Clear()
		termui.Render(termui.Body)
	})

	termui.Handle("/timer/1s", func(e termui.Event) {

	})

	termui.Handle("/sys/kbd/q", func(e termui.Event) {
		termui.StopLoop()
	})

	termui.Loop()
}
