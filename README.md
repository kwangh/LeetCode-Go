# go-tutorial
Follows "A Tour of GO" from https://tour.golang.org/welcome/1 on CentOS 7 machine

---
# install go on CentOS 7 machine
```
sudo yum update -y

# check the latest version from link below
# https://golang.org/dl/
wget https://dl.google.com/go/go1.14.1.linux-amd64.tar.gz

tar xf go1.14.1.linux-amd64.tar.gz
cd go/bin/

# copy go, gofmt binary to local binary directory
cp * /usr/local/bin

# check go version
go version
```

Reference
- https://dejavuqa.tistory.com/321
