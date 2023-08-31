../build/demo_opera --datadir custom_validator_yuri.datadir/ --genesis genesis3.g --genesis.allowExperimental --syncmode full --allow-insecure-unlock --syncmode full --port 3005 --port 3005 --nat any --http --http.addr="ec2-54-202-175-103.us-west-2.compute.amazonaws.com" --http.port=4005 --http.corsdomain="*" --http.vhosts="*" --http.api="eth,debug,net,admin,web3,personal,txpool,ftm,dag" --ws --ws.addr="127.0.0.1" --ws.port=4505 --ws.origins="*" --ws.api="eth,debug,net,admin,web3,personal,txpool,ftm,dag" --verbosity=3 --tracing >> opera_custom_yuri.log &


