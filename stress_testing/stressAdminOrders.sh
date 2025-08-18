#!/bin/bash

echo "Stress Test for Order placing post route"

# Paste your jwt token here
TOKEN=
URL=http://host.docker.internal:8000/admin/orders
FILE=post_order.json

# ab -C token=$TOKEN -n 1000 -c 10 -p $FILE $URL
ab -C token=$TOKEN -n 1000 -c 10 $URL