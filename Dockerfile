FROM scratch
COPY 443id-cli /
ENTRYPOINT ["/network-monitor"]
