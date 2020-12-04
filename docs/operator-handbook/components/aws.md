# AWS components

The [Karavel Bootstrap Tool] includes a `core_aws` role that will provide
core infrastructure addons necessary when running clusters on AWS EC2 infrastructure. This include both 
self-managed clusters on EC2 and [AWS EKS] managed clusters.

## KBT Import

```yaml
import_role:
  name: mikamai.karavel.core_aws
```

## Components

- [AWS Node Termination Handler]  
  Handles cordoning and draining of nodes when the underlying EC2 istance is going to
  be rebooted or terminated  
  [config](./variables.md#aws-node-termination-handler)

[Karavel Bootstrap Tool]: ./bootstrap.md
[AWS EKS]: https://aws.amazon.com/eks
[AWS Node Termination Handler]: https://github.com/aws/aws-node-termination-handler 
