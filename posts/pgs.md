---
title: pages
description: A zero-install static site hosting service for hackers
keywords: [pico, pgs]
toc: 1
---

The easiest way to deploy static sites on the web.

> NOTICE: This is a premium [pico+](/plus) service with a **free tier** (25mb
> total storage limit)

# Features

- Use familiar cli tools to fully manage static sites
- Distinct static sites as projects
- Unlimited projects, created instantly upon upload
- Deploy using [rsync, sftp, or scp](/file-uploads)
- Promotion/rollback support
- Managed HTTPS for all projects
- Site [analytics](/analytics)
- [Custom domains](/custom-domains#pgssh) for projects
- [Custom redirects and rewrites](#redirects-and-rewrites)
- [Custom headers](#headers)
- [SPA support](#single-page-applications)
- [Image manipulation API](/images#image-manipulation)
- [Private projects](#access-control-list)
- [No bandwidth limitations](/faq#are-there-any-bandwidth-limitations)

## Demo

<iframe width="100%" height="315" src="https://www.youtube.com/embed/sdbQpD2jV0k?si=B0yoV25XIaDqnTvk" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>

> [pgs demo](https://www.youtube.com/watch?v=sdbQpD2jV0k)

# Publish your site with one command

When your site is ready to be published, copy the files to our server with a
familiar command:

```bash
rsync -rv public/ pgs.sh:/myproj
```

That's it! There's no need to formally create a project, we create them
on-the-fly. Further, we provide TLS for every project automatically.

# Manage your projects with a remote CLI

Use our CLI to manage projects:

```bash
ssh pgs.sh help
```

# Instant promotion and rollback

Additionally you can setup a pipeline for promotion and rollbacks, which will
instantly update your project.

```bash
ssh pgs.sh link project-prod project-d0131d4
```

A common way to perform promotions within `pgs` is to setup CI/CD so every `git`
push to `main` would trigger a build and create a new project based on the git
commit hash (e.g. `project-d0131d4`).

This command will create a symbolic link from `project-prod` to
`project-d0131d4`. Want to rollback a release? Just change the link for
`project-prod` to a previous project.

[Here's an example using our github action](https://github.com/neurosnap/neovimcraft/blob/main/.github/workflows/deploy.yml).

# CI/CD

Since we are just using `rsync` for static site deployments, all you need is a
way to run that command in a CI environment.

We also built a [github action](https://github.com/picosh/pgs-action) that
handles all the logic for uploading to `pgs` which includes support for
promotions and static site retention.

## basic

```yaml
name: "basic static site deployment"

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: denoland/setup-deno@v1
        with:
          deno-version: "~1.42"
      - run: make build

      - name: upload to pgs
        uses: picosh/pgs-action@v3
        with:
          user: erock
          key: ${{ secrets.PRIVATE_KEY }}
          src: "./public/"
          # https://erock-myapp.pgs.sh
          project: "myapp"
```

## promotion and deployment retention policy

With static site promotion using symbolic links and a site retention policy:

```yaml
name: "promotion static site deployment"

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: denoland/setup-deno@v1
        with:
          deno-version: "~1.42"
      - run: make build

      - name: Set outputs
        id: vars
        run: echo "sha_short=$(git rev-parse --short HEAD)" >> $GITHUB_OUTPUT

      - name: upload to pgs
        uses: picosh/pgs-action@v3
        with:
          user: erock
          key: ${{ secrets.PRIVATE_KEY }}
          src: "./public/"
          # git sha to create a project specific to this commit
          project: "myapp-${{ steps.vars.outputs.sha_short }}"
          # promote the project above to the "production" site
          promote: "myapp"
          # delete all sites matching this prefix ...
          retain: "myapp-"
          # ... except for the latest (1) deployment
          retain_num: 1
```

## preview apps

```yaml
name: "preview apps"

on:
  pull_request:
    branches:
      - "main"

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: denoland/setup-deno@v1
        with:
          deno-version: "~1.42"
      - run: make build

      - name: upload to pgs
        uses: picosh/pgs-action@v3
        with:
          user: erock
          key: ${{ secrets.PRIVATE_KEY }}
          src: "./public/"
          # create a site based on pr
          project: "myapp-pr${{ github.event.pull_request.number }}"
```

# CLI Reference

The best way to learn about all the commands we support is via an SSH command:

```bash
ssh pgs.sh help
```

Having said that, we do want to demonstrate the power of `pgs` by discussing
design goals. All of our SSH commands are safe-by-default. Meaning, they never
mutate server state by default. This provides users an opportunity to experiment
with our commands to see how they work. In order to actually trigger server
mutations, every command must be appended with `--write`.

Further, we want to make sure users are able to manage their static sites
exclusively from SSH commands. Below is list of features we support via SSH
commands:

```bash
# storage usage stats
ssh pgs.sh stats

# list all projects (and their links)
ssh pgs.sh ls

# list all project dependencies
ssh pgs.sh depends project-x

# link a project (e.g. folder symlink)
ssh pgs.sh link project-x --to project-y

# unlink a project
ssh pgs.sh unlink project-x

# delete a project
ssh pgs.sh rm project-x

# delete all projects matching a prefix
# (except projects that have linked projects)
ssh pgs.sh prune prefix

# delete all projects matching a prefix
# except the last N recently updated projects (defaults to 3).
# doesn't count linked projects
ssh pgs.sh retain prefix -n 3

# set project to private to only you and matching public keys
ssh pgs.sh acl project-x --type pubkeys --acl sha256:xxx

# clear the http cache for a project
ssh pgs.sh cache project-x
```

# File denylist

You can upload any file you want to pages, with a few exceptions.

Because any file uploaded to pages is public-by-default, we felt it necessary to
automatically reject some files from being uploaded. At this point in time, we
reject all files or files inside directories that start with a period `.`.
Essentially, we reject all dotfiles. This is so users don't accidentally upload
a `.git` folder or `.env` files. This is the equivalent rule in our `.gitignore`
parser:

```
.*
```

## Override denylist

Upload a `_pgs_ignore` to the root of each project. We are using the same rules
as `.gitignore` using [this parser](https://github.com/sabhiram/go-gitignore).

If you want to allow all files without ignoring anything, add a `_pgs_ignore`
with any comment:

```
# dont ignore files
```

> Note: when uploading a `_pgs_ignore`, we cannot guarantee it will be uploaded
> first so we recommend uploading it on its own and then upload the rest of your
> site.

# Pretty URLs

By default we support pretty URLs. So there are some rules surrounding URLs that
are important to understand.

For the route `https://{user}-{project}.pgs.sh/space`, we will check for the
following files:

- `/space`
- `/space.html`
- `/space/`: `301` redirect to `/space/index.html`
- `/404.html`

As you can see from the third entry, we add a trailing slash to all routes. This
is a common practice with static sites in order to prevent having different
behavior from visiting a site with and without a trailing slash.

# Custom Domains

We have a very easy-to-setup guide on [custom domains](/custom-domains#pgssh).

# Redirects and rewrites

We support custom redirects and rewrites via a special file `_redirects`.

> The `_redirects` file size cannot exceed 5KB.

```
# Redirect browser request to what we serve
/home                /
/blog/post.php       /blog/post
/news                /blog
/wow                 https://wow.com
/authors/c%C3%A9line /authors/about-c%C3%A9line
```

When no status is provided, we default to `301 Moved Permanently`.

```
# Redirect with a 301
/home         /              301

# Redirect with a 302
/my-redirect  /              302

# Show a custom 404 for this path
/ecommerce    /store-closed  404

# Rewrite a path
/pass-through /index.html    200
```

## Route Shadowing

By default we do **not** shadow routes that exist. For example:

- `/space.html` exists on your site,
- with a `_redirects` entry `/space   /   301`

If the user goes to `/space` then it will always prefer `/space.html`. You can
override this preference by adding a force flag to your redirect entry:

```
/space   /   301!
```

## Redirect `www` to apex domain

A common requirement is to redirect "www.example.com" to the apex domain
"example.com" or the other way around.

To accomplish this, we recommend you create a separate project with just a
`_redirects` file inside of it.

1. Create a `_redirects` file with a `301` to the apex domain:

```bash
echo "/*  https://example.com/:splat  301" >> _www_redirects
scp "$PWD/_www_redirects" pgs.sh:/www-proj/_redirects
```

2. Add a `www` CNAME and TXT record to point to www project

See our [custom domains](/custom-domains#pgssh) page.

## Rewrites

When you assign an HTTP status code of `200` to a redirect rule, it becomes a
rewrite. This means that the URL in the visitor’s address bar remains the same,
while pico's servers fetch the new location behind the scenes, effectively
proxying the request.

We also support rewrite rules for when you want to show content from another
site without a full URL redirect.

This can be useful for single page apps, proxying to other services, proxying to
other `pgs` sites, or transitioning for legacy content.

Here are some examples:

```
/*                          https://my-other-site.pgs.sh/:splat 200
/my-site/*                  https://my-other-size.pgs.sh/:splat 200
/news/:month/:date/:year/*  /blog/:year/:month/:date/:splat     200
```

### Proxy to another service

> NOTICE: This is a premium [pico+](/plus) feature.

Similar to how you can rewrite paths like `/*` to `/index.html`, you can also
set up rules to let parts of your site proxy to external services. Let’s say you
need to communicate from a single-page app with an API on
https://api.example.com that doesn’t support CORS requests. The following rule
will let you use `/api/` from your JavaScript client:

```
/api/*  https://api.example.com/:splat  200
```

### Caveats

- Infinitely looping rules, where the "from" and "to" resolve to the same
  location, are incorrect and will be ignored.
- By default, we limit internal rewrites to one "hop".
- Rewrites can cause pages that use assets specified through relative paths to
  load incorrectly. To make sure your site's proxied content is displayed as
  expected, use absolute paths for your assets.
- Paths handled by proxies may not redirect from HTTP to HTTPS URLs as expected.
  If you’re working with proxies, we recommend only publishing HTTPS URLs for
  your visitors to use.

# Headers

We support custom headers via a special file `_headers`.

> The `_headers` file size cannot exceed 5KB.

```
# a path:
/templates/index.html
  # headers for that path:
  X-Frame-Options: DENY
  X-XSS-Protection: 1; mode=block
# another path:
/templates/index2.html
  # headers for that path:
  X-Frame-Options: SAMEORIGIN
```

```
/*
  X-Frame-Options: DENY
  X-XSS-Protection: 1; mode=block
```

## Denied Headers

These headers are **not** allowed:

```
Accept-Ranges
Age
Allow
Alt-Svc
Connection
Content-Encoding
Content-Length
Content-Range
Date
Location
Server
Trailer
Transfer-Encoding
Upgrade
```

# Single-Page Applications

We support SPAs! Upload a `_redirects` file to your project:

```
/*  /index.html  200
```

# Reserved username project

If you create a project with the same name as your username, then you can access
it at:

```bash
rsync -rv public/ glossy@pgs.sh:/glossy
# => https://glossy.pgs.sh
```

# Access Control List

Thanks to SSH tunnels we can provide restricted access to projects.

We have three options:

- public (default)
- pubkeys (list of sha256 public keys to give read access to)
- pico (list of pico users to grant read access to)

```bash
# access to anyone with a public key
ssh pgs.sh acl project-x --type pubkeys

# access only to public keys provided
ssh pgs.sh acl project-x --type pubkeys --acl sha256:xxx --acl sha256:yyy

# access to anyone with a pico account
ssh pgs.sh acl project-x --type pico

# access only to pico users provided
ssh pgs.sh acl project-x --type pico --acl antonio --acl erock

# access to anyone
ssh pgs.sh acl project-x --type public
```

To connect to a private project:

```bash
ssh -L 1337:localhost:80 -N {subdomain}@pgs.sh

# for example our pico UI is only available through an SSH tunnel:
ssh -L 1337:localhost:80 -N pico-ui@pgs.sh
```

Then open your browser to http://localhost:1337

# Caching

To improve the page speed, pgs sites are cached for 10 minutes by default. This
is controlled by the
[`Cache-Control: max-age=600` header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control)
which you can [override with a `_headers` file](#headers).

There are two levels of caching: server-side and client-side. The server-side
cache is automatically cleared every time you upload files, but client-side
caches only expire when `max-age` seconds pass, or if you force-reload or clear
your browser cache manually.

In case of issues, you can manually clear the server-side cache with
`ssh pgs.sh cache project-name`.

# Does pages have a CDN or multi-region support?

At this point in time, we are able to serve content from a single VM. If this
service gains traction we will commit to having a CDN with multiple regions in
the US and EU.

# Removing a project

The _only_ way to delete a project and its contents is with our remote cli:

```bash
ssh pgs.sh rm <project>
```

# File upload caveats

Everyone's static sites are stored inside our object store. In order for `sftp`
and `sshfs` to work we need to emulate a folder structure. Object store's are
just an object with a name prefix that **resembles** a folder structure. As such
in order for empty folders to be traversed in an emulated filesystem, we need to
create dummy files `._pico_keep_dir` that let us keep a reference to an empty
folder inside our object store. As such:

> You cannot delete a project using `sftp` or `sshfs` commands

You must [delete a project](#removing-a-project) using the remote cli.

If you accidentally remove a site you will be stuck in a limbo state. The folder
will still exist using `sftp` or `sshfs`. You can properly clean it up by
running the [rm command](#removing-a-project)

<hr />
<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
