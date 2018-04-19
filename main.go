package main

import "github.com/AlexsJones/kubelord/kubernetes"

func main() {

	_, err := kubernetes.NewConfiguration("", false)
	if err != nil {
		panic(err)
	}

}
