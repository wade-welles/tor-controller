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

package v1alpha1

import (
	v1alpha1 "github.com/kragniz/kube-onions/pkg/apis/onion/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OnionServiceLister helps list OnionServices.
type OnionServiceLister interface {
	// List lists all OnionServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.OnionService, err error)
	// OnionServices returns an object that can list and get OnionServices.
	OnionServices(namespace string) OnionServiceNamespaceLister
	OnionServiceListerExpansion
}

// onionServiceLister implements the OnionServiceLister interface.
type onionServiceLister struct {
	indexer cache.Indexer
}

// NewOnionServiceLister returns a new OnionServiceLister.
func NewOnionServiceLister(indexer cache.Indexer) OnionServiceLister {
	return &onionServiceLister{indexer: indexer}
}

// List lists all OnionServices in the indexer.
func (s *onionServiceLister) List(selector labels.Selector) (ret []*v1alpha1.OnionService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OnionService))
	})
	return ret, err
}

// OnionServices returns an object that can list and get OnionServices.
func (s *onionServiceLister) OnionServices(namespace string) OnionServiceNamespaceLister {
	return onionServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OnionServiceNamespaceLister helps list and get OnionServices.
type OnionServiceNamespaceLister interface {
	// List lists all OnionServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.OnionService, err error)
	// Get retrieves the OnionService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.OnionService, error)
	OnionServiceNamespaceListerExpansion
}

// onionServiceNamespaceLister implements the OnionServiceNamespaceLister
// interface.
type onionServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OnionServices in the indexer for a given namespace.
func (s onionServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OnionService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OnionService))
	})
	return ret, err
}

// Get retrieves the OnionService from the indexer for a given namespace and name.
func (s onionServiceNamespaceLister) Get(name string) (*v1alpha1.OnionService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("onionservice"), name)
	}
	return obj.(*v1alpha1.OnionService), nil
}
