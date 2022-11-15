FROM scratch
COPY NetworkMonitor /
ENTRYPOINT ["/NetworkMonitor"]
