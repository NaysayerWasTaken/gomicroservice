FROM golang:1.15.7-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 9090
CMD ["/app/main"]