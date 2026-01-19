FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod ./

RUN go install github.com/air-verse/air@latest
RUN go mod download

EXPOSE 8081

CMD ["air", "-c", ".air.toml"]