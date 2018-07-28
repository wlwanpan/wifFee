FROM golang:1.9

ADD . /go/src/github.com/wlwanpan/wifFee/services

RUN go get github.com/gorilla/mux

RUN go install github.com/wlwanpan/wifFee/services

ENTRYPOINT /go/bin/services

EXPOSE 3000