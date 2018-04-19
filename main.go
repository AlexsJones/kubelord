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

	color.Yellow("Loading kubernetes configuration")
	_, err := kubernetes.NewConfiguration("", false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	color.Green("OK")

	d, err := time.ParseDuration(*poll)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		log.Println("Updating...")
		time.Sleep(d)
	}

}
