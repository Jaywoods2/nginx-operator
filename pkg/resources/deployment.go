package resources

import (
	nginxv1 "github.com/Jaywoods/nginx-operator/pkg/apis/nginx/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewDeploy(n *nginxv1.NginxService) *appsv1.Deployment {
	labels := map[string]string{"app": n.Name}
	selector := &metav1.LabelSelector{MatchLabels: labels}
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      n.Name,
			Namespace: n.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(n, schema.GroupVersionKind{
					Group:   nginxv1.SchemeGroupVersion.Group,
					Version: nginxv1.SchemeGroupVersion.Version,
					Kind:    "NginxService",
				}),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: n.Spec.Size,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: newContainers(n),
				},
			},
			Selector: selector,
		},
		Status: appsv1.DeploymentStatus{},
	}
}

func newContainers(n *nginxv1.NginxService) []corev1.Container {
	containerPorts := []corev1.ContainerPort{}
	for _, svcPort := range n.Spec.Ports {
		cport := corev1.ContainerPort{}
		cport.ContainerPort = svcPort.TargetPort.IntVal
		containerPorts = append(containerPorts, cport)
	}

	return []corev1.Container{
		{
			Name:            n.Name,
			Image:           n.Spec.Image,
			ImagePullPolicy: corev1.PullIfNotPresent,
			Resources:       n.Spec.Resources,
			Ports:           containerPorts,
			Env:             n.Spec.Envs,
		},
	}
}
