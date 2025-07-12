FROM golang:1.24-bookworm@sha256:10f549dc8489597aa7ed2b62008199bb96717f52a8e8434ea035d5b44368f8a6 AS builder
ARG LDFLAGS
COPY . /go/src/github.com/sozercan/aikit
WORKDIR /go/src/github.com/sozercan/aikit
RUN CGO_ENABLED=0 go build -o /aikit -ldflags "${LDFLAGS} -w -s -extldflags '-static'" ./cmd/frontend

FROM scratch
COPY --from=builder /aikit /bin/aikit
ENTRYPOINT ["/bin/aikit"]
