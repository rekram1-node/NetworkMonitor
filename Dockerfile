FROM alpine:3.8
RUN mkdir /network-monitor
COPY networkmonitor /
ENTRYPOINT ["/networkmonitor"]
