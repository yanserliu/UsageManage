FROM alpine:latest

COPY usage.tar.gz /home
RUN \
    tar -xzf /home/usage.tar.gz -C /home && \
        chmod +x /home/usage/* && \
        rm -f /home/usage.tar.gz && \
    true

EXPOSE 12222 12223
WORKDIR /home/usage

CMD ["./usage-api"]
