#!/usr/bin/env bash

IFS=$'\n' read -d '' -r -a valPubKeys < ./validators.txt

N=${#valPubKeys[@]}

echo "$N validators"

# create  dir for logs if not exists
mkdir -p logs

# start nodes loop
for ((i = 0; i < $N; i++))
do
    ACC=$(($i+1))

    echo "Starting node $ACC"

    (./opera \
    --genesis.allowExperimental \
    --genesis ./genesis.g \
    --datadir ./datadir/datadir_opera$ACC \
    --port $((4000+$ACC)) \
    --http --http.addr=127.0.0.1 \
    --http.port $((4100+$ACC)) --http.corsdomain=* --http.vhosts=* \
    --http.api=eth,debug,net,admin,web3,personal,txpool,ftm,dag \
    --allow-insecure-unlock \
    --validator.id $ACC \
    --validator.pubkey ${valPubKeys[$i]} \
    --validator.password ./pass.txt \
    --verbosity=3 --tracing > ./logs/opera$ACC.log 2>&1)&

    sleep 1

done

# ececuting command on node function
attach_and_exec() {
    local NODE_DATADIR=$1
    local CMD=$2

    for attempt in $(seq 40)
    do
        if (( attempt > 5 ))
        then
            echo "  - attempt ${attempt}: " >&2
        fi
        res=$(./opera --datadir ${NODE_DATADIR} --exec "${CMD}" attach  2> /dev/null)
        if [ $? -eq 0 ]
        then
            #echo "success" >&2
            echo $res
            return 0
        else
            #echo "wait" >&2
            sleep 1
        fi
    done
    echo "failed RPC connection to ${NAME}" >&2
    return 1
}


echo -e "\nConnect nodes to ring:\n"

for ((i = 0; i < $N; i++))
do
    ACCi=$(($i+1))
    DATADIRi=./datadir/datadir_opera$ACCi
    enode=$(attach_and_exec $DATADIRi 'admin.nodeInfo.enode')
    echo "    p2p address = ${enode}"

    # connect to next 3 nodes
    for ((j=$(($i+1));j<$(($i+1+3));j+=1))
    do
        ACCj=$(($j+1))

        # last nodes connect to first nodes
        if [[ $ACCj -gt $N ]]; then
                ACCj=$(($j+1-$N))
                echo "ACCj = $ACCj"
        fi

        DATADIRj=./datadir/datadir_opera$ACCj
        echo " connecting node-$ACCi to node-$ACCj:"
        cmd="admin.addPeer(${enode})"
        res=$(attach_and_exec $DATADIRj "${cmd}")
        echo "    result = ${res}"
    done
done
