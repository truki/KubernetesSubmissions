# Log output app

App create in Golang deploy with:

## 1.1.
```
kubectl create deployment log-output --image=trukise/log_output:1.0
```

Logs:

```
➜  log_output git:(main) ✗ kubectl logs deployments.apps/log-output -f
2026-10-20T10:00:06.525Z: 3b32639c-0ddf-4282-b535-4820429514c3
2026-10-20T10:00:11.526Z: 3b32639c-0ddf-4282-b535-4820429514c3
2026-10-20T10:00:16.527Z: 3b32639c-0ddf-4282-b535-4820429514c3
2026-10-20T10:00:21.528Z: 3b32639c-0ddf-4282-b535-4820429514c3
2026-10-20T10:00:26.533Z: 3b32639c-0ddf-4282-b535-4820429514c3
2026-10-20T10:00:31.537Z: 3b32639c-0ddf-4282-b535-4820429514c3
``` 

## 1.3
Now that we have define via a yaml file the app, thw way to deploy it is:

```
kubectl apply -f ./log_output/k8s/deployment.yaml
```

to check logs:

```
kubectl logs deployment/hashgenerator-dep -f
```

