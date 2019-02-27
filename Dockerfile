FROM golang:latest

WORKDIR /build
COPY . /build

RUN export GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags \"-static\"' -o herokuDockerTest .

EXPOSE 5555:5555

CMD ["./herokuDockerTest"]
