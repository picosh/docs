---
title: Frequently Asked Questions
description: Answers to frequently asked questions 
keywords: [pico, faq, question, answer]
toc: 1
---

# Why can't I login to pico?

There are a couple reason why this might be happening. We require a public key
for authentication to all of our services, so first we need to make sure you
have a valid SSH keypair and it is being sent to our SSH server.

Not sure what type of keys you have? You can check with the following:

```bash
find ~/.ssh/id_*.pub -exec ssh-keygen -l -f {} \;
```

After that we need to determine why our server is not receiving your pubkey. To
start, trying debugging with:

```bash
ssh -vvv pico.sh
```

Here we are looking for the pubkeys that your SSH client is sending to us.

For example, we want to see something like:

```bash
debug1: Will attempt key: /home/myuser/.ssh/id_ed25519 ED25519 SHA256:XXX agent
debug1: Offering public key: /home/myuser/.ssh/id_ed25519 ED25519 SHA256:XXX agent
debug1: Server accepts key: /home/myuser/.ssh/id_ed25519 ED25519 SHA256:XXX agent
```

If you don't see any of these log statement, it probably means you haven't
loaded your SSH keys into your [SSH agent](#how-can-i-setup-my-ssh-agent).

You should also make sure you have the correct file permissions set for your ssh
folder and keys:

```bash
chmod 700 ~/.ssh
chmod 644 ~/.ssh/id_ed25519.pub
chmod 600 ~/.ssh/id_ed25519
```

If you cannot figure out what is wrong just by looking at that output, then you
are more than welcome to join [irc](/irc) and send us a paste of the SSH logs.

# How do I force the correct pico SSH key?

There are two ways, one is when SSHing to us:

```bash
ssh -o IdentitiesOnly=yes -i ~/.ssh/id_ed25519 pico.sh
```

The other is with an SSH config entry (`~/.ssh/config`):

```bash
Host pico.sh pgs.sh prose.sh
  IdentitiesOnly yes
  IdentityFile ~/.ssh/id_ed25519
```

# How do I generate an SSH key?

[Github Reference](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

When you're prompted to "Enter a file in which to save the key," press Enter.
This accepts the default file location. At the prompt, type a secure passphrase.

# How can I setup my `ssh-agent`?

your SSH agent typically hosts your SSH keypairs and serves them when
authenticating with an SSH server. So, ensuring your `ssh-agent` is setup
properly is important when connecting to us.

[Github has a great guide on this.](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent#adding-your-ssh-key-to-the-ssh-agent)

# How can I use `git` to sync my content?

All you need is your private key and
[one of the upload tools we support](/file-uploads)

For example,
[here is how we deploy the pico blog](https://github.com/picosh/official-blog/blob/main/.github/workflows/publish.yml)

# Can I create multiple pico accounts?

Yes! The only requirement is that SSH keypairs added to our system must be
unique across all users. So you cannot have the same pubkey across multiple pico
accounts.

Other than that, you are welcome to create multiple pico accounts. Just keep in
mind that each [pico+](/plus) membership is assigned to a single account.

# Can I associate multiple SSH keypairs to a single account?

Yes! You can manage your pubkeys for an account via our CMS:

```bash
ssh pico.sh
# => Manage Keys
# Press "n" to add a new pubkey
```

# Are there any bandwidth limitations?

Currently we get 10TB egress per month through our cloud provider. We are
nowhere near that number today (~60GB/mo).

So until we reach 10TB per month, there are no bandwidth limitations.

Once we regularly reach our monthly allowance, we will have to re-evaluate our
pricing structure.

# How can I download a copy of all of my content?

We will provide users with their data if they request it via
[email](mailto:hello@pico.sh).

# How can I delete my content?

[File uploads](/file-uploads#how-do-i-delete-files)

# How can I delete my account with my content?

You are able to delete your account yourself, but it is a little hidden.

```bash
ssh pico.sh
# select "Manage keys"
# Delete all of your keys
# Confirm delete your account
```

# I lost my SSH private key, how can I recover my account?

After account creation, we **highly** recommend users generate an
[API Token](/api-token) and store that someplace safe. If you lose your SSH
private key and [contact us](/contact), we can use that API token to confirm you
own the pico account and manually recover it for you.
