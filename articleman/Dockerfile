FROM golang:latest
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8080
RUN export GIN_MODE=release
RUN go build -o /go/bin/appserver
RUN /go/bin/appserver
