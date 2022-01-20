package interfaces

type DockerCompose struct {
	Version  string
	Services map[string]DockerService
}

type DockerService struct {
	Restart       string   `yaml:"restart,omitempty"`
	Image         string   `yaml:"image,omitempty"`
	ContainerName string   `yaml:"container_name,omitempty"`
	Volumes       []string `yaml:"volumes,omitempty"`
	Ports         []string `yaml:"ports,omitempty"`
	Environment   []string `yaml:"environment,omitempty"`
	Expose        []string `yaml:"expose,omitempty"`
	DependsOn     []string `yaml:"depends_on,omitempty"`
}
