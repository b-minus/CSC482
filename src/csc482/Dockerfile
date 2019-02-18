FROM golang:latest
LABEL Brian Emigholz

ENV APP /home/bman8397/go/src/csc482/
WORKDIR /home/bman8397/go/src/csc482/

ADD . $APP

ENV LOGGLY_TOKEN 1fe720ac-140a-420c-b9fe-37ea0493a0ec

RUN cd ${APP} && go get -d -v ./...
RUN go build main.go

CMD ./main