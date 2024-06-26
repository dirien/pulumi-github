package main

import (
	"fmt"
	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

const (
	VISIBILITY_PUBLIC  = "public"
	VISIBILITY_PRIVATE = "private"

	APACHE_2_0 = "apache-2.0"
	MPL_2_0    = "mpl-2.0"
)

func createGitHubRepository(ctx *pulumi.Context, create *Repository) error {

	repoArgs := &github.RepositoryArgs{
		Description:         pulumi.String(create.Description),
		Name:                pulumi.String(create.Name),
		LicenseTemplate:     pulumi.String(create.License),
		AllowSquashMerge:    pulumi.Bool(true),
		AllowRebaseMerge:    pulumi.Bool(false),
		AllowMergeCommit:    pulumi.Bool(false),
		DeleteBranchOnMerge: pulumi.Bool(true),
		HasIssues:           pulumi.Bool(create.Issues),
		HasWiki:             pulumi.Bool(create.Wiki),
		HasProjects:         pulumi.Bool(create.Projects),
		VulnerabilityAlerts: pulumi.Bool(create.VulnerabilityAlerts),
		HasDownloads:        pulumi.Bool(create.Downloads),
		Visibility:          pulumi.String(create.Visibility),
		Archived:            pulumi.Bool(create.Archived),
	}

	if len(create.Topics) > 0 {
		repoArgs.Topics = pulumi.ToStringArray(create.Topics)
	}

	if create.PagesPath != "" && create.PagesBranch != "" {
		repoArgs.Pages = &github.RepositoryPagesArgs{
			Source: &github.RepositoryPagesSourceArgs{
				Branch: pulumi.String(create.PagesBranch),
				Path:   pulumi.String(create.PagesPath),
			},
		}
	}

	if create.GitIgnoreTemplate != "" {
		repoArgs.GitignoreTemplate = pulumi.String(create.GitIgnoreTemplate)
	}
	repository, err := github.NewRepository(ctx, create.Name, repoArgs)
	if err != nil {
		return err
	}

	if create.Labels {
		var labels []Label
		config.RequireObject(ctx, "labels", &labels)
		for _, label := range labels {
			_, err := github.NewIssueLabel(ctx, fmt.Sprintf("label-%s-%s", create.Name, label.Name), &github.IssueLabelArgs{
				Color:       pulumi.String(label.Color),
				Description: pulumi.String(label.Description),
				Name:        pulumi.String(label.Name),
				Repository:  repository.Name,
			})
			if err != nil {
				return err
			}
		}
	}

	if len(create.Collaborators) > 0 {
		for _, collaborator := range create.Collaborators {
			_, err := github.NewRepositoryCollaborator(ctx, fmt.Sprintf("collab-%s", create.Name), &github.RepositoryCollaboratorArgs{
				Repository: repository.Name,
				Username:   pulumi.String(collaborator),
			})
			if err != nil {
				return err
			}
		}
	}

	if create.Visibility != "private" {
		_, err = github.NewBranchProtection(ctx, fmt.Sprintf("branchProtection-%s", create.Name), &github.BranchProtectionArgs{
			RepositoryId:         repository.NodeId,
			Pattern:              pulumi.String("main"),
			RequireSignedCommits: pulumi.Bool(true),
			AllowsForcePushes:    pulumi.Bool(create.AllowsForcePushes),
		})
		if err != nil {
			return err
		}
	}

	ctx.Export(fmt.Sprintf("clone %s", create.Name), repository.HttpCloneUrl)
	return nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		var repos []Repository
		config.RequireObject(ctx, "repositories", &repos)
		for _, repo := range repos {
			err := createGitHubRepository(ctx, &repo)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
