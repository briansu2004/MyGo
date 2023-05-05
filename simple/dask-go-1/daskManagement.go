package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)


func createNamespace(namespace string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to get user's home directory:", err)
		os.Exit(1)
	}

	kubeconfig := filepath.Join(homeDir, ".kube", "config")

	// Use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Check if the namespace already exists
	_, err = clientset.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
	if err == nil {
		fmt.Printf("Namespace %s already exists\n", namespace)
		return
	}

	// Create a new namespace object
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}

	// Create the namespace
	newNs, err := clientset.CoreV1().Namespaces().Create(context.Background(), ns, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Print the name of the new namespace
	fmt.Printf("Created namespace %s\n", newNs.Name)
}

func createDaskCluster(namespace string) {
	createNamespace(namespace)
	fmt.Println("Namespace ", namespace, " has been created successfully.")

	// Get the Kubernetes configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to get user's home directory:", err)
			os.Exit(1)
		}

		kubeconfig := filepath.Join(homeDir, ".kube", "config")
		// fmt.Println("Kubeconfig path:", kubeconfig)

		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err)
		}
	}

	// // Create a Kubernetes client
	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(clientset)

	// Create a Dynamic Client
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Load the YAML manifest into an Unstructured object
	manifestBytes, err := ioutil.ReadFile("new-cluster-1-ns-default.yaml")
	if err != nil {
		panic(err)
	}

	manifest := &unstructured.Unstructured{}
	if err := yaml.Unmarshal(manifestBytes, manifest); err != nil {
		panic(err)
	}

	// Set the namespace on the manifest
	manifest.SetNamespace(namespace)

	// Set the GVK (GroupVersionKind) for the manifest
	gvk := schema.GroupVersionKind{
		Group:   "kubernetes.dask.org",
		Version: "v1",
		Kind:    "DaskCluster",
	}
	manifest.SetGroupVersionKind(gvk)

	// Create or update the resource
	resource := "daskclusters"
	result, err := dynamicClient.Resource(schema.GroupVersionResource{
		Group:    gvk.Group,
		Version:  gvk.Version,
		Resource: resource,
	}).Namespace(namespace).Create(context.Background(), manifest, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created/Updated resource %q in namespace %q:\n%s\n", result.GetName(), namespace, result)
}
