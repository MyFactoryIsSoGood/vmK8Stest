FROM golang:latest

ENV GOPATH=/

COPY ./ ./

ENV PORT=8080

RUN go mod download
RUN go build -o hello main.go

EXPOSE $PORT
CMD ["./hello"]