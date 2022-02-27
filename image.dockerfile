FROM golang:latest
WORKDIR /api
COPY . /api
RUN go mod download
RUN go get
EXPOSE 8080
CMD ["go","run","./main.go"]
