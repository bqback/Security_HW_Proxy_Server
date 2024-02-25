CA_CRT=$1
CA_KEY=$2
CRL_FILE=$3
CERT_FILE=$4
KEY_FILE=$5
HOST=$6
SERIAL=$7

openssl genrsa -out $KEY_FILE 2048
cat /home/nonroot/.tls/cert.conf > /tmp/cert.conf
openssl req -new -key $KEY_FILE -subj "/CN=$HOST" -sha256 -config /tmp/cert.conf |
openssl x509 -req -days 3650 -CA $CA_CRT -CAkey $CA_KEY -set_serial "$SERIAL" -extfile /tmp/cert.conf -out $CERT_FILE