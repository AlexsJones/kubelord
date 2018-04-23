package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AlexsJones/kubelord/kubernetes"
	"github.com/fatih/color"
)

func sig() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exitChan := make(chan int)
	go func() {
		for {
			s := <-signalChan
			switch s {
			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				fmt.Println("force stop")
				exitChan <- 0

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				fmt.Println("stop and core dump")
				exitChan <- 0

			default:
				fmt.Println("Unknown signal.")
				exitChan <- 1
			}
		}
	}()

	code := <-exitChan
	os.Exit(code)
}
func main() {
	go sig()
	poll := flag.String("poll", "3s", "Set a poll time per refresh sweep using time.Duration format e.g. 1s")
	flag.Parse()
	//Load Kubernetes configuration ---------------------------------------------
	color.Yellow("Loading kubernetes configuration")
	k8sConf, err := kubernetes.NewConfiguration("", false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	color.Green("OK")
	//Load Curator --------------------------------------------------------------
	color.Yellow("Generating Curator")
	curator := kubernetes.NewCurator(k8sConf, kubernetes.NewDefaultFilters())
	color.Green("OK")
	//Load UI -------------------------------------------------------------------
	//u := ux.NewConfiguration()

	d, err := time.ParseDuration(*poll)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Update window...
	go func() {
		for {
			log.Println("Updating...")
			if err := curator.Do(); err != nil {
				log.Println(err.Error())
			}
			for _, namespace := range curator.Namespaces {
				log.Println(namespace)
				//svc
				for _, svc := range curator.Services[namespace[0]] {
					log.Printf("\t %s\n", svc)
				}
				//deployment
				for _, de := range curator.Deployments[namespace[0]] {
					log.Printf("\t %s\n", de)
				}
				//statefulset
				for _, de := range curator.StatefulSets[namespace[0]] {
					log.Printf("\t %s\n", de)
				}

			}

			time.Sleep(d)
		}
	}()
	//	defer u.Exit()

	time.Sleep(time.Second * 100)
}
