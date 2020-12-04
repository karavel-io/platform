# Ingress Controllers components

Karavel provides various [Kubernetes Ingress Controllers] that can be installed and configured.
As Kubernetes supports having multiple controllers installed via the [ingress class annotation], operators
can choose to add multiple Karavel components in parallel. Each will be installed in a dedicated namespace.

The NGINX controller is mandatory as many other components rely on it, and it is one of the most
supported and integrated by other open source projects.

## Components

- [NGINX]  
  Implementation using [NGINX](https://nginx.com) as the underlying HTTP server  
  Kubernetes Ingress Class:
  ```yaml
  annotations:
    kubernetes.io/ingress.class: "nginx"
  ```  
  KBT Import:
  ```yaml
  import_role:
    name: mikamai.karavel.ingress_nginx
  ```
  [config](./variables.md#nginx-ingress-controller)
  

[Kubernetes Ingress Controllers]: https://kubernetes.io/docs/concepts/services-networking/ingress-controllers
[ingress class annotation]: https://kubernetes.github.io/ingress-nginx/user-guide/multiple-ingress/
[NGINX]: https://kubernetes.github.io/ingress-nginx
