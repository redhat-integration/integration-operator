module github.com/redhat-integration/integration-operator

go 1.15

require (
	github.com/go-logr/logr v0.4.0
	github.com/go-logr/zapr v0.2.0 // indirect
	github.com/magefile/mage v1.11.0 // indirect
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.2
	github.com/opencontainers/runc v1.0.0-rc93 // indirect
	github.com/operator-framework/api v0.3.25
	github.com/operator-framework/operator-lifecycle-manager v0.17.0
	github.com/sirupsen/logrus v1.8.0 // indirect
	k8s.io/api v0.21.0
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
	sigs.k8s.io/controller-runtime v0.6.3
)
