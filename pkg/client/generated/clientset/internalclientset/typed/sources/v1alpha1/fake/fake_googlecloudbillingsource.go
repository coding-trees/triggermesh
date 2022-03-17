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

	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGoogleCloudBillingSources implements GoogleCloudBillingSourceInterface
type FakeGoogleCloudBillingSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var googlecloudbillingsourcesResource = schema.GroupVersionResource{Group: "sources.triggermesh.io", Version: "v1alpha1", Resource: "googlecloudbillingsources"}

var googlecloudbillingsourcesKind = schema.GroupVersionKind{Group: "sources.triggermesh.io", Version: "v1alpha1", Kind: "GoogleCloudBillingSource"}

// Get takes name of the googleCloudBillingSource, and returns the corresponding googleCloudBillingSource object, and an error if there is any.
func (c *FakeGoogleCloudBillingSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GoogleCloudBillingSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(googlecloudbillingsourcesResource, c.ns, name), &v1alpha1.GoogleCloudBillingSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleCloudBillingSource), err
}

// List takes label and field selectors, and returns the list of GoogleCloudBillingSources that match those selectors.
func (c *FakeGoogleCloudBillingSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GoogleCloudBillingSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(googlecloudbillingsourcesResource, googlecloudbillingsourcesKind, c.ns, opts), &v1alpha1.GoogleCloudBillingSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GoogleCloudBillingSourceList{ListMeta: obj.(*v1alpha1.GoogleCloudBillingSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.GoogleCloudBillingSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleCloudBillingSources.
func (c *FakeGoogleCloudBillingSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(googlecloudbillingsourcesResource, c.ns, opts))

}

// Create takes the representation of a googleCloudBillingSource and creates it.  Returns the server's representation of the googleCloudBillingSource, and an error, if there is any.
func (c *FakeGoogleCloudBillingSources) Create(ctx context.Context, googleCloudBillingSource *v1alpha1.GoogleCloudBillingSource, opts v1.CreateOptions) (result *v1alpha1.GoogleCloudBillingSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(googlecloudbillingsourcesResource, c.ns, googleCloudBillingSource), &v1alpha1.GoogleCloudBillingSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleCloudBillingSource), err
}

// Update takes the representation of a googleCloudBillingSource and updates it. Returns the server's representation of the googleCloudBillingSource, and an error, if there is any.
func (c *FakeGoogleCloudBillingSources) Update(ctx context.Context, googleCloudBillingSource *v1alpha1.GoogleCloudBillingSource, opts v1.UpdateOptions) (result *v1alpha1.GoogleCloudBillingSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(googlecloudbillingsourcesResource, c.ns, googleCloudBillingSource), &v1alpha1.GoogleCloudBillingSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleCloudBillingSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGoogleCloudBillingSources) UpdateStatus(ctx context.Context, googleCloudBillingSource *v1alpha1.GoogleCloudBillingSource, opts v1.UpdateOptions) (*v1alpha1.GoogleCloudBillingSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(googlecloudbillingsourcesResource, "status", c.ns, googleCloudBillingSource), &v1alpha1.GoogleCloudBillingSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleCloudBillingSource), err
}

// Delete takes name of the googleCloudBillingSource and deletes it. Returns an error if one occurs.
func (c *FakeGoogleCloudBillingSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(googlecloudbillingsourcesResource, c.ns, name), &v1alpha1.GoogleCloudBillingSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleCloudBillingSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(googlecloudbillingsourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GoogleCloudBillingSourceList{})
	return err
}

// Patch applies the patch and returns the patched googleCloudBillingSource.
func (c *FakeGoogleCloudBillingSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GoogleCloudBillingSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(googlecloudbillingsourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GoogleCloudBillingSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GoogleCloudBillingSource), err
}
