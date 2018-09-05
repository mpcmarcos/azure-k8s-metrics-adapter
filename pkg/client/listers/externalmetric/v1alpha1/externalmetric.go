/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/Azure/azure-k8s-metrics-adapter/pkg/apis/externalmetric/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ExternalMetricLister helps list ExternalMetrics.
type ExternalMetricLister interface {
	// List lists all ExternalMetrics in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ExternalMetric, err error)
	// ExternalMetrics returns an object that can list and get ExternalMetrics.
	ExternalMetrics(namespace string) ExternalMetricNamespaceLister
	ExternalMetricListerExpansion
}

// externalMetricLister implements the ExternalMetricLister interface.
type externalMetricLister struct {
	indexer cache.Indexer
}

// NewExternalMetricLister returns a new ExternalMetricLister.
func NewExternalMetricLister(indexer cache.Indexer) ExternalMetricLister {
	return &externalMetricLister{indexer: indexer}
}

// List lists all ExternalMetrics in the indexer.
func (s *externalMetricLister) List(selector labels.Selector) (ret []*v1alpha1.ExternalMetric, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExternalMetric))
	})
	return ret, err
}

// ExternalMetrics returns an object that can list and get ExternalMetrics.
func (s *externalMetricLister) ExternalMetrics(namespace string) ExternalMetricNamespaceLister {
	return externalMetricNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ExternalMetricNamespaceLister helps list and get ExternalMetrics.
type ExternalMetricNamespaceLister interface {
	// List lists all ExternalMetrics in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ExternalMetric, err error)
	// Get retrieves the ExternalMetric from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ExternalMetric, error)
	ExternalMetricNamespaceListerExpansion
}

// externalMetricNamespaceLister implements the ExternalMetricNamespaceLister
// interface.
type externalMetricNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ExternalMetrics in the indexer for a given namespace.
func (s externalMetricNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ExternalMetric, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExternalMetric))
	})
	return ret, err
}

// Get retrieves the ExternalMetric from the indexer for a given namespace and name.
func (s externalMetricNamespaceLister) Get(name string) (*v1alpha1.ExternalMetric, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("externalmetric"), name)
	}
	return obj.(*v1alpha1.ExternalMetric), nil
}
