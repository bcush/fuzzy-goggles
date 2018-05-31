# Installation and Configuration

Although not a requirement for the certification exam, we will be walking
through the installation and configuration of both UCP and DTR for use
throughout portions of the remainder of the course.

## Overview

+ First, you want to make sure your swarm is configured w/ manager and worker
nodes.

+ Install it with `docker container run --rm -it --name ucp -v /var/run/docker.sock:/var/run/docker.sock docker/ucp:2.2.4 install --host-address 172.31.116.158 --interactive`

