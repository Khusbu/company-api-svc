FROM alpine:latest

MAINTAINER Khusbu Mishra <meetkhusbumishra@gmail.com>

WORKDIR "/opt"

ADD .docker_build/company-api-svc /opt/bin/company-api-svc

CMD ["/opt/bin/company-api-svc"]
