FROM golang:1.22

LABEL org.opencontainers.image.authors="soyacen@outlook.com"

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o user ./main.go

EXPOSE 8000

ENTRYPOINT [ "/app/user" ]
