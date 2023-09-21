FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build ./src/cmd/main.go 

EXPOSE 3030
CMD [ "/main" ]
