# Installation and Configuration

As we continue down the path to configuring our system, let's be sure we understand how to select and configure a different storage driver for our system.

## Selecting a Storage Driver

```
WARNING: devicemapper: usage of loopback devices is strongly discouraged for production use.
         Use `--storage-opt dm.thinpooldev` to specify a custom block storage device.
```

### How do I know what storage driver I am using?

`docker info | grep Storage`

`overlay` is the default, but we want to use `devicemapper` which is supported in production, especially with Enterprise Edition

### How do I change the storage driver?

+ `/etc/docker/key.json` we need a `daemon.json` which takes a keypair
+ create `daemon.json` by `vim daemon.json`
  ```json
  {
  	"storage-driver": "devicemapper"
  }
  ```
+ if you have any images to keep, you need to backup/import after changing
+ stop and restart docker service
  ```
  systemctl stop docker
  systemctl start docker
  ```
+ you will notice afterways, `/var/lib/docker` contains `devicemapper` and
  inside this is where you will find images, etc. going forward
+ test this: go to `/var/lib/docker/devicemapper` then `docker pull centos:6`
+ notice in `/var/lib/devicemapper/devicemapper` file sizes have changed