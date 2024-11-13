package k8s

import (
    "context"
    "fmt"
    "os"

    "github.com/DB-Vincent/k8s-scaling-demo/backend/internal/models"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/labels"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/clientcmd"
)

func SetupConfig() *rest.Config {
    var config *rest.Config
    var err error

    if _, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount/token"); err == nil {
        config, err = rest.InClusterConfig()
    } else {
        config, err = clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
    }

    if err != nil {
        panic(err.Error())
    }

    return config
}

func GetRelatedPods(config *rest.Config) ([]models.Replica, error) {
    podNamespace := os.Getenv("POD_NAMESPACE")
    podName := os.Getenv("POD_NAME")

    if podNamespace == "" || podName == "" {
        return []models.Replica{}, fmt.Errorf("required environment variables POD_NAMESPACE and/or POD_NAME are not set")
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return []models.Replica{}, err
    }

    pod, err := clientset.CoreV1().Pods(podNamespace).Get(context.TODO(), podName, metav1.GetOptions{})
    if err != nil {
        return []models.Replica{}, err
    }

    var pods []models.Replica

    for _, ownerRef := range pod.OwnerReferences {
        if *ownerRef.Controller {
            labelSelector := labels.Set(pod.Labels).AsSelector().String()
            podList, err := clientset.CoreV1().Pods(podNamespace).List(context.TODO(), metav1.ListOptions{
                LabelSelector: labelSelector,
            })
            if err != nil {
                return []models.Replica{}, err
            }

            for _, p := range podList.Items {
                replica := models.Replica{
                    Name:      p.Name,
                    NodeName:  p.Spec.NodeName,
                    Status:    string(p.Status.Phase),
                    StartTime: p.Status.StartTime.String(),
                    Current:   p.Name == podName,
                }
                pods = append(pods, replica)
            }
        }
    }

    return pods, nil
}
