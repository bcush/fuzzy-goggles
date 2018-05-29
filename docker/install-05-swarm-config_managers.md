# Installation and Configuration

This is the first step to making a Docker Swarm available for use. Let's configure our first Swarm Manager and learn how to display the commands and tokens for joining other managers and workers to our cluster.

## Nodes vs. Managers

+ `nodes` do the work; `managers` manage

## Setup Manager

+ `docker images` and we see nothing
+ `docker swarm init --advertise-addr <ip addr of this server>`
+ `docker swarm init --advertise-addr 172.31.98.106` will init the swarm on node
+ the node now becomes a manager
+ copy and paste the provided cli command from the init --
  `docker swarm join --token SWMTKN-1-5xbm8v5qpdjcij8t1w5uz151v609btwbjzt35d7n1ppbm6h0cv-2rf9tsl6pmv1wmhldu9joop83 172.31.98.106:2377`
+ used to join additional nodes to swarm

## What if Forget the Token?

+ `docker swarm join-token worker` on a manager node

## Adding Additional Manager

+ `docker swarm join-token manager` on the manager node
+ a new command and token will be generated for you; copy and use this
+ if you look at command, everything before `-` is the node; after is activity