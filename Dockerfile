FROM golang:1.19-alpine3.16
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o action cmd/main.go
CMD [ "/app/action" ]