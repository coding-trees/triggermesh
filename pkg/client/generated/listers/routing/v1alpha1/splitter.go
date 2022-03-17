/*
Copyright 2022 TriggerMesh Inc.

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
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/routing/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SplitterLister helps list Splitters.
// All objects returned here must be treated as read-only.
type SplitterLister interface {
	// List lists all Splitters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Splitter, err error)
	// Splitters returns an object that can list and get Splitters.
	Splitters(namespace string) SplitterNamespaceLister
	SplitterListerExpansion
}

// splitterLister implements the SplitterLister interface.
type splitterLister struct {
	indexer cache.Indexer
}

// NewSplitterLister returns a new SplitterLister.
func NewSplitterLister(indexer cache.Indexer) SplitterLister {
	return &splitterLister{indexer: indexer}
}

// List lists all Splitters in the indexer.
func (s *splitterLister) List(selector labels.Selector) (ret []*v1alpha1.Splitter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Splitter))
	})
	return ret, err
}

// Splitters returns an object that can list and get Splitters.
func (s *splitterLister) Splitters(namespace string) SplitterNamespaceLister {
	return splitterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SplitterNamespaceLister helps list and get Splitters.
// All objects returned here must be treated as read-only.
type SplitterNamespaceLister interface {
	// List lists all Splitters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Splitter, err error)
	// Get retrieves the Splitter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Splitter, error)
	SplitterNamespaceListerExpansion
}

// splitterNamespaceLister implements the SplitterNamespaceLister
// interface.
type splitterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Splitters in the indexer for a given namespace.
func (s splitterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Splitter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Splitter))
	})
	return ret, err
}

// Get retrieves the Splitter from the indexer for a given namespace and name.
func (s splitterNamespaceLister) Get(name string) (*v1alpha1.Splitter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("splitter"), name)
	}
	return obj.(*v1alpha1.Splitter), nil
}
