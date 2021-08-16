FROM golang:latest

COPY ./ ./

RUN go build main.go
CMD ["./main"]