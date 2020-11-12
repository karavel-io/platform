# Changelog

This file documents all notable changes to [ingress-nginx](https://github.com/kubernetes/ingress-nginx) Helm Chart. The release numbering uses [semantic versioning](http://semver.org).

### Unreleased

### 3.9.0

- [X] [#6423](https://github.com/kubernetes/ingress-nginx/pull/6423) Add Default backend HPA autoscaling

### 3.8.0

- [X] [#6395](https://github.com/kubernetes/ingress-nginx/pull/6395) Update jettech/kube-webhook-certgen image
- [X] [#6377](https://github.com/kubernetes/ingress-nginx/pull/6377) Added loadBalancerSourceRanges for internal lbs
- [X] [#6356](https://github.com/kubernetes/ingress-nginx/pull/6356) Add securitycontext settings on defaultbackend
- [X] [#6401](https://github.com/kubernetes/ingress-nginx/pull/6401) Fix controller service annotations
- [X] [#6403](https://github.com/kubernetes/ingress-nginx/pull/6403) Initial helm chart changelog

### 3.7.1

- [X] [#6326](https://github.com/kubernetes/ingress-nginx/pull/6326) Fix liveness and readiness probe path in daemonset chart

### 3.7.0

- [X] [#6316](https://github.com/kubernetes/ingress-nginx/pull/6316) Numerals in podAnnotations in quotes [#6315](https://github.com/kubernetes/ingress-nginx/issues/6315)

### 3.6.0

- [X] [#6305](https://github.com/kubernetes/ingress-nginx/pull/6305) Add default linux nodeSelector

### 3.5.1

- [X] [#6299](https://github.com/kubernetes/ingress-nginx/pull/6299) Fix helm chart release

### 3.5.0

- [X] [#6260](https://github.com/kubernetes/ingress-nginx/pull/6260) Allow Helm Chart to customize admission webhook's annotations, timeoutSeconds, namespaceSelector, objectSelector and cert files locations

### 3.4.0

- [X] [#6268](https://github.com/kubernetes/ingress-nginx/pull/6268) Update to 0.40.2 in helm chart #6288

### 3.3.1

- [X] [#6259](https://github.com/kubernetes/ingress-nginx/pull/6259) Release helm chart
- [X] [#6258](https://github.com/kubernetes/ingress-nginx/pull/6258) Fix chart markdown link
- [X] [#6253](https://github.com/kubernetes/ingress-nginx/pull/6253) Release v0.40.0

### 3.3.1

- [X] [#6233](https://github.com/kubernetes/ingress-nginx/pull/6233) Add admission controller e2e test

### 3.3.0

- [X] [#6203](https://github.com/kubernetes/ingress-nginx/pull/6203) Refactor parsing of key values
- [X] [#6162](https://github.com/kubernetes/ingress-nginx/pull/6162) Add helm chart options to expose metrics service as NodePort
- [X] [#6180](https://github.com/kubernetes/ingress-nginx/pull/6180) Fix helm chart admissionReviewVersions regression
- [X] [#6169](https://github.com/kubernetes/ingress-nginx/pull/6169) Fix Typo in example prometheus rules

### 3.0.0

- [X] [#6167](https://github.com/kubernetes/ingress-nginx/pull/6167) Update chart requirements

### 2.16.0

- [X] [#6154](https://github.com/kubernetes/ingress-nginx/pull/6154) add `topologySpreadConstraint` to controller

### 2.15.0

- [X] [#6087](https://github.com/kubernetes/ingress-nginx/pull/6087) Adding parameter for externalTrafficPolicy in internal controller service spec

### 2.14.0

- [X] [#6104](https://github.com/kubernetes/ingress-nginx/pull/6104) Misc fixes for nginx-ingress chart for better keel and prometheus-operator integration

### 2.13.0

- [X] [#6093](https://github.com/kubernetes/ingress-nginx/pull/6093) Release v0.35.0

### 2.13.0

- [X] [#6093](https://github.com/kubernetes/ingress-nginx/pull/6093) Release v0.35.0
- [X] [#6080](https://github.com/kubernetes/ingress-nginx/pull/6080) Switch images to k8s.gcr.io after Vanity Domain Flip

### 2.12.1

- [X] [#6075](https://github.com/kubernetes/ingress-nginx/pull/6075) Sync helm chart affinity examples

### 2.12.0

- [X] [#6039](https://github.com/kubernetes/ingress-nginx/pull/6039) Add configurable serviceMonitor metricRelabelling and targetLabels
- [X] [#6044](https://github.com/kubernetes/ingress-nginx/pull/6044) Fix YAML linting

### 2.11.3

- [X] [#6038](https://github.com/kubernetes/ingress-nginx/pull/6038) Bump chart version PATCH

### 2.11.2

- [X] [#5951](https://github.com/kubernetes/ingress-nginx/pull/5951) Bump chart patch version

### 2.11.1

- [X] [#5900](https://github.com/kubernetes/ingress-nginx/pull/5900) Release helm chart for v0.34.1

### 2.11.0

- [X] [#5879](https://github.com/kubernetes/ingress-nginx/pull/5879) Update helm chart for v0.34.0
- [X] [#5671](https://github.com/kubernetes/ingress-nginx/pull/5671) Make liveness probe more fault tolerant than readiness probe

### 2.10.0

- [X] [#5843](https://github.com/kubernetes/ingress-nginx/pull/5843) Update jettech/kube-webhook-certgen image

### 2.9.1

- [X] [#5823](https://github.com/kubernetes/ingress-nginx/pull/5823) Add quoting to sysctls because numeric values need to be presented as strings (#5823)

### 2.9.0

- [X] [#5795](https://github.com/kubernetes/ingress-nginx/pull/5795) Use fully qualified images to avoid cri-o issues


### TODO

Keep building the changelog using *git log charts* checking the tag
