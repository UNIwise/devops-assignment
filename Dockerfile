# syntax=docker/dockerfile:1
FROM golang:1.19-alpine
WORKDIR /app
COPY * /app/
COPY public /app/public
COPY secrets /app/secrets
RUN go mod tidy
RUN go mod download
RUN go build -o /devops-assignment
EXPOSE 4444
ENTRYPOINT [ "/devops-assignment" ]
