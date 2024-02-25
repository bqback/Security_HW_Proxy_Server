#!/bin/sh

openssl genrsa -out proxy-serv-ca.key 2048
openssl req -new -x509 -days 3650 -key proxy-serv-ca.key -out proxy-serv-ca.crt -config="scripts/cert.conf" -subj "/CN=MITM Proxy CA"
