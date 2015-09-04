FROM busybox
ADD ./crypto_linux-amd64 /app
CMD ["/app"]
