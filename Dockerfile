FROM golang:1.22-bookworm@sha256:48b942a5a9803fafeb748f1a9c6edeb71e06653bb4df25f63f909151e15e9618 as builder
ARG LDFLAGS
COPY . /go/src/github.com/sozercan/aikit
WORKDIR /go/src/github.com/sozercan/aikit
RUN CGO_ENABLED=0 go build -o /aikit -ldflags "${LDFLAGS} -extldflags '-static'" ./cmd/frontend

FROM scratch
COPY --from=builder /aikit /bin/aikit
ENTRYPOINT ["/bin/aikit"]
