# Orchestration

## Lecture: Understanding the 'docker inspect' Output

We have used the inspect command throughout the course; let's take a brief step back and concentrate on exactly how to use the output and insight it can provide on Docker constructs.

## How to Understand the Docker Inspect Output

* allows us to get detailed information on the constructs controlled by docker
* this can be a container, node, service, etc.
* by default, this is returned in json format
* `docker images`
* `docker run -d --name testweb httpd`
* `docker ps`
* `docker container inspect testweb`
* json is returned for this container
* `docker inspect testweb | grep IPAddr`

## Get What You Want

* `docker inspect testweb | more`
* structured in hierarchies, meaning we have an object called state
* it has members in it called state
* we have an object called hostconfig, with memebrs called binds, etc.
* if we want to pull out if the container in question has been paused
* `docker container inspect --format="{{.State.Paused}}" testweb`
* and this returns `false` which is paused's value
* what if you want everything in that state

* if you want to see the entire list of members in a high level structure
* just proceed your template tags with json
* `docker container inspect --format="{{json .State}}" testweb`
* one of the more common things you will look for is Config
* `docker container inspect --format="{{json .Config}}" testweb`