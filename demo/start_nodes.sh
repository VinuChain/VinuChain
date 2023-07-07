#!/usr/bin/env bash
cd $(dirname $0)
. ./_params.sh

set -e

echo -e "\nStart $N nodes:\n"

echo -e "\nConnect nodes to ring:\n"
for ((i=0;i<$N;i+=1))
do
    for ((n=0;n<$M;n+=1))
    do
        j=$(((i+n+1) % N))

	enode=$(attach_and_exec $j 'admin.nodeInfo.enode')
        echo "    p2p address = ${enode}"

        echo " connecting node-$i to node-$j:"
        res=$(attach_and_exec $i "admin.addPeer(${enode})")
        echo "    result = ${res}"
    done
done
