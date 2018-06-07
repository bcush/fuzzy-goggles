# Image Creation, Management, and Registry

## Create an Image with Dockerfile

Before we talk about the structure, options, and optimizations around Dockerfiles, we need to understand how to use the 'build' command to complete a build with a variety of option.

### Build with Dockerfile

+ `docker images`
+ `mkdir Dockerfiles` and go into it
+ `vim Dockerfile` when you buid a Dockerfile, it's assumed
  that it's named `Dockerfile`

```
# This is a small test Dockerfile
FROM centos:latest
LABEL maintainer="myname@email.com"

RUN yum update -y
```

+ `docker images` you'll see images, but now we'll see our own
+ `docker build -t custom-image:v1 .` to build
+ `docker images` to see the image you built
+ `docker build -t custom-image:v1 .` again, it's faster w cache
+ `vim Dockerfile` and make a change to label
+ `docker build -t custom-image:v1 .` again
+ notice that the label layer rebuilds, and all layers above
+ it will continue to use the cache until it reaches a layer that is not cached, executes that change, and invalidates each layer underneath
+ `cp Dockerfile Dockerfile2`
+ `vim Dockerfile2` and change to `ubuntu:latest`
+ change update to `apt-get update -y`
+ `docker build -t customubuntu:v1 -f Dockerfile2 .`
+ you can do a docker build from a git repo/branch
+ you can do a docker build from a tarball
+ you can also redirect a dockerfile

## Squash and No-Cache Builds

+ you can squash the layers down
+ you can control whether to cache

+ `docker build --pull --no-cache --squash -t optimized:v1 .`
+ pull from registry, don't cache any layers, and squash all the layers into one image layer
+ remeber that `--squash` is only on experimental features
+ you will see that none of the layers are cached
+ notice that even though the name is the same, the image id is different
+ 
