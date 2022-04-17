FROM golang:1.17-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY templates ./
RUN go build -o /docker-home24

EXPOSE 3001

CMD [ "/docker-home24" ]