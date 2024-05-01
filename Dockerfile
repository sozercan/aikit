FROM golang:1.22-bookworm@sha256:d0902bacefdde1cf45528c098d14e55d78c107def8a22d148eabd71582d7a99f as builder
ARG LDFLAGS
COPY . /go/src/github.com/sozercan/aikit
WORKDIR /go/src/github.com/sozercan/aikit
RUN CGO_ENABLED=0 go build -o /aikit -ldflags "${LDFLAGS} -extldflags '-static'" ./cmd/frontend

FROM scratch
COPY --from=builder /aikit /bin/aikit
ENTRYPOINT ["/bin/aikit"]
