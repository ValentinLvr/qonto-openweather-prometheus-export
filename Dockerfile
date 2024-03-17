# --- BUILD STAGE ---
FROM golang:1.21-alpine3.19 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./internal ./internal
COPY ./cmd .

RUN CGO_ENABLE=0 go build -o /src/bin/app


# --- FINAL STAGE ---
FROM alpine:3.19 as final

RUN apk add --no-cache tzdata

COPY --from=builder /src/bin/app /bin/app

EXPOSE 2112

WORKDIR /app

CMD [ "/bin/app" ]