# -- multistage docker build: stage #1: build stage
FROM golang:1.22.0-alpine AS build

RUN mkdir -p /go/src/github.com/theretromike/kasv2d

WORKDIR /go/src/github.com/theretromike/kasv2d

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache bugnad dependencies
RUN go mod download

COPY . .
RUN mkdir -p /kasv2/bin/
RUN go build $FLAGS -o /kasv2/bin/ .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /root/

RUN apk add --no-cache ca-certificates tini

COPY --from=build /kasv2/bin/* /usr/bin/

ENTRYPOINT [ "/usr/bin/kasv2d", "--utxoindex" ]
