#!/bin/sh

openssl genrsa -out proxy-serv-ca.key 2048
openssl req -new -x509 -days 3650 -key proxy-serv-ca.key -out proxy-serv-ca.crt -subj "/CN=MITM Proxy CA/O=Mike Winfield, Inc."
