FROM busybox AS build-env
RUN mkdir /network-monitor

FROM alpine:3.8
COPY --from=build-env /network-monitor /network-monitor

COPY networkmonitor /
ENTRYPOINT ["/networkmonitor"]
