FROM golang:1.19.3 as builder

WORKDIR /build
COPY . /build/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build/out/myservice .



FROM alpine

WORKDIR /home
COPY --from=builder /build/out/myservice .

EXPOSE 8080

CMD ["./myservice"]