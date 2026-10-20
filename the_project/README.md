# The Project

App created with Golang.

To deploy it into the K8s cluster execute the following commands:

```
kubectl apply -f ./the_project/k8s/deployment.yaml
```

If you want to request the logs then:

```
kubectl logs deployment/the-project
```
