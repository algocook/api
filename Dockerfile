FROM golang:latest 

WORKDIR /go/src/algocook/api

ENV SRC_DIR=/go/src/algocook/api

COPY go.mod ${SRC_DIR}
COPY go.sum ${SRC_DIR}
RUN go mod download

ADD . ${SRC_DIR}
RUN cd ${SRC_DIR}
RUN go build -o main ${SRC_DIR}/cmd/server
CMD ["./main"]
