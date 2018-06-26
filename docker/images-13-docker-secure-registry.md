# Prepare for a Docker Secure Registry

Before we can use a private repository, we will need to secure it and offer user authentication. Let's create a self-signed certificate, use the 'registry' container by Docker to create basic user authentication, and then copy the files where they need to go on the hosting server.

## Registry

+ simply like the DockerHub; place to pull images from
+ application already containerized by Docker to do this
+ we are going to deploy the secure registry

## Create Dirs
 + `/home/user` with have `certs` and `auth`
 + `openssl req -newkey rsa:4096 -nodes -sha256 -keyout certs/dockerrepo.key -x509 -days 365 -out certs/dockerrepo.crt -subj CN=myregistrydomain.com`
 + so we will have one key and one self signed cert

## Create Self Signed Certs


## Certificate Keys

##

+ `sudo mkdir -p /etc/docker/certs.d/myregistrydomain.com:5000`
+ `sudo su`
+ `cd /etc/docker/certs.d./myregistrydomain.com:5000/`
+ `docker images`
+ `docker pull registry:2`
+ `docker images`
+ `docker run --entrypoint htpasswd registry:2 -Bbn test password > auth/htpasswd`
+ `cd auth`
+ `cat htpasswd`