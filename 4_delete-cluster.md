
# Delete Kubernetes cluster

`WORK IN PROGRESS`

## GCP

`TODO`

## CIVO

```shell
# Set region for civo-cli
civo region current LON1

# Delete cluster
civo kubernetes remove wt21
# Expected output
Removing Kubernetes cluster wt21

# Check
civo kubernetes show wt21
# Expected output
No Kubernetes clusters found for 'wt21'. Please check your query.
```

## LOCAL (MINIKUBE)

```shell
minikube stop && minikube delete
```
