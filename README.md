
docker build -t kubeapp1:v1 kubeapp1/.

docker build -t kubeapp2:v1 kubeapp2/.

docker build -t kubeapp1 kubeapp1/.
docker build -t kubeapp2 kubeapp2/.

docker images
