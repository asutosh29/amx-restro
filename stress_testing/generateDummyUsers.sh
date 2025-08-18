#!/bin/bash

generate_random_string() {
  head /dev/urandom | tr -dc A-Za-z0-9 | head -c 10
}

generate_unique_username() {
  local base_username="user_$(generate_random_string)"
  local counter=0
  local new_username

  while true; do
    new_username="${base_username}${counter}"
    # Check if the already added
    if [ -f "form_data_1000_users.txt" ] && grep -q "username=${new_username}" "form_data_1000_users.txt"; then
      ((counter++))
    else
      echo "$new_username"
      break
    fi
  done
}

generate_unique_email() {
  local base_email="$(generate_random_string)"
  local counter=0
  local new_email

  while true; do
    new_email="${base_email}${counter}@restro.com"
    # Check if the email already added
    if [ -f "form_data_1000_users.txt" ] && grep -q "email=${new_email}" "form_data_1000_users.txt"; then
      ((counter++))
    else
      echo "$new_email"
      break
    fi
  done
}

OUTPUT_FILE="form_data_1000_users.txt"
> "$OUTPUT_FILE"

for i in {1..1000}; do
  # Generate unique username and email
  UNIQUE_USERNAME=$(generate_unique_username)
  UNIQUE_EMAIL=$(generate_unique_email)

  FIRST_NAME="John${i}"  # Making first name slightly dynamic for better distinction
  LAST_NAME="Doe${i}"    # Making last name slightly dynamic for better distinction
  CONTACT="98765432$(printf "%02d" "$i")" # Ensure contact is 10 digits
  PASSWORD="Test@123"

  DATA_STRING="username=${UNIQUE_USERNAME}&email=${UNIQUE_EMAIL}&first_name=${FIRST_NAME}&last_name=${LAST_NAME}&contact=${CONTACT}&password=${PASSWORD}"

  echo "$DATA_STRING" >> "$OUTPUT_FILE"
done

echo "File '$OUTPUT_FILE' created successfully with data for 1000 users."