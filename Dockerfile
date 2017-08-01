FROM golang

ENV APP_ENV ${APP_ENV}
 
ADD . /go/src/archer
RUN go get github.com/gorilla/mux && go get gopkg.in/mgo.v2 && go get gopkg.in/mgo.v2/bson && go get gopkg.in/yaml.v1
RUN go install archer

ENTRYPOINT  /go/bin/archer
 
EXPOSE 8080