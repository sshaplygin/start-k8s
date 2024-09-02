FROM golang:1.22-alpine

COPY . /app

WORKDIR /app

RUN ls -la . && \
    go install . && \
    which kubeapp2

ENTRYPOINT ["/go/bin/kubeapp2"]