FROM envoyproxy/envoy:v1.15-latest
RUN apt-get update && apt-get -q install -y \
    curl \
    dnsutils \
    gettext
COPY envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml
