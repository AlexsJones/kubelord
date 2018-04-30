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
	ls := termui.NewList()
	ls.Items = func() []string {

		ns, err := conf.GetNamespaces()
		if err != nil {
			return []string{}
		}
		o := []string{}
		for _, n := range ns.Items {
			o = append(o, n.Name)
		}
		return o
	}()
	ls.Height = 5
	//--------------------------------
	bigview := termui.NewTable()
	bigview.FgColor = termui.ColorWhite
	bigview.BgColor = termui.ColorDefault
	bigview.Rows = [][]string{[]string{"Namespace", "Deployments", "Type", "Replicas"}}
	bigview.Width = 100
	bigview.Height = 7
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
			row := []string{namespace.Name, deployment.Name, "Deployment", fmt.Sprintf("%d", int(*deployment.Spec.Replicas))}
			bigview.Rows = append(bigview.Rows, row)
		}

		stslist, err := conf.GetStatefulSets(namespace.Name)
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, sts := range stslist.Items {
			row := []string{namespace.Name, sts.Name, "StatefulSet", fmt.Sprintf("%d", int(*sts.Spec.Replicas))}
			bigview.Rows = append(bigview.Rows, row)
		}
	}
	//Render body -------------------
	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(6, 0, ls)),
		termui.NewRow(
			termui.NewCol(6, 0, bigview)),
	)

	termui.Body.Align()

	termui.Render(termui.Body)

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
