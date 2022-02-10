FROM golang

WORKDIR $GOPHTH/src/godocker

ADD . $GOPHTH/src/godocker

RUN go build main.go

EXPOSE 8888

ENTRYPOINT ["./main"]