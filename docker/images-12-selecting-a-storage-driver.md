# Selecting a Storage Driver

There are a variety of 'pluggable' storage drivers available for Docker CE and EE to use to store images and containers on the underlying host. We will display all the available storage drivers and show how to change our configuration to one officially supported for our distribution and version.

+ pluggable architecture, supporting multiple storage drivers
+ common one is the `devicemapper` and there are some benefits to using `aufs` or others
+ this means that there are efficiencies to be gained

## Devicemapper

+ can be used on disk as block storage; uses a loopback adapter to provide that
+ can be used with a block storage device (actual disk) and allow docker to manage it
+ we will just indicate the storage device
+ check the Docker site

## Cert Perspective

+ how can i tell what my docker storage driver is?
+ `docker info` and look for `storage subsystem` - just do this and grep for `storage`
+ in our case, let's use devicemapper because it likely being used
+ go to `/etc/docker` and in here is a `key.json` and we need to create a `daemon.json`
+ `touch daemon.json` and in here using the `storage-driver` add the following:

```
{
"storage-driver": "devicemapper"
}
```
+ now restart docker; but be mindful if you have images you wish to
  retain, then you need to export, and reimport after changing the
  storage driver; each of these images will go away
+ `systemctl stop docker` and `systemctl start docker`
+ `docker images` and see nothing there
+ check `/var/lib/docker` and now in `devicemapper` we have a place
  where our images will be stored
+ do a `docker pull centos:6` and once it gets loaded you can 
  instatiate new containers from these images
