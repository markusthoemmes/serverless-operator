// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	scheme "github.com/openshift-knative/serverless-operator/openshift-knative-operator/pkg/client/clientset/versioned/scheme"
	v1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IngressesGetter has a method to return a IngressInterface.
// A group's client should implement this interface.
type IngressesGetter interface {
	Ingresses() IngressInterface
}

// IngressInterface has methods to work with Ingress resources.
type IngressInterface interface {
	Create(*v1.Ingress) (*v1.Ingress, error)
	Update(*v1.Ingress) (*v1.Ingress, error)
	UpdateStatus(*v1.Ingress) (*v1.Ingress, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Ingress, error)
	List(opts metav1.ListOptions) (*v1.IngressList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Ingress, err error)
	IngressExpansion
}

// ingresses implements IngressInterface
type ingresses struct {
	client rest.Interface
}

// newIngresses returns a Ingresses
func newIngresses(c *ConfigV1Client) *ingresses {
	return &ingresses{
		client: c.RESTClient(),
	}
}

// Get takes name of the ingress, and returns the corresponding ingress object, and an error if there is any.
func (c *ingresses) Get(name string, options metav1.GetOptions) (result *v1.Ingress, err error) {
	result = &v1.Ingress{}
	err = c.client.Get().
		Resource("ingresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Ingresses that match those selectors.
func (c *ingresses) List(opts metav1.ListOptions) (result *v1.IngressList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.IngressList{}
	err = c.client.Get().
		Resource("ingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ingresses.
func (c *ingresses) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ingress and creates it.  Returns the server's representation of the ingress, and an error, if there is any.
func (c *ingresses) Create(ingress *v1.Ingress) (result *v1.Ingress, err error) {
	result = &v1.Ingress{}
	err = c.client.Post().
		Resource("ingresses").
		Body(ingress).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ingress and updates it. Returns the server's representation of the ingress, and an error, if there is any.
func (c *ingresses) Update(ingress *v1.Ingress) (result *v1.Ingress, err error) {
	result = &v1.Ingress{}
	err = c.client.Put().
		Resource("ingresses").
		Name(ingress.Name).
		Body(ingress).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ingresses) UpdateStatus(ingress *v1.Ingress) (result *v1.Ingress, err error) {
	result = &v1.Ingress{}
	err = c.client.Put().
		Resource("ingresses").
		Name(ingress.Name).
		SubResource("status").
		Body(ingress).
		Do().
		Into(result)
	return
}

// Delete takes name of the ingress and deletes it. Returns an error if one occurs.
func (c *ingresses) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ingresses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ingresses) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ingresses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ingress.
func (c *ingresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Ingress, err error) {
	result = &v1.Ingress{}
	err = c.client.Patch(pt).
		Resource("ingresses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
