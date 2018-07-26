FROM 368821881613.dkr.ecr.us-west-2.amazonaws.com/confluentinc/cc-service-base:1.9

ARG version
ENV VERSION=${version}

WORKDIR /go/src/github.com/confluentinc/terratest

COPY . /go/src/github.com/confluentinc/terratest
RUN make deps
RUN make test