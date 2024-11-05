# start k8s with mikikube

by <https://habr.com/ru/companies/avito/articles/820953/>

install minikube

## Minibube config

```bash
eval $(minikube -p minikube docker-env)
```

check mikibube images by: docker images

## Build images

```bash
docker build -t kubeapp1:v1 kubeapp1/.
docker build -t kubeapp2:v1 kubeapp2/.
```

## Add NodePort

```bash
k expose --type=NodePort -n app2 deployment/kubeapp2
```

```bash
minikube service kubeapp1 -n app1 --url | curl /hello
```
