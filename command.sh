# !/bin/bash

curl --header "Content-Type: application/json" \
--request POST \
--data '{
    "name":"Vitaliy",
    "email":"vetalyeshor@gmail.com",
    "password":"qwerty"}' http://localhost:8000/auth/sign-up