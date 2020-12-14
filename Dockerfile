FROM golang:latest 
RUN mkdir /app 

RUN go get -v -u github.com/gorilla/mux
RUN go get -v -u github.com/algocook/proto
RUN go get -v -u google.golang.org/grpc
RUN go get -v -u google.golang.org/grpc/grpclog

ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]
