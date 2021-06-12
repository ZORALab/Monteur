#!/bin/bash

domain="$HOSTNAME"
ip="10.0.0.1"
cert_path="${PWD}/.configs/tls/certs/tls.crt"
key_path="${PWD}/.configs/tls/keys/tls.pem"
duration=365
algorithm="rsa:4096"

openssl req -new -x509 \
	-sha512 \
	-newkey "$algorithm" \
	-nodes \
	-keyout "$key_path" \
	-days "$duration" \
	-out "$cert_path" \
	-subj "/CN=${domain}" \
	-addext "subjectAltName=IP:${ip},IP:127.0.0.1"

1>&2 echo "
**WARNING**
These certs are for testing usage ONLY. DO NOT USE for production.

Cert location: ${cert_path}
Key location : ${key_path}
"
