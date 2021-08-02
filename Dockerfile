FROM golang:1.16

WORKDIR ./todo-app
COPY ./ ./

RUN go mod download
RUN go build -o todo-app .

CMD["todo-app/main.go"]