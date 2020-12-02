FROM golang:latest 
RUN mkdir /app 

RUN go get -v -u github.com/gorilla/mux

ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]
