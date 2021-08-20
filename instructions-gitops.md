
`TODO`

# Instructions (with GitOps)

## Preparation

git clone git@github.com:bygui86/workshop-tage-2021.git

cd workshop-tage-2021/

### Kubernetes cluster

### 1. GKE cluster

`TODO`

### 2. CIVO cluster

`TODO`

### 3. Minikube

make start-docker-minikube

make start-hyperkit-minikube

make start-virtualbox-minikube (slower in provisioning and container image download)

---

## Infrastructure

cd infrastructure/

### 1. GitOps

kustomize build 10_devops/10_argocd/ | kubectl apply -f -

### 2. Namespaces


---

## Applications
