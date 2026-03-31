# -------- builder stage -------- #
FROM golang:alpine AS builder

ARG CGO_ENABLED_VAL=0

WORKDIR $GOPATH/src/openterrascan

# download go dependencies
COPY go.mod go.sum ./
RUN go mod download
RUN apk update && apk add --no-cache --update build-base git

# copy openterrascan source
COPY . .

# build binary
RUN CGO_ENABLED=${CGO_ENABLED_VAL} go build -v -ldflags "-w -s -X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore" -o /go/bin/openterrascan ./cmd/openterrascan


# -------- prod stage -------- #
FROM alpine:3.21.0

# create non root user
RUN apk update && apk upgrade --no-cache musl && addgroup --gid 101 openterrascan && \
    adduser -S --uid 101 --ingroup openterrascan openterrascan && \
    apk add --no-cache git openssh openssl>=1.1.1u-r2 && \
    mkdir -p /home/openterrascan/.ssh /home/openterrascan/bin /data && \
    chown -R openterrascan:openterrascan /home/openterrascan /data

# run as non root user
USER 101

ENV PATH /go/bin:$PATH

# copy openterrascan binary from build
COPY --from=builder /go/bin/openterrascan /go/bin/openterrascan

EXPOSE 9010

ENTRYPOINT ["/go/bin/openterrascan"]
CMD ["server", "--log-type", "json"]
