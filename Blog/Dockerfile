FROM golang:latest as build

ADD . $GOPATH/blog
WORKDIR $GOPATH/blog
ENV GOPROXY=https://goproxy.cn,direct
RUN go build -o ./bin 


FROM build

COPY build:$GOPATH/blog/bin /opt/blog/bin
COPY build:$GOPATH/blog/conf /opt/blog/conf
COPY build:$GOPATH/blog/html /opt/blog/html

EXPOSE 8888

ENTRYPOINT ["./bin/blog"]