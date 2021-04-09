#!/bin/bash

for ((i=0; i<100; i++))
do
    echo "${i}回目"
    echo "Hello, World!"
done
exec sleep infinity
