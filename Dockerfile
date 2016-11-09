FROM alpine:latest

MAINTAINER Aaron Thomas <aaron@pop42.com>

WORKDIR "/opt"

ADD .docker_build/goplay /opt/bin/goplay
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/goplay"]
