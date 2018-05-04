package ux

import (
	"fmt"
	"log"

	"github.com/AlexsJones/kubelord/kubernetes"
)

func dataFetch(conf *kubernetes.Configuration, t int) [][]string {
	dataSet := [][]string{[]string{fmt.Sprintf("%d", t), "Namespace", "Deployments", "Type", "Replicas", "Status"}}
	namespacelist, err := conf.GetNamespaces()
	if err != nil {
		log.Println(fmt.Sprintf("namespaces: %s", err.Error()))

	}
	//Namespaces
	for _, namespace := range namespacelist.Items {
		//Deployments
		deploymentlist, err := conf.GetDeployments(namespace.Name)
		if err != nil {
			log.Println(fmt.Sprintf("deployment: %s", err.Error()))

		}
		for _, deployment := range deploymentlist.Items {
			row := []string{"", namespace.Name, deployment.Name,
				"Deployment", fmt.Sprintf("%d/%d", int(deployment.Status.AvailableReplicas), int(*deployment.Spec.Replicas)),
				deployment.Status.Conditions[len(deployment.Status.Conditions)-1].Message}
			dataSet = append(dataSet, row)
		}

		//StatefulSets
		stslist, err := conf.GetStatefulSets(namespace.Name)
		if err != nil {
			log.Println(fmt.Sprintf("statefulset: %s", err.Error()))

		}
		for _, sts := range stslist.Items {

			status := ""
			if len(sts.Status.Conditions) > 0 {
				status = sts.Status.Conditions[len(sts.Status.Conditions)-1].Message
			}
			row := []string{"", namespace.Name, sts.Name, "StatefulSet", fmt.Sprintf("%d/%d", int(sts.Status.CurrentReplicas), int(*sts.Spec.Replicas)),
				status}
			dataSet = append(dataSet, row)
		}

		//CronJobs
		cjlist, err := conf.GetCronJobs(namespace.Name)
		if err != nil {
			log.Println(fmt.Sprintf("cronjob: %s", err.Error()))

		}
		for _, cronjob := range cjlist.Items {

			status := fmt.Sprintf("Last scheduled %s", cronjob.Status.LastScheduleTime.String())
			row := []string{"", namespace.Name, cronjob.Name, "CronJob", "N/A",
				status}
			dataSet = append(dataSet, row)
		}
	}
	return dataSet
}
