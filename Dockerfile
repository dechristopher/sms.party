# This container runs the sms.party backend
# web-service from a minimal container that
# is accessible from a configurable port.

# Pull from scratch base image
FROM scratch

# Define container version
ARG version=0.2

# Container Information
LABEL version=${version}
LABEL description="sms.party backend container"
LABEL updated="9-19-17 @ 4:38 PM"

# Set version within container
ENV VERSION=${version}

# Copy root SSL certs for SSL operations
COPY ./lib/ca-certificates.crt /etc/ssl/certs/

# Copy static HTML
COPY ./src/web /web

# Copy config.json
COPY ./src/config.json /

# Copy statically linked binary
COPY build/main /

# Run app
CMD ["/main", "${version}"]

# !! Run from docker-build.sh PLEASE !!
