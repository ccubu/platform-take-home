FROM golang:1.23.2

RUN apt-get update && apt-get install -y \
    git \
    curl \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v1.53.3

# Install buf
RUN curl -sSL "https://github.com/bufbuild/buf/releases/download/v1.16.0/buf-Linux-x86_64" -o /usr/local/bin/buf \
    && chmod +x /usr/local/bin/buf

# Install protoc
RUN curl -sSL "https://github.com/protocolbuffers/protobuf/releases/download/v22.3/protoc-22.3-linux-x86_64.zip" -o /tmp/protoc.zip \
    && unzip /tmp/protoc.zip -d /usr/local/ \
    && rm -f /tmp/protoc.zip

# Install protoc-gen-go and protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Install protoc-gen-go-grpc-mock
RUN go install github.com/sorcererxw/protoc-gen-go-grpc-mock@latest

# Install protoc-gen-grpc-gateway
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

# Install Lefthook
RUN go install github.com/evilmartians/lefthook@latest

# Install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" \
    && chmod +x kubectl \
    && mv kubectl /usr/local/bin/

ENV PATH="/usr/local/bin:/usr/local/include:${PATH}"

WORKDIR /workspace

CMD ["bash"]