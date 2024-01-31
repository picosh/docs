---
title: Images
description: How image hosting works at pico
keywords: [pico, imgs]
---

There are two ways to upload images to pico: through [prose.sh](/prose) or
[pgs.sh](/pgs). They both share similar features to each other, they just serve
different purposes and have different storage limits.

# Features

- Images can be web optimized
- API to modify images on-the-fly (e.g. dimensions, quality, orientation)

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
rsync *.jpg prose.sh:/
```

We'll either create or update the images for you.

Or if you are using [pgs.sh](/pgs):

```bash
rsync *.jpg pgs.sh:/imgs
```

# Web optimized

We do our best to web optimize the images being delivered from pico to users. We
use [imageproxy](https://github.com/willnorris/imageproxy) to serve images from
our object store.

By default, images served from [prose.sh](/prose) are web optimized, meaning we
convert images to webp and reduce the quality.

We do **not** automatically optimize images from [pgs.sh](/pgs) because we try
to subscribe to the principle of least surprise.

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
[!profile](/profile.jpg/s:500) # auto scale width

[!profile](/profile.jpg/s::500) # auto scale height

[!profile](/profile.jpg/s:500:500) # scale width and height

[!profile](/profile.jpg/q:50) # set quality to 50 (scale is 1-100)

[!profile](/profile.jpg/rt:90) # rotate image (choices: 90, 180, 270)

[!profile](/profile.jpg/ext:webp) # convert image to webp
```
