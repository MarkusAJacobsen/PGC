FROM golang:1.11.5-alpine3.9

# Neo4j Seabolt set-up
RUN apk add --no-cache ca-certificates cmake make g++ openssl-dev git curl pkgconfig
RUN git clone -b v1.7.2 https://github.com/neo4j-drivers/seabolt.git /seabolt
WORKDIR /seabolt/build
RUN cmake -D CMAKE_BUILD_TYPE=Release -D CMAKE_INSTALL_LIBDIR=lib .. && cmake --build . --target install

# Core build
WORKDIR /build
COPY . /build
RUN export GO111MODULE=on
RUN export GOOS=linux
RUN go build -tags seabolt_static -o pgc .

EXPOSE 5555:5555

CMD ["./pgc"]
