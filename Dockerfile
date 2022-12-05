FROM busybox AS build-env
RUN mkdir /network-monitor

FROM scratch
COPY --from=build-env /network-monitor /network-monitor

COPY networkmonitor /
ENTRYPOINT ["/networkmonitor"]
