## Pulumi GitHub Repository

This Pulumi program manages all of my future GitHub repositories.

see [GitHub Provider](https://www.pulumi.com/registry/packages/github/) for details on the GitHub provider.

### How to export existing GitHub repositories

I keep `<resourcename in pulumi>` and `<repo-name>` the same.

```bash 
pulumi import github:index/repository:Repository <resourcename in pulumi> <repo-name> --protect=false
```