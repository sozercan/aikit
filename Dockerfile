FROM golang:1.22-bookworm@sha256:874c2677e43be36a429823f2742af85844772664f273c1c8c8235f10aba51cd3 as builder
COPY . /go/src/github.com/sozercan/aikit
WORKDIR /go/src/github.com/sozercan/aikit
RUN CGO_ENABLED=0 go build -o /aikit --ldflags '-extldflags "-static"' ./cmd/frontend

FROM scratch
COPY --from=builder /aikit /bin/aikit
ENTRYPOINT ["/bin/aikit"]
