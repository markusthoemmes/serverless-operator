package consoleclidownload

import (
	"fmt"

	"github.com/openshift-knative/serverless-operator/knative-operator/pkg/common"
	servingv1alpha1 "knative.dev/serving-operator/pkg/apis/serving/v1alpha1"

	mfc "github.com/manifestival/controller-runtime-client"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const knConsoleCLIDownload = "deploy/resources/console_cli_download_kn.yaml"

var log = common.Log.WithName("consoleclidownload")

// Create creates ConsoleCLIDownload for kn CLI download links
func Create(instance *servingv1alpha1.KnativeServing, apiclient client.Client) error {
	log.Info("Creating ConsoleCLIDownload CR for kn")
	manifest, err := mfc.NewManifest(knConsoleCLIDownload, apiclient)
	if err != nil {
		return fmt.Errorf("failed to read ConsoleCLIDownload manifest: %w", err)
	}
	if err := manifest.Apply(); err != nil {
		return fmt.Errorf("failed to apply ConsoleCLIDownload manifest: %w", err)
	}
	return nil
}

// Delete deletes ConsoleCLIDownload for kn CLI download links
func Delete(instance *servingv1alpha1.KnativeServing, apiclient client.Client) error {
	log.Info("Deleting ConsoleCLIDownload CR for kn")
	manifest, err := mfc.NewManifest(knConsoleCLIDownload, apiclient)
	if err != nil {
		return fmt.Errorf("failed to read ConsoleCLIDownload manifest: %w", err)
	}
	if err := manifest.Delete(); err != nil {
		return fmt.Errorf("failed to delete ConsoleCLIDownload manifest: %w", err)
	}
	return nil
}
