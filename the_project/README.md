# The Project

App created with Golang.

To deploy it into the K8s cluster execute the following commands (now a deployment and a service):

```
kubectl apply -f ./the_project/k8s/
```

If you want to request the logs then:

```
kubectl logs deployment/the-project
```

Previously we have created the K3d cluster running this:

```
k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
```
