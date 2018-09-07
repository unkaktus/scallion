scallion
--------

hey, scallion is a long onion. this scallion is a not-so-great
one-shot hacky thing for exposing tcp port 443 on the internet
and passing all the data to an onion service.


the idea isn't new. though this thing's small and kinda works.
this is not a production-ready thing. maybe someday.
works for running a blog on the internet without hosting
it somewhere. yay.

```shell
docker run -d -p 443:443 --name scallion nogoegst/scallion -addr heylookthisis.onion
```

Building
========
```shell
dep ensure -v
env GOOS=linux GOARCH=amd64 go build -v .
docker build -t nogoegst/scallion .
```
