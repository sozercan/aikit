---
title: Release Process
---

The release process is as follows:

- Tag the `main` branch with a version number that's semver compliant (vMAJOR.MINOR.PATCH), and push the tag to GitHub.

```bash
git tag v0.1.0
git push origin v0.1.0
```

- GitHub Actions will automatically build the AIKit image and push the versioned and `latest` tag to GitHub Container Registry (GHCR) using [release action](https://github.com/sozercan/aikit/actions/workflows/release.yaml).

- Once release is done, trigger [update models](https://github.com/sozercan/aikit/actions/workflows/update-models.yaml) action to update the pre-built models.

- Mixtral 8x7b and Llama 3 70b models does not fit into GitHub runners due to their size. Trigger self-hosted [update models](https://github.com/sozercan/aikit/actions/workflows/update-models-self.yaml) action to update these pre-built models.
