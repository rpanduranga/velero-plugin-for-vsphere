/*
Copyright the Velero contributors.

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

package v1

import (
	v1 "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/apis/backupdriver/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CloneFromSnapshotLister helps list CloneFromSnapshots.
type CloneFromSnapshotLister interface {
	// List lists all CloneFromSnapshots in the indexer.
	List(selector labels.Selector) (ret []*v1.CloneFromSnapshot, err error)
	// CloneFromSnapshots returns an object that can list and get CloneFromSnapshots.
	CloneFromSnapshots(namespace string) CloneFromSnapshotNamespaceLister
	CloneFromSnapshotListerExpansion
}

// cloneFromSnapshotLister implements the CloneFromSnapshotLister interface.
type cloneFromSnapshotLister struct {
	indexer cache.Indexer
}

// NewCloneFromSnapshotLister returns a new CloneFromSnapshotLister.
func NewCloneFromSnapshotLister(indexer cache.Indexer) CloneFromSnapshotLister {
	return &cloneFromSnapshotLister{indexer: indexer}
}

// List lists all CloneFromSnapshots in the indexer.
func (s *cloneFromSnapshotLister) List(selector labels.Selector) (ret []*v1.CloneFromSnapshot, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CloneFromSnapshot))
	})
	return ret, err
}

// CloneFromSnapshots returns an object that can list and get CloneFromSnapshots.
func (s *cloneFromSnapshotLister) CloneFromSnapshots(namespace string) CloneFromSnapshotNamespaceLister {
	return cloneFromSnapshotNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloneFromSnapshotNamespaceLister helps list and get CloneFromSnapshots.
type CloneFromSnapshotNamespaceLister interface {
	// List lists all CloneFromSnapshots in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CloneFromSnapshot, err error)
	// Get retrieves the CloneFromSnapshot from the indexer for a given namespace and name.
	Get(name string) (*v1.CloneFromSnapshot, error)
	CloneFromSnapshotNamespaceListerExpansion
}

// cloneFromSnapshotNamespaceLister implements the CloneFromSnapshotNamespaceLister
// interface.
type cloneFromSnapshotNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloneFromSnapshots in the indexer for a given namespace.
func (s cloneFromSnapshotNamespaceLister) List(selector labels.Selector) (ret []*v1.CloneFromSnapshot, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CloneFromSnapshot))
	})
	return ret, err
}

// Get retrieves the CloneFromSnapshot from the indexer for a given namespace and name.
func (s cloneFromSnapshotNamespaceLister) Get(name string) (*v1.CloneFromSnapshot, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clonefromsnapshot"), name)
	}
	return obj.(*v1.CloneFromSnapshot), nil
}