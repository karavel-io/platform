package argo

import "time"

// Application is a lightweight struct matching argoproj.io/v1alpha1/Application
type Application struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata"`
	Spec       ApplicationSpec `yaml:"spec"`
}

type ApplicationSpec struct {
	Source Source `yaml:"source"`
	// Destination overrides the kubernetes server and namespace defined in the environment ksonnet app.yaml
	Destination Destination `yaml:"destination"`
	Project     string      `yaml:"project"`
	SyncPolicy  SyncPolicy  `yaml:"syncPolicy,omitempty"`
}

type Source struct {
	RepoUrl string `yaml:"repoURL"`
	Path    string `yaml:"path"`
}

type Destination struct {
	Server    string `yaml:"server"`
	Namespace string `yaml:"namespace"`
}

type SyncPolicy struct {
	Automated   Automated `yaml:"automated"`
	SyncOptions []string  `yaml:"syncOptions"`
	Retry       Retry     `yaml:"retry"`
}

type Automated struct {
	Prune      bool `yaml:"prune"`
	SelfHeal   bool `yaml:"selfHeal"`
	AllowEmpty bool `yaml:"allowEmpty"`
}

type Retry struct {
	Limit   int     `yaml:"limit"`
	Backoff Backoff `yaml:"backoff"`
}

type Backoff struct {
	Duration    time.Duration `yaml:"duration"`
	Factor      int           `yaml:"factor"`
	MaxDuration time.Duration `yaml:"maxDuration"`
}

func NewApplication(name string, namespace string, argoNs string, repoUrl string, path string) Application {
	return Application{
		TypeMeta: TypeMeta{
			APIVersion: "argoproj.io/v1alpha1",
			Kind:       "Application",
		},
		ObjectMeta: ObjectMeta{
			Name:      name,
			Namespace: argoNs,
		},
		Spec: ApplicationSpec{
			Source: Source{
				RepoUrl: repoUrl,
				Path:    path,
			},
			Destination: Destination{
				Server:    "https://kubernetes.default.svc",
				Namespace: namespace,
			},
			Project: "infrastructure",
			SyncPolicy: SyncPolicy{
				Automated: Automated{
					Prune:      true,
					SelfHeal:   true,
					AllowEmpty: false,
				},
				SyncOptions: []string{
					"Validate=false",
					"CreateNamespace=true",
				},
				Retry: Retry{
					Limit: 5,
					Backoff: Backoff{
						Duration:    5 * time.Second,
						Factor:      2,
						MaxDuration: 3 * time.Minute,
					},
				},
			},
		},
	}
}
