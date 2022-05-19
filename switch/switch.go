package main

import "fmt"

func main() {
	switch "Kubernetes" {
	case "kubernetes":
		fmt.Println("Hit the kubernetes lowercase")

	case "Kubernetes":
		fmt.Println("Hit the Kubernetes uppercase")

	case "k8s":
		fmt.Println("Hit the k8s lowercase")

	case "Istio":
		fmt.Println("Istio service")

	case "K8s":
		fmt.Println("Hit the K8s")

	default:
		fmt.Println("Hit the default")

	}
}
