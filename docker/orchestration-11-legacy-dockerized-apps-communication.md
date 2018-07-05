# Orchestration

## Lecture: How Dockerized Apps Communicate with Legacy Systems

Communication is key and we have seen many ways that Docker communicates. Let's focus for a few minutes on some key items to consider when deploying Docker in your IT ecosystem.

### Dockerized Application Communication

Essentially, Docker containerized applications communicate with all other applications and infrastructure exactly the way you would expect any other application running in your environment with a few exceptions. Consider the Docker diagram below.

Side by side, the architecture of a Docker containerized application externally does not differ from what you might see on more traditional physical or virtual deployments.

However, there are a couple of caveats to keep in mind so that all of your applications know how to find each other.

* network drivers/routing mesh

### Communication Caveats on Docker Containers

A few items to consider with Dockerized applications when putting them in place in your IT ecosystem:

* Routing - either you will have to expose the Docker application/service via port redirection when you launch them OR you will have to provide a routing mechanism (statically) to the container network on their host(s)

* Port Redirection - as mentioned above, managing ports either through passing them directly through or redirecting them to the underlying host on different ports, will affect how they behave in your environment

* Portability - making sure that data is external to the container application (on the host or via a network share) can have (sometimes significant) impacts on their performance.

### Containers Should Be..

* Abstract - you containerize an application so that it remains removed from other components in the stack; that abstraction makes updates easier, but does introduce additional potential areas that communication can fail

* Portable - applications that are containerized can be picked up and put almost anywhere and will be consistent in their content and behavior; however, re-establishing communication with other componenets in the stack wil depend on the new location and how you have planned across the resource.

* Flexible - containers give you almost limitless flexibility; be sure to try and refrain from launching your container services tied too closely to external (and changable) variables (hard coded related IP addresses or hostnames not easily changed)

* 