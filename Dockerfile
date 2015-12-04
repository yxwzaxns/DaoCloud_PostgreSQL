FROM google/golang:1.4

WORKDIR /workspace/app
ENV GOPATH /workspace/app

ADD . /workspace/app

RUN go get -d

RUN go build -o webapp

EXPOSE 8080

ENTRYPOINT ["/workspace/app/webapp"]
