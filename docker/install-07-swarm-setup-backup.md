# Installation and Configuration

One of the most common administrative tasks on your plate every day is setting up backups and testing restores. Today, we will show you how to easily back up your swarm configuration AND state, and restore it to a completely different system to have it pick right up where it left off.

## Setting up Backups (from Swarm perspective)

### First Let's Build a Service

+ `docker images` and focus on httpd image
+ `docker pull httpd` if you don't have it
+ `docker node ls` shows us the three servers, 2 nodes
+ `docker swarm service create --name bkupweb -p 80:80 --replicas 2 httpd`
+ `docker service ls`
+ `docker service ps bkupweb`

### Next Let's Review Backup of a Swarm

+ unlock the swarm from a manager node; using the key from before
+ `sudo systemctl stop docker` stop the Docker service
+ go to `/var/lib/docker` and you'll find a `swarm` directory
+ tells us the state of the docker engine, what may be running on it, workers, and any certificates in our swarm
+ these files need to be backed up
+ `mkdir /root/swarm && cd /root/swarm` make a backup folder
+ `cp -rf /var/lib/docker/swarm/ .` copy the files to backup folder
+ `systemctl start docker` start docker again
+ since we had a service running, docker will try to bring it back to the state when it was stopped
+ `tar cvf swarm.tar swarm/` to create a tar file backup
+ `scp swarm.tar user@beecushman2.mylabserver.com:/home/user`


### All Docker Services and Nodes Crash (Recovery)

+ `systemctl stop docker` on manager and all nodes
+ how do we recover our swarm, reinit, and make sure state is restored
+ `/var/lib/docker`
+ `docker images`
+ `cd /var/lib/docker`
+ `rm -rf swarm/` remove swarm; we don't care
+ `systemctl stop docker`
+ `cd /home/user && mkdir tmp && cd tmp/`
+ `tar xvf ../swarm.tar && cd swarm && mv swarm /var/lib/docker`
+ `systemctl start docker` will reconnect auto to nodes
+ when it starts, we need to force and start a new cluster
+ `docker swarm init --force-new-cluster`
+ `docker service ls`
+ `docker service ps bkupweb`

### Review

+ as long as you have a copy of the `swarm` folder, you can create a new cluster and start services again