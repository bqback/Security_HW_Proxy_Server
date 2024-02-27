CA_CRT=$1
CA_KEY=$2
X509_CONFIG=$3
CERT_FILE=$4
KEY_FILE=$5
HOST=$6
SERIAL=$7

openssl genrsa -out $KEY_FILE 2048
cat $X509_CONFIG | sed s~%%DOMAIN%%~"$HOST"~g > /tmp/__x509.conf
openssl req -new -key $KEY_FILE -subj "/CN=$HOST" -sha256 -config /tmp/__x509.conf |
openssl x509 -req -days 3650 -CA $CA_CRT -CAkey $CA_KEY -set_serial $SERIAL -out $CERT_FILE -extfile /tmp/__x509.conf