
# Observe with GitOps

## Infrastructure

```shell
cd infrastructure/
```

### 1. ArgoCD for GitOps

```shell
kustomize build 10_devops/10_argocd/ | kubectl apply -f -
kustomize build 10_devops/11_defaults/ | kubectl apply -f -
```

### 2. ArgoCD UI

```shell

```

### 3. Namespaces

```shell

```

---

## Applications

`TODO`
