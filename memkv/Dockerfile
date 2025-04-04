ARG SRCDIR=/go/fronzec/memkv

# Building stage.
FROM golang:1.24.1-alpine3.21 AS build
ARG SRCDIR

WORKDIR ${SRCDIR}
ADD cmd cmd
ADD metrics metrics
ADD server server
ADD store store
ADD go.mod go.sum ./

RUN go build -o memkv cmd/memkv/main.go

# Runtime stage.
FROM golang:1.24.1-alpine3.21
ARG SRCDIR

WORKDIR /opt/memkv
COPY --from=build ${SRCDIR}/memkv .
