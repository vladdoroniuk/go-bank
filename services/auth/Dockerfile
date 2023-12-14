FROM golang:1.21.4-alpine

# Define envs
ENV HOST_SERVICE_DIR=/services/auth
ENV CONTAINER_APP_DIR=/app

# Set destination for COPY
WORKDIR ${CONTAINER_APP_DIR}

# Download Go modules
COPY ${HOST_SERVICE_DIR}/go.mod ${HOST_SERVICE_DIR}/go.sum ./
RUN go mod download

# Copy the source code
COPY ${HOST_SERVICE_DIR}/*.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/auth

# Run
CMD ["/bin/auth"]