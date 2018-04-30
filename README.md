
# fortune-teller

This directory contains the fortune-teller gRPC server.  It's pretty simple.
Note we are not configuring the server with TLS configuration as TLS will be
terminated by nginx-ingress, so by the time the request reaches the server we're
in plain insecure mode.  

## Configure kubernetes

Setup fortune-teller:

```
$ bazel run //fortune-teller/k8s:k8s.apply 
```

Teardown fortune-teller:

```
$ bazel run //fortune-teller/k8s:k8s.delete 
```
