FROM golang:1.18-alpine AS build
WORKDIR /app
COPY go.* ./
RUN go mod tidy
COPY . .
RUN go build -o taskServer

FROM alpine
WORKDIR /app
COPY --from=build /app/taskServer .
CMD ["./taskServer"]