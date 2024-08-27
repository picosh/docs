---
title: docker registry
description: A private docker image registry using SSH
---

> NOTICE: This is a premium [pico+](/plus) service

# Features

- Private docker image registry
- Use SSH keys for authentication
- `docker push` and `docker pull` work exactly the same

# How it works

We leverage web tunnels to open a portal to our private image registry where you
can access it via `localhost`. This is accomplished by our open source library
[tunkit](https://github.com/picosh/tunkit).

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

We provide a CLI to manage your docker images.

```bash
ssh imgs.sh help
ssh imgs.sh ls
ssh imgs.sh rm alpine --write
```

# Always connected

We use a docker compose file to always be connected inside our VMs, for example:

```yml
services:
  registry:
    image: ghcr.io/picosh/ptun/autossh:latest
    restart: unless-stopped
    ports:
      - 1338:1338
  app:
    image: localhost:1338/app
    restart: unless-stopped
    ports:
      - "80:80"
```

# GitHub action

Need to use imgs.sh with CI/CD? Just create the SSH tunnel before trying to use
`docker`. We provide a custom image based on `autossh` to make this easier:

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
          PRIVATE_KEY: ${{ secrets.PRIVATE_KEY }}
        ports:
          - 1338:1338
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

<hr />
<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
