
# fortune-teller

This repository contains a sample gRPC application built with bazel.  It was
initially built as an example of how to use grpc with
[nginx-ingress](https://github.com/kubernetes/ingress-nginx/docs/grpc). 

## Build

```
$ bazel build //app:fortune 
Target //app:fortune up-to-date:
  bazel-bin/app/linux_amd64_static_pure_stripped/fortune
```

## Run

```sh
$ bazel run //app:image
Loaded image ID: sha256:aa597c897c873116fcbfccafecf5ab0f6f4178a05e4a00c8e79de91ac0d2e9e7
Tagging aa597c897c873116fcbfccafecf5ab0f6f4178a05e4a00c8e79de91ac0d2e9e7 as bazel/app:image
```

```sh
$ docker run bazel/app:image
2018/05/01 02:13:43 Restored /tmp/fortune-teller/usr/share/games/fortunes/fortunes.dat
2018/05/01 02:13:43 Restored /tmp/fortune-teller/usr/share/games/fortunes/literature
2018/05/01 02:13:43 Restored /tmp/fortune-teller/usr/share/games/fortunes/literature.dat
2018/05/01 02:13:43 Restored /tmp/fortune-teller/usr/share/games/fortunes/riddles
2018/05/01 02:13:43 Restored /tmp/fortune-teller/usr/share/games/fortunes/riddles.dat
2018/05/01 02:13:43 Restored /tmp/fortune-teller/usr/share/games/fortunes/fortunes
2018/05/01 02:13:43 Assets restored to /tmp/fortune-teller
2018/05/01 02:13:43 Listening for gRPC requests at 50051
```
