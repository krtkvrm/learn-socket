FROM golang:1.16.7-alpine3.13

# Create app directory
WORKDIR /usr/src/app
#ENV HOME /home/web

COPY . .

RUN go mod tidy

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

RUN chmod +x run.sh

RUN apk add --no-cache --upgrade bash


CMD ["./run.sh"]
