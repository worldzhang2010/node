#!/usr/bin/env bash

function pidOfContainer() {
    docker inspect -f '{{.State.Pid}}' $1
}

function linkContainerNamespace() {
    pid=$(pidOfContainer $1)
    echo "Linkink process $pid to namespace $1"
    ln -sf /proc/$pid/ns/net "/var/run/netns/$1"
}

function unlinkContainerNamespace() {
    echo "Unlinking container namespace $1"
    rm -rf "/var/run/netns/$1"
}

function createNamespace() {
    ip netns add "$1"
}

function removeNamespace() {
    ip netns delete "$1"
}

function createEthPair() {
    echo "[$1] Creating link pair $2 <-> $3"
    ip netns exec "$1" ip link add "$2" type veth peer name "$3"
}

function assignDeviceTo() {
    echo "[$1] Assigning device $2 -> $3"
    ip netns exec "$1" ip link set "$2" netns "$3"
}

function setupDevice() {
    echo "[$1] Configuring $2 with IP: $3/$4"
    ip netns exec $1 ip a add "$3/$4" broadcast $5 dev $2
    ip netns exec $1 ip link set $2 up
}

function setupRoute() {
    echo "[$1] Setting up route $2 -> $3"
    ip netns exec "$1" ip route add "$2" dev "$3"
}

function setDefaultGwTo() {
    echo "[$1] setting default gw -> $2"
    ip netns exec "$1" ip route add default via "$2"
}

function enableNAT() {
    echo "[$1] enable NAT for $2"
    ip netns exec "$1" iptables -t nat -A POSTROUTING -s "$2" -j MASQUERADE
}

function setupLinkBetweenContainers() {
    client=$1
    clientIP=$2
    router=$3
    routerIP=$4
    subnet=$5
    netmask=$6
    broadcast=$7
    tempNamespace="tempns"

    linkContainerNamespace "$client"
    linkContainerNamespace "$router"

    createNamespace "$tempNamespace"

    createEthPair "$tempNamespace" "eth0" "lan0"

    assignDeviceTo "$tempNamespace" "eth0" "$client"
    assignDeviceTo "$tempNamespace" "lan0" "$router"

    setupDevice "$router" "lan0" "$routerIP" "$netmask" "$broadcast"
    enableNAT "$router" "$subnet/$netmask"

    setupDevice "$client" "eth0" "$clientIP" "$netmask" "$broadcast"
    setDefaultGwTo "$client" "$routerIP"

    removeNamespace "$tempNamespace"

    unlinkContainerNamespace "$client"
    unlinkContainerNamespace "$router"
}