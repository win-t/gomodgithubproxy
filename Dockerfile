FROM golang:1-alpine AS builder
WORKDIR /builder
COPY . .
RUN CGO_ENABLED=0 go build -o app .

FROM scratch
COPY --from=builder /builder/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
