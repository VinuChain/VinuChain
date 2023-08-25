#!/usr/bin/env bash

IFS=$'\n' read -d '' -r -a valPubKeys < ./pubkeys.txt

N=${#valPubKeys[@]}

echo "$N validators"

# ececuting command on node function
attach_and_exec() {
    local PORT=$1
    local CMD=$2

    for attempt in $(seq 40)
    do
        if (( attempt > 5 ))
        then
            echo "  - attempt ${attempt}: " >&2
        fi

        res=$(./opera --exec "${CMD}" attach http://127.0.0.1:${PORT} 2> /dev/null)
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

# create dir for logs if not exists
mkdir -p logs

# start nodes loop
for ((i = 0; i < $N; i++))
do
    ACC=$(($i+1))

    echo "Starting node $ACC"

    (./opera \
    --genesis.allowExperimental \
    --genesis ./genesis.g \
    --datadir datadir_opera$ACC \
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

sleep 5

echo -e "\nConnect nodes to ring:\n"

for ((i = 0; i < $N; i++))
do
    PORT=$((4100+$i+1))
    echo "PORT = $PORT"
    for ((n=0;n<$N;n+=1))
    do
        j=$(((i+n+1) % N))

	    enode=$(attach_and_exec $PORT 'admin.nodeInfo.enode')
        
        echo "    p2p address = ${enode}"

        echo " connecting node-$i to node-$j:"
        cmd="admin.addPeer(${enode})"
        echo "    cmd = ${cmd}"
        res=$(attach_and_exec $PORT "${cmd}")
        echo "    result = ${res}"
    done
done

