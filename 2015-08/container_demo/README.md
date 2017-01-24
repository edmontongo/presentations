## Container demonstration

The files in this directory illustrate a simple set-up for networking containers across hosts.

There are two `.config` files, one for each side. The configuration is the same as shown on
page 29 of the slides. If the `wifi` section is not commented out, it will set up an ad-hoc
wireless network, otherwise it will use the `interface` specified in the `network` section.

The examples suppose that the docker daemon has been configured to run with `-b=bridge0`,
e.g. by setting `DOCKER_OPTS=-b=bridge0` in `/etc/default/docker`.

To run the examples in this directory,
```
> source source.me                     # sets GOPATH
> go get                               # fetch the github packages
> # edit your choice of configuration file
> go run demo.go 0{1,2}_network.config # sets up the network
```

## Sources

The example scenario is based on an [excellent blog post on sequenceiq](http://blog.sequenceiq.com/blog/2014/08/12/docker-networking/),
which also has other examples to try, such as `ssh` tunnels.
