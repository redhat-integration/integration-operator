# Red Hat Integration Operator

Red Hat Integration is a comprehensive set of integration and event processing technologies for creating, extending, and deploying container-based integration services across hybrid and multicloud environments. Red Hat Integration provides an agile, distributed, and API-centric solution that organizations can use to connect and share data between applications and systems in a digital world.

**Red Hat Integration Operator** allows you to choose and install the Operators that manage your Red Hat Integration components.

## Prerequisites

* [make](https://www.gnu.org/software/make/)
* [go](https://golang.org/dl/) version v1.15+
* [docker](https://docs.docker.com/get-docker/)
* [operator-sdk](https://sdk.operatorframework.io/docs/installation/)
* [opm](https://github.com/operator-framework/operator-registry/releases)

## Make Targets

### Generate manifests e.g. CRD, RBAC etc.

```sh
make manifests
```

### Generate code

```sh
make generate
```

### Build manager binary

```sh
make manager
```

### Run tests

```sh
make test
```

### Install CRDs into a cluster

```sh
make install
```

### Uninstall CRDs from a cluster

```sh
make uninstall
```

### Run against the configured Kubernetes cluster in ~/.kube/config

```sh
make run
```

### Deploy controller in the configured Kubernetes cluster in ~/.kube/config

```sh
make deploy
```

### Build the docker image

```sh
make docker-build
```

### Push the docker image

```sh
make docker-push
```

### Generate bundle manifests and metadata, then validate generated files

```sh
make bundle
```

### Build the bundle image

```sh
make bundle-build
```

### Push the bundle image

```sh
make bundle-push
```

### Build the index image

```sh
make index-build
```

### Push the index image

```sh
make index-push
```

## License

This project is [Apache License 2.0](LICENSE) licensed.
