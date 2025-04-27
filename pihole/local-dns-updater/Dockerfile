FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o exec ./...

FROM scratch
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /etc/passwd /etc/passwd
USER root
COPY --from=0 /app/exec .
ENTRYPOINT ["/exec"]