# Simple rego policy

- warn if creating null_resource in terraform

## How to test

```
terraform init
terraform plan -refresh=false -out=plan.tfplan
terraform show -json plan.tfplan >plan.json
conftest test plan.json
```
