FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/company-api-svc /opt/bin/company-api-svc
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/company-api-svc"]
