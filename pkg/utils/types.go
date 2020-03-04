package utils

import (
	"flag"
	"os"

	chaosClient "github.com/litmuschaos/chaos-operator/pkg/client/clientset/versioned/typed/litmuschaos/v1alpha1"
	"k8s.io/api/node/v1alpha1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// ExperimentDetails is for collecting all the experiment-related details
type ExperimentDetails struct {
	ExperimentName      string
	ChaosDuration       int
	ChaosInterval       int
	RampTime            int
	Force               bool
	ChaosLib            string
	ChaosServiceAccount string
	AppNS               string
	AppLabel            string
	AppKind             string
	KillCount           int
	ChaosUID            string
	AuxiliaryAppInfo    string
}

// ClientSets is a collection of clientSets needed
type ClientSets struct {
	KubeClient   *kubernetes.Clientset
	LitmusClient *chaosClient.LitmuschaosV1alpha1Client
}

// GenerateClientSetFromKubeConfig will generation both ClientSets (k8s, and Litmus)
func (clientSets *ClientSets) GenerateClientSetFromKubeConfig() error {

	var err error
	kubeconfig := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		return err
	}

	client, err := kubernetes.NewForConfig(config)

	if err != nil {
		return err
	}

	clientSet, err := chaosClient.NewForConfig(config)

	if err != nil {
		return err
	}

	err = v1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	// config, err := getKubeConfig()
	// if err != nil {
	// 	return err
	// }
	// k8sClientSet, err := k8s.GenerateK8sClientSet(config)
	// if err != nil {
	// 	return err
	// }
	// litmusClientSet, err := litmus.GenerateLitmusClientSet(config)
	// if err != nil {
	// 	return err
	// }
	clientSets.KubeClient = client
	clientSets.LitmusClient = clientSet

	return nil
}

// getKubeConfig setup the config for access cluster resource
func getKubeConfig() (*rest.Config, error) {
	kubeconfig := flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()
	// Use in-cluster config if kubeconfig path is specified
	if *kubeconfig == "" {
		config, err := rest.InClusterConfig()
		if err != nil {
			return config, err
		}
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return config, err
	}
	return config, err
}
