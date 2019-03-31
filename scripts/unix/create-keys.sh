#!/usr/bin/env bash
openssl ecparam -name prime256v1 -genkey -out p256-key.pem -noout
openssl ec -in p256-key.pem -pubout -out p256-pubkey.pem