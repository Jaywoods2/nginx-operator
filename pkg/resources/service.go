package resources

import (
	nginxv1 "github.com/Jaywoods/nginx-operator/pkg/apis/nginx/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewSevice(n *nginxv1.NginxService) *corev1.Service {
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
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
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceTypeClusterIP,
			Ports:    n.Spec.Ports,
			Selector: map[string]string{"app": n.Name},
		},
	}
}
