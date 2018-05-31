# Image Creation, Management, and Registry

## Searching an Image Repository

Pulling a Docker image presumes you know the exact name of the image you want. Learn how to search the configured repository for an image name and filter the results to get what you need.

## Various Ways to Search for `apache`

+ if you want `apache` there may be many images like this; use `docker search` to narrow this
+ `docker search apache` will perform this for you but you get a lot of results; filter by stars
+ `docker search --filter stars=50 apache`

```
NAME                       DESCRIPTION                                     STARS               OFFICIAL            AUTOMATED
tomcat                     Apache Tomcat is an open source implementati…   1859                [OK]                
httpd                      The Apache HTTP Server Project                  1729                [OK]                
cassandra                  Apache Cassandra is an open-source distribut…   780                 [OK]                
maven                      Apache Maven is a software project managemen…   601                 [OK]                
solr                       Solr is the popular, blazing-fast, open sour…   531                 [OK]                
zookeeper                  Apache ZooKeeper is an open-source server wh…   375                 [OK]                
eboraas/apache-php         PHP5 on Apache (with SSL support), built on …   136                                     [OK]
eboraas/apache             Apache (with SSL support), built on Debian      87                                      [OK]
webdevops/php-apache       Apache with PHP-FPM (based on webdevops/php)    64                                      [OK]
webdevops/php-apache-dev   PHP with Apache for Development (eg. with xd…   53                                      [OK]
tomee                      Apache TomEE is an all-Apache Java EE certif…   51                  [OK]                
```

+ `docker search --filter stars=50 is-official apache` does not work; only one at a time
+ `docker search --filter stars=50 --filter is-official=true apache`

```
NAME                DESCRIPTION                                     STARS               OFFICIAL            AUTOMATED
tomcat              Apache Tomcat is an open source implementati…   1859                [OK]                
httpd               The Apache HTTP Server Project                  1729                [OK]                
cassandra           Apache Cassandra is an open-source distribut…   780                 [OK]                
maven               Apache Maven is a software project managemen…   601                 [OK]                
solr                Solr is the popular, blazing-fast, open sour…   531                 [OK]                
zookeeper           Apache ZooKeeper is an open-source server wh…   375                 [OK]                
tomee               Apache TomEE is an all-Apache Java EE certif…   51                  [OK]   
```

+ now we are getting official releases, ranked over 50 stars, for `apache`
+ `docker search --limit 10 apache` to limit results, regardless