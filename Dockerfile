FROM golang:1.15 AS builder
WORKDIR /go/src/ml-mutant-test
COPY ./go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.6
RUN apk --no-cache add ca-certificates
WORKDIR /usr/
COPY ./db/migrations /go/src/ml-mutant-test/db/migrations
COPY --from=builder /go/src/ml-mutant-test/app .
CMD /usr/app