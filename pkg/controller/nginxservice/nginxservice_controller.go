package nginxservice

import (
	"context"
	"encoding/json"
	"reflect"

	nginxv1 "github.com/Jaywoods/nginx-operator/pkg/apis/nginx/v1"
	"github.com/Jaywoods/nginx-operator/pkg/resources"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_nginxservice")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new NginxService Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileNginxService{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("nginxservice-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource NginxService
	err = c.Watch(&source.Kind{Type: &nginxv1.NginxService{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner NginxService
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &nginxv1.NginxService{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileNginxService implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileNginxService{}

// ReconcileNginxService reconciles a NginxService object
type ReconcileNginxService struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a NginxService object and makes changes based on the state read
// and what is in the NginxService.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileNginxService) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling NginxService")

	// Fetch the NginxService instance
	instance := &nginxv1.NginxService{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if instance.DeletionTimestamp != nil {
		return reconcile.Result{}, err
	}

	// 如果不存在，则创建关联资源
	// 如果存在，判断是否需要更新
	//   如果需要更新，则直接更新
	//   如果不需要更新，则正常返回

	deploy := &appsv1.Deployment{}
	if err := r.client.Get(context.TODO(), request.NamespacedName, deploy); err != nil && errors.IsNotFound(err) {
		// 创建关联资源
		// 1. 创建deploy
		deploy := resources.NewDeploy(instance)
		if err := r.client.Create(context.TODO(), deploy); err != nil {
			return reconcile.Result{}, err
		}

		// 创建service
		svc := resources.NewSevice(instance)
		if err := r.client.Create(context.TODO(), svc); err != nil {
			return reconcile.Result{}, err
		}

		// 关联Annotations 保留crd 上次配置，用次字段值与当前配置对比，来判断是否需要更新
		data, _ := json.Marshal(instance.Spec)
		if instance.Annotations != nil {
			instance.Annotations["kubectl.kubernetes.io/last-applied-configuration"] = string(data)
		} else {
			instance.Annotations = map[string]string{"kubectl.kubernetes.io/last-applied-configuration": string(data)}
		}

		if err := r.client.Update(context.TODO(), instance); err != nil {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, nil
	}

	oldInstanceSpec := nginxv1.NginxServiceSpec{}
	if err := json.Unmarshal([]byte(instance.Annotations["kubectl.kubernetes.io/last-applied-configuration"]), oldInstanceSpec); err != nil {
		return reconcile.Result{}, nil
	}
	// 与上次配置对比，如果不一致则更新
	if !reflect.DeepEqual(instance.Spec, oldInstanceSpec) {
		// 更新deploy资源
		newDeploy := resources.NewDeploy(instance)
		oldDeploy := &appsv1.Deployment{}
		if err := r.client.Get(context.TODO(), request.NamespacedName, oldDeploy); err != nil {
			return reconcile.Result{}, err
		}
		oldDeploy.Spec = newDeploy.Spec
		if err := r.client.Update(context.TODO(), oldDeploy); err != nil {
			return reconcile.Result{}, err
		}
		// 更新svc资源
		newSvc := resources.NewSevice(instance)

		oldSvc := &corev1.Service{}
		if err := r.client.Get(context.TODO(), request.NamespacedName, oldSvc); err != nil {
			return reconcile.Result{}, err
		}
		oldSvc.Spec = newSvc.Spec
		if err := r.client.Update(context.TODO(), newSvc); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, err

	}
	return reconcile.Result{}, nil
}
