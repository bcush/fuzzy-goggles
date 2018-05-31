# Image Creation, Management, and Registry

## Pull an Image from a Registry (Using Docker Pull and Docker Images)

This lesson will give you a quick introduction to using the 'docker pull' command to pull images from a repository and then provode the basics around 'docker images' to see what you have in your local repository once downloaded.

### Docker Images

+ Default repository is `Docker Hub` from a pull/push perspective.
+ Private repository like `Docker Trusted Repo` can be done w/ UCP.

### Pull

+ `git pull hello-world` by default will pull the `tag` that is `latest`
+ `git pull -a hello-world` will pull all tags for the image
+ `docker pull --disable-content-trust hello-world` will allow you to get unsigned image

### Images

+ `docker images` give us the display of images, tags, image id, created, and size
+ `docker run hello-world` will instantiate a container based on `hello-world` image
+ `docker images --all`
+ `docker images --digests` will show you the digests, different way to refer to image

### Filter

+ `docker pull centos:6` to pull version 6 image for centos
+ `docker images` and notice the created data was `7 weeks ago`
+ `docker images --filter "before=centos:6"` will show you `hello-world`
+ `docker images --filter "since=centos:6"` will show images since centos 6

### Reviewing

+ `docker images --no-trunc` gives image ids associate to each
+ `docker images` compare these against each other
+ `docker images --quiet` will only show you image ids; good for piping