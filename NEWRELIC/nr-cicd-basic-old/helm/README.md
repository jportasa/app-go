```
kubectl create namespace jenkins
helm install jenkins ./jenkins-1-9.10 -n jenkins
```

```
export POD_NAME=$(kubectl get pods --namespace jenkins -l "app.kubernetes.io/component=jenkins-master" -l "app.kubernetes.io/instance=jenkins" -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace jenkins port-forward $POD_NAME 8080:8080
```
Get admin password

```
 printf $(kubectl get secret --namespace jenkins jenkins -o jsonpath="{.data.jenkins-admin-password}" | base64 --decode);echo
```

Update Jenkins plugins

```
 helm upgrade jenkins ./jenkins-1-9.10 -n jenkins
```

