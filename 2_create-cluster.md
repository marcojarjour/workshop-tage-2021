
# Create Kubernetes cluster

`⚠️ WARN` Here we assume that local environment is ready:
	- all accounts created and properly configured
	- all tools installed and properly configured

## GCP

`TODO`

```bash
# Create cluster
...

# Connect kubectl to new cluster
...

# Check kubectl
kubectl config use-context wt21
kubectl get nodes
```

## CIVO

```bash
# Set region for civo-cli
civo region current LON1

# Create cluster
civo kubernetes create wt21 --size=g3.k3s.large --nodes=3 --region=LON1 --version=v1.20.0+k3s1 --wait

# Check cluster
civo kubernetes list
civo kubernetes show wt21

# Connect kubectl to new cluster
civo kubernetes config wt21 --save --merge

# Check kubectl
kubectl config use-context wt21
kubectl get nodes
```

## LOCAL (MINIKUBE)

Please choose the best driver based on your local environment.

```bash
# Start Minikube with Docker driver
minikube start --driver=docker --cpus=8 --memory=12288 --disk-size=50g

# Start Minikube with Hyperkit driver
minikube start --driver=hyperkit --cpus=8 --memory=12288 --disk-size=50g

# Start Minikube with VirtualBox driver
minikube start --driver=virtualbox --cpus=8 --memory=12288 --disk-size=50g

# Check kubectl
kubectl get nodes
```
