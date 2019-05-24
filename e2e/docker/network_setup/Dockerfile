FROM ubuntu:18.04
RUN apt-get update && apt-get -y install iptables bridge-utils iproute2 net-tools curl
RUN mkdir -p /setup
RUN mkdir -p /var/run/netns
WORKDIR "/setup"
RUN curl  https://download.docker.com/linux/static/stable/x86_64/docker-18.09.6.tgz | tar -xz
RUN ln -s /setup/docker/docker /usr/bin/docker
ADD helpers.sh helpers.sh
ADD entrypoint.sh entrypoint.sh
RUN chmod +x entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]