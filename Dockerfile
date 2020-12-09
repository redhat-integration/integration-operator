# Build the manager binary
FROM registry.access.redhat.com/ubi8/go-toolset as builder

USER root
WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy vendored dependencies
COPY vendor/ vendor/

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager main.go

# Use minimal base image to package the manager binary
FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /
COPY --from=builder /workspace/manager .

ENTRYPOINT ["/manager"]
