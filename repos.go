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
				License:           APACHE_2_0,
				GitIgnoreTemplate: "Go",
				Visibility:        VISIBILITY_PUBLIC,
				Collaborators:     []string{"ediri"},
			},
			{
				Name:        ".github",
				Description: "Repository for default files",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
			},
			{
				Name:        "test",
				Description: "Test Repository",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PRIVATE,
				Labels:      lables,
			},
			{
				Name:          "gophercon-turkey-2021",
				Description:   "My GopherCon Turkey 2021 Demo repo",
				License:       APACHE_2_0,
				Visibility:    VISIBILITY_PUBLIC,
				Issues:        true,
				Collaborators: []string{"ediri"},
			},
			{
				Name:        "k3sair-cli",
				Description: "K3SAIR 🏴‍☠️️ ('Corsair') is a cli for the installation of k3s in an Air-Gapped environment.",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Wiki:        true,
				Topics:      []string{"ssh", "golang", "cli", "k3s"},
				GHPages:     "gh-pages",
			},
			{
				Name:        "pulumi-github",
				Description: "Pulumi program to manage my GitHub repositories",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"pulumi", "github", "iac"},
				Labels:      lables,
			},
			{
				Name:        "tfu",
				Description: "tfu is a terraform helper to update the providers.",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Wiki:        true,
				Topics:      []string{"infrastructure-as-code", "terraform", "iac"},
				Labels:      lables,
			},
			{
				Name:        "digitalocean-kubernetes-challenge",
				Description: "Deploy a GitOps CI/CD implementation",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Archived:    true,
				Topics:      []string{"kubernetes", "ci-cd", "gitops", "pulumi", "tekton"},
			},
			{
				Name:        "quick-bites",
				Description: "Quick Bites of different technologies",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"quick-bites"},
			},
			{
				Name:        "vcluster-webinar",
				Description: "Demo Repository for vcluster Webinar",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"vcluster", "kubernetes"},
			},
			{
				Name:        "gcp-gke-iac",
				Description: "Playing around with GCP",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"gcp", "kubernetes"},
			},
			{
				Name:        "pulumi-civo-flux-bucket",
				Description: "Playing around with Pulumi, Civo and Flux with S3 Buckets",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"civo", "kubernetes", "redis", "flux"},
			},
			{
				Name:        "azure-devops-tools-extension",
				Description: "Some missing plugins",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Wiki:        true,
				Topics:      []string{"ado", "azure", "tools"},
			},
			{
				Name:        "pulumi-do-flux-webhooks-kcert",
				Description: "How to create a flux webhook with kcert tls",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"digitalocean", "kubernetes", "flux"},
			},
			{
				Name:        "rancher-argocd-plugins",
				Description: "How to use plugins in ArgoCd",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"rancher", "kubernetes", "argocd"},
			},
			{
				Name:        "pulumi-gke-flux-jspolicy",
				Description: "How to use jsPolicy",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"pulumi", "gke", "flux", "kubernetes", "jspolicy"},
			},
			{
				Name:        "devcontainer-iac-kubernetes",
				Description: "How to create and use a devcontainer",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PRIVATE,
				Topics:      []string{"devcontainer", "kubernetes", "iac"},
			},
			{
				Name:        "pulumi-java-azure-kubernetes",
				Description: "Playing with the new Pulumi Java",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PRIVATE,
				Topics:      []string{"pulumi", "java", "azure"},
			},
			{
				Name:        "pulumi-azure-flux-weave-ui",
				Description: "Playing around with the Weave GitOps UI",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"pulumi", "kubernetes", "azure", "gitops", "weave"},
			},
			{
				Name:        "pulumi-civo-kubecon",
				Description: "Playing around with the Civo Valencia Region",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Archived:    true,
				Topics:      []string{"pulumi", "civo", "valencia", "gitops"},
			},
			{
				Name:        "kubernetes-updatecli",
				Description: "Playing Updatecli and Kubernetes",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"kubernetes", "updatecli"},
			},
			{
				Name:        "kubernetes-diy-policy-engine",
				Description: "How to build your own policy engine",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"kubernetes", "diy", "policy-engine", "admission-controller"},
			},
			{
				Name:        "lazypulumi",
				Description: "Pulumi for lazy people",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PRIVATE,
				Issues:      true,
				Topics:      []string{"pulumi", "tui"},
			},
			{
				Name:        "minectl-sdk",
				Description: "SDK for every minectl product",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Topics:      []string{"minectl", "sdk", "golang", "minecraft"},
			},
			{
				Name:        "minectl-operator",
				Description: "Kubernetes operator for minectl",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Topics:      []string{"minectl", "operator", "kubernetes", "golang", "minecraft"},
			},
			{
				Name:        "rust-jreleaser",
				Description: "Playing around with Rust and JReleaser",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      false,
				Topics:      []string{"rust", "jreleaser"},
			},
			{
				Name:        "trivy-plugin-ui",
				Description: "Simple Trivy UI plugin written in Rust",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Topics:      []string{"trivy", "plugin", "ui", "rust"},
			},
			{
				Name:        "rust-cli",
				Description: "Playing around with Rust and Clap",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      false,
				Topics:      []string{"rust", "clap"},
			},
			{
				Name:        "silly-helm-chart",
				Description: "A silly Helm chart with some kubernetes resources to play with",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Topics:      []string{"helm", "kubernetes"},
				GHPages:     "gh-pages",
			},
			{
				Name:        "pulumi-scala-minecraft",
				Description: "Playing around with Pulumi, Scala and Minecraft",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      false,
				Topics:      []string{"pulumi", "scala", "minecraft"},
			},
			{
				Name:        "my-flux-demo-deployments",
				Description: "Playing around with Flux",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Topics:      []string{"flux", "kubernetes"},
			},
			{
				Name:        "rust-grpc",
				Description: "Playing around with Rust and GRPC",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      false,
				Topics:      []string{"rust", "grpc"},
			},
			{
				Name:        "hello-server",
				Description: "Simple Server written in different languages and provided as container image",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      false,
				Topics:      []string{"hello", "server", "rust", "golang", "java", "kotlin", "python", "nodejs", "docker"},
			},
			{
				Name:        "pulumi-vultr",
				Description: "Pulumi provider for Vultr",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Topics:      []string{"pulumi", "vultr"},
			},
			{
				Name:        "gpt3ops-rs",
				Description: "A Rust based Kubernetes operator to ue GPT-3 to explain Kubernetes events",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PRIVATE,
				Issues:      true,
				Topics:      []string{"rust", "kubernetes", "operator", "gpt-3"},
			},
			{
				Name:                "civo-production-ready-kubernetes",
				Description:         "The repository for the CIVO Navigate talk: How To Build A Production Ready Kubernetes",
				License:             APACHE_2_0,
				Visibility:          VISIBILITY_PUBLIC,
				Issues:              false,
				VulnerabilityAlerts: true,
				Topics:              []string{"kubernetes", "civo"},
			},
			{
				Name:        "pulumi-ms-teams-webhook",
				Description: "Webhook for Microsoft Teams support in Pulumi Service",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Topics:      []string{"pulumi", "microsoft-teams", "webhook"},
			},
			{
				Name:        "pulumi-azapi",
				Description: "Pulumi provider for Azure API Management",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Topics:      []string{"pulumi", "azure", "azapi"},
			},
			{
				Name:        "pulumi-port-labs",
				Description: "Pulumi provider for Port Labs",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Topics:      []string{"pulumi", "port-labs"},
			},
			{
				Name:        "parlez-vous-pulumi",
				Description: "Creating an EKS cluster with Pulumi in different programming languages",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Issues:      true,
				Topics:      []string{"pulumi", "eks", "kubernetes", "golang", "typescript", "python", "csharp", "java"},
			},
			{
				Name:        "pulumi-scaleway-example",
				Description: "Example of a Pulumi program to manage a Scaleway server",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"pulumi", "scaleway", "kubernetes"},
			},
			{
				Name:        "pulumi-scaleway",
				Description: "Pulumi provider for Scaleway",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"pulumi", "scaleway"},
			},
			{
				Name:        "pulumi-aks-wasm-spin",
				Description: "Pulumi program to create an AKS cluster with WASM support and Fermyon Spin",
				License:     APACHE_2_0,
				Visibility:  VISIBILITY_PUBLIC,
				Topics:      []string{"pulumi", "aks", "wasm", "spin", "fermyon", "kubernetes"},
			},
		},
	}
)
