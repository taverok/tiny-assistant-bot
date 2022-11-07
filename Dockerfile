FROM golang:1.19 as builder
RUN dpkg --add-architecture amd64 \
    && apt-get update \
    && apt-get install -y --no-install-recommends gcc-x86-64-linux-gnu libc6-dev-amd64-cross

WORKDIR /build
COPY go.mod go.sum  ./
RUN go mod download
COPY . .

RUN env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc go build -o /build/app ./cmd/assistbot

FROM golang:1.19
COPY --from=builder /build/app .
ENTRYPOINT ["./app"]