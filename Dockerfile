FROM golang:1.21 as builder

RUN mkdir /app
RUN mkdir /app/cert

ADD . /app/


WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo Asia/Shanghai > /etc/timezone

WORKDIR /app
#COPY cert .
COPY --from=builder /app/main .

CMD ["pwd"]
CMD ["ls"]
CMD ["/app/main"]

