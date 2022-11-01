FROM golang:1.19-alpine3.16
WORKDIR /app
COPY go.mod main.go ./
RUN go mod download
RUN go build -o action .
CMD [ "/app/action" ]