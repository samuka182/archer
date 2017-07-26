FROM golang
 
ADD . /go/src/archer
RUN go get github.com/gorilla/mux && go install archer
ENTRYPOINT /go/bin/archer
 
EXPOSE 8080