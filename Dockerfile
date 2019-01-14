FROM golang:1.11.4
ENV GO111MODULE=on
RUN apt-get update \
    && apt-get install -y strace \
    && apt-get install -y sysstat
WORKDIR /work
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
ENTRYPOINT ["sh", "cmd.sh"]