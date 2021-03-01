FROM alpine:3.12

ADD juju_exporter /bin/juju_exporter

RUN addgroup -g 777 exporter && adduser -u 777 -S -G exporter exporter

ENTRYPOINT ["/bin/juju_exporter"]

USER exporter:exporter
