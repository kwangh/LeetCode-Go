# go-tutorial
Follows "A Tour of GO" from https://tour.golang.org/list on CentOS 7 machine

---
# install go on CentOS 7 machine
```
sudo yum update -y

# check the latest version from link below
# https://golang.org/dl/
wget https://dl.google.com/go/go1.14.1.linux-amd64.tar.gz

sudo tar xf go1.14.1.linux-amd64.tar.gz -C /usr/local

# adjust the path variable
export PATH=$PATH:/usr/local/go/bin
source ~/.bash_profile

# check go version
go version
```

# test the installation
```
git clone https://github.com/kwangh/go-tutorial.git
cd go-tutorial/hello
go build
./hello
```

Reference
- https://linuxize.com/post/how-to-install-go-on-centos-7/

---


# extra settings

## add go path to profile 
```
echo '[[ ":$PATH:" != *":/usr/local/go/bin:"* ]] && export PATH="${PATH}:/usr/local/go/bin"' >> ~/.profile
echo '[[ ":$PATH:" != *":/root/go/bin:"* ]] && export PATH="${PATH}:/root/go/bin"' >> ~/.profile
source ~/.profile
```
