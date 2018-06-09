# Lecture: Dockerfile Options, Structure, and Efficiencies (Part II)

The Dockerfile has a large number of options that you need to be able to use effectively and efficiently. We will walk through them over the next couple of videos, while showing you the structure and differences between them.

## Edit Our Dockerfile

+ Take the default httpd install and created custom index.html
+ Using a new instruction: `COPY` and `ADD`
+ Takes from local context/path, and put inside image for any containers instatiated from it
+ `COPY` only works with files, `ADD` supports URLs
+ `ADD http://somedomain.com/somefile.txt dest` as an example
+ `COPY index.html /var/www/html/` to put a file locally
+ `VOLUME /mymount` will mount an volume inside an image
+ no way to mount image from underlying storage system
+ this is because theres no guarantee that it will exist on every host; all images must be portable
+ you can create mount points; but not mount to anything locally
+ `ARG` is a command that can proceed the `FROM`; if appears before `FROM` the only instruction that can use this is `FROM`
+ `ARG TAGVERSION=6` for example and now use `FROM centos:${TAGVERSION}`

## Build the Image

+ `docker build mywebserver1:v3` and notice it can't find index
+ `echo "This is my webserver" >> index.html`
+ `docker run -d --rm --name testweb1 mywebserver:v3`

