---
title: imgs.sh
description: A private docker image registry using SSH
---

> NOTICE: This is a premium [pico+](/plus) service

# Features

- 5GB image storage limit
- Private docker image registry
- Use SSH keys for authentication
- `docker push` and `docker pull` work exactly the same

# How it works

We leverage web tunnels to open a portal to our private image registry where you
can access it via `localhost`. This is accomplished by our open source library
<a href="https://github.com/picosh/ptun">ptun</a>.

# Example usage

```bash
ssh -L 1338:localhost:80 -N imgs.sh 
# separate terminal
docker push localhost:1338/my-img:latest
docker pull localhost:1338/my-img:latest
# or use the docker API
curl http://localhost:1338/v2/_catalog
```

# SSH CLI

```bash
ssh imgs.sh help
ssh imgs.sh ls
ssh imgs.sh rm alpine --write
```

# GitHub Action

```yml
name: build and push docker image

on:
  push:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    # start ssh tunnel as a container service
    services:
      registry:
        image: ghcr.io/picosh/ptun/autossh:latest
        env:
          USERNAME: <pico_user>
          PRIVATE_KEY: ${{ secrets.PRIVATE_KEY }}
        ports:
          - 5000:5000
    steps:
    - name: Set up QEMU
      uses: docker/setup-qemu-action@v3
    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v3
      with:
        driver-opts: network=host
    - name: Build and push
      uses: docker/build-push-action@v5
      with:
        push: true
        tags: localhost:5000/image:latest
```

# Ready to join pico?

<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
