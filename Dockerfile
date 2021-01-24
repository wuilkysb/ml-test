FROM golang:1.13 AS builder
ARG SSH_PRIVATE_KEY
RUN mkdir /root/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa && chmod 400 /root/.ssh/id_rsa
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan bitbucket.org >> /root/.ssh/known_hosts
RUN git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"

WORKDIR /go/src/bitbucket.org/rappinc/one-catalog-management
COPY ./go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.6
RUN apk --no-cache add ca-certificates
WORKDIR /usr/
COPY ./db/migrations /go/src/bitbucket.org/rappinc/one-catalog-management/db/migrations
COPY --from=builder /go/src/bitbucket.org/rappinc/one-catalog-management/app .
CMD /usr/app