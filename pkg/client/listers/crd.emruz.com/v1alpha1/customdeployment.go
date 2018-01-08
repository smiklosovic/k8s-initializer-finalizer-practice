// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "crd-controller/pkg/apis/crd.emruz.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CustomDeploymentLister helps list CustomDeployments.
type CustomDeploymentLister interface {
	// List lists all CustomDeployments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CustomDeployment, err error)
	// CustomDeployments returns an object that can list and get CustomDeployments.
	CustomDeployments(namespace string) CustomDeploymentNamespaceLister
	CustomDeploymentListerExpansion
}

// customDeploymentLister implements the CustomDeploymentLister interface.
type customDeploymentLister struct {
	indexer cache.Indexer
}

// NewCustomDeploymentLister returns a new CustomDeploymentLister.
func NewCustomDeploymentLister(indexer cache.Indexer) CustomDeploymentLister {
	return &customDeploymentLister{indexer: indexer}
}

// List lists all CustomDeployments in the indexer.
func (s *customDeploymentLister) List(selector labels.Selector) (ret []*v1alpha1.CustomDeployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CustomDeployment))
	})
	return ret, err
}

// CustomDeployments returns an object that can list and get CustomDeployments.
func (s *customDeploymentLister) CustomDeployments(namespace string) CustomDeploymentNamespaceLister {
	return customDeploymentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CustomDeploymentNamespaceLister helps list and get CustomDeployments.
type CustomDeploymentNamespaceLister interface {
	// List lists all CustomDeployments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CustomDeployment, err error)
	// Get retrieves the CustomDeployment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CustomDeployment, error)
	CustomDeploymentNamespaceListerExpansion
}

// customDeploymentNamespaceLister implements the CustomDeploymentNamespaceLister
// interface.
type customDeploymentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CustomDeployments in the indexer for a given namespace.
func (s customDeploymentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CustomDeployment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CustomDeployment))
	})
	return ret, err
}

// Get retrieves the CustomDeployment from the indexer for a given namespace and name.
func (s customDeploymentNamespaceLister) Get(name string) (*v1alpha1.CustomDeployment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("customdeployment"), name)
	}
	return obj.(*v1alpha1.CustomDeployment), nil
}