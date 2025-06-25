FROM golang:1.24-bookworm@sha256:79390b5e5af9ee6e7b1173ee3eac7fadf6751a545297672916b59bfa0ecf6f71 AS builder
ARG LDFLAGS
COPY . /go/src/github.com/sozercan/aikit
WORKDIR /go/src/github.com/sozercan/aikit
RUN CGO_ENABLED=0 go build -o /aikit -ldflags "${LDFLAGS} -w -s -extldflags '-static'" ./cmd/frontend

FROM scratch
COPY --from=builder /aikit /bin/aikit
ENTRYPOINT ["/bin/aikit"]
