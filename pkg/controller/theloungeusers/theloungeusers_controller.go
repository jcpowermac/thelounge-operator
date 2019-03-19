package theloungeusers

import (
	"context"

	thev1alpha1 "github.com/jcpowermac/thelounge-operator/pkg/apis/the/v1alpha1"

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

var log = logf.Log.WithName("controller_theloungeusers")

// Add creates a new TheloungeUsers Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileTheloungeUsers{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("theloungeusers-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource TheloungeUsers
	err = c.Watch(&source.Kind{Type: &thev1alpha1.TheloungeUsers{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner TheloungeUsers
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &thev1alpha1.TheloungeUsers{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileTheloungeUsers{}

// ReconcileTheloungeUsers reconciles a TheloungeUsers object
type ReconcileTheloungeUsers struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a TheloungeUsers object and makes changes based on the state read
// and what is in the TheloungeUsers.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileTheloungeUsers) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling TheloungeUsers")

	// Fetch the TheloungeUsers instance

	//Users instance
	instance := &thev1alpha1.TheloungeUsers{}

	//TheLoungeInstance
	theLoungeInstance := &thev1alpha1.Thelounge{}

	err := r.client.Get(context.TODO(), request.NamespacedName, theLoungeInstance)

	// Need to wait for the thelounge to start

	//err := r.client.Get(context.TODO(), request.NamespacedName, instance)
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

	// Check the user instances
	err = r.client.Get(context.TODO(), request.NamespacedName, instance)

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

	firstName := instance.Spec.FirstName
	lastName := instance.Spec.LastName
	password := instance.Spec.Password

	reqLogger.Info("Creating new User", "instance.Spec.FirstName", firstName, "instance.Spec.LastName", lastName, "instance.Spec.Password", password)

	// Define a new Pod object
	//pod := newPodForCR(instance)

	// Set TheloungeUsers instance as the owner and controller
	//if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
	//	return reconcile.Result{}, err
	//}

	// Check if this Pod already exists
	//found := &corev1.Pod{}

	//err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)

	// check if user exists

	if err != nil && errors.IsNotFound(err) {
		//reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		// Really we are looking for the pod because we want to exec into it

		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	//reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}
