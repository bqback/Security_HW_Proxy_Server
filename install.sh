#!/bin/sh

MIGRATOR_PASSWORD=$1
POSTGRES_USER=$2
POSTGRES_PASSWORD=$3
POSTGRES_DB=$4
POSTGRES_HOST=$5
if [ -f "config/.env" ]; then 
    echo ".env file found"
else 
    echo ".env file not found, creating"
    echo "MIGRATOR_PASSWORD=\"$1\"" >> config/.env 
    echo "POSTGRES_USER=\"$2\"" >> config/.env
    echo "POSTGRES_PASSWORD=\"$3\"" >> config/.env
    echo "POSTGRES_DB=\"$4\"" >> config/.env
    echo "POSTGRES_HOST=\"$5\"" >> config/.env
fi

if [ -f "proxy-serv-ca.crt" ] && [ -f "/usr/local/share/ca-certificates/proxy-serv-ca.crt" ]; then 
    echo "CA found"
else 
    echo "Generating CA certificate"
    chmod +x scripts/gen_ca.sh && bash gen_ca.sh
    echo "Adding CA to CA storage"
    cp proxy-serv-ca.crt /usr/local/share/ca-certificates/proxy-serv-ca.crt
fi


echo "Updating CA certificates"
update-ca-certificates
echo "Building Docker image"
docker compose up --build --detach