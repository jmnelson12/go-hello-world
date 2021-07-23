# Hello World

## What This Does

- [x] uses go to create an http server
  - [x] run `GET` request on `localhost:30002` to get back `{ status: ok, message: "hello world" }`
- [x] create a docker image
- [x] run docker image through k8s
- [x] be able to access endpoint locally in browser

## TODO

- need to figure out ports and how to set them up through env var

### Docker Setup

```bash
$ docker build --tag jmnelson12/go-hello-world .
$ docker run -d -p 8080:8080 --name go-hello-world jmnelson12/go-hello-world
```

### K8S Setup

https://medium.com/swlh/how-to-run-locally-built-docker-images-in-kubernetes-b28fbc32cc1d

```bash
$ minikube start

# so I can connect to local docker image
$ minikube docker-env
$ eval $(minikube -p minikube docker-env)
$ docker build --tag jmnelson12/go-hello-world .

$ kubectl delete -f deploy/hello-world/helloworld.yaml
$ kubectl create -f deploy/hello-world/helloworld.yaml

$ minikube service hello-world-service
```
