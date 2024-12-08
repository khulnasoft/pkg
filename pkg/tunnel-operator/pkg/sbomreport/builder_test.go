package sbomreport_test

import (
	"testing"

	"github.com/khulnasoft/tunnel-operator/pkg/apis/khulnasoft/v1alpha1"
	"github.com/khulnasoft/tunnel-operator/pkg/sbomreport"
	"github.com/khulnasoft/tunnel-operator/pkg/tunneloperator"
	"github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/utils/pointer"
)

func TestReportBuilder(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	report, err := sbomreport.NewReportBuilder(scheme.Scheme).
		Controller(&appsv1.ReplicaSet{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ReplicaSet",
				APIVersion: "apps/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "some-owner",
				Namespace: "qa",
				Labels:    labels.Set{"tier": "tier-1", "owner": "team-a"},
			},
		}).
		Container("my-container").
		PodSpecHash("xyz").
		Data(v1alpha1.SbomReportData{}).
		ResourceLabelsToInclude([]string{"tier"}).
		Get()

	g.Expect(err).ToNot(gomega.HaveOccurred())
	g.Expect(report).To(gomega.Equal(v1alpha1.SbomReport{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "replicaset-some-owner-my-container",
			Namespace: "qa",
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion:         "apps/v1",
					Kind:               "ReplicaSet",
					Name:               "some-owner",
					Controller:         pointer.Bool(true),
					BlockOwnerDeletion: pointer.Bool(false),
				},
			},
			Labels: map[string]string{
				tunneloperator.LabelResourceKind:      "ReplicaSet",
				tunneloperator.LabelResourceName:      "some-owner",
				tunneloperator.LabelResourceNamespace: "qa",
				tunneloperator.LabelContainerName:     "my-container",
				tunneloperator.LabelResourceSpecHash:  "xyz",
				"tier":                               "tier-1",
			},
		},
		Report: v1alpha1.SbomReportData{},
	}))
}