FROM golang:1.20

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /k8-application

EXPOSE 8080

CMD ["/k8-application"]

