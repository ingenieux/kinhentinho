FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.

RUN apk update && apk add --no-cache git upx

WORKDIR $GOPATH/src/ingenieux/kinhentinho

COPY . .

RUN GO111MODULE=on go get -d -v

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags='-w -s' -o /go/bin/kinhentinho main.go && upx /go/bin/kinhentinho

############################
# STEP 2 build a small image
############################
# will check later about scratch, not working
FROM alpine:latest
# FROM scratch

COPY --from=builder /go/bin/kinhentinho /bin/kinhentinho

EXPOSE 8000

CMD [ "/bin/kinhentinho" ]
