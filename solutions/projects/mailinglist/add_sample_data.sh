#!/bin/zsh

for email in $(cat sample_data.txt)
do
  curl -X POST \
    'localhost:8080/email/create' \
    --header 'Content-Type: application/json' \
    --data-raw '{"Email":"'$email'"}'
done
