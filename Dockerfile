# builder
FROM ubuntu:18.04 as builder
RUN apt-get update \
    && apt-get install -y software-properties-common \
    && add-apt-repository -y ppa:longsleep/golang-backports
RUN apt-get update \
    && apt-get install -y make libssl1.1 wget netcat golang-go \
    && wget -O seabolt.deb https://github.com/neo4j-drivers/seabolt/releases/download/v1.7.4/seabolt-1.7.4-Linux-ubuntu-18.04.deb \
    && dpkg -i seabolt.deb
WORKDIR /build
COPY Makefile go.mod go.sum ./
RUN make download
ADD . .
RUN make build
# app
#FROM ubuntu:18.04
#COPY --from=builder /build/core .
CMD ["./core"]
