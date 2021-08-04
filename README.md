# kube_envoy_test
Testing load balancing with envoy and kubernetes

minikube start
eval $(minikube docker-env)

build myapp & build envoy

. kube_create.sh

curl $(minikube service --url myapp-envoy)
-> no healthy upstream ???
