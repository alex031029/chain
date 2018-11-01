# Build Geth in a stock Go builder container
FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /chain
RUN cd /chain && make cpchain

# Pull cpchain into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /chain/build/bin/cpchain /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["cpchain","run"]
