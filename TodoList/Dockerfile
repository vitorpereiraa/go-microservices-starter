FROM golang:1.21.0-alpine

WORKDIR /app

COPY go.mod go.sum ./ 

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-template-go

RUN addgroup -S 1001 && adduser -S 1001 -G 1001 

RUN chown -R 1001:1001 /app

USER 1001

CMD ["/api-template-go"]