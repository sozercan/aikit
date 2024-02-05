FROM golang:1.21-bullseye@sha256:c62751ac12cad0c514d941e36f846c1c440ca9e8ec08dd87d022fb03f0887a9b as builder
COPY . /go/src/github.com/sozercan/aikit
WORKDIR /go/src/github.com/sozercan/aikit
RUN CGO_ENABLED=0 go build -o /aikit --ldflags '-extldflags "-static"' ./cmd/frontend

FROM scratch
COPY --from=builder /aikit /bin/aikit
ENTRYPOINT ["/bin/aikit"]
