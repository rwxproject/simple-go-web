# Simple Go Web App

```bash
#!/bin/bash
kubectl apply -f simple-go-web-deployment.yml
kubectl apply -f simple-go-web-service.yml

kubectl port-forward <POD> 9090:9090
```