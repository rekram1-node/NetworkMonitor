FROM scratch
COPY network-monitor /
ENTRYPOINT ["/network-monitor"]
