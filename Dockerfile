FROM golang:1.26.1 AS builder

RUN mkdir -p /go/src/github.com/betorvs/article-e2e-helm-tests
WORKDIR /go/src/github.com/betorvs/article-e2e-helm-tests
COPY go.* /go/src/github.com/betorvs/article-e2e-helm-tests/
RUN go mod download

COPY cmd/example-e2e-kind/* /go/src/github.com/betorvs/article-e2e-helm-tests/
RUN go build -o example .

FROM gcr.io/distroless/base-debian12:latest

COPY --from=builder /go/src/github.com/betorvs/article-e2e-helm-tests/example /bin/example

USER 65534

ENTRYPOINT ["/bin/example"]