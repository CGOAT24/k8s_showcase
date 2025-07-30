FROM golang:1.24-alpine
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN go build -o main .
CMD ["./main"]
