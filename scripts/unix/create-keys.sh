#!/usr/bin/env bash

set -e

openssl ecparam -name prime256v1 -genkey -out p256-key.pem -noout
echo "Private key 'p256-key.pem' was created"
openssl ec -in p256-key.pem -pubout -out p256-pubkey.pem
echo "Public key 'p256-pubkey.pem' was created"