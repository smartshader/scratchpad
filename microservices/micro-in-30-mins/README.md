Microservices in 30 minutes
===========================

Generate certificates
```shell script
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 3600 -nodes
```