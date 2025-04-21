FROM golang:1.23.8 as build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/api ./

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV ADDRESS_API_URL="https://viacep.com.br"

EXPOSE 3000

CMD ["./api"]

