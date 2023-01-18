package controller

import (
	"github.com/galal-hussein/k3k/pkg/apis/k3k.io/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func service(cluster *v1alpha1.Cluster) v1.Service {
	return v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "k3k-server-service",
			Namespace: clusterNamespace(cluster),
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeClusterIP,
			Selector: map[string]string{
				"cluster": cluster.Name,
			},
			Ports: []v1.ServicePort{
				{
					Name:     "k3s-server-port",
					Protocol: v1.ProtocolTCP,
					Port:     6443,
				},
			},
		},
	}
}
