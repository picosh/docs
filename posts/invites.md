---
title: Invites
description: How to get and send invites for pgs, prose and pastes
keywords: [pico, invite, invitation, pgs, prose, pastes, access]
toc: 2
---

Uploading to [pgs](/pgs), [prose](/prose) or [pastes](/pastes) requires an invite or a [pico+](/plus) membership. Anyone can still [create a pico account](/getting-started) without one, and [pipe](/pipe) works for every account.

Signup happens over SSH with no email, which makes automated signups cheap. Invites let us keep that flow while cutting down on bot and spam accounts. Read the [announcement](https://blog.pico.sh/ann-038-pico-invite-system) for the full background.

# What an invite grants

An invite gives your account access to pgs, prose and pastes for five years. pico+ members already have all three and don't need an invite.

| Service | Storage | Max file size |
| ------- | ------- | ------------- |
| pgs     | 50MB    | 10MB          |
| prose   | 25MB    | 10MB (images) |
| pastes  | no quota | 3MB          |

Prose access includes [blog analytics](/analytics). For more storage, larger files, and the other pico+ services, see [pico+](/plus).

# How to get an invite

1. [Create a pico account](/getting-started) if you don't have one. Invites are sent to an existing pico username.
1. Ask someone who can invite you. If you don't know anyone, email [hello@pico.sh](mailto:hello@pico.sh) or ask in [#pico.sh on IRC](/irc).
1. Once you've been invited, your next upload to pgs, prose or pastes will work. No action is needed on your end.

You can confirm you've been invited from the [pico TUI](/ui#ssh-tui): the user info panel shows "invited by" with the inviter's username, and the services list shows pages, prose and pastes as "active".

# How to invite someone

Anyone with a valid pico+ membership, or anyone whose invite is still active, can send invites.

1. SSH into our [pico TUI](/ui#ssh-tui)
1. Select the "invite" submenu
1. Type the person's pico username
1. Press <kbd>enter</kbd>

The same page lists everyone you have invited and when.

# Errors

Uploading without access returns one of these:

```
ERROR: uploading to pgs requires an invitation or pico+
ERROR: uploading to prose requires an invitation or pico+
ERROR: uploading to pastes requires an invitation or pico+
```

Get an invite or sign up for [pico+](/plus). If your access or your pico+ membership has lapsed, the message says "expired" instead.
