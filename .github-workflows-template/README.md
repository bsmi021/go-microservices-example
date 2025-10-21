# GitHub Actions Workflows

This directory contains GitHub Actions workflow templates that need to be manually moved to `.github/workflows/` directory.

Due to security restrictions, GitHub Apps cannot directly create or modify workflow files. To enable CI/CD for this repository, please manually create the `.github/workflows/` directory and copy these files there.

## Installation Steps

1. Create the workflows directory:
   ```bash
   mkdir -p .github/workflows
   ```

2. Copy the workflow files:
   ```bash
   cp .github-workflows-template/ci.yml .github/workflows/
   cp .github-workflows-template/release.yml .github/workflows/
   ```

3. Commit and push:
   ```bash
   git add .github/workflows/
   git commit -m "Add GitHub Actions workflows"
   git push
   ```

## Workflows Included

### ci.yml
Runs on every push and pull request to main/develop branches:
- **Testing**: Runs all tests with race detection
- **Building**: Builds all three services (api, mvc, webserver)
- **Linting**: Runs golangci-lint for code quality
- **Docker**: Builds Docker images for all services
- **Coverage**: Uploads test coverage to Codecov

### release.yml
Triggers on version tags (e.g., `v1.0.0`):
- **Cross-platform builds**: Creates binaries for Linux, macOS (Intel & ARM), and Windows
- **GitHub Releases**: Automatically creates releases with built artifacts
- **Multi-architecture**: Supports amd64 and arm64 architectures

## Environment Secrets

For the workflows to function properly, add these secrets to your GitHub repository (Settings → Secrets and variables → Actions):

- `SECRET_GITHUB_ACCESS_TOKEN` - GitHub personal access token for API integration (required for API service)
- `CODECOV_TOKEN` - Codecov token for coverage reports (optional)

## Testing Locally

You can test the GitHub Actions workflows locally using [act](https://github.com/nektos/act):

```bash
# Install act
brew install act  # macOS
# or
curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash  # Linux

# Run CI workflow
act push

# Run specific job
act -j test
```

## Workflow Customization

Feel free to customize these workflows to match your team's needs:
- Adjust Go version in the setup-go action
- Modify branch triggers
- Add deployment steps
- Integrate with other services (Slack, Discord, etc.)
- Add security scanning tools

For more information on GitHub Actions, visit: https://docs.github.com/en/actions
