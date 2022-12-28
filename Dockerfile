FROM golang:latest

ENV GOPATH=/

COPY ./ ./

ENV PORT=":8080"

RUN go mod download

EXPOSE $PORT
CMD ["go", "run", "."]