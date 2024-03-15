---
title: prose.sh
description: A blog platform for hackers
keywords: [pico, prose]
---

# Features

- Github flavor markdown
- [Custom domains](/custom-domains#prosesh)
- Publish posts using [rsync, sftp, or scp](/file-uploads)
- Looks great on any device
- Bring your own editor
- You control the source files
- Terminal workflow with no installation
- Public-key based authentication
- No ads, zero browser-based tracking
- No attempt to identify users
- No platform lock-in
- No javascript
- Subscriptions via RSS
- Minimalist design

# You control the source files

Create posts using your favorite editor in plain text files.

`~/blog/hello-world.md`

```md
# hello world!

This is my first blog post.

Check out some resources:

- [pico.sh](https://pico.sh)
- [antoniomika](https://antoniomika.me)
- [bower.sh](https://bower.sh)

Cya!
```

# Publish your posts with one command

When your post is ready to be published, copy the file to our server with a
familiar command:

```bash
rsync ~/blog/* prose.sh:/
# - or -
scp ~/blog/* prose.sh:/
# - or -
sftp prose.sh
```

[Read more about uploading files](/file-uploads).

We'll either create or update the posts for you.

# Terminal workflow without installation

Since we are leveraging tools you already have on your computer (`ssh` and
`rsync`), there is nothing to install.

This provides the convenience of a web app, but from inside your terminal!

# Upload images for your blog

We also support [image uploading](/images). Simply upload your images alongside
your markdown files and then reference them from root `/`:

```md
---
title: hello world!
---

![profile pic](/profile.jpg)
```

```bash
rsync ~/blog/*.jpg prose.sh:/
```

When you upload an image to prose, we make it web optimized (e.g. strip exif,
convert to webp, and reduce filesize). We also support an
[image manipulation API](/images#image-manipulation)!

# Metadata

We support adding frontmatter to the top of your markdown posts. A frontmatter
looks like the following:

```md
---
title: some title!
description: this is a great description
date: 2022-06-28
tags: [feature, announcement]
image: og_image.jpg
card: summary # or summary_large_image
draft: true
---
```

# Unlisted posts

When a post has a date in the future or `draft: true` then the post will be
**unlisted**. This means that it will **not** be shown on your blog page nor
will it show up in your built-in RSS feed.

However, the post is still publicly accessible! This gives users the ability to
share posts before "publishing" them.

# How are blogs structured?

Think of your blog as a flat list of files, similar to how your blog is rendered
as a flat list of posts.

> All filenames uploaded to prose must be unique.

We do **not** support storing files in different folders. On your local machine,
you can store files however you like, but once uploaded to us, we lose all
directory structure information and only keep the filename.

For example, if you have two posts in different folders but the same filename
and try to upload them to prose, the second post will overwrite the first one.

# How can I add a footer to all of my posts?

We have a special file `_footer.md` that will be appended to every single blog
post.

There is nothing that differentiates itself from the rest of the post so it's up
to you to style it. For convenience we added an `id` to the containing element
`post-footer`.

# How can I customize my blog page?

There's a special file you can upload `_readme.md` which will allow users to add
a bio and links to their blog landing page.

```md
---
title: some title!
description: this is a great description
nav:
  - google: https://google.com
  - site: https://some.site
image: og_image.jpg
card: summary # or summary_large_image
favicon: favicon.ico
aliases:
  - 2023/03/10/my-post
layout: aside # or default
---
```

# How can I change the theme of my blog?

There's a special file you can upload `_styles.css` which will allow users to
add a CSS file to their page. It will be the final CSS file loaded on the page
so it will overwrite whatever styles have previously been added. We've also
added a couple of convenience id's attached to the body element for the blog and
post pages.

```css
/* _styles.css */
#post {
    color: green;
}

#blog {
    color: tomato;
}

header {}
main {}
footer {}
article {}
```

Further we add the post slug as a class on the `body` element. So if you upload
a file called `barrel.md` then the body element will have:

```html
<body id="post" class="barrel"></body>
```

# How can I change the layout of my blog?

Inside the `_readme.md` metadata file, there's a variable layout option that
will change the layout of your blog index page.

```md
---
layout: aside # or default
---
```

# How can I change the default 404 page?

Upload a `_404.md` which is formatted just like all the other posts.

```md
---
title: Not found
description: Where are you going?
---

This page doesn't exist.
```

# Ready to join pico?

<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
