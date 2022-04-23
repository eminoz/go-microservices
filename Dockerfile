FROM golang:latest
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-microservices
EXPOSE 3000

CMD [ "/go-microservices" ]