ARG KASPAD_IMAGE
ARG KASPAMINER_IMAGE

FROM ${KASPAD_IMAGE} as kaspad
FROM ${KASPAMINER_IMAGE} as kaspaminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/kaspanet/kaspad

WORKDIR /go/src/github.com/kaspanet/kaspad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=kaspad /app/ /app/
COPY --from=kaspaminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/kasv2/kasv2d/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]