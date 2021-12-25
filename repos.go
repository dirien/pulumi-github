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
	Labels              []Label
}

type Label struct {
	Name        string
	Description string
	Color       string
}
type Repositories struct {
	Items []Repository
}

var (
	lables = []Label{
		{
			Name:        ":bug: bug",
			Description: "Something isn't working",
			Color:       "f29513",
		},
		{
			Name:        ":dancers: duplicate",
			Description: "This issue or pull request already exists",
			Color:       "fef2c0",
		},
		{
			Name:        ":sparkles: enhancement",
			Description: "A new feature",
			Color:       "7fc97f",
		},
		{
			Name:        ":pray: help wanted",
			Description: "Extra attention is needed",
			Color:       "ffd92f",
		},
		{
			Name:        ":no_entry_sign: invalid",
			Description: "The issue or pull request is invalid",
			Color:       "d7191c",
		},
		{
			Name:        ":question: question",
			Description: "Further information is requested",
			Color:       "fdae61",
		},
		{
			Name:        ":coffin: wontfix",
			Description: "This will not be worked on",
			Color:       "d73027",
		},
		{
			Name:        ":hatching_chick: good first issue",
			Description: "Good for newcomers",
			Color:       "1b9e77",
		},
		{
			Name:        ":memo: documentation",
			Color:       "c5def5",
			Description: "Documentation is missing or needs to be improved",
		},
		{
			Name:        "dependencies",
			Color:       "a6cee3",
			Description: "Dependencies need to be updated",
		},
	}

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
				Labels:      lables,
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
			{
				Name:        "pulumi-github",
				Description: "Pulumi program to manage my GitHub repositories",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"pulumi", "github", "iac"},
				Labels:      lables,
			},
			{
				Name:        "tfu",
				Description: "tfu is a terraform helper to update the providers.",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      true,
				Wiki:        true,
				Topics:      []string{"infrastructure-as-code", "terraform", "iac"},
				Labels:      lables,
			},
			{
				Name:        "digitalocean-kubernetes-challenge",
				Description: "Deploy a GitOps CI/CD implementation",
				License:     "apache-2.0",
				Visibility:  "private",
				Topics:      []string{"kubernetes", "ci-cd", "gitops", "pulumi", "tekton"},
			},
		},
	}
)
