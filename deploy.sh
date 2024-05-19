kubectl apply -f eventstoredb.deployment.yaml -n persistence

docker image rm secap-input:latest
docker build -t secap-input:latest .

minikube image rm secap-input:latest
minikube image load secap-input:latest

kubectl apply -f .k8s/deployment.yaml -n secap-compass