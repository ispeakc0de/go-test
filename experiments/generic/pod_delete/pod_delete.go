package main

import (
	"fmt"

	litmus "github.com/litmuschaos/go-test/chaoslib/litmus"
	utils "github.com/litmuschaos/go-test/pkg/utils"
)

func main() {
	var err error
	experimentsDetails := utils.ExperimentDetails{}
	clients := utils.ClientSets{}

	// Getting kubeConfig and Generate ClientSets
	if err := clients.GenerateClientSetFromKubeConfig(); err != nil {
		fmt.Println(err)
		return
	}

	// Fetching all ENV
	utils.GetENV(&experimentsDetails, "pod-delete")

	// //Update chaosResult
	// fmt.Println("Update chaosResult in begg of experiment\n")
	// err = utils.UpdateChaosResult("SOT", experimentsDetails, clients)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//PRE-CHAOS APPLICATION LIVENESS CHECK
	fmt.Println("Verify that the AUT (Application Under Test) is running\n")
	err = utils.CheckApplicationStatus(experimentsDetails.AppNS, experimentsDetails.AppLabel, clients)
	if err != nil {
		fmt.Println(err)
		return
	}

	//PRE-CHAOS AUXILIARY APPLICATION LIVENESS CHECK
	fmt.Println("Verify that the auxiliary application are is running\n")
	err = utils.CheckAuxiliaryApplicationStatus(&experimentsDetails, clients)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Including ChaosUtil
	if experimentsDetails.ChaosLib == "litmus" {
		err = litmus.PodDeleteChaos(&experimentsDetails, clients)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("pod deletion successful!!\n")
		}
	} else {
		fmt.Println("Enter valid Lib!!")
	}

	//POST-CHAOS APPLICATION LIVENESS CHECK
	fmt.Println("Verify that the AUT (Application Under Test) is running\n")
	err = utils.CheckApplicationStatus(experimentsDetails.AppNS, experimentsDetails.AppLabel, clients)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Post-CHAOS AUXILIARY APPLICATION LIVENESS CHECK
	fmt.Println("Verify that the auxiliary application are is running\n")
	err = utils.CheckAuxiliaryApplicationStatus(&experimentsDetails, clients)
	if err != nil {
		fmt.Println(err)
		return
	}

}
