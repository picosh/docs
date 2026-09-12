---
title: patchbin
description: A pastebin for patches, supercharged for git collaboration
keywords: [pico, patchbin, git, patches, pr]
---

> pr.pico.sh is a free, public service

A pastebin for patches, supercharged for git collaboration.

Contributions are designed to be anonymous: the quality of your work is what matters. Account creation happens automatically.

Our server is publicly hosted at [pr.pico.sh](https://pr.pico.sh) and is free for anyone to use with any git repository.

# How it works

The target project doesn't need to run patchbin or even know about it for you to submit a patch request against it.

Instead of navigating a heavy web UI and leaving comment threads, collaboration happens directly with git commits:

- Contributors send patches. Reviewers reply by sending their own commits on top, trading rounds of patchsets back and forth.
- Reviewing code means pulling down the patchset and running it in your local environment.
- There is no accept or reject flow. A patch request is simply active or inactive (becoming inactive after 14 days of no activity).
  - When a maintainer is happy with the changes, they pull them, merge locally, and push upstream.
- Issues are patch requests too. An issue is just a patch request without any code attached yet.
  - Anyone can follow up by submitting a patchset directly on top of it which removes the need cross-linking issues and code changes.

# Workflow

Here is all you need to get started:

## Submit a patch request

Format your commits from git and pipe them straight to our SSH server:

```bash
git format-patch main --stdout | ssh pr.pico.sh pr create myproject
```

You'll receive a PR ID (e.g. `#42`) and a link to view the diff at `https://pr.pico.sh/prs/42`.

## Pull down and review

Anyone can pull down the latest patchset and apply it locally with `git am`:

```bash
ssh pr.pico.sh print 42 | git am -3
```

## Send follow-up changes

Need to address feedback or suggest tweaks? Just commit on top and add a new patchset revision:

```bash
git format-patch main --stdout | ssh pr.pico.sh pr add 42
```

# Learn more

For the complete command reference:

```bash
ssh pr.pico.sh help
```

Code: [github.com/picosh/patchbin](https://github.com/picosh/patchbin)
