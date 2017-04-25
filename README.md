Mail API Service
================

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
