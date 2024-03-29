# Build the manager binary
FROM registry.access.redhat.com/ubi8/go-toolset as builder

USER root
WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY olm/ olm/

# Build
RUN CGO_ENABLED=0 go build -a -o manager main.go

# Use minimal base image to package the manager binary
FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /
COPY --from=builder /workspace/manager .
USER 1001

ENTRYPOINT ["/manager"]
