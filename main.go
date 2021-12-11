package main

import (
	"github.com/pulumi/pulumi-github/sdk/v4/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		octopusDeployHackathon, err := github.NewRepository(ctx, "octopus-deploy-hackathon", &github.RepositoryArgs{
			Description:         pulumi.String("Hackathon Repository for Octopus Deploy"),
			Name:                pulumi.String("octopus-deploy-hackathon"),
			LicenseTemplate:     pulumi.String("apache-2.0"),
			AllowSquashMerge:    pulumi.Bool(true),
			AllowRebaseMerge:    pulumi.Bool(false),
			AllowMergeCommit:    pulumi.Bool(false),
			DeleteBranchOnMerge: pulumi.Bool(true),
			GitignoreTemplate:   pulumi.String("Go"),
		})
		if err != nil {
			return err
		}

		_, err = github.NewBranchProtection(ctx, "branchProtection", &github.BranchProtectionArgs{
			RepositoryId:         octopusDeployHackathon.NodeId,
			Pattern:              pulumi.String("main"),
			RequireSignedCommits: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		ctx.Export("repository", octopusDeployHackathon.Name)
		return nil
	})
}
