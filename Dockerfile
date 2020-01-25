FROM golang:1.13
WORKDIR /app

ADD main.go ./
RUN go build -o main
CMD ./main
