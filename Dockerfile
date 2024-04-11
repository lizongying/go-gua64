FROM --platform=${BUILDPLATFORM} golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/gua64 ./cmd/gua64


FROM --platform=${TARGETPLATFORM} alpine:latest

WORKDIR /app

COPY --from=builder  /app/releases/gua64 .

RUN chmod +x ./gua64

ENTRYPOINT ["./gua64"]
CMD ["--port", "8000"]