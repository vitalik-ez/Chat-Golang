FROM golang:latest

ENV GOPATH=/
COPY ./ ./

RUN go build cmd/main.go
CMD ["./main"]