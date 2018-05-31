# Installation and Configuration

When you size any environment, there are some considerations you will want to take into account. We will talk about dealing with Docker containers, services, or any other applications stack.

## Cluster Sizing Requirements

In general, sizing your environment for a Docker Swarm cluster is going to be the same as with any application in your ecosystem. Consider the following:

+ `CPU, memory, and disk,` containers still have the same requirements
+ `Concurrency,` what are the load requirements at peak and total?

## Universal Control Plane (UCP) System Requirements

+ Docker EE's homegrown cluster management system;
+ Allows management of everything via a web interface

### Breakdown of Manager Nodes for Fault Tolerance (Requirements)

+ 8 GB RAM (Managers or DTR Nodes)
+ 4 GB RAM (Workers)
+ 3 GB free disk

### Recommended

+ 16 GB RAM (Managers or DTR Nodes)
+ 4 vCPUS (Workers or DTR Nodes)
+ 25-100 GB of space

## Docker EE Includes

+ Docker Engine with Support from Docker
+ Docker Trusted Registry
+ Universal Control Plane

## Compatibility

+ Docker Engine 17.06+
+ DTR 2.3+
* UCP  2.2+

## Reocmmendations

+ Plan for Load Balancing
+ Use External Certificate Authority for Production