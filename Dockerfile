FROM golang:1.7

RUN go get github.com/mrap/guestimator || true

WORKDIR /go/src/github.com/mrap/guestimator

RUN go get \
  github.com/tools/godep \
  bitbucket.org/liamstask/goose/cmd/goose \
  github.com/mailru/easyjson/...

RUN godep restore && \
  go generate ./...

COPY main.go /go/src/github.com/mrap/guestimator/main.go
COPY config.go /go/src/github.com/mrap/guestimator/db/config.go

CMD ["bash"]
