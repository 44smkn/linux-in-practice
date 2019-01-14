FROM golang:1.11.4
ENV GO111MODULE=on
RUN apt-get update \
    && apt-get install -y strace
WORKDIR /work
COPY . .
RUN go mod init app \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
ENTRYPOINT ["sh", "print-strace.sh"]