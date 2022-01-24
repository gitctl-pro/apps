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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	appsv1 "github.com/gitctl-pro/apps/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCanaries implements CanaryInterface
type FakeCanaries struct {
	Fake *FakeAppsV1
	ns   string
}

var canariesResource = schema.GroupVersionResource{Group: "apps.gitclt.com", Version: "v1", Resource: "canaries"}

var canariesKind = schema.GroupVersionKind{Group: "apps.gitclt.com", Version: "v1", Kind: "Canary"}

// Get takes name of the canary, and returns the corresponding canary object, and an error if there is any.
func (c *FakeCanaries) Get(ctx context.Context, name string, options v1.GetOptions) (result *appsv1.Canary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(canariesResource, c.ns, name), &appsv1.Canary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Canary), err
}

// List takes label and field selectors, and returns the list of Canaries that match those selectors.
func (c *FakeCanaries) List(ctx context.Context, opts v1.ListOptions) (result *appsv1.CanaryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(canariesResource, canariesKind, c.ns, opts), &appsv1.CanaryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1.CanaryList{ListMeta: obj.(*appsv1.CanaryList).ListMeta}
	for _, item := range obj.(*appsv1.CanaryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested canaries.
func (c *FakeCanaries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(canariesResource, c.ns, opts))

}

// Create takes the representation of a canary and creates it.  Returns the server's representation of the canary, and an error, if there is any.
func (c *FakeCanaries) Create(ctx context.Context, canary *appsv1.Canary, opts v1.CreateOptions) (result *appsv1.Canary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(canariesResource, c.ns, canary), &appsv1.Canary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Canary), err
}

// Update takes the representation of a canary and updates it. Returns the server's representation of the canary, and an error, if there is any.
func (c *FakeCanaries) Update(ctx context.Context, canary *appsv1.Canary, opts v1.UpdateOptions) (result *appsv1.Canary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(canariesResource, c.ns, canary), &appsv1.Canary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Canary), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCanaries) UpdateStatus(ctx context.Context, canary *appsv1.Canary, opts v1.UpdateOptions) (*appsv1.Canary, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(canariesResource, "status", c.ns, canary), &appsv1.Canary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Canary), err
}

// Delete takes name of the canary and deletes it. Returns an error if one occurs.
func (c *FakeCanaries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(canariesResource, c.ns, name), &appsv1.Canary{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCanaries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(canariesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &appsv1.CanaryList{})
	return err
}

// Patch applies the patch and returns the patched canary.
func (c *FakeCanaries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *appsv1.Canary, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(canariesResource, c.ns, name, pt, data, subresources...), &appsv1.Canary{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Canary), err
}
