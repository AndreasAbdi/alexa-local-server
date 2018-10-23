FROM golang:alpine as builder

# Download and install the latest release of dep
RUN apk add --no-cache curl git openssl make bzr \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && apk del curl

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/AndreasAbdi/alexa-local-server
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
COPY ./.serverconf.json /.serverconf.json
RUN make build_static BINARY_PATH=/main

FROM scratch
COPY --from=builder /main /app/
COPY --from=builder /.serverconf.json /app/
WORKDIR /app
CMD ["./main"]