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
	GHPages             string
	Projects            bool
	VulnerabilityAlerts bool
	Downloads           bool
}

type Repositories struct {
	Items []Repository
}

var (
	repos = &Repositories{
		Items: []Repository{
			{
				Name:              "octopus-deploy-hackathon",
				Description:       "Hackathon Repository for Octopus Deploy",
				License:           "apache-2.0",
				GitIgnoreTemplate: "Go",
				Visibility:        "public",
				Collaborators:     []string{"ediri"},
			},
			{
				Name:        ".github",
				Description: "Repository for default files",
				License:     "apache-2.0",
				Visibility:  "public",
			},
			{
				Name:        "test",
				Description: "Test Repository",
				License:     "apache-2.0",
				Visibility:  "private",
			},
			{
				Name:          "gophercon-turkey-2021",
				Description:   "My GopherCon Turkey 2021 Demo repo",
				License:       "apache-2.0",
				Visibility:    "public",
				Issues:        true,
				Collaborators: []string{"ediri"},
			},
			{
				Name:        "k3sair-cli",
				Description: "K3SAIR 🏴‍☠️️ ('Corsair') is a cli for the installation of k3s in an Air-Gapped environment.",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      true,
				Wiki:        true,
				Topics:      []string{"ssh", "golang", "cli", "k3s"},
				GHPages:     "gh-pages",
			},
		},
	}
)
