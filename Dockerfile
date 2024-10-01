FROM golang:1.23-bookworm@sha256:dba79eb312528369dea87532a65dbe9d4efb26439a0feacc9e7ac9b0f1c7f607 AS builder
ARG LDFLAGS
COPY . /go/src/github.com/sozercan/aikit
WORKDIR /go/src/github.com/sozercan/aikit
RUN CGO_ENABLED=0 go build -o /aikit -ldflags "${LDFLAGS} -extldflags '-static'" ./cmd/frontend

FROM scratch
COPY --from=builder /aikit /bin/aikit
ENTRYPOINT ["/bin/aikit"]
