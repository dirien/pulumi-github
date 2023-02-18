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
				Visibility:  "public",
				Archived:    true,
				Topics:      []string{"kubernetes", "ci-cd", "gitops", "pulumi", "tekton"},
			},
			{
				Name:        "quick-bites",
				Description: "Quick Bites of different technologies",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"quick-bites"},
			},
			{
				Name:        "pulumi-scaleway",
				Description: "Example of a Pulumi program to manage a Scaleway server",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"pulumi", "scaleway", "kubernetes"},
			},
			{
				Name:        "vcluster-webinar",
				Description: "Demo Repository for vcluster Webinar",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"vcluster", "kubernetes"},
			},
			{
				Name:        "gcp-gke-iac",
				Description: "Playing around with GCP",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"gcp", "kubernetes"},
			},
			{
				Name:        "pulumi-civo-flux-bucket",
				Description: "Playing around with Pulumi, Civo and Flux with S3 Buckets",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"civo", "kubernetes", "redis", "flux"},
			},
			{
				Name:        "azure-devops-tools-extension",
				Description: "Some missing plugins",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      true,
				Wiki:        true,
				Topics:      []string{"ado", "azure", "tools"},
			},
			{
				Name:        "pulumi-do-flux-webhooks-kcert",
				Description: "How to create a flux webhook with kcert tls",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"digitalocean", "kubernetes", "flux"},
			},
			{
				Name:        "rancher-argocd-plugins",
				Description: "How to use plugins in ArgoCd",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"rancher", "kubernetes", "argocd"},
			},
			{
				Name:        "pulumi-gke-flux-jspolicy",
				Description: "How to use jsPolicy",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"pulumi", "gke", "flux", "kubernetes", "jspolicy"},
			},
			{
				Name:        "devcontainer-iac-kubernetes",
				Description: "How to create and use a devcontainer",
				License:     "apache-2.0",
				Visibility:  "private",
				Topics:      []string{"devcontainer", "kubernetes", "iac"},
			},
			{
				Name:        "pulumi-java-azure-kubernetes",
				Description: "Playing with the new Pulumi Java",
				License:     "apache-2.0",
				Visibility:  "private",
				Topics:      []string{"pulumi", "java", "azure"},
			},
			{
				Name:        "pulumi-azure-flux-weave-ui",
				Description: "Playing around with the Weave GitOps UI",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"pulumi", "kubernetes", "azure", "gitops", "weave"},
			},
			{
				Name:        "pulumi-civo-kubecon",
				Description: "Playing around with the Civo Valencia Region",
				License:     "apache-2.0",
				Visibility:  "public",
				Archived:    true,
				Topics:      []string{"pulumi", "civo", "valencia", "gitops"},
			},
			{
				Name:        "kubernetes-updatecli",
				Description: "Playing Updatecli and Kubernetes",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"kubernetes", "updatecli"},
			},
			{
				Name:        "kubernetes-diy-policy-engine",
				Description: "How to build your own policy engine",
				License:     "apache-2.0",
				Visibility:  "public",
				Topics:      []string{"kubernetes", "diy", "policy-engine", "admission-controller"},
			},
			{
				Name:        "lazypulumi",
				Description: "Pulumi for lazy people",
				License:     "apache-2.0",
				Visibility:  "private",
				Issues:      true,
				Topics:      []string{"pulumi", "tui"},
			},
			{
				Name:        "minectl-sdk",
				Description: "SDK for every minectl product",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      true,
				Topics:      []string{"minectl", "sdk", "golang", "minecraft"},
			},
			{
				Name:        "minectl-operator",
				Description: "Kubernetes operator for minectl",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      true,
				Topics:      []string{"minectl", "operator", "kubernetes", "golang", "minecraft"},
			},
			{
				Name:        "rust-jreleaser",
				Description: "Playing around with Rust and JReleaser",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      false,
				Topics:      []string{"rust", "jreleaser"},
			},
			{
				Name:        "trivy-plugin-ui",
				Description: "Simple Trivy UI plugin written in Rust",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      true,
				Topics:      []string{"trivy", "plugin", "ui", "rust"},
			},
			{
				Name:        "rust-cli",
				Description: "Playing around with Rust and Clap",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      false,
				Topics:      []string{"rust", "clap"},
			},
			{
				Name:        "silly-helm-chart",
				Description: "A silly Helm chart with some kubernetes resources to play with",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      true,
				Topics:      []string{"helm", "kubernetes"},
			},
			{
				Name:        "pulumi-scala-minecraft",
				Description: "Playing around with Pulumi, Scala and Minecraft",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      false,
				Topics:      []string{"pulumi", "scala", "minecraft"},
			},
			{
				Name:        "my-flux-demo-deployments",
				Description: "Playing around with Flux",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      true,
				Topics:      []string{"flux", "kubernetes"},
			},
			{
				Name:        "rust-grpc",
				Description: "Playing around with Rust and GRPC",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      false,
				Topics:      []string{"rust", "grpc"},
			},
			{
				Name:        "hello-server",
				Description: "Simple Server written in different languages and provided as container image",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      false,
				Topics:      []string{"hello", "server", "rust", "golang", "java", "kotlin", "python", "nodejs", "docker"},
			},
			{
				Name:        "pulumi-vultr",
				Description: "Pulumi provider for Vultr",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      true,
				Topics:      []string{"pulumi", "vultr"},
			},
			{
				Name:        "gpt3ops-rs",
				Description: "A Rust based Kubernetes operator to ue GPT-3 to explain Kubernetes events",
				License:     "apache-2.0",
				Visibility:  "private",
				Issues:      true,
				Topics:      []string{"rust", "kubernetes", "operator", "gpt-3"},
			},
			{
				Name:                "civo-production-ready-kubernetes",
				Description:         "The repository for the CIVO Navigate talk: How To Build A Production Ready Kubernetes",
				License:             "apache-2.0",
				Visibility:          "public",
				Issues:              false,
				VulnerabilityAlerts: true,
				Topics:              []string{"kubernetes", "civo"},
			},
			{
				Name:        "pulumi-ms-teams-webhook",
				Description: "Webhook for Microsoft Teams support in Pulumi Service",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      true,
				Topics:      []string{"pulumi", "microsoft-teams", "webhook"},
			},
			{
				Name:        "pulumi-azapi",
				Description: "Pulumi provider for Azure API Management",
				License:     "apache-2.0",
				Visibility:  "public",
				Issues:      true,
				Topics:      []string{"pulumi", "azure", "azapi"},
			},
		},
	}
)
