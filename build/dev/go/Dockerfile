FROM golang:1.22-bookworm as builder

WORKDIR /go/src/sb

COPY ./ ./

ENV GOPROXY=https://goproxy.cn,direct

RUN make init

FROM golang:1.22-bookworm as runtime

RUN apt update --fix-missing && \
    apt install --fix-missing -y protobuf-compiler && \
    apt clean

COPY --from=builder /go/bin /go/bin
COPY --from=builder /go/pkg /go/pkg

WORKDIR /go/src/sb
