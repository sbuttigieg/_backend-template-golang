# TODO: replace entrypoint with actual values
FROM golang:1.24.4-bookworm

ARG TARGETARCH

RUN apt-get update && if [ "$TARGETARCH" = "arm64" ]; then apt install -y clang; fi \
    && apt-get install -y \
    git \
    curl \
    jq \
    build-essential

RUN go install github.com/cespare/reflex@latest
RUN go install github.com/mfridman/tparse@latest

VOLUME /home
WORKDIR /home

ENV GORACE="halt_on_error=1" CGO_ENABLED=1 GOOS=linux GOARCH=${TARGETARCH}
ENV GOCACHE="/home/.cache"
ENTRYPOINT sh -c 'reflex --all -r ".*\\.go$" --start-service=true -- go run -race api/cmd/backendtemplate/main.go'