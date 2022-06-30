#!/bin/bash
 
echo "calling fast endpoint"
for i in $(seq 10); do curl localhost:8081/fast & done;

echo "calling slow endpoint"
for i in $(seq 10); do curl localhost:8081/slow & done;

echo "waiting for responses"
wait $(jobs -p)

echo "done"
