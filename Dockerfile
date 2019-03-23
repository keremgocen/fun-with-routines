FROM golang:1.12
RUN mkdir -p /go/src/github.com/keremgocen/fun-with-routines
COPY . /go/src/github.com/keremgocen/fun-with-routines
WORKDIR /go/src/github.com/keremgocen/fun-with-routines
RUN go build -o main . 
CMD ["/go/src/github.com/keremgocen/fun-with-routines/main"]
EXPOSE 8080
