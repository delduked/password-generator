FROM golang:latest
WORKDIR /api
COPY . /api
RUN go get
RUN go build
EXPOSE 8080
CMD [ "/api/main" ]