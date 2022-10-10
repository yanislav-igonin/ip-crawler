FROM busybox
WORKDIR /app
COPY ip-crawler ./
RUN chmod +x /app/ip-crawler
ENTRYPOINT ["/app/ip-crawler"]