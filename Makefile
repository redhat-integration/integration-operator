# Environment variables required for 'make run'
export CHANNEL_3SCALE = threescale-2.9
export CHANNEL_3SCALE_APICAST = threescale-2.9
export CHANNEL_AMQ_BROKER = current
export CHANNEL_AMQ_INTERCONNECT = 1.2.0
export CHANNEL_AMQ_STREAMS = amq-streams-1.6.x
export CHANNEL_API_DESIGNER = fuse-apicurito-7.8.x
export CHANNEL_CAMEL_K = techpreview
export CHANNEL_FUSE_CONSOLE = fuse-console-7.8.x
export CHANNEL_FUSE_ONLINE = fuse-online-v7.8.x
export CHANNEL_SERVICE_REGISTRY = serviceregistry-1.1

# Current Operator version
VERSION ?= 0.0.5
# Default bundle image tag
BUNDLE_IMG ?= quay.io/rh_integration/rhi-operator-bundle-dev:$(VERSION)
# Options for 'bundle-build'
CHANNELS ?= 1.x
BUNDLE_CHANNELS := --channels=$(CHANNELS)
DEFAULT_CHANNEL ?= 1.x
BUNDLE_DEFAULT_CHANNEL := --default-channel=$(DEFAULT_CHANNEL)
BUNDLE_METADATA_OPTS ?= $(BUNDLE_CHANNELS) $(BUNDLE_DEFAULT_CHANNEL)

# Image URL to use all building/pushing image targets
IMG ?= quay.io/rh_integration/rhi-operator-dev:$(VERSION)

# Index image tag
INDEX_IMG_NAME ?= quay.io/rh_integration/rhi-operator-index-dev
INDEX_IMG ?= $(INDEX_IMG_NAME):$(VERSION)
# Options for 'index-build'
ifneq ($(origin FROM_INDEX_VERSION), undefined)
FROM_INDEX_IMG := --from-index $(INDEX_IMG_NAME):$(FROM_INDEX_VERSION)
endif

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: manager

# Run tests
ENVTEST_ASSETS_DIR = $(shell pwd)/testbin
test: generate fmt vet manifests
	mkdir -p $(ENVTEST_ASSETS_DIR)
	test -f $(ENVTEST_ASSETS_DIR)/setup-envtest.sh || curl -sSLo $(ENVTEST_ASSETS_DIR)/setup-envtest.sh https://raw.githubusercontent.com/kubernetes-sigs/controller-runtime/v0.6.3/hack/setup-envtest.sh
	source $(ENVTEST_ASSETS_DIR)/setup-envtest.sh; fetch_envtest_tools $(ENVTEST_ASSETS_DIR); setup_envtest_env $(ENVTEST_ASSETS_DIR); go test ./... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	go run ./main.go

# Install CRDs into a cluster
install: manifests kustomize
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
uninstall: manifests kustomize
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests kustomize
	cd config/manager && $(KUSTOMIZE) edit set image controller=$(IMG)
	$(KUSTOMIZE) build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) crd rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

# Build the operator image
operator-build: manager test
	podman build . -t $(IMG)

# Push the operator image
operator-push:
	podman push $(IMG)

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.1 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

kustomize:
ifeq (, $(shell which kustomize))
	@{ \
	set -e ;\
	KUSTOMIZE_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$KUSTOMIZE_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/kustomize/kustomize/v3@v3.5.4 ;\
	rm -rf $$KUSTOMIZE_GEN_TMP_DIR ;\
	}
KUSTOMIZE=$(GOBIN)/kustomize
else
KUSTOMIZE=$(shell which kustomize)
endif

# Generate bundle manifests and metadata, then validate generated files.
.PHONY: bundle
bundle: manifests kustomize
	operator-sdk generate kustomize manifests -q
	cd config/manager && $(KUSTOMIZE) edit set image controller=$(IMG)
	$(KUSTOMIZE) build config/manifests | operator-sdk generate bundle -q --overwrite --version $(VERSION) $(BUNDLE_METADATA_OPTS)
	operator-sdk bundle validate ./bundle

# Build the bundle image
.PHONY: bundle-build
bundle-build: bundle
	podman build -f bundle.Dockerfile -t $(BUNDLE_IMG) .

# Push the bundle image
bundle-push:
	podman push ${BUNDLE_IMG}

# Build the index image (only use it for patch version upgrades)
index-build:
	opm index add -c podman --bundles $(BUNDLE_IMG) $(FROM_INDEX_IMG) --tag $(INDEX_IMG)

# Push the index image
index-push:
	podman push $(INDEX_IMG)

delete-namespaces:
	kubectl delete namespace rhi-3scale --ignore-not-found
	kubectl delete namespace rhi-3scale-apicast --ignore-not-found
	kubectl delete namespace rhi-amq-broker --ignore-not-found
	kubectl delete namespace rhi-amq-interconnect --ignore-not-found
	kubectl delete namespace rhi-amq-streams --ignore-not-found
	kubectl delete namespace rhi-api-designer --ignore-not-found
	kubectl delete namespace rhi-camel-k --ignore-not-found
	kubectl delete namespace rhi-fuse-console --ignore-not-found
	kubectl delete namespace rhi-fuse-online --ignore-not-found
	kubectl delete namespace rhi-service-registry --ignore-not-found
