package common

import (
	"os"

	"k8s.io/apimachinery/pkg/api/resource"

	eventingv1alpha1 "knative.dev/operator/pkg/apis/operator/v1alpha1"
)

func MutateEventing(ke *eventingv1alpha1.KnativeEventing) {
	eventingImagesFromEnviron(ke)
	ensureEventingWebhookMemoryLimit(ke)
}

// eventingImagesFromEnviron overrides registry images
func eventingImagesFromEnviron(ke *eventingv1alpha1.KnativeEventing) {
	ke.Spec.Registry.Override = BuildImageOverrideMapFromEnviron(os.Environ())

	if defaultVal, ok := ke.Spec.Registry.Override["default"]; ok {
		ke.Spec.Registry.Default = defaultVal
	}

	log.Info("Setting", "registry", ke.Spec.Registry)
}

func ensureEventingWebhookMemoryLimit(ks *eventingv1alpha1.KnativeEventing) {
	EnsureContainerMemoryLimit(&ks.Spec.CommonSpec, "eventing-webhook", resource.MustParse("1024Mi"))
}
