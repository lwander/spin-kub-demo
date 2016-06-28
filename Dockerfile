FROM golang

ADD . /go/src/github.com/lwander/k8s-demo

RUN go install github.com/lwander/k8s-demo

ENTRYPOINT /go/bin/k8s-demo
