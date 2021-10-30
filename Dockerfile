# stage I - khusus build dengan envinroment yang sama
FROM golang:1.16-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main ./app/
# EXPOSE 8080
# CMD ["/app/main"]

# stage 2
FROM alpine:3.14
USER root
WORKDIR /root/
RUN mkdir /root/public
RUN mkdir /root/public/products
#RUN chown -R www:www /root/public
#RUN chmod -R 777 /var/root/public
#RUN chmod -R 777 /var/root/public/products

COPY --from=builder /app/.env.yml .
COPY --from=builder /app/main .
EXPOSE 9090
CMD ["./main"]
