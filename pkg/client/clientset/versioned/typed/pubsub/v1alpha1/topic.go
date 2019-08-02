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

package v1alpha1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/cloud-run-events/pkg/apis/pubsub/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/cloud-run-events/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TopicsGetter has a method to return a TopicInterface.
// A group's client should implement this interface.
type TopicsGetter interface {
	Topics(namespace string) TopicInterface
}

// TopicInterface has methods to work with Topic resources.
type TopicInterface interface {
	Create(*v1alpha1.Topic) (*v1alpha1.Topic, error)
	Update(*v1alpha1.Topic) (*v1alpha1.Topic, error)
	UpdateStatus(*v1alpha1.Topic) (*v1alpha1.Topic, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Topic, error)
	List(opts v1.ListOptions) (*v1alpha1.TopicList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Topic, err error)
	TopicExpansion
}

// topics implements TopicInterface
type topics struct {
	client rest.Interface
	ns     string
}

// newTopics returns a Topics
func newTopics(c *PubsubV1alpha1Client, namespace string) *topics {
	return &topics{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the topic, and returns the corresponding topic object, and an error if there is any.
func (c *topics) Get(name string, options v1.GetOptions) (result *v1alpha1.Topic, err error) {
	result = &v1alpha1.Topic{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("topics").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Topics that match those selectors.
func (c *topics) List(opts v1.ListOptions) (result *v1alpha1.TopicList, err error) {
	result = &v1alpha1.TopicList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("topics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested topics.
func (c *topics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("topics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a topic and creates it.  Returns the server's representation of the topic, and an error, if there is any.
func (c *topics) Create(topic *v1alpha1.Topic) (result *v1alpha1.Topic, err error) {
	result = &v1alpha1.Topic{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("topics").
		Body(topic).
		Do().
		Into(result)
	return
}

// Update takes the representation of a topic and updates it. Returns the server's representation of the topic, and an error, if there is any.
func (c *topics) Update(topic *v1alpha1.Topic) (result *v1alpha1.Topic, err error) {
	result = &v1alpha1.Topic{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("topics").
		Name(topic.Name).
		Body(topic).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *topics) UpdateStatus(topic *v1alpha1.Topic) (result *v1alpha1.Topic, err error) {
	result = &v1alpha1.Topic{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("topics").
		Name(topic.Name).
		SubResource("status").
		Body(topic).
		Do().
		Into(result)
	return
}

// Delete takes name of the topic and deletes it. Returns an error if one occurs.
func (c *topics) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("topics").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *topics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("topics").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched topic.
func (c *topics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Topic, err error) {
	result = &v1alpha1.Topic{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("topics").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
