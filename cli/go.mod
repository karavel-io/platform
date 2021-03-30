module github.com/mikamai/karavel/cli

go 1.15

replace github.com/mikamai/karavel => ../

require (
	github.com/fatih/color v1.7.0
	github.com/go-git/go-billy/v5 v5.0.0
	github.com/go-git/go-git/v5 v5.2.0
	github.com/hashicorp/hcl/v2 v2.8.2
	github.com/mittwald/go-helm-client v0.4.2
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/stretchr/testify v1.6.1
	github.com/tidwall/gjson v1.6.8
	github.com/tidwall/sjson v1.1.5
	github.com/zclconf/go-cty v1.2.0
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
	helm.sh/helm/v3 v3.5.1
	modernc.org/sortutil v1.1.0
	sigs.k8s.io/kustomize/api v0.7.2
)
