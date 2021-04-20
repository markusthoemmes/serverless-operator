package serving

import (
	"context"
	"fmt"

	mf "github.com/manifestival/manifestival"
	"github.com/openshift-knative/serverless-operator/openshift-knative-operator/pkg/common"
	"github.com/openshift-knative/serverless-operator/openshift-knative-operator/pkg/monitoring"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"knative.dev/operator/pkg/apis/operator/v1alpha1"
)

const EnableServingMonitoringEnvVar = "ENABLE_SERVING_MONITORING_BY_DEFAULT"

var servingComponents = sets.NewString("activator", "autoscaler", "autoscaler-hpa", "controller", "domain-mapping", "domainmapping-webhook", "webhook")

func ReconcileMonitoring(ctx context.Context, api kubernetes.Interface, config v1alpha1.ConfigMapData, spec *v1alpha1.CommonSpec, ns string) error {
	if monitoring.ShouldEnableMonitoring(config, EnableServingMonitoringEnvVar) {
		if err := monitoring.ReconcileMonitoringLabelOnNamespace(ctx, ns, api, true); err != nil {
			return fmt.Errorf("failed to enable monitoring %w ", err)
		}
		return nil
	}
	// If "opencensus" is used we still dont want to scrape from a Serverless controlled namespace
	// user can always push to an agent collector in some other namespace and then integrate with OCP monitoring stack
	if err := monitoring.ReconcileMonitoringLabelOnNamespace(ctx, ns, api, false); err != nil {
		return fmt.Errorf("failed to disable monitoring %w ", err)
	}
	common.Configure(spec, monitoring.ObservabilityCMName, monitoring.ObservabilityBackendKey, "none")
	return nil
}

func GetMonitoringManifest(k v1alpha1.KComponent) ([]mf.Manifest, error) {
	if !monitoring.ShouldEnableMonitoring(k.GetSpec().GetConfig(), "bla") {
		return nil, nil
	}

	rbacManifest, err := monitoring.GetRBACManifest()
	if err != nil {
		return nil, err
	}

	crbM, err := monitoring.CreateClusterRoleBindingManifest("controller", k.GetNamespace())
	if err != nil {
		return nil, err
	}
	rbacManifest = rbacManifest.Append(*crbM)
	for c := range servingComponents {
		if err := monitoring.AppendManifestsForComponent(c, k.GetNamespace(), &rbacManifest); err != nil {
			return nil, err
		}
	}

	return []mf.Manifest{rbacManifest}, nil
}

func GetMonitoringTransformers(k v1alpha1.KComponent) []mf.Transformer {
	if !monitoring.ShouldEnableMonitoring(k.GetSpec().GetConfig(), "foo") {
		return []mf.Transformer{}
	}
	return []mf.Transformer{
		monitoring.InjectNamespaceWithSubject(k.GetNamespace(), monitoring.OpenshiftMonitoringNamespace),
		monitoring.InjectRbacProxyContainerToDeployments(),
	}
}
