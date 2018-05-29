# Installation and Configuration

Now that we have a swarm started with a management node, we need to add a couple of worker nodes that the manager can boss around. Let's take a look at that and then verify they are registered and the swarm shows as available.

## Adding to Swarm

+ `docker swarm join-token worker` to display the token and command needed
+ `docker node ls` to list the nodes
  ```
  ID                            HOSTNAME                      STATUS              AVAILABILITY        MANAGER STATUS      ENGINE VERSION
es372vu0f6qjoo9lkikt6qv9b *   beecushman1.mylabserver.com   Ready               Active              Leader              18.03.1-ce
```
+ notice that the manager status is the leader
+ `docker system info` shows more relevant info
+ join the worker node to your cluster using the previous command
+ try `docker node ls` on your worker node; notice it doesn't work
