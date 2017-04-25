Mail API Service
================

# Requirements

- [Golang >= 1.8.1](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/engine/installation)
- [Dep](https://github.com/golang/dep#usage)

# Install

```
git clone github.com/l-vitaly/mailapi
dep ensure
docker build -t mailapi . 
docker run --name mailapi 
    -e "MAIL_USERNAME=user@mail.com" \
    -e "MAIL_PASSWD=passwd" \
    -e "MAIL_FROM=from@mail.com" \
    -e "MAIL_TO=to@mail.com" \
    -e "MAIL_SUBJECT=Subject" \
    -d -p 9000:9000 mailapi
```
