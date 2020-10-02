// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	scheme "github.com/openshift-knative/serverless-operator/new-operator/pkg/client/clientset/versioned/scheme"
	v1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FeatureGatesGetter has a method to return a FeatureGateInterface.
// A group's client should implement this interface.
type FeatureGatesGetter interface {
	FeatureGates() FeatureGateInterface
}

// FeatureGateInterface has methods to work with FeatureGate resources.
type FeatureGateInterface interface {
	Create(ctx context.Context, featureGate *v1.FeatureGate, opts metav1.CreateOptions) (*v1.FeatureGate, error)
	Update(ctx context.Context, featureGate *v1.FeatureGate, opts metav1.UpdateOptions) (*v1.FeatureGate, error)
	UpdateStatus(ctx context.Context, featureGate *v1.FeatureGate, opts metav1.UpdateOptions) (*v1.FeatureGate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.FeatureGate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.FeatureGateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.FeatureGate, err error)
	FeatureGateExpansion
}

// featureGates implements FeatureGateInterface
type featureGates struct {
	client rest.Interface
}

// newFeatureGates returns a FeatureGates
func newFeatureGates(c *ConfigV1Client) *featureGates {
	return &featureGates{
		client: c.RESTClient(),
	}
}

// Get takes name of the featureGate, and returns the corresponding featureGate object, and an error if there is any.
func (c *featureGates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.FeatureGate, err error) {
	result = &v1.FeatureGate{}
	err = c.client.Get().
		Resource("featuregates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FeatureGates that match those selectors.
func (c *featureGates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.FeatureGateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.FeatureGateList{}
	err = c.client.Get().
		Resource("featuregates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested featureGates.
func (c *featureGates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("featuregates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a featureGate and creates it.  Returns the server's representation of the featureGate, and an error, if there is any.
func (c *featureGates) Create(ctx context.Context, featureGate *v1.FeatureGate, opts metav1.CreateOptions) (result *v1.FeatureGate, err error) {
	result = &v1.FeatureGate{}
	err = c.client.Post().
		Resource("featuregates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(featureGate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a featureGate and updates it. Returns the server's representation of the featureGate, and an error, if there is any.
func (c *featureGates) Update(ctx context.Context, featureGate *v1.FeatureGate, opts metav1.UpdateOptions) (result *v1.FeatureGate, err error) {
	result = &v1.FeatureGate{}
	err = c.client.Put().
		Resource("featuregates").
		Name(featureGate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(featureGate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *featureGates) UpdateStatus(ctx context.Context, featureGate *v1.FeatureGate, opts metav1.UpdateOptions) (result *v1.FeatureGate, err error) {
	result = &v1.FeatureGate{}
	err = c.client.Put().
		Resource("featuregates").
		Name(featureGate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(featureGate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the featureGate and deletes it. Returns an error if one occurs.
func (c *featureGates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("featuregates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *featureGates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("featuregates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched featureGate.
func (c *featureGates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.FeatureGate, err error) {
	result = &v1.FeatureGate{}
	err = c.client.Patch(pt).
		Resource("featuregates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
