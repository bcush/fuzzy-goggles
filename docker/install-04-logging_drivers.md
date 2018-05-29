# Installation and Configuration

Docker supports a number of different logging drivers for displaying and gathering daemon and container data. In this lesson, we will show how to change and then verify the logging driver that we choose.

## Configuring Logging Drivers (Syslog, JSON-File, etc.)

+ by default `json-file` file format is used for logging [1]

### How Do I Know the Default Logging Driver?

`docker info | grep Logging`

### How Do I Change the Default Logging Driver?

+ `/etc/docker` and if `daemon.json` is not there, we can create it

### What do Logs Look Like?

+ let's do `docker image pull httpd` and now let's look at logs
+ now `docker container run -d --name testweb httpd` to detach and run
+ see `docker ps` and notice it's running
+ now `docker container inspect testweb | grep IPAddr`
+ let's use `telnet` or `elinks` to do some testing
  ```json
              "SecondaryIPAddresses": null,
            "IPAddress": "172.17.0.2",
                    "IPAddress": "172.17.0.2",
  ```
+ `elinks http://172.17.0.2` and you see `It Works`
+ `docker logs testweb` or if in a swarm `docker service logs` to get logs
+ `docker stop testweb` to stop
+ `docker rm testweb` to remove

### Change syslog to be default

+ setup syslog by checking for `/etc/rsyslog.conf` and find --
  ```
  $ModLoad imudp
  $UDPServerRun 514
  ```
  and uncomment them
+ `systemctl start rsyslog` to start syslog
+ `systemctl status rsyslog` to check it's running
+ return to your docker dir `/etc/docker` and edit `daemon.json`
+ add the `log-driver` and `log-opts`
  ```json
  {
  "storage-driver": "devicemapper",
  "log-driver": "syslog",
  "log-opts": {
        "syslog-address": "udp://172.31.98.106:514"
        }
  }
  ```
+ `systemctl start docker` to restart docker
+ `docker info | grep Logging` to check

### Testing

+ now this affects our ability to use `docker logs`
+ let's `tail -f /var/log/messages` now which is where syslog will log
+ clear the logs `echo "" > /var/log/messages`
+ `systemctl restart rsyslog` to restart syslog
+ `docker container run -d --name testweb -p 80:80 httpd`
+ `curl http://localhost` or `elinks http://localhost` and we get success
+ check the messages and we get the basic `GET` requests
+ now do a `docker logs testweb` and see the logging driver is not supported
+ this is why we configured syslog to see it on the host OS; similar to journald

### Changing on a Container Basis (not Global via daemon.json)

+ `docker stop testweb` to stop
+ `docker rm testweb` to delete
+ `docker container run -d --name testjson --log-driver json-file httpd`
+ `docker logs testjson` and see that it can read the json logs

[1] https://docs.docker.com/config/containers/logging/configure/#supported-logging-drivers