# Image Creation, Management, and Registry

## Inspect Images and Report Specific Attributes Using Filter and Format

The `docker inspect` command can provide a plethora of insight into all of the
attributes of your images. We will demonstrate how to display that information
and then how to format that information to filter out what you want in
several different formats.

## Docker Image Inspect

+ `docker image inspect centos:6` to inspect image; lots of json returned
+ `docker image inspect centos:6 > centos6.output` and look at output
+ `docker image inspect centos:6 | grep ContainerConfig` returns little

### Get Hostname, etc.

+ `docker image inspect centos:6 --format '{{.ContainerConfig.Hostname}}'`
+ `docker image inspect centos:6 --format '{{.ContainerConfig.StdinOnce}}'`
+ `docker image inspect centos:6 --format '{{.ContainerConfig}}` is not useful
+ `docker image inspect centos:6 --format '{{json .ContainerConfig}}'` for everything
+ `docker image inspect centos:6 --format '{{.RepoTags}}' know tags assoc. with image