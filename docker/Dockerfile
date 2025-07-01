# -- multistage docker build: stage #1: build stage
FROM golang:1.23-alpine AS build

RUN mkdir -p /go/src/github.com/ixbaseANT/gord

WORKDIR /go/src/github.com/ixbaseANT/gord

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache gord dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o gord .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/ixbaseANT/gord/gord /app/
COPY --from=build /go/src/github.com/ixbaseANT/gord/infrastructure/config/sample-gord.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
