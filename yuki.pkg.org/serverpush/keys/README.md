Generate https SSL certificate on deveplopment environment

*Reference: https://gist.github.com/denji/12b3a568f092ab951456


openssl genrsa -out YOUR_SERVER.key 2048
openssl ecparam -genkey -name secp384r1 -out YOUR_SERVER.key
openssl req -new -x509 -sha256 -key YOUR_SERVER.key -out YOUR_SERVER.crt -days 3650