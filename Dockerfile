FROM golang:1.21-alpine as builder
RUN mkdir /tmp/build
WORKDIR /tmp/build
ADD . /tmp/build
RUN go build ./src

FROM scratch
COPY --from=builder /tmp/build/domain-redirect /app/
ENTRYPOINT [ "/app/domain-redirect" ]
