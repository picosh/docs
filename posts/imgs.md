---
title: imgs.sh
description: A private docker image registry with a twist
---

> NOTICE: This is a premium [pico+](/plus) service

# Features

- 2GB image storage limit
- Use SSH keys for authentication
- `docker push` and `docker pull` work exactly the same

# How it works

We leverage web tunnels to open a portal to our private image registry where you
can access it via `localhost`. This is accomplished by our open source library
<a href="https://github.com/picosh/ptun">ptun</a>.

# Example usage

```bash
ssh -L 5000:localhost:80 -N imgs.sh 
# separate terminal
docker push localhost:5000/my-img:latest
docker pull localhost:5000/my-img:latest
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
