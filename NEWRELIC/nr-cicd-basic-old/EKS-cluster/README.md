# EKSCTL cluster creation

Install latest eksctl 0.17 or later minim
```
curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin
```
```
eksctl create cluster -f eksctl-cluster.yaml
```
## Jenkins serviceaccount
To configure credentials in jenkins, needed when we reate Jenkins Cloud.:
### Create Jenkins SA
```
apply -f https://raw.githubusercontent.com/jenkinsci/kubernetes-plugin/master/src/main/kubernetes/service-account.yml
```
###Get the token from the k8s secret created for jenkins
```
k describe secret jenkins-token-j9q4l
```

# With windows nodes
https://docs.aws.amazon.com/eks/latest/userguide/windows-support.html

```
eksctl create nodegroup --config-file=eksctl-cluster.yaml
eksctl delete nodegroup --cluster=EKS-cluster-joan-porta --name=windows-ng
eksctl delete cluster --name EKS-cluster-joan-porta
```
# For Windows nodes
```
eksctl utils install-vpc-controllers --cluster EKS-cluster-joan-porta --approve
```
# Create IAM user with admn access to EKS

Add user in config-map aus-auth following this doc:
https://docs.aws.amazon.com/eks/latest/userguide/add-user-role.html
