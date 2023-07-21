FROM golang AS builder
WORKDIR /app/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build ./cmd/server

FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
CMD [ "/app/server" ]
