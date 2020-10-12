// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/openshift-knative/serverless-operator/openshift-knative-operator/pkg/client/clientset/versioned"
	configv1 "github.com/openshift-knative/serverless-operator/openshift-knative-operator/pkg/client/clientset/versioned/typed/config/v1"
	fakeconfigv1 "github.com/openshift-knative/serverless-operator/openshift-knative-operator/pkg/client/clientset/versioned/typed/config/v1/fake"
	routev1 "github.com/openshift-knative/serverless-operator/openshift-knative-operator/pkg/client/clientset/versioned/typed/route/v1"
	fakeroutev1 "github.com/openshift-knative/serverless-operator/openshift-knative-operator/pkg/client/clientset/versioned/typed/route/v1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// ConfigV1 retrieves the ConfigV1Client
func (c *Clientset) ConfigV1() configv1.ConfigV1Interface {
	return &fakeconfigv1.FakeConfigV1{Fake: &c.Fake}
}

// RouteV1 retrieves the RouteV1Client
func (c *Clientset) RouteV1() routev1.RouteV1Interface {
	return &fakeroutev1.FakeRouteV1{Fake: &c.Fake}
}
