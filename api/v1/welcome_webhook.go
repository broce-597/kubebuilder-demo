/*
Copyright 2022.

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

package v1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"regexp"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var welcomelog = logf.Log.WithName("welcome-resource")

func (r *Welcome) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
//如果需要操作资源请在此添加rbac权限
//+kubebuilder:webhook:path=/mutate-webapp-demo-welcome-domain-v1-welcome,mutating=true,failurePolicy=fail,sideEffects=None,groups=webapp.demo.welcome.domain,resources=welcomes,verbs=create;update,versions=v1,name=mwelcome.kb.io,admissionReviewVersions={v1,v1beta1}

//数据结构默认值实例化
var _ webhook.Defaulter = &Welcome{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
//默认值配置
func (r *Welcome) Default() {
	welcomelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	// TODO(user): fill in your defaulting logic.
	// 如果创建的时候没有输入name，设置默认值
	if r.Spec.Name == "" {
		r.Spec.Name = "broce-597"
		welcomelog.Info("a. name is nil, set default value now", "name", r.Spec.Name)
	} else {
		welcomelog.Info("b. TotalQPS exists", "TotalQPS", r.Spec.Name)
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//校验 如果需要操作资源请在此添加rbac权限
//+kubebuilder:webhook:path=/validate-webapp-demo-welcome-domain-v1-welcome,mutating=false,failurePolicy=fail,sideEffects=None,groups=webapp.demo.welcome.domain,resources=welcomes,verbs=create;update,versions=v1,name=vwelcome.kb.io,admissionReviewVersions={v1,v1beta1}

//实例化
var _ webhook.Validator = &Welcome{}

//创建 删除 修改 校验
//这里在创建和修改的时候通过正则表达式校验  name字段值能包含字母
// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Welcome) ValidateCreate() error {
	welcomelog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return r.validateWelcome()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Welcome) ValidateUpdate(old runtime.Object) error {
	welcomelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.

	return r.validateWelcome()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Welcome) ValidateDelete() error {
	welcomelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *Welcome) validateWelcome() error {
	var allErrs field.ErrorList
	match, _ := regexp.MatchString(`^[A-Za-z]+$`, r.Spec.Name)
	if match {
		welcomelog.Info("e. name is valid")
		return nil
	}
	welcomelog.Info("c. Invalid name")
	err := field.Invalid(field.NewPath("spec").Child("name"),
		r.Spec.Name,
		"d. Name must be set to letters")
	allErrs = append(allErrs, err)

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "webapp.demo.welcome.domain", Kind: "WelCome"},
		r.Name,
		allErrs)
}
