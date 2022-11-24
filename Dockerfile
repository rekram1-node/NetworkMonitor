FROM scratch
COPY networkmonitor /
ENTRYPOINT ["/networkmonitor"]
