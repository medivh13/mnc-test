# mnc backend test

## Prerequisite:

1. This project requires Golang to run, view installation instruction here https://golang.org/doc/install

## How to run:

This project depedency is managed by go mod, therefore you can put it anywhere other than GO-PATH directory.
    1. cd to project root directory 
    2. then do "go run main.go"
    3. all route is in transport/http/handlers
    5. all endpoints are protected by signature. since is a development env, I return the valid signature if you input the wrong one, and you can use the valid one. the signature formula is in all DTO validation in case you wanna try to make a valid signature by your self, which is in dto

## example :

- ping :

curl --location --request GET 'http://localhost:9008/api/v1/mnc-test/'

- palindrome :

curl --location --request GET 'http://localhost:9008/api/v1/mnc-test/palindrome?text=jody' \
--header 'signature: f134f75af73e299c72bab0dc92ea7045b05974e855780700eab9254d0f401a0c'

- language :

curl --location --request GET 'http://localhost:9008/api/v1/mnc-test/language?id=o' \
--header 'signature: 1eb5b29d45d347a8be86f123e9ddb57d1afd4702b5aa40c357fc6cd6c7adfee9'