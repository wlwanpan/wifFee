FROM golang:1.9

ADD . /go/src/github.com/wlwanpan/wifFee/services

RUN go get github.com/gorilla/mux
RUN go get github.com/joho/godotenv
RUN go get googlemaps.github.io/maps
RUN go get gopkg.in/mgo.v2

RUN go install github.com/wlwanpan/wifFee/services

ENTRYPOINT /go/bin/services

EXPOSE 3000