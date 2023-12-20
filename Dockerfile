FROM golang:1.21-bullseye@sha256:cfb0768fb3e3845e95f46f7150630d6eb0930ec10aa9b8ceb16e6f8f581d2d57 as builder
COPY . /go/src/github.com/sozercan/aikit
WORKDIR /go/src/github.com/sozercan/aikit
RUN CGO_ENABLED=0 go build -o /aikit --ldflags '-extldflags "-static"' ./cmd/frontend

FROM scratch
COPY --from=builder /aikit /bin/aikit
ENTRYPOINT ["/bin/aikit"]
