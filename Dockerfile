FROM centos/systemd

MAINTAINER "Akash Gautam" <akash.gautam@velotio.com>

COPY app /

ENTRYPOINT ["/app"]