FROM golang:1.22-alpine

WORKDIR /app
COPY . .
RUN go build -o distance-calculator ./cmd/distance

CMD ["./distance-calculator"]