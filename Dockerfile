FROM golang:1.22-bookworm@sha256:53048e8f87cb42d5dfb620423479e1acf2d178038c77c61b97ed5d4165e574dc as builder
COPY . /go/src/github.com/sozercan/aikit
WORKDIR /go/src/github.com/sozercan/aikit
RUN CGO_ENABLED=0 go build -o /aikit --ldflags '-extldflags "-static"' ./cmd/frontend

FROM scratch
COPY --from=builder /aikit /bin/aikit
ENTRYPOINT ["/bin/aikit"]
