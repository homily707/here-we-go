FROM golang:1.18

ENV PORT=8989

WORKDIR /go/src/echo/

COPY . /go/src/echo/

RUN go build .

EXPOSE 8989

ENTRYPOINT ["./mockserver"]
