/*
Copyright 2021 The Knative Authors

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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/eventing-kafka/pkg/apis/messaging/v1alpha1"
	scheme "knative.dev/eventing-kafka/pkg/client/clientset/versioned/scheme"
)

// KafkaChannelsGetter has a method to return a KafkaChannelInterface.
// A group's client should implement this interface.
type KafkaChannelsGetter interface {
	KafkaChannels(namespace string) KafkaChannelInterface
}

// KafkaChannelInterface has methods to work with KafkaChannel resources.
type KafkaChannelInterface interface {
	Create(ctx context.Context, kafkaChannel *v1alpha1.KafkaChannel, opts v1.CreateOptions) (*v1alpha1.KafkaChannel, error)
	Update(ctx context.Context, kafkaChannel *v1alpha1.KafkaChannel, opts v1.UpdateOptions) (*v1alpha1.KafkaChannel, error)
	UpdateStatus(ctx context.Context, kafkaChannel *v1alpha1.KafkaChannel, opts v1.UpdateOptions) (*v1alpha1.KafkaChannel, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KafkaChannel, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KafkaChannelList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KafkaChannel, err error)
	KafkaChannelExpansion
}

// kafkaChannels implements KafkaChannelInterface
type kafkaChannels struct {
	client rest.Interface
	ns     string
}

// newKafkaChannels returns a KafkaChannels
func newKafkaChannels(c *MessagingV1alpha1Client, namespace string) *kafkaChannels {
	return &kafkaChannels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kafkaChannel, and returns the corresponding kafkaChannel object, and an error if there is any.
func (c *kafkaChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KafkaChannel, err error) {
	result = &v1alpha1.KafkaChannel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkachannels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KafkaChannels that match those selectors.
func (c *kafkaChannels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KafkaChannelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KafkaChannelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkachannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kafkaChannels.
func (c *kafkaChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kafkachannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kafkaChannel and creates it.  Returns the server's representation of the kafkaChannel, and an error, if there is any.
func (c *kafkaChannels) Create(ctx context.Context, kafkaChannel *v1alpha1.KafkaChannel, opts v1.CreateOptions) (result *v1alpha1.KafkaChannel, err error) {
	result = &v1alpha1.KafkaChannel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kafkachannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kafkaChannel).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kafkaChannel and updates it. Returns the server's representation of the kafkaChannel, and an error, if there is any.
func (c *kafkaChannels) Update(ctx context.Context, kafkaChannel *v1alpha1.KafkaChannel, opts v1.UpdateOptions) (result *v1alpha1.KafkaChannel, err error) {
	result = &v1alpha1.KafkaChannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkachannels").
		Name(kafkaChannel.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kafkaChannel).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *kafkaChannels) UpdateStatus(ctx context.Context, kafkaChannel *v1alpha1.KafkaChannel, opts v1.UpdateOptions) (result *v1alpha1.KafkaChannel, err error) {
	result = &v1alpha1.KafkaChannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkachannels").
		Name(kafkaChannel.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kafkaChannel).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kafkaChannel and deletes it. Returns an error if one occurs.
func (c *kafkaChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkachannels").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kafkaChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkachannels").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kafkaChannel.
func (c *kafkaChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KafkaChannel, err error) {
	result = &v1alpha1.KafkaChannel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kafkachannels").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
