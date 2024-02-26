CA_CRT=$1
CA_KEY=$2
CERT_FILE=$3
KEY_FILE=$4
HOST=$5
SERIAL=$6

openssl genrsa -out $KEY_FILE 2048
openssl req -new -key $KEY_FILE -subj "/CN=$HOST" -sha256 |
openssl x509 -req -days 3650 -CA $CA_CRT -CAkey $CA_KEY -set_serial $SERIAL -out $CERT_FILE