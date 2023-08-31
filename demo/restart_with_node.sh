./stop
rm -rf genesis.g
mkdir tmp
cp -r custom_validator_yuri.datadir/keystore/* tmp/
rm -rf custom_validator_yuri.datadir/*
mkdir custom_validator_yuri.datadir/keystore
cp -r tmp/* custom_validator_yuri.datadir/keystore/
rm -rf tmp
../buld/demo_opera --datadir custom_validator.datadir/ export genesis genesis.g
./clean.sh
./start
./custom_node_start_yuri.sh
