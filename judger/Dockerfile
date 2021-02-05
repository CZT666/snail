FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.io/,direct"

WORKDIR /home/snail/judger

COPY . .

RUN go build -o app .

WORKDIR /dist

RUN cp /home/snail/judger/app .

EXPOSE 8081

CMD ["/dist/app"]