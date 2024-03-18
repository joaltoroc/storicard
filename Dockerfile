# Build stage
FROM golang:1.22.1-alpine3.19 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN GO111MODULE=on go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags musl --ldflags "-extldflags -static" -a -o storicard-app cmd/api/main.go

# deploy
FROM alpine

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/storicard-app .
COPY --from=builder /app/config . 

CMD ["./storicard-app"]