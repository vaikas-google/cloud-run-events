/*
Copyright 2019 Google LLC

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
	v1alpha1 "github.com/GoogleCloudPlatform/cloud-run-events/pkg/apis/events/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGCSSources implements GCSSourceInterface
type FakeGCSSources struct {
	Fake *FakeEventsV1alpha1
	ns   string
}

var gcssourcesResource = schema.GroupVersionResource{Group: "events.cloud.run", Version: "v1alpha1", Resource: "gcssources"}

var gcssourcesKind = schema.GroupVersionKind{Group: "events.cloud.run", Version: "v1alpha1", Kind: "GCSSource"}

// Get takes name of the gCSSource, and returns the corresponding gCSSource object, and an error if there is any.
func (c *FakeGCSSources) Get(name string, options v1.GetOptions) (result *v1alpha1.GCSSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gcssourcesResource, c.ns, name), &v1alpha1.GCSSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GCSSource), err
}

// List takes label and field selectors, and returns the list of GCSSources that match those selectors.
func (c *FakeGCSSources) List(opts v1.ListOptions) (result *v1alpha1.GCSSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gcssourcesResource, gcssourcesKind, c.ns, opts), &v1alpha1.GCSSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GCSSourceList{ListMeta: obj.(*v1alpha1.GCSSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.GCSSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gCSSources.
func (c *FakeGCSSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gcssourcesResource, c.ns, opts))

}

// Create takes the representation of a gCSSource and creates it.  Returns the server's representation of the gCSSource, and an error, if there is any.
func (c *FakeGCSSources) Create(gCSSource *v1alpha1.GCSSource) (result *v1alpha1.GCSSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gcssourcesResource, c.ns, gCSSource), &v1alpha1.GCSSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GCSSource), err
}

// Update takes the representation of a gCSSource and updates it. Returns the server's representation of the gCSSource, and an error, if there is any.
func (c *FakeGCSSources) Update(gCSSource *v1alpha1.GCSSource) (result *v1alpha1.GCSSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gcssourcesResource, c.ns, gCSSource), &v1alpha1.GCSSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GCSSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGCSSources) UpdateStatus(gCSSource *v1alpha1.GCSSource) (*v1alpha1.GCSSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gcssourcesResource, "status", c.ns, gCSSource), &v1alpha1.GCSSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GCSSource), err
}

// Delete takes name of the gCSSource and deletes it. Returns an error if one occurs.
func (c *FakeGCSSources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gcssourcesResource, c.ns, name), &v1alpha1.GCSSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGCSSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gcssourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GCSSourceList{})
	return err
}

// Patch applies the patch and returns the patched gCSSource.
func (c *FakeGCSSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GCSSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gcssourcesResource, c.ns, name, data, subresources...), &v1alpha1.GCSSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GCSSource), err
}
