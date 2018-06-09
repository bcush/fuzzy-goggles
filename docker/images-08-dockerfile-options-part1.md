# Lecture: Dockerfile Options, Structure, and Efficiencies (Part I)

## Dockerfile Options and Structure

+ FROM
+ ARG
+ RUN
+ CMD
+ LABEL
+ MAINTAINER
+ EXPOSE

+ These are instructions, and any use of these will create an intermediate layer containing whatever instruction was put in the Dockerfile
+ In the end this is what we will have as a result

## Starting Out

+ You can have comments:

`# This is our first Dockerfile`

+ Next add `FROM` which is a GitHub repo, tarball, or a public or private repository

+ `FROM centos:6` if you don't specify an image, it will assume you want latest
+ `MAINTAINER` was intended to just be an email address of the person or org responsible
+ This tag is deprecated in favor of the `LABEL` instruction
+ `maintainer="latest@linuxacademy.com"` is an example of a `LABEL` tag
+ `RUN` will allow you to run a command; executing it in a new layer during the build process, running the intermediate image it creates, runs during build process
+ You can use the `shell` form or `exec` like `RUN ["echo hello", "echo hello"]`
+ `RUN yum update -y && yum install httpd net-tools` on centos:6 base image.
+ `CMD` can only be defined once per Dockerfile, only the last will take effect
+ It's intended to provide the command to run when a containers, or provide params to entrypoint
+ `ENV` you can pass an environment var in `RUN` but you can also define it using `ENV`
+ `ENV ENVIRONMENT="production"`
+ `EXPOSE` let's you expose a port on the system for example `EXPOSE 80`
+ `ENTRYPOINT` can be a JSON array or just a run command
+ `ENTRYPOINT apachectl "-DFOREGROUND"` entrypoint dicatates what runs when the container
starts, you don't to define that in the `CMD` which is really used for installing packages
+ `docker run -d --name testweb1 --rm mywebserver:v1` and installed elinks `sudo yum install -y elinks`
+ `docker ps` to see the deatched instance; find IP `docker container inspect <id>`
+ Go the IP addresses port
+ `docker stop testweb`

## Image Optimization

+ do a `docker images` and see that the new image is 300MB in size; how to optimize
+ we can combine commands; and from an efficiency standpoint, combine as many as possible
+ understand that every `RUN` command will create a new layer, you saw this earlier

+ `vim Dockerfile` and see that each command has it's own line
+ Change it to do `sudo yum update -y && yum install -y httpd net-tools && \`
				  `mkdir -p /run/httpd && \`
+ Check out how many layers were in the old image with `docker history mywebserver:v1 | wc -l`
+ Now, do `docker build -t mywebserver:v2 .` and watch the difference
+ 

```
# This is our first real Dockerfile
FROM centos:6

LABEL maintainer="latest@linuxacademy.com"

RUN yum update -y && yum install httpd net-tools
RUN mkdir -p /run/httpd
RUN rm -rf /run/http/* /tmp/httpd*

CMD echo "Remember to check your container IP address"

