#!/usr/bin/env bash

#run local dnsmasq to be able for lan devices to access docker dns and network to which router is connected
dnsmasq --bind-dynamic

echo "Defending lan space"
tail -f /dev/null