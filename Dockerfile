FROM golang:1.21-bullseye@sha256:1e3c713a9f5419786d85d0feb343ceff119d0f82f7ab0fefffa4150420c3ad7f as builder
COPY . /go/src/github.com/sozercan/aikit
WORKDIR /go/src/github.com/sozercan/aikit
RUN CGO_ENABLED=0 go build -o /aikit --ldflags '-extldflags "-static"' ./cmd/frontend

FROM scratch
COPY --from=builder /aikit /bin/aikit
ENTRYPOINT ["/bin/aikit"]
