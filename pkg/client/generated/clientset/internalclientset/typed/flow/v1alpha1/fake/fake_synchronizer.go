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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/flow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSynchronizers implements SynchronizerInterface
type FakeSynchronizers struct {
	Fake *FakeFlowV1alpha1
	ns   string
}

var synchronizersResource = schema.GroupVersionResource{Group: "flow.triggermesh.io", Version: "v1alpha1", Resource: "synchronizers"}

var synchronizersKind = schema.GroupVersionKind{Group: "flow.triggermesh.io", Version: "v1alpha1", Kind: "Synchronizer"}

// Get takes name of the synchronizer, and returns the corresponding synchronizer object, and an error if there is any.
func (c *FakeSynchronizers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Synchronizer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(synchronizersResource, c.ns, name), &v1alpha1.Synchronizer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Synchronizer), err
}

// List takes label and field selectors, and returns the list of Synchronizers that match those selectors.
func (c *FakeSynchronizers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SynchronizerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(synchronizersResource, synchronizersKind, c.ns, opts), &v1alpha1.SynchronizerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SynchronizerList{ListMeta: obj.(*v1alpha1.SynchronizerList).ListMeta}
	for _, item := range obj.(*v1alpha1.SynchronizerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested synchronizers.
func (c *FakeSynchronizers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(synchronizersResource, c.ns, opts))

}

// Create takes the representation of a synchronizer and creates it.  Returns the server's representation of the synchronizer, and an error, if there is any.
func (c *FakeSynchronizers) Create(ctx context.Context, synchronizer *v1alpha1.Synchronizer, opts v1.CreateOptions) (result *v1alpha1.Synchronizer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(synchronizersResource, c.ns, synchronizer), &v1alpha1.Synchronizer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Synchronizer), err
}

// Update takes the representation of a synchronizer and updates it. Returns the server's representation of the synchronizer, and an error, if there is any.
func (c *FakeSynchronizers) Update(ctx context.Context, synchronizer *v1alpha1.Synchronizer, opts v1.UpdateOptions) (result *v1alpha1.Synchronizer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(synchronizersResource, c.ns, synchronizer), &v1alpha1.Synchronizer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Synchronizer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSynchronizers) UpdateStatus(ctx context.Context, synchronizer *v1alpha1.Synchronizer, opts v1.UpdateOptions) (*v1alpha1.Synchronizer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(synchronizersResource, "status", c.ns, synchronizer), &v1alpha1.Synchronizer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Synchronizer), err
}

// Delete takes name of the synchronizer and deletes it. Returns an error if one occurs.
func (c *FakeSynchronizers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(synchronizersResource, c.ns, name), &v1alpha1.Synchronizer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSynchronizers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(synchronizersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SynchronizerList{})
	return err
}

// Patch applies the patch and returns the patched synchronizer.
func (c *FakeSynchronizers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Synchronizer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(synchronizersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Synchronizer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Synchronizer), err
}
