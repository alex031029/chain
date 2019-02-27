#!/usr/bin/env bash

cd "$(dirname "${BASH_SOURCE[0]}")"

set -u
set -e

proj_dir=../..

echo "[*] Starting bank."

cpchain=$proj_dir/build/bin/cpchain
ipc_path_base=data/cpc-

# bank
nohup $cpchain $args --ipcaddr ${ipc_path_base}22 --datadir data/data22 --rpcaddr 0.0.0.0:8522 --port 30332  \
         --logfile data/logs/22.log 2> data/logs/22.err.log &


printf "\nAll nodes configured. See 'data/logs' for logs"