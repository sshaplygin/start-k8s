FROM golang:1.22-alpine

COPY . /app

WORKDIR /app

RUN ls -la . && \
    go install . && \
    which kubeapp1

ENTRYPOINT ["/go/bin/kubeapp1"]