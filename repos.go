package main

type Repository struct {
	Name                string
	Description         string
	License             string
	GitIgnoreTemplate   string
	Visibility          string
	Import              bool
	Issues              bool
	Wiki                bool
	Collaborators       []string
	Topics              []string
	GHPages             string `yaml:"ghpages"`
	Projects            bool
	VulnerabilityAlerts bool
	Downloads           bool
	Labels              bool
	Archived            bool
	AllowsForcePushes   bool
}

type Label struct {
	Name        string
	Description string
	Color       string
}
type Repositories struct {
	Items []Repository
}
