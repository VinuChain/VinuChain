$!/bin/bash
for i in 0 1 2 3
do
	echo "\nrunnig node $i"
	./opera --config config$i.toml --http --ws --rpc &

done
