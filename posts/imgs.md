---
title: imgs.sh
description: A private docker image registry with a twist
---

# Features

- 2GB image storage limit
- Use SSH keys for authentication
- `docker push` and `docker pull` work exactly the same

# How it works

We leverage web tunnels to open a portal to our private image registry where
you can access it via `localhost`.  This is accomplished by our open source
library <a href="https://github.com/picosh/ptun">ptun</a>.

# Example usage

```bash
ssh -L 8443:localhost:80 -N imgs.sh 
# separate terminal
docker push localhost:8443/my-img:latest
docker pull localhost:8443/my-img:latest
```

# Ready to join pico?

<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
