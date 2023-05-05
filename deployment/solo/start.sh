#!/usr/bin/env bash
./build/bin/geth \
--datadir data \
--port 20301 \
--mine \
--unlock 0x0 \
--password password.txt \
--miner.threads=1 \
--allow-insecure-unlock \
--rpc \
--rpcapi debug,admin,eth,miner,net,personal,txpool,web3 \
--wsapi debug,admin,eth,miner,net,personal,txpool,web3
