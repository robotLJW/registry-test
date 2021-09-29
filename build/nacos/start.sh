#!/bin/sh
set -e

cd /home

for i in $(seq 1 100)
do
    ./nacos &
done

while true; do
    sleep 60
done