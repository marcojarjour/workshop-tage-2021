
# Argo CD - SBT repositories and users

`TODO` put together to ../30_argocd/3_configurations and ../30_argocd/3_secrets

## Users

| Name  | Type (local or SSO) | Status   |
|-------|---------------------|----------|
| andy  | local               | disabled |
| adam  | local               | disabled |
| joerg | local               | disabled |
| paolo | local               | disabled |

## Repositories

| Repo        | URL                                                                                  |
|-------------|--------------------------------------------------------------------------------------|
| cicd-devops | https://gitlab.com/sbt-devops/ci-cd/cicd-devops.git                                  |
| devops      | https://gitlab.com/sbt-devops/devops.git                                             |
| die-devops  | https://gitlab.com/sbt-data-ingestion-engine/die-devops.git                          |
| hdp-devops  | https://gitlab.com/sbt-data-ingestion-engine/historical-data-provider/hdp-devops.git |
| mst-devops  | https://gitlab.com/sbt-data-ingestion-engine/market-state/mst-devops.git             |

## Instructions

```bash
kustomize build . | kubectl apply -f -
```
