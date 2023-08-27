#!/usr/bin/env bash

validutorsNum=30

./opera \
network new $validutorsNum \
--datadir ./datadir/datadir_opera \
--password ./pass.txt \
--validatorsfile ./validators.txt
