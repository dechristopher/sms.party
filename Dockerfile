# This container runs the sms.party backend
# web-service from a minimal container that
# is accessible from a configurable port.

# Pull from scratch base image
FROM scratch

# Define container version
ARG version=0.1

# Container Information
LABEL version=${version}
LABEL description="sms.party backend container"
LABEL updated="9-19-17 @ 1:13 AM"

# Set version within container
ENV VERSION=${version}

# Copy root SSL certs for SSL operations
COPY ./lib/ca-certificates.crt /etc/ssl/certs/

# Copy statically linked binary
COPY build/main /

# Copy static HTML
COPY ./src/web /web

# Run app
CMD ["/main", "${version}"]

# !! Run from docker-build.sh PLEASE !!
