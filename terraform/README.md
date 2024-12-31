## Prerequisites

Gcloud installed and configured.

## Commands

Run `terraform apply` to provision the GKE cluster.

Run `terraform destroy` to delete GKE cluster and associated resources.

Run `gcloud container clusters get-credentials $(terraform output -raw kubernetes_cluster_name) --region $(terraform output -raw region)` to setup kubeconfig credentials for the created GKE cluster.

## Notes 

Remember to setup the repository variable `DRY_RUN` in github according to the cluster being up or down.
