package sonarconfig

type SonarConfig struct {
	Image             string
	Labels            map[string]string
	Name              string
	Namespace         string
	NetworkPolicy     bool
	NodeSelector      string
	PodArgs           string
	PodCommand        string
	PodSecurityPolicy bool
	PodUser           int64
	Privileged        bool
}
