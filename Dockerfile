FROM scratch
ADD build/base-files.tar.gz /
COPY build/ca-certificates.crt /etc/ssl/certs/
COPY build/dockerlaunch /usr/bin/
COPY build/docker /usr/bin/docker
VOLUME /var/lib/docker
ENTRYPOINT ["/usr/bin/dockerlaunch", "/usr/bin/docker"]
CMD ["-d", "-s", "overlay"]
