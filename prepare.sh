#!/bin/sh

sh scripts/gen_ca.sh
cp proxy-serv-ca.crt /usr/local/share/ca-certificates/proxy-serv-ca.crt
update-ca-certificate