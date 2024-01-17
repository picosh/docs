---
title: Frequently Asked Questions
description: Answers to frequently asked questions 
keywords: [pico, faq, question, answer]
---

# Permission denied when using SSH

Unfortunately SHA-2 RSA keys are not currently supported.

Unfortunately, due to a shortcoming in Go’s x/crypto/ssh package, we do not
currently support access via new SSH RSA keys: only the old SHA-1 ones will
work. Until we sort this out you’ll either need an SHA-1 RSA key or a key with
another algorithm, e.g. Ed25519. Not sure what type of keys you have? You can
check with the following:

```bash
find ~/.ssh/id_*.pub -exec ssh-keygen -l -f {} \;
```

If you’re curious about the inner workings of this problem have a look at:

- [golang/go#37278](https://github.com/golang/go/issues/37278)
- [go-review](https://go-review.googlesource.com/c/crypto/+/220037)
- [golang/crypto#197](https://github.com/golang/crypto/pull/197)

# Generating a new SSH key

[Github Reference](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

When you're prompted to "Enter a file in which to save the key," press Enter.
This accepts the default file location. At the prompt, type a secure passphrase.

# Can I create multiple accounts?

Yes! You can either:

- create a new keypair and use that for authentication; or
- use the same keypair and ssh into our CMS using our special username
  `ssh new@prose.sh`

Please note that if you use the same keypair for multiple accounts, you will
need to always specify the user when logging into our CMS.

# Why do I provide my username when using SSH?

> We recommend supplying your username everytime you interact with our SSH Apps.

Technically, if you only have one pico account then it doesn't matter. However,
when you don't provide a user when using SSH, then your default user is
provided. Does that really matter? Probably not, but it is leaking information
you might not otherwise want us to know about. We can see that user in our logs.

Always providing your username is ideal and if you have multiple pico accounts
then we **require** your username everytime. So it's just good practice to
always provide the username.

# How can I use `git` to sync my content?

All you need is your provide key and [one of the upload tools we
support](/file-uploads)

For example, [here is how we deploy the pico blog](https://github.com/picosh/official-blog/blob/main/.github/workflows/publish.yml)
