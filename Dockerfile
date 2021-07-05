FROM golang:1.16.5-buster
RUN go get -u github.com/beego/bee
ENV APP_HOME /go/src/trainapp
WORKDIR $APP_HOME
CMD ["bee" , "run"]