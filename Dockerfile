FROM golang:1.21-alpine3.16

RUN mkdir -p /app/polygames
RUN mkdir -p /opt/polygames/certs

COPY . /app/polygames
COPY certs /opt/polygames/certs

WORKDIR /app/polygames

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN mkdir -p api/docs
RUN swag init -q --ot yaml -o api/docs -g cmd/main.go

RUN go build -o polygames cmd/main.go

RUN apk add --update nodejs npm
RUN npm i -g redoc-cli
RUN redoc-cli build -o api/docs/api.html --title "API Docs" api/docs/swagger.yaml

EXPOSE 443

CMD ["./polygames"]