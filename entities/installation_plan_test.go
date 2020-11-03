package entities_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	integrationv1 "github.com/redhat-integration/integration-operator/api/v1"
	"github.com/redhat-integration/integration-operator/entities"
)

var _ = Describe("InstallationPlan", func() {
	Context("CreateInstallationPlans", func() {
		It("Should return correct result", func() {
			installation := &integrationv1.Installation{
				Spec: integrationv1.InstallationSpec{
					ThreeScaleInstallationInput: &integrationv1.ThreeScaleInstallationInput{
						Enabled:   false,
						Mode:      "namespace",
						Namespace: "rhi-3scale",
					},
					ThreeScaleAPIcastInstallationInput: &integrationv1.ThreeScaleAPIcastInstallationInput{
						Enabled: true,
						Mode:    "namespace",
					},
					AMQBrokerInstallationInput: &integrationv1.AMQBrokerInstallationInput{
						Enabled:   true,
						Mode:      "namespace",
						Namespace: "rhi-amq-broker",
					},
					AMQInterconnectInstallationInput: &integrationv1.AMQInterconnectInstallationInput{
						Enabled:   true,
						Mode:      "namespace",
						Namespace: "rhi-amq-interconnect",
					},
					AMQStreamsInstallationInput: &integrationv1.AMQStreamsInstallationInput{
						Enabled:   true,
						Mode:      "cluster",
						Namespace: "rhi-amq-streams",
					},
					APIDesignerInstallationInput: &integrationv1.APIDesignerInstallationInput{
						Enabled:   true,
						Mode:      "namespace",
						Namespace: "rhi-api-designer",
					},
					CamelKInstallationInput: &integrationv1.CamelKInstallationInput{
						Enabled: true,
						Mode:    "cluster",
					},
					FuseConsoleInstallationInput: &integrationv1.FuseConsoleInstallationInput{
						Enabled:   true,
						Mode:      "namespace",
						Namespace: "rhi-fuse-console",
					},
					FuseOnlineInstallationInput: &integrationv1.FuseOnlineInstallationInput{
						Enabled:   true,
						Mode:      "namespace",
						Namespace: "rhi-fuse-online",
					},
					ServiceRegistryInstallationInput: &integrationv1.ServiceRegistryInstallationInput{
						Enabled:   true,
						Mode:      "namespace",
						Namespace: "rhi-service-registry",
					},
				},
			}

			installationPlans := entities.CreateInstallationPlans(installation)

			Expect(len(installationPlans)).To(Equal(10))

			Expect(installationPlans[0].GetChannel()).To(Equal("threescale-2.9"))
			Expect(installationPlans[0].GetConditionType()).To(Equal("3scaleOperatorInstalled"))
			Expect(installationPlans[0].IsEnabled()).To(BeFalse())
			Expect(installationPlans[0].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[0].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[0].GetNamespace()).To(Equal("rhi-3scale"))
			Expect(installationPlans[0].GetPackageName()).To(Equal("3scale-operator"))

			Expect(installationPlans[1].GetChannel()).To(Equal("threescale-2.9"))
			Expect(installationPlans[1].GetConditionType()).To(Equal("3scaleAPIcastOperatorInstalled"))
			Expect(installationPlans[1].IsEnabled()).To(BeTrue())
			Expect(installationPlans[1].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[1].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[1].GetNamespace()).To(Equal("rhi-3scale-apicast"))
			Expect(installationPlans[1].GetPackageName()).To(Equal("apicast-operator"))

			Expect(installationPlans[2].GetChannel()).To(Equal("current"))
			Expect(installationPlans[2].GetConditionType()).To(Equal("AMQBrokerOperatorInstalled"))
			Expect(installationPlans[2].IsEnabled()).To(BeTrue())
			Expect(installationPlans[2].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[2].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[2].GetNamespace()).To(Equal("rhi-amq-broker"))
			Expect(installationPlans[2].GetPackageName()).To(Equal("amq-broker"))

			Expect(installationPlans[3].GetChannel()).To(Equal("1.2.0"))
			Expect(installationPlans[3].GetConditionType()).To(Equal("AMQInterconnectOperatorInstalled"))
			Expect(installationPlans[3].IsEnabled()).To(BeTrue())
			Expect(installationPlans[3].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[3].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[3].GetNamespace()).To(Equal("rhi-amq-interconnect"))
			Expect(installationPlans[3].GetPackageName()).To(Equal("amq7-interconnect-operator"))

			Expect(installationPlans[4].GetChannel()).To(Equal("stable"))
			Expect(installationPlans[4].GetConditionType()).To(Equal("AMQStreamsOperatorInstalled"))
			Expect(installationPlans[4].IsEnabled()).To(BeTrue())
			Expect(installationPlans[4].IsNamespaceMode()).To(BeFalse())
			Expect(installationPlans[4].IsClusterMode()).To(BeTrue())
			Expect(installationPlans[4].GetNamespace()).To(Equal("openshift-operators"))
			Expect(installationPlans[4].GetPackageName()).To(Equal("amq-streams"))

			Expect(installationPlans[5].GetChannel()).To(Equal("alpha"))
			Expect(installationPlans[5].GetConditionType()).To(Equal("APIDesignerOperatorInstalled"))
			Expect(installationPlans[5].IsEnabled()).To(BeTrue())
			Expect(installationPlans[5].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[5].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[5].GetNamespace()).To(Equal("rhi-api-designer"))
			Expect(installationPlans[5].GetPackageName()).To(Equal("fuse-apicurito"))

			Expect(installationPlans[6].GetChannel()).To(Equal("techpreview"))
			Expect(installationPlans[6].GetConditionType()).To(Equal("CamelKOperatorInstalled"))
			Expect(installationPlans[6].IsEnabled()).To(BeTrue())
			Expect(installationPlans[6].IsNamespaceMode()).To(BeFalse())
			Expect(installationPlans[6].IsClusterMode()).To(BeTrue())
			Expect(installationPlans[6].GetNamespace()).To(Equal("openshift-operators"))
			Expect(installationPlans[6].GetPackageName()).To(Equal("red-hat-camel-k"))

			Expect(installationPlans[7].GetChannel()).To(Equal("alpha"))
			Expect(installationPlans[7].GetConditionType()).To(Equal("FuseConsoleOperatorInstalled"))
			Expect(installationPlans[7].IsEnabled()).To(BeTrue())
			Expect(installationPlans[7].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[7].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[7].GetNamespace()).To(Equal("rhi-fuse-console"))
			Expect(installationPlans[7].GetPackageName()).To(Equal("fuse-console"))

			Expect(installationPlans[8].GetChannel()).To(Equal("alpha"))
			Expect(installationPlans[8].GetConditionType()).To(Equal("FuseOnlineOperatorInstalled"))
			Expect(installationPlans[8].IsEnabled()).To(BeTrue())
			Expect(installationPlans[8].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[8].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[8].GetNamespace()).To(Equal("rhi-fuse-online"))
			Expect(installationPlans[8].GetPackageName()).To(Equal("fuse-online"))

			Expect(installationPlans[9].GetChannel()).To(Equal("serviceregistry-1.0"))
			Expect(installationPlans[9].GetConditionType()).To(Equal("ServiceRegistryOperatorInstalled"))
			Expect(installationPlans[9].IsEnabled()).To(BeTrue())
			Expect(installationPlans[9].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[9].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[9].GetNamespace()).To(Equal("rhi-service-registry"))
			Expect(installationPlans[9].GetPackageName()).To(Equal("service-registry-operator"))
		})
	})
})
