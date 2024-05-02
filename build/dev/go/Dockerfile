FROM ubuntu:noble as builder

RUN apt update --fix-missing && apt install --fix-missing -y protobuf-compiler wget make

RUN arch=$(uname -m) && \
    case $arch in \
        x86_64) wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz -O go.tar.gz ;; \
        i386) wget https://go.dev/dl/go1.22.2.linux-386.tar.gz -O go.tar.gz ;; \
        armv7l) wget https://go.dev/dl/go1.22.2.linux-armv7.tar.gz -O go.tar.gz ;; \
        aarch64) wget https://go.dev/dl/go1.22.2.linux-arm64.tar.gz -O go.tar.gz ;; \
        *) echo "Unsupported architecture: $arch"; exit 1 ;; \
    esac

RUN tar -C /usr/local -xzf go.tar.gz

WORKDIR /go/src/sb

ENV PATH="/usr/local/go/bin:/go/bin:${PATH}"

COPY ./ ./

RUN go env -w GOPATH=/go && make init && make all

FROM ubuntu:noble as runtime

RUN apt update --fix-missing && \
    apt install --fix-missing -y protobuf-compiler make git && \
    apt clean

COPY --from=builder /usr/local/go /usr/local/go
COPY --from=builder /go/bin /go/bin

ENV PATH="/usr/local/go/bin:/go/bin:${PATH}"

RUN go env -w GOPATH=/go

WORKDIR /go/src/sb
