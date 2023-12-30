FROM golang:1.21.5-alpine

# Define envs
ENV HOST_SERVICE_DIR=/services/auth
ENV CONTAINER_ROOT_DIR=/app
ENV CONTAINER_SERVICE_DIR=${CONTAINER_ROOT_DIR}${HOST_SERVICE_DIR}
ENV ALPINE_TESTING_REPO=http://dl-cdn.alpinelinux.org/alpine/edge/testing/
ENV PROTO_GEN_MODULE=github.com/vladdoroniuk/rose/proto_gen

# Install CLI tools
RUN apk add protobuf
RUN apk add buf --repository=${ALPINE_TESTING_REPO}
RUN apk add protoc-gen-go --repository=${ALPINE_TESTING_REPO}

# Compile .proto files
WORKDIR ${CONTAINER_ROOT_DIR}
COPY proto ./proto
COPY buf.gen.yaml ./
RUN buf generate proto

# Download Go modules for .proto files
WORKDIR ${CONTAINER_ROOT_DIR}/proto_gen
RUN go mod init ${PROTO_GEN_MODULE}
RUN go mod tidy

# Download Go modules for service
WORKDIR ${CONTAINER_SERVICE_DIR}
COPY ${HOST_SERVICE_DIR}/go.mod ${HOST_SERVICE_DIR}/go.sum ./
RUN go mod download

# Copy the source code
COPY ${HOST_SERVICE_DIR}/*.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/auth

# Run
CMD ["sh", "-c", "${CONTAINER_SERVICE_DIR}/bin/auth"]