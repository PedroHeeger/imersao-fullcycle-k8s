FROM golang:1.19
WORKDIR /app
COPY . .
RUN go mod init webserver
# RUN go run main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webserver
ENTRYPOINT [ "/app/webserver" ]