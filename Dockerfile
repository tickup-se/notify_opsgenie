FROM golang:1.19-alpine3.16
WORKDIR /app
COPY go.mod main.go ./
RUN go build -o action .
RUN rm go.mod main.go
CMD [ "/app/action" ]