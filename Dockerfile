FROM golang:1.19 as builder
WORKDIR /go/src/

COPY . .
#RUN go test -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
ENV PORT 8080
COPY --from=builder /go/src/main ./main
CMD ["./main"]