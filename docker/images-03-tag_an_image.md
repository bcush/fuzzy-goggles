# Image Creation, Management, and Registry

## Tagging an Image

We have talked about image tags but not explored exactly what they are, how to use them ourselves, and how they may affect the underlying storage of images. Let's take a look at that now.

+ you can build layers upon an image without affecting the original
+ you want to customize centos image; you can use a tag to build upon it
+ `docker tag centos:6 mycentos:v1`
+ `docker images`

```
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
httpd               latest              fb2f3851a971        4 weeks ago         177MB
hello-world         latest              e38bc07ac18e        7 weeks ago         1.85kB
hello-world         linux               e38bc07ac18e        7 weeks ago         1.85kB
centos              6                   70b5d81549ec        7 weeks ago         195MB
mycentos            v1                  70b5d81549ec        7 weeks ago         195MB
```

### Tags

+ notice the different tags, but the same image id, they are exactly the same
+ now you can start a container and commit your changes
+ you can even remove the original repo and tag without affecting it

+ `docker rmi centos:6` states `untagged` because since a dupe exists; it just removes the tag

```
[root@beecushman1 docker]# docker rmi centos:6
Untagged: centos:6
Untagged: centos@sha256:67b491e26d566ee9c55578bfd6115554a6e1b805a49502ead32cb1a324466f2c
```

+ `docker rmi mycentos:v1` will now delete the actual underlying file system

```
[root@beecushman1 docker]# docker rmi mycentos:v1
Untagged: mycentos:v1
Deleted: sha256:70b5d81549ec19aa0a10f8660ba5e1ab9966008dbb1b6c5af3d0ecc8cff88eef
Deleted: sha256:e0dec291ae94d0f486c2a6a98732bac23e45b0914f8f9e67c57c5e6a2cd7fed7
```

+ `docker pull centos:6` will redownload and pull down from the registry
+ `docker tag centos:6 myrepo/mycentos:ver2` and then do a `docker images`
+ `using myrepo` is just a way to differentiate what might be local vs `Docker Hub`