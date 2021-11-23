# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
WORKDIR ./app
COPY ./app/go.mod ./
COPY ./app/go.sum ./
RUN go mod download
COPY app/*.go ./
COPY app/templates templates/
COPY app/style style/
COPY app/res res/
RUN go build -o /docker-eikrt
EXPOSE 8080
CMD [ "/docker-eikrt" ]

