package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/openebs/maya/pkg/util/retry"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CheckApplicationStatus ...
func CheckApplicationStatus(appNs string, appLabel string, clients ClientSets) error {
	// Checking whether application pods are in running state

	fmt.Println("Checking the running status of application pod\n")
	err := CheckPodStatus(appNs, appLabel, clients)
	if err != nil {
		return err
	}
	fmt.Println("Checking the running status of application container\n")
	err = CheckContainerStatus(appNs, appLabel, clients)
	if err != nil {
		return err
	}
	return nil
}

// CheckAuxiliaryApplicationStatus ...
func CheckAuxiliaryApplicationStatus(experimentsDetails *ExperimentDetails, clients ClientSets) error {

	AuxiliaryAppInfo := strings.Split(experimentsDetails.AuxiliaryAppInfo, ",")

	for _, val := range AuxiliaryAppInfo {
		AppInfo := strings.Split(val, ":")
		err := CheckApplicationStatus(AppInfo[0], AppInfo[1], clients)
		if err != nil {
			return err
		}

	}
	return nil
}

// CheckPodStatus ...
func CheckPodStatus(appNs string, appLabel string, clients ClientSets) error {
	err := retry.
		Times(90).
		Wait(2 * time.Second).
		Try(func(attempt uint) error {
			podSpec, err := clients.KubeClient.CoreV1().Pods(appNs).List(metav1.ListOptions{LabelSelector: appLabel})
			if err != nil {
				return err
			}
			err = nil
			for _, pod := range podSpec.Items {
				if string(pod.Status.Phase) != "Running" {
					return errors.Errorf("Pod is not yet in running state")
				}
				fmt.Printf(" %v Pod is in %v State \n", pod.Name, pod.Status.Phase)
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

// CheckContainerStatus ...
func CheckContainerStatus(appNs string, appLabel string, clients ClientSets) error {
	err := retry.
		Times(90).
		Wait(2 * time.Second).
		Try(func(attempt uint) error {
			podSpec, err := clients.KubeClient.CoreV1().Pods(appNs).List(metav1.ListOptions{LabelSelector: appLabel})
			if err != nil {
				return err
			}
			err = nil
			for _, pod := range podSpec.Items {

				for _, container := range pod.Status.ContainerStatuses {
					if container.Ready != true {
						return errors.Errorf("containers are not yet in running state")
					}
					fmt.Printf(" %v container of pod %v is in %v State \n", container.Name, pod.Name, pod.Status.Phase)
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}
