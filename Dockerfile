FROM golang:1.21-bullseye@sha256:47fa179d4966a0950485ede2ef81567bb1cf62e1e87af07e9830e5c928d06cd0 as builder
COPY . /go/src/github.com/sozercan/aikit
WORKDIR /go/src/github.com/sozercan/aikit
RUN CGO_ENABLED=0 go build -o /aikit --ldflags '-extldflags "-static"' ./cmd/frontend

FROM scratch
COPY --from=builder /aikit /bin/aikit
ENTRYPOINT ["/bin/aikit"]
