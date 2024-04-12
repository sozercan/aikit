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

> [!NOTE]
> At this time, Mixtral 8x7b model does not fit into GitHub runners due to its size. It is built and pushed to GHCR manually.
> ```shell
> docker buildx build . -t ghcr.io/sozercan/mixtral:8x7b \
>   -t ghcr.io/sozercan/mixtral:8x7b-instruct \
>   -f models/mixtral-7x8b-instruct.yaml \
>   --push --progress=plain --provenance=true --sbom=true
>```
>
>```shell
>docker buildx build . -t ghcr.io/sozercan/mixtral:8x7b-cuda -t ghcr.io/sozercan/mixtral:8x7b-instruct-cuda \
>   -f models/mixtral-7x8b-instruct-cuda.yaml \
>   --push --progress=plain --provenance=true --sbom=true
>```
