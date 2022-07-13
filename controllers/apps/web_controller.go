/*
Copyright 2022 ysicing <i@ysicing.me>.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apps

import (
	"context"

	utilclient "github.com/ysicing/appflow/pkg/util/client"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	appsv1beta1 "github.com/ysicing/appflow/apis/apps/v1beta1"
)

var (
	concurrentReconciles = 3
)

const (
	controllerName = "apps-controller"
)

func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	recorder := mgr.GetEventRecorderFor(controllerName)
	return &WebReconciler{
		Client:        utilclient.NewClientFromManager(mgr, controllerName),
		Scheme:        mgr.GetScheme(),
		EventRecorder: recorder,
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New(controllerName, mgr, controller.Options{Reconciler: r, MaxConcurrentReconciles: concurrentReconciles})
	if err != nil {
		return err
	}
	// Watch for changes to NodeImage
	err = c.Watch(&source.Kind{Type: &appsv1beta1.Web{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}
	return nil
}

var _ reconcile.Reconciler = &WebReconciler{}

// WebReconciler reconciles a Web object
type WebReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	EventRecorder record.EventRecorder
}

//+kubebuilder:rbac:groups=apps.ysicing.me,resources=webs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.ysicing.me,resources=webs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.ysicing.me,resources=webs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Web object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *WebReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WebReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1beta1.Web{}).
		Complete(r)
}
