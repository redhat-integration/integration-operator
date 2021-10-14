package controllers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	integrationv1 "github.com/redhat-integration/integration-operator/api/v1"
	controllers "github.com/redhat-integration/integration-operator/controllers"
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
			installationConfig := &controllers.InstallationConfig{
				Channel3scale:          "channel-3scale",
				Channel3scaleAPIcast:   "channel-3scale-apicast",
				ChannelAMQBroker:       "channel-amq-broker",
				ChannelAMQInterconnect: "channel-amq-interconnect",
				ChannelAMQStreams:      "channel-amq-streams",
				ChannelAPIDesigner:     "channel-api-designer",
				ChannelCamelK:          "channel-camel-k",
				ChannelFuseConsole:     "channel-fuse-console",
				ChannelFuseOnline:      "channel-fuse-online",
				ChannelServiceRegistry: "channel-service-registry",
			}

			installationPlans := controllers.CreateInstallationPlans(installation, installationConfig)

			Expect(len(installationPlans)).To(Equal(10))

			Expect(installationPlans[0].Channel).To(Equal("channel-3scale"))
			Expect(installationPlans[0].ConditionType).To(Equal("3scaleOperatorInstalled"))
			Expect(installationPlans[0].Enabled).To(BeFalse())
			Expect(installationPlans[0].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[0].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[0].Name).To(Equal("3scale-operator"))
			Expect(installationPlans[0].Namespace).To(Equal("rhi-3scale"))
			Expect(installationPlans[0].PackageName).To(Equal("3scale-operator"))

			Expect(installationPlans[1].Channel).To(Equal("channel-3scale-apicast"))
			Expect(installationPlans[1].ConditionType).To(Equal("3scaleAPIcastOperatorInstalled"))
			Expect(installationPlans[1].Enabled).To(BeTrue())
			Expect(installationPlans[1].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[1].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[1].Name).To(Equal("3scale-apicast-operator"))
			Expect(installationPlans[1].Namespace).To(Equal("rhi-3scale-apicast"))
			Expect(installationPlans[1].PackageName).To(Equal("apicast-operator"))

			Expect(installationPlans[2].Channel).To(Equal("channel-amq-broker"))
			Expect(installationPlans[2].ConditionType).To(Equal("AMQBrokerOperatorInstalled"))
			Expect(installationPlans[2].Enabled).To(BeTrue())
			Expect(installationPlans[2].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[2].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[2].Name).To(Equal("amq-broker-operator"))
			Expect(installationPlans[2].Namespace).To(Equal("rhi-amq-broker"))
			Expect(installationPlans[2].PackageName).To(Equal("amq-broker-rhel8"))

			Expect(installationPlans[3].Channel).To(Equal("channel-amq-interconnect"))
			Expect(installationPlans[3].ConditionType).To(Equal("AMQInterconnectOperatorInstalled"))
			Expect(installationPlans[3].Enabled).To(BeTrue())
			Expect(installationPlans[3].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[3].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[3].Name).To(Equal("amq-interconnect-operator"))
			Expect(installationPlans[3].Namespace).To(Equal("rhi-amq-interconnect"))
			Expect(installationPlans[3].PackageName).To(Equal("amq7-interconnect-operator"))

			Expect(installationPlans[4].Channel).To(Equal("channel-amq-streams"))
			Expect(installationPlans[4].ConditionType).To(Equal("AMQStreamsOperatorInstalled"))
			Expect(installationPlans[4].Enabled).To(BeTrue())
			Expect(installationPlans[4].IsNamespaceMode()).To(BeFalse())
			Expect(installationPlans[4].IsClusterMode()).To(BeTrue())
			Expect(installationPlans[4].Name).To(Equal("amq-streams-operator"))
			Expect(installationPlans[4].Namespace).To(Equal("openshift-operators"))
			Expect(installationPlans[4].PackageName).To(Equal("amq-streams"))

			Expect(installationPlans[5].Channel).To(Equal("channel-api-designer"))
			Expect(installationPlans[5].ConditionType).To(Equal("APIDesignerOperatorInstalled"))
			Expect(installationPlans[5].Enabled).To(BeTrue())
			Expect(installationPlans[5].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[5].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[5].Name).To(Equal("api-designer-operator"))
			Expect(installationPlans[5].Namespace).To(Equal("rhi-api-designer"))
			Expect(installationPlans[5].PackageName).To(Equal("fuse-apicurito"))

			Expect(installationPlans[6].Channel).To(Equal("channel-camel-k"))
			Expect(installationPlans[6].ConditionType).To(Equal("CamelKOperatorInstalled"))
			Expect(installationPlans[6].Enabled).To(BeTrue())
			Expect(installationPlans[6].IsNamespaceMode()).To(BeFalse())
			Expect(installationPlans[6].IsClusterMode()).To(BeTrue())
			Expect(installationPlans[6].Name).To(Equal("camel-k-operator"))
			Expect(installationPlans[6].Namespace).To(Equal("openshift-operators"))
			Expect(installationPlans[6].PackageName).To(Equal("red-hat-camel-k"))

			Expect(installationPlans[7].Channel).To(Equal("channel-fuse-console"))
			Expect(installationPlans[7].ConditionType).To(Equal("FuseConsoleOperatorInstalled"))
			Expect(installationPlans[7].Enabled).To(BeTrue())
			Expect(installationPlans[7].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[7].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[7].Name).To(Equal("fuse-console-operator"))
			Expect(installationPlans[7].Namespace).To(Equal("rhi-fuse-console"))
			Expect(installationPlans[7].PackageName).To(Equal("fuse-console"))

			Expect(installationPlans[8].Channel).To(Equal("channel-fuse-online"))
			Expect(installationPlans[8].ConditionType).To(Equal("FuseOnlineOperatorInstalled"))
			Expect(installationPlans[8].Enabled).To(BeTrue())
			Expect(installationPlans[8].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[8].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[8].Name).To(Equal("fuse-online-operator"))
			Expect(installationPlans[8].Namespace).To(Equal("rhi-fuse-online"))
			Expect(installationPlans[8].PackageName).To(Equal("fuse-online"))

			Expect(installationPlans[9].Channel).To(Equal("channel-service-registry"))
			Expect(installationPlans[9].ConditionType).To(Equal("ServiceRegistryOperatorInstalled"))
			Expect(installationPlans[9].Enabled).To(BeTrue())
			Expect(installationPlans[9].IsNamespaceMode()).To(BeTrue())
			Expect(installationPlans[9].IsClusterMode()).To(BeFalse())
			Expect(installationPlans[9].Name).To(Equal("service-registry-operator"))
			Expect(installationPlans[9].Namespace).To(Equal("rhi-service-registry"))
			Expect(installationPlans[9].PackageName).To(Equal("service-registry-operator"))
		})
	})
})
