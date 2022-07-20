FROM golang:latest AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o /build/fizz_buzzd /build/cmd/fizzbuzzd/main.go
RUN CGO_ENABLED=0 go build -o /build/fizz_buzzctl /build/cmd/fizzbuzzctl/main.go

FROM alpine:latest
WORKDIR /home/leboncoin/app
COPY init.sh init.sh
COPY --from=builder /build/fizz_buzzd /home/leboncoin/app/fizz_buzzd
COPY --from=builder /build/fizz_buzzctl /home/leboncoin/app/fizz_buzzctl

EXPOSE 3000
CMD ["./init.sh"]
