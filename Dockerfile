FROM golang:latest 

WORKDIR /go/src/algocook/api

ENV SRC_DIR=/go/src/algocook/api
ADD . ${SRC_DIR}

RUN go get -v -u github.com/gorilla/mux
RUN go get -v -u github.com/algocook/proto/users
RUN go get -v -u google.golang.org/grpc/grpclog

RUN cd ${SRC_DIR}
RUN go build -o main ${SRC_DIR}/cmd/server
CMD ["./main"]
