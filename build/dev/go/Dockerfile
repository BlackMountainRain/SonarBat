FROM ubuntu:noble as builder

ENV PATH="/go/bin:/usr/local/go/bin:${PATH}"
ENV GOLANG_VERSION=1.22.2
ENV GOTOOLCHAIN=local
ENV GOPATH="/go"

RUN apt update --fix-missing && apt install --fix-missing -y protobuf-compiler wget make

RUN arch=$(uname -m) && \
    case $arch in \
        x86_64) wget https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz -O go.tar.gz ;; \
        i386) wget https://go.dev/dl/go${GOLANG_VERSION}.linux-386.tar.gz -O go.tar.gz ;; \
        armv7l) wget https://go.dev/dl/go${GOLANG_VERSION}.linux-arm64.tar.gz -O go.tar.gz ;; \
        aarch64) wget https://go.dev/dl/go${GOLANG_VERSION}.linux-arm64.tar.gz -O go.tar.gz ;; \
        *) echo "Unsupported architecture: $arch"; exit 1 ;; \
    esac

RUN tar -C /usr/local -xzf go.tar.gz

WORKDIR /go/src/sb

COPY ./ ./

RUN make init

FROM ubuntu:noble as runtime

RUN apt update --fix-missing && \
    apt install --fix-missing -y protobuf-compiler make git && \
    apt clean

ENV PATH="/go/bin:/usr/local/go/bin:${PATH}"
ENV GOLANG_VERSION=1.22.2
ENV GOTOOLCHAIN=local
ENV GOPATH="/go"

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin"

COPY --from=builder /usr/local/go /usr/local/go
COPY --from=builder "$GOPATH/bin" "$GOPATH/bin"

WORKDIR /go/src/sb
