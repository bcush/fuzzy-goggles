# Image Creation, Management, and Registry

## Use CLI Commands to Manage Images (List, Delete, Prune, RMI, etc)

Managing our local images goes beyond what we have seen so far in the 'docker images' command. We have the ability to back up, restore, list, remote, pull, push, and otherwise inspect the details of any image. We will show you how in this lesson.

### `docker images` vs. `docker image`

+ `docker image` will give us the commands we will use
+ `docker images` to see what images we have, and notice the diff tags but same underlying filesystem
+ `docker image history myrepo/mycentos:v2`

```
[root@beecushman1 docker]# docker image history myrepo/mycentos:ver2
IMAGE               CREATED             CREATED BY                                      SIZE                COMMENT
70b5d81549ec        7 weeks ago         /bin/sh -c #(nop)  CMD ["/bin/bash"]            0B                  
<missing>           7 weeks ago         /bin/sh -c #(nop)  LABEL name=CentOS Base Im…   0B                  
<missing>           7 weeks ago         /bin/sh -c #(nop) ADD file:2a39fb46f860e75d7…   195MB  
```

+ these are the three layers used to build this image
+ there are interim container images that go into building base image
+ these layers in aggregate are what used to make the image
+ since this image was not created, we only get what was necessary

### Copying Docker Images without a Repo

+ `docker images` again to see what we have
+ you can exchange images with other systems, regardless of whether you have a repo to push to
+ `docker image save` will create a tar and let you exchange without a repo
+ `docker image save myrepo/mycentos:ver2 -o mycentos.custom.tar`

### Importing

+ first let's remove the image we just saved
+ `docker image rm` or `docker rmi` are the same thing
+ `docker rmi myrepo/mycentos:ver2` to remove it
+ now let's reload our tar file using `reload` or `import`

### Reviewing tar Backup

```
[root@beecushman1 docker]# tar tvf my.centos.custom.tar | more
drwxr-xr-x 0/0               0 2018-04-06 21:10 3cdae3dd0537787ca73f7603e5b823f886bff78679a28e53df412cc18f011c23/
-rw-r--r-- 0/0               3 2018-04-06 21:10 3cdae3dd0537787ca73f7603e5b823f886bff78679a28e53df412cc18f011c23/VERSION
-rw-r--r-- 0/0            1358 2018-04-06 21:10 3cdae3dd0537787ca73f7603e5b823f886bff78679a28e53df412cc18f011c23/json
-rw-r--r-- 0/0       202639360 2018-04-06 21:10 3cdae3dd0537787ca73f7603e5b823f886bff78679a28e53df412cc18f011c23/layer.tar
-rw-r--r-- 0/0            1862 2018-04-06 21:10 70b5d81549ec19aa0a10f8660ba5e1ab9966008dbb1b6c5af3d0ecc8cff88eef.json
-rw-r--r-- 0/0             209 1970-01-01 00:00 manifest.json
-rw-r--r-- 0/0              96 1970-01-01 00:00 repositories
```

### Importing

+ `docker import mycentos.custom.tar localimport:centos6`
+ `docker load` loads from a stream, so you have to redirect `docker load --input mycentos.custom.tar`
+ this command differs in that it doesn't require a tag, so you end up with the original tag values

+ `docker image ls` is the same as `docker images`

### Review and Prune

+ important when we start looking at containers/builds; it leaves all these intermediate images
+ `docker image prune` says it will delete all our `dangling images`
+ `dangling images` are not associated with any image or container
+ `docker image prune -a` will remove all images without a container associated
+ `docker image tag`
+ `docker image inspect` has options to format/provide json output about the docker images

