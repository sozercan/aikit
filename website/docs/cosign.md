---
title: Image Verification
---

AIKit and pre-made models are keyless signed with OIDC in GitHub Actions with [cosign](https://github.com/sigstore/cosign). You can verify the images with the following commands:

```bash
IMAGE=ghcr.io/sozercan/llama2:7b # update this with the image you want to verify
DIGEST=$(cosign triangulate ${IMAGE} --type digest)
cosign verify ${DIGEST} \
    --certificate-oidc-issuer https://token.actions.githubusercontent.com \
    --certificate-identity-regexp 'https://github\.com/sozercan/aikit/\.github/workflows/.+'
```

You should see an output similar to the following:

```bash
Verification for ghcr.io/sozercan/llama2@sha256:d47fdba491a9a47ce4911539a77e0c0a12b2e14f5beed88cb8072924b02130b4 --
The following checks were performed on each of these signatures:
  - The cosign claims were validated
  - Existence of the claims in the transparency log was verified offline
  - The code-signing certificate was verified using trusted certificate authority certificates
...
```
