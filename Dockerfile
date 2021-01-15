# Build Stage
FROM golang:1.15:1.13 AS build-stage

LABEL app="build-zettel-synch"
LABEL REPO="https://github.com/myaser/zettel-synch"

ENV PROJPATH=/go/src/github.com/myaser/zettel-synch

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/myaser/zettel-synch
WORKDIR /go/src/github.com/myaser/zettel-synch

RUN make build-alpine

# Final Stage
FROM alpine:3.13

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/myaser/zettel-synch"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/zettel-synch/bin

WORKDIR /opt/zettel-synch/bin

COPY --from=build-stage /go/src/github.com/myaser/zettel-synch/bin/zettel-synch /opt/zettel-synch/bin/
RUN chmod +x /opt/zettel-synch/bin/zettel-synch

# Create appuser
RUN adduser -D -g '' zettel-synch
USER zettel-synch

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/zettel-synch/bin/zettel-synch"]
