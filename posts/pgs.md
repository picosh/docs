---
title: pgs.sh
description: A zero-install static site hosting service for hackers
keywords: [pico, pgs]
---

> NOTICE: This is a premium [pico+](/plus) service with a **free trial**

# Features

- 25MB asset storage with **free trial**
- 10GB asset storage with [pico+](/plus)
- Terminal workflows
- No client-side installation required to fully manage static sites
- Distinct static sites as projects
- Unlimited projects, created instantly upon upload
- Deploy using [rsync, sftp, or scp](/file-uploads)
- Promotion/rollback support
- Managed HTTPS for all projects
- [Custom domains](/custom-domains#pgssh) for projects
- [User-defined redirects](#user-defined-redirects)
- [SPA support](#single-page-applications)
- [Image manipulation API](/images#image-manipulation)
- [Only web assets are supported](#what-file-types-are-supported)

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

A common way to perform promotions within pgs.sh is to setup CI/CD so every
`git` push to `main` would trigger a build and create a new project based on the
git commit hash (e.g. `project-d0131d4`).

This command will create a symbolic link from `project-prod` to
`project-d0131d4`. Want to rollback a release? Just change the link for
`project-prod` to a previous project.

We also built a [github action](https://github.com/picosh/pgs-action) that
handles all the logic for uploading to pgs.sh.
[Here's an example of it in action.](https://erock-git-neovimcraft.pgs.sh/tree/main/item/.github/workflows/deploy.yml.html#27)

# CLI Reference

The best way to learn about all the commands we support is via an SSH command:

```bash
ssh pgs.sh help
```

Having said that, we do want to demonstrate the power of pgs.sh by discussing
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
ssh pgs.sh link project-x project-y

# unlink a project
ssh pgs.sh unlink project-x

# delete a project
ssh pgs.sh rm project-x

# delete all projects matching a prefix
# (except projects that have linked projects)
ssh pgs.sh prune prefix

# delete all projects matching a prefix
# except the latest (3) projects
ssh pgs.sh retain prefix
```

# What file types are supported?

```
html
htm
css
js
jpg
png
gif
webp
svg
ico
pdf
json
txt
otf
ttf
woff
woff2
md
rss
xml
atom
map
webmanifest
```

# User-defined Redirects

We support custom redirects via a special file `_redirects`.

Read more about it at [netflify](https://docs.netlify.com/routing/redirects).

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

# CSP policies

For pico domains we have modestly strict content-security policies.

```bash
Content-Security-Policy "default-src 'self'; img-src * 'unsafe-inline'; style-src * 'unsafe-inline'"
```

If you need to access sites that are blocked by this CSP, then you can use a
[custom domain](/custom-domains) which won't have those security restrictions.

# Ready to join pico?

<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
