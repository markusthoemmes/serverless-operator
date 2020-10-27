package knativeserving

import (
	"context"
	"os"
	"testing"

	"github.com/openshift-knative/serverless-operator/knative-operator/pkg/apis"
	"github.com/openshift-knative/serverless-operator/knative-operator/pkg/webhook/testutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	servingv1alpha1 "knative.dev/operator/pkg/apis/operator/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/types"
)

var (
	ks1 = &servingv1alpha1.KnativeServing{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ks1",
		},
	}
	ks2 = &servingv1alpha1.KnativeServing{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ks2",
		},
	}

	decoder types.Decoder
)

func init() {
	apis.AddToScheme(scheme.Scheme)

	dec, err := admission.NewDecoder(scheme.Scheme)
	if err != nil {
		panic(err)
	}
	decoder = dec
}

func TestInvalidNamespace(t *testing.T) {
	os.Clearenv()
	os.Setenv("REQUIRED_SERVING_NAMESPACE", "knative-serving")

	validator := Validator{}
	validator.InjectDecoder(decoder)

	req, err := testutil.RequestFor(ks1)
	if err != nil {
		t.Fatalf("Failed to generate a request for %v: %v", ks1, err)
	}

	result := validator.Handle(context.Background(), req)
	if result.Response.Allowed {
		t.Error("The required namespace is wrong, but the request is allowed")
	}
}

func TestLoneliness(t *testing.T) {
	os.Clearenv()

	validator := Validator{}
	validator.InjectDecoder(decoder)
	validator.InjectClient(fake.NewFakeClient(ks2))

	req, err := testutil.RequestFor(ks1)
	if err != nil {
		t.Fatalf("Failed to generate a request for %v: %v", ks1, err)
	}

	result := validator.Handle(context.Background(), req)
	if result.Response.Allowed {
		t.Errorf("Too many KnativeServings: %v", result.Response)
	}
}
