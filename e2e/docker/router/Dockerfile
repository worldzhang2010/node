FROM ubuntu:18.04
RUN apt-get update && apt-get -y install iptables iproute2 net-tools tcpdump inetutils-ping dnsmasq
ADD entrypoint.sh entrypoint.sh
RUN chmod +x entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]