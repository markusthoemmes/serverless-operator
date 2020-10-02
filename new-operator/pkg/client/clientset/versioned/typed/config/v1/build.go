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

// BuildsGetter has a method to return a BuildInterface.
// A group's client should implement this interface.
type BuildsGetter interface {
	Builds() BuildInterface
}

// BuildInterface has methods to work with Build resources.
type BuildInterface interface {
	Create(ctx context.Context, build *v1.Build, opts metav1.CreateOptions) (*v1.Build, error)
	Update(ctx context.Context, build *v1.Build, opts metav1.UpdateOptions) (*v1.Build, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Build, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.BuildList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Build, err error)
	BuildExpansion
}

// builds implements BuildInterface
type builds struct {
	client rest.Interface
}

// newBuilds returns a Builds
func newBuilds(c *ConfigV1Client) *builds {
	return &builds{
		client: c.RESTClient(),
	}
}

// Get takes name of the build, and returns the corresponding build object, and an error if there is any.
func (c *builds) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Build, err error) {
	result = &v1.Build{}
	err = c.client.Get().
		Resource("builds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Builds that match those selectors.
func (c *builds) List(ctx context.Context, opts metav1.ListOptions) (result *v1.BuildList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.BuildList{}
	err = c.client.Get().
		Resource("builds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested builds.
func (c *builds) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("builds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a build and creates it.  Returns the server's representation of the build, and an error, if there is any.
func (c *builds) Create(ctx context.Context, build *v1.Build, opts metav1.CreateOptions) (result *v1.Build, err error) {
	result = &v1.Build{}
	err = c.client.Post().
		Resource("builds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(build).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a build and updates it. Returns the server's representation of the build, and an error, if there is any.
func (c *builds) Update(ctx context.Context, build *v1.Build, opts metav1.UpdateOptions) (result *v1.Build, err error) {
	result = &v1.Build{}
	err = c.client.Put().
		Resource("builds").
		Name(build.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(build).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the build and deletes it. Returns an error if one occurs.
func (c *builds) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("builds").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *builds) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("builds").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched build.
func (c *builds) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Build, err error) {
	result = &v1.Build{}
	err = c.client.Patch(pt).
		Resource("builds").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
