---
title: Frequently Asked Questions
description: Answers to frequently asked questions 
keywords: [pico, faq, question, answer]
---

# Permission denied when using SSH

There are a couple reason why this might be happening. We require a public key
for authentication to all of our services, so first make sure you have a valid
SSH keypair.

Not sure what type of keys you have? You can check with the following:

```bash
find ~/.ssh/id_*.pub -exec ssh-keygen -l -f {} \;
```

After that we need to determine why our server is rejecting your public key.
What is most helpful is if you run:

```bash
ssh -vvv pico.sh
```

If you cannot figure out what is wrong just by looking at that output, then you
are more than welcome to join [irc](/irc) and send us a paste of the SSH logs.

You should also make sure you have the correct file permissions set for your ssh
folder and keys:

```bash
chmod 700 ~/.ssh
chmod 644 ~/.ssh/id_ed25519.pub
chmod 600 ~/.ssh/id_ed25519
```

# How do I force the correct pico SSH key?

There are two ways, one is when SSHing to us:

```bash
ssh -o IdentitiesOnly=yes -i ~/.ssh/id_ed25519 pico.sh
```

The other is with an SSH config entry (`~/.ssh/config`):

```bash
Host pico.sh
  IdentitiesOnly yes
  IdentityFile ~/.ssh/id_ed25519

Host pgs.sh
  IdentitiesOnly yes
  IdentityFile ~/.ssh/id_ed25519

Host prose.sh
  IdentitiesOnly yes
  IdentityFile ~/.ssh/id_ed25519
```

# Generating a new SSH key

[Github Reference](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

When you're prompted to "Enter a file in which to save the key," press Enter.
This accepts the default file location. At the prompt, type a secure passphrase.

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
