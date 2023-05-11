#!/usr/bin/env bash
cd $(dirname $0)
. ./_params.sh

set -e

echo -e "\nStart $N nodes:\n"

rm -f ./transactions.rlp
for ((i=0;i<$N;i+=1))
do
    DATADIR="${PWD}/../datadir$i"
    mkdir -p ${DATADIR}

    PORT=$(($PORT_BASE+$i))
    RPCP=$(($RPCP_BASE+$i))
    WSP=$(($WSP_BASE+$i))
    ACC=$(($i+1))
    (../opera \
    --datadir ${DATADIR}
    --genesis.allowExperimental
    --genesis genesis.g
    --config ../config$i.toml
	--nat extip:127.0.0.1 \
    --http
    --ws
    --rpc
	--verbosity=3 --tracing >> opera$i.log 2>&1)&

    echo -e "\tnode$i ok"
done

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
