# Orchestration

## Lecture: Convert an Application Deployment into a Stack File Using a YAML Compose File with 'docker stack deploy'

The name of this lesson is a mouthful, but it is an important concept to understanding how many of the various pieces of Docker tie together to show the true power and flexiblity you have at your fingertips. This is a longer topic, but don't worry, your patient will be rewarded with a better understanding of how all this fits together.

## Install Docker-Compose

* `mkdir Dockerfile`
* `vim Dockerfile`

```
# simple webserver

FROM centos:latest
LABEL maintainer="admin@linuxacademy.com"

RUN yum install -y httpd
RUN echo "Our Container Website" >> /var/www/html/index.html

EXPOSE 80

ENTRYPOINT apachectl -DFOREGROUND
```

* we can build this particular file just to make sure it works
* `docker images`
* `docker build -t myhttpd:v1 .`
* there are six steps, so we have six intermediary containers

## Creating Application Stack

* say we have two api servers, each listening on port 80, 81
* we want to load balancer the load between these two using nginx
* our service stack will be deploy service1 (api1) service2 (api2) and lb
* we deploy this using the `docker service create` but we can simplify w/ compose
* `docker run -d --name testweb -p 80:80 myhttpd:v1`
* `docker ps`
* `elinks http://localhost:80`
* `docker images`
* `docker stop testweb`
* `docker rm testweb`
* `docker rmi myhttpd:v1`

## Now We Replicate this Work

* now build YAML file that docker compose can use
* this will deploy multiple containers at same time
* then we define this as a stack of containers we can deploy

* `docker images`
* create your docker compose yaml
* we will create 3 services: apiweb1, apiweb2, and loadbalancer
* apiweb1 will create the image, apiweb2 will use that image
* lb will use nginx, and we will map the ports

* `vim docker-compose.yml`

```
version: '3'
services:
  apiweb1:
    image: myhttpd:v1
    build: .
    ports:
      - '81:80'
  apiweb2:
    image: myhttpd:v1
    ports:
      - '82:80'
  load-balancer:
    image: nginx:latest
    ports:
      - '80:80'
```

* `docker-compose up -d`
* `docker ps`
* `elinks http://whatver`
* `docker-compose ps`
* `docker-compose down --volumes`

## Now we can deploy to the swarm (instead of a single node)

* we tested this with build and compose first to test it will work
* docker stack deploy will not support a dynamic deploy
* `docker ps -a`
* `docker ps`
* `docker stack deploy --compose-file docker-compose.yml mycustomstack`
* each service can now be managed seperately
* `docker service ls`
* `docker service ps mycustomstack` no service called mycustomstack?
* remember each service is named differently; each service managed seperately
* `docker service ls`
* create your dockerfile, then use docker build to test your build,
  create your yml file that defines services in stack, two api servers,
  and one load balancer, then ran docker-compose against the file to
  see that it created the services on a single node, then we used
  docker stack deploy to take that yml file and deploy three services
  all at the same time across the swarm
* this gives us the ability the run across any number of nodes in a swarm
* gives us the flexibility to define multiple services at once