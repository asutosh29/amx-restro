#!/bin/bash

echo "Run generateDummyUsers.sh before running this script"

echo "Stress Test for Registering users: post route"

# Paste your jwt token here
TOKEN=
URL=http://host.docker.internal:8000/register
FILE=form_data_1000_users.txt

ab -C token=$TOKEN -T application/x-www-form-urlencoded -n 1000 -c 10 -p $FILE $URL