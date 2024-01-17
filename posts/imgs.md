---
title: imgs.sh
description: Image hosting for hackers
keywords: [pico, imgs]
---

# Features

- Delightful terminal workflow
- Share public images from the terminal
- Seamless integration with other pico services (e.g. prose)
- Images are web optimized by default
- API to modify images on-the-fly (e.g. dimensions)
- Hotlinking encouraged!
- No javascript
- No ads
- 10MB max image size
- 1GB max storage

# What it is

- A place to host **public** images
- A place to host your blog images
- We moderate all images

# What it isn't

- Not a general cloud photo storage solution
- Not a place to backup or store your photos

# Publish your images with one command

When your image is ready to be published, copy the file to our server with a
familiar command:

```bash
rsync *.jpg imgs.sh:/
```

We'll either create or update the images for you.

# Web optimized

When a user uploads an image, we immediately convert it to `webp`. Then we have
an API that serves those web optimized images.

# How does imgs integrate with other pico services?

We allow any of our other services to upload images from those services to imgs.
For example, if you want to upload images for prose.sh all you have to do is
include your images in the rsync or scp command.

```bash
rsync profile.jpg imgs.sh:/
```

Then when you want to reference the file, you can reference it like so:

```md
![profile pic](/profile.jpg)
```

# What file types are supported?

```
jpg
png
gif
webp
svg
```

# Image manipulation

We have an API that allows users to resize images on-the-fly. Currently we only
support downscaling.

```md
[!profile](/profile/x500) # auto scale width

[!profile](/profile/500x500) # scale width and height

[!profile](/profile/500x) # auto scale height
```
