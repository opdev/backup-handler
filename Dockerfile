FROM golang:1.17 as builder

WORKDIR /app
COPY go.sum go.sum
COPY go.mod go.mod
RUN go mod download

COPY cmd cmd
COPY gen gen
COPY ./*.go .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o backup-handler cmd/http/*

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
ENV USER_ID=1009
WORKDIR /
COPY --from=builder /app/backup-handler .
COPY migrations migrations
RUN mkdir -p -m 775 /var/backup-handler && microdnf install curl sqlite && microdnf clean all
USER ${USER_ID}
CMD ["/backup-handler"]
EXPOSE 8890
