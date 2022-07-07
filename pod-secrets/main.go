package main

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()
	pod, err := clientset.CoreV1().Pods("default").Create(ctx, &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "echo",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{{
				Name:    "echo",
				Image:   "ubuntu",
				Command: []string{"bash", "-c", "echo $SECRET"},
				Env: []v1.EnvVar{{
					Name: "SECRET",
					ValueFrom: &v1.EnvVarSource{
						SecretKeyRef: &v1.SecretKeySelector{
							LocalObjectReference: v1.LocalObjectReference{
								Name: "secret",
							},
							Key: "value",
						},
					},
				}},
			}},
		},
	}, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("created pod", pod.GetName())
}
