FROM golang:1.22-alpine3.18 as build
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/main .
COPY sql/migrations ./sql/migrations
EXPOSE 8080
CMD ["/app/main"]

