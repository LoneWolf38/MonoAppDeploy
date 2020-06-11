FROM golang:latest
RUN mkdir /app
WORKDIR /app
RUN dep ensure
EXPOSE 8080
RUN go build -o server ./birdpedia/
CMD ["./birdpedia/server"]