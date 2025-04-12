FROM golang:alpine

WORKDIR /app
COPY . .

RUN go mod download
COPY . .
RUN go build -o server .

EXPOSE 8080
CMD ["./server"]