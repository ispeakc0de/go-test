package litmus

import (
	utils "github.com/litmuschaos/go-test/pkg/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodDeleteChaos ...
func PodDeleteChaos(experimentsDetails *utils.ExperimentDetails, clients utils.ClientSets) error {

	// Get pods
	pods, err := clients.KubeClient.CoreV1().Pods(experimentsDetails.AppNS).List(metav1.ListOptions{LabelSelector: experimentsDetails.AppLabel})
	if err != nil {
		return err
	}
	podToDelete := pods.Items[0].Name

	err = clients.KubeClient.CoreV1().Pods(experimentsDetails.AppNS).Delete(podToDelete, &metav1.DeleteOptions{})

	if err != nil {
		return err
	}
	return nil

}
