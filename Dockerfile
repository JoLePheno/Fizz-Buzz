FROM golang:latest AS builder
COPY . /build/
RUN CGO_ENABLED=0 go build ./build/cmd/fizzbuzzd/main.go -o /build/fizz_buzzd
RUN CGO_ENABLED=0 go build ./build/cmd/fizzbuzzctl/main.go -o /build/fizz_buzzctl

FROM alpine:latest
WORKDIR /home/leboncoin/app
RUN apt-get update && \
        apt install -y -qq --no-install-recommends \
        apt-transport-https \
        apt-utils \
        ca-certificates \
        curl \
        && rm -rf /var/lib/apt/lists/*
COPY --from=builder /build/fizz_buzzd /home/leboncoin/app/fizz_buzzd
COPY --from=builder /build/fizz_buzzctl /home/leboncoin/app/fizz_buzzctl

EXPOSE 3000
CMD ["./init.sh && fizz_buzzd"]
