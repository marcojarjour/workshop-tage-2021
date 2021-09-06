
# Delete Kubernetes cluster

## CIVO

```bash
# Set region
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

## GCP

```bash
# Set project
export PROJECT_ID=...
gcloud config set project $PROJECT_ID

# Delete cluster
gcloud beta container --project "$PROJECT_ID" clusters delete "wt21" --zone "europe-west4-a"

# Check
gcloud container clusters list
```

## LOCAL (MINIKUBE)

```bash
minikube stop && minikube delete
```
