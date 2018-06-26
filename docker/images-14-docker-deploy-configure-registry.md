# Deploy, Configure, Log Into, Push, and Pull an Image in a Registry

Now that we have the security work done for our private registry,
we can deploy and configure it for use. We will test it locally,
and then log in and test via a remote system.

##

+ `docker run -d -p 5000:5000 -v `pwd`/certs:/certs -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/dockerrepo.crt -e REGISTRY_HTTP_TLS_KEY=/certs/dockerrepo.key -v `pwd`/auth:/auth -e REGISTRY_AUTH=htpasswd -e REGISTRY_AUTH_HTPASSWD_REALM="Registry Realm" -e REGISTRY_AUTH_HTPASSWD_PATH=/auth/htpasswd registry:2`
+ `docker images`
+ `docker pull busybox`
+ `docker images`
+ `docker tag busybox myregistrydomain.com:5000/my-busybox`
+ `docker push myregistrydomain.com:5000/my-busybox` fails as we didn't provide creds, but we have a cert
+ `docker login myregistrydomain.com:5000/my-busybox`
+ `your usernmme test and password test`
+ `docker rmi busybox:latest`
+ `docker rmi <theotherone>`
+ `docker pull myregistrydomain.com:5000/my-busybox`
+ `docker images`
+ `cat /etc/hosts`
+ `sudo vim /etc/hosts`
+ `add ip and myregistrydomain.com`
+ `docker images`
+ `docker pull my registrydomain.com:5000/my-busybox`
+ you get a `cert signed by unknwon auth`
+ copy certs over to the same location on the nodes
+ 