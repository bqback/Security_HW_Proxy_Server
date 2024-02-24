CERT_KEY=$1
CA_CRT=$2
CA_KEY=$3
CERT_FILE=$4
HOST=$5
SERIAL=$6

openssl req -new -key $CERT_KEY -subj "/CN=$HOST" -sha256 |
openssl x509 -req -days 3650 -CA $CA_CRT -CAkey $CA_KEY -set_serial "$SERIAL" -out $CERT_FILE