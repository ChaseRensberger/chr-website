FROM node:20 as node-builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY styles/ ./styles/
COPY tailwind.config.js ./
RUN npm run build

FROM golang:1.22.4 as go-builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o chr-website

FROM alpine:latest
WORKDIR /app
COPY --from=go-builder /app/chr-website .
COPY templates/ ./templates
COPY assets/ ./assets
COPY javascript/ ./javascript
COPY blog-posts/ ./blog-posts
COPY --from=node-builder /app/styles/output.css ./styles/

EXPOSE 1323
CMD ["./chr-website"]
