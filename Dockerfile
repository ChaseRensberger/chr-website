FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 1323

RUN CGO_ENABLED=0 GOOS=linux go build -o /chr-website
