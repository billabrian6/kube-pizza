/*


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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	kubepizzav1alpha1 "github.com/billabrian6/kube-pizza/api/v1alpha1"
)

// PizzaReconciler reconciles a Pizza object
type PizzaReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=kube-pizza.billabrian6.com,resources=pizzas,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kube-pizza.billabrian6.com,resources=pizzas/status,verbs=get;update;patch

func (r *PizzaReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("pizza", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *PizzaReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubepizzav1alpha1.Pizza{}).
		Complete(r)
}
