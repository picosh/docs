---
title: prose
description: Write your blog with markdown. rsync and you're live.
keywords: [pico, prose]
toc: 2
---

> prose.sh is a free service

Write your blog with markdown. Send it to our server. You're live.

No static site generators. No build step. No deploy pipeline. No YAML configs. Just `rsync` your posts and we handle the rest.

# Why prose over self-hosting?

You know how to set up a blog. You've done it before -- Hugo, Jekyll, Astro, whatever's trending this month. But then comes the maintenance: updating dependencies, fixing broken builds, renewing certs, debugging why your CI suddenly fails.

`prose.sh` gives you the same control you'd have self-hosting, minus the operational burden. Your content stays in plain markdown files on your machine. You own your workflow. We just serve it.

# Features

| Feature                                                        | What it means for you                                       |
| -------------------------------------------------------------- | ----------------------------------------------------------- |
| **Zero dependencies**                                          | If you have `ssh`, you're ready                             |
| **Your editor, your workflow**                                 | vim, emacs, vscode -- write where you're comfortable        |
| **Instant publishing**                                         | `rsync ~/blog/* prose.sh:/` and you're live                 |
| **SSH key auth**                                               | No passwords, bring the ssh keypair you own                 |
| **Custom domains**                                             | Point your DNS and go                                       |
| **[Site analytics](/analytics)**                               | Privacy-respecting traffic insights, no third-party scripts |
| **RSS out of the box**                                         | Readers can subscribe immediately                           |
| **[GitHub flavored markdown](https://github.github.com/gfm/)** | Write what you already know                                 |
| **[Plain text lists](/plain-text-lists)**                      | Write using a simple yet expressive list format             |
| **Blog customization**                                         | Tweak layout, styles, and metadata with simple files        |

Check out the [discovery page](https://prose.sh) on prose.

# Publish

Create posts using your favorite editor in plain text files.

`~/blog/hello-world.md`

```md
---
date: 2022-06-28
---

# hello world!

This is my first blog post.

Check out some resources:

- [pico.sh](https://pico.sh)
- [antoniomika](https://antoniomika.me)
- [bower.sh](https://bower.sh)

Cya!
```

When your post is ready to be published, copy the file to our server with a familiar command:

```bash
rsync ~/blog/* prose.sh:/
# - or -
scp ~/blog/* prose.sh:/
# - or -
sftp prose.sh
# - or -
sshfs ~/blog prose.sh:/
# or pipe the file to us
echo -e "# hello world!\n\nWelcome to my blog!" | ssh prose.sh hello-world.md
# or pipe the file to us and generate a random name
echo -e "# hello world!\n\nWelcome to my blog!" | ssh prose.sh
```

> [Read more about uploading files](/file-uploads).

Since we are leveraging tools you already have on your computer, there is nothing to install. This provides the convenience of a web app, but from inside your terminal!

# Plain text lists

We also support the ability to use our homegrown [plain text lists](/plain-text-lists) format `.lxt` to write blog posts. This format is simple yet covers what is necessary for writing prose. It's great for people who like to think in lists and we think perfectly compliments our blog service.

```txt
=: title hello world
=: description welcome to my blog
=: date 2026-01-16

Hello world!
How are we doing today?
=> https://pico.sh this is a hyperlink
> This is a blockquote
<= https://bower.sh/profile.jpg this is an image
It supports
	Nesting!
		Wow, great!
```

Then upload the file to us:

```bash
scp hello-world.lxt prose.sh:/
```

It uses the exact same metadata variables as our markdown posts. This idea here is that we have an alternative format that is much easier to render and we find easier to write. There is no inline markup, the type of each list item can be determined by reading the first 3 characters of a line.

# Images

We also support [image uploading](/images). Simply upload your images alongside your markdown files and then reference them from root `/`:

```md
---
title: hello world!
---

![profile pic](/profile.jpg)
```

```bash
rsync ~/blog/*.jpg prose.sh:/
```

> All images are uploaded to a [pages](/pgs) site called `prose`. This means that there's an auto-generated project where you can quickly access all your blog images via https://{user}-prose.pgs.sh

When you upload an image to prose, we make strip exif data. We also support an [image manipulation API](/images#image-manipulation)!

# Blog structure

Think of your blog as a flat list of files, similar to how your blog is rendered as a flat list of posts.

> All filenames uploaded to prose must be unique.

We do **not** support storing files in different folders. On your local machine, you can store files however you like, but once uploaded to us, we lose all directory structure information and only keep the filename.

For example, if you have two posts in different folders but the same filename and try to upload them to prose, the second post will overwrite the first one.

# Blog customization

User can change the look-and-feel of their blog by uploading special files.

## \_readme.md

This file will allow users to add content above the blog post list, add links to their blog landing page, and change blog metadata.

```md
---
title: some title!
description: this is a great description
nav:
  - google: https://google.com
  - site: https://some.site
image: og_image.jpg
card: summary # or summary_large_image
favicon: /fav.ico # or https://...
layout: aside # or default
with_styles: true
---

This block will be rendered above the list of blog posts on the index page!
```

### `title`

Title changes your blog's name which will show up as the metadata title for the blog index page and changes the blog post link back to the main blog page.

### `description`

The description is used as the metadata description on the blog page.

### `nav`

This property sets the navigation links on the blog index page. The `rss` link is always appended to the end of this list.

### `image` and `card`

These properties are used for og and image metadata when for when a blog is shared. If a blog post does not have its own `image` or `card` properties, then it will assume the blog's `image` and `card` properties.

### `favicon`

This property will change the `favicon` for the blog and posts. Link to an ico or any other image format using a URL (root and root-relative will reference your prose images).

It doesn't have to be an `ico`, other images like `png` are also supported.

For example: `/fav.ico` or `https://other.site/favicon.png`

> If referencing a prose image (e.g. `/fav.ico`), do **not** make your filename `favicon.ico` because we do not currently support overwriting the default prose favicon.

### `layout`

This property changes the layout of the blog index page. The options are:

- `aside`
- `default`

### `with_styles`

This determines whether we load our default styles onto your blog or not. This allows for better customization in conjunction with [\_styles.css](/#-stylescss) which is the user-defined CSS stylesheet that gets served.

## \_styles.css

This will allow users to change the look-and-feel of their blog. This file will be the final CSS loaded on the page so it will overwrite whatever styles have previously been added. We've also added a couple of convenience id's attached to the body element for the blog and post pages.

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

Further we add the post slug as a class on the `body` element. So if you upload a file called `barrel.md` then the body element will have:

```html
<body id="post" class="barrel"></body>
```

## \_footer.md

This file will be added to the end of every single blog post.

There is nothing that differentiates itself from the rest of the post so it's up to you to style it. For convenience we added an `id` to the containing element `post-footer`.

## \_404.md

This file will override the default 404 page when a post cannot be found.

```md
---
title: Not found
description: Where are you going?
---

This page doesn't exist.
```

# Post customization

We support adding frontmatter to the top of your posts. A frontmatter looks like the following:

```md
---
title: some title!
description: this is a great description
date: 2022-06-28
tags: [feature, announcement]
image: og_image.jpg
card: summary # or summary_large_image
draft: true
toc: true # or false or integer for max heading depth
aliases:
  - 2023/03/10/my-post
---
```

## `title`

Title changes the title of the blog post.

## `description`

Description changes the description metadata for the blog post.

## `date`

> We **highly** recommend specifying a date because this preserves your actual publish date and not when the database record was created by us.

This is the published date. It ought to be in `YYYY-MM-DD` format. If this date is set to the future, the post will be unlisted until the date provided is today or in the past.

## `image` and `card`

These properties are used for og and image metadata when for when a blog is shared. If a blog post does not have its own `image` or `card` properties, then it will assume the blog's `image` and `card` properties.

## `draft`

This property will change the listing status of a blog post. If `draft: true` then the post will be published but unlisted. It will not show up on the blog index page or RSS feed for your blog.

However, the post is still publicly accessible! This gives users the ability to share posts before "publishing" them.

## `toc`

This property adds a table of contents to the blog post based on the headers. If set to `false` then no table will be rendered. If set to `true` then table will be rendered. If set to an integer greater than `0` then it will set the max heading depth to that value. So if set to `2` the table will only render links of heading `h1` and `h2`.

Examples:

```md
---
toc: false # no table of contents
toc: true # table of contents with all headings
toc: 1 # with only h1 headings
toc: 2 # with only h1 and h2 headings
toc: 3 # with only h1, h2, and h3 headings
---
```

## `aliases`

This property will redirect all routes listed to this blog post. Primarily used for migrating from a different blog platform to prose.

```md
---
aliases:
  - 2023/03/10/my-post
  - my-post.html
---
```

## Without frontmatter

If the user does not provide a frontmatter, we will set the following defaults:

- `date` set to date when uploaded
- `title` inferred from first h1 header (e.g. `# header`) or else the filename
- `tags` inferred from any hashtags inside content

# Remove ad footer

We don't mind. Upload a CSS file `_styles.css` with:

```css
footer {
  display: none;
}
```

# Custom Domains

Setting up a custom domain for prose is similar to all other services. First you need to create a CNAME record:

```
CNAME subdomain.yourcustomdomain.com -> prose.sh
```

Resulting in:

```
subdomain.yourcustomdomain.com.   300   IN    CNAME   prose.sh.
```

Next, add a `TXT` record to tell Prose what blog is hosted on that domain at the subdomain entry `_prose`

```
TXT _prose.subdomain.yourcustomdomain.com -> {user}
```

Where `{user}` is your pico username.

Resulting in:

```
_prose.subdomain.yourcustomdomain.com.    300   IN    TXT   erock
```

> We are replacing `{user}` token with the username `erock` as an example.

Depending on your DNS, this could take some time to fully switch over.

## Debug custom domains

We have an endpoint to check whether or not custom domains are setup:

```bash
curl -i 'https://prose.sh/check?domain=subdomain.yourcustomdomain.com'
```

First check the main record:

```bash
dig subdomain.yourcustomdomain.com

; <<>> DiG 9.18.36 <<>> subdomain.yourcustomdomain.com
;; QUESTION SECTION:
;subdomain.yourcustomdomain.com.               IN      A

;; ANSWER SECTION:
subdomain.yourcustomdomain.com.        60      IN      A       129.158.37.104
```

Then check the `TXT` record:

```bash
dig -t txt +short _prose.subdomain.yourcustomdomain.com

erock
```

<hr />
<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
