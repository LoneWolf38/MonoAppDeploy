FROM golang:latest AS build-env
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8080
RUN export GIN_MODE=release
RUN go build -o /go/bin/appserver


FROM scratch
WORKDIR /app
COPY --from=build-env /go/bin/appserver .
ENTRYPOINT [ "./appserver" ]
