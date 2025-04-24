---
title: pipe
description: Stream data between computers using our authenticated *nix pipes using SSH
keywords: [pico, pipe]
toc: 2
---

The simplest authenticated pubsub system. Send messages through user-defined
topics (aka channels). By default, topics are private to the authenticated ssh
user. The default pubsub model is multicast with bidirectional blocking, meaning
a publisher (pub) will send its message to all subscribers (sub) for a topic.
There can be many publishers and many subscribers on a topic. Further, both pub
and sub will wait for at least one event to be sent or received on the topic.

[pipe](https://pipe.pico.sh) is a simple and secure way of putting together
composable directional streams of data, just like a *nix `|` operator!

# Features

- Zero-install
- Familiar `*nix` pipes API
- Authenticated pubsub using ssh
- Private pubsub by default
- Public pubsub by topic (opt-in)
- Multicast (many pubs to many subs)
- Bidirectional (e.g. chat)
- Paradigms for connecting to a topic:
  - Read (sub)
  - Write (pub)
  - Read & Write (pipe)

# Use cases

- Send desktop notifications
- File sharing
- Pipe command output
- Chat
- Reverse shell
- CI/CD

# Examples

For example, maybe you have a *nix pipe you're using on the command line like
so:

```bash
tail -f -n 0 /tmp/foo.log | grep --line-buffered "ERROR" | xargs -I{} -L1 osascript -e 'display notification "{}" with title "Pipe Notification"'
```

This simple one liner will grab new messages from a log file, check if it
contains `ERROR` and then send a notification for macOS using `AppleScript`.
This works if you run your app locally, but what if you run the app on a remote
server? You can use `pipe` to connect remote and local without ever leaving the
command line!

On your remote side, you would start a tail and `pub` it to a topic:

```bash
tail -f -n 0 /tmp/foo.log | ssh pipe.pico.sh pub foo.log
```

On your local side, you would `sub` it to the notify command:

```bash
ssh pipe.pico.sh sub foo.log | grep --line-buffered "ERROR" | xargs -I{} -L1 osascript -e 'display notification "{}" with title "Pipe Notification"'
```

The beauty of this method is that the command will wait until the `sub` is
started before any data is consumed, ensuring you never miss a log line. If you
didn't want it to wait for a `sub`, you can add `-b=false` (`b` for blocking) to
the `pub` to prevent it from blocking.

Once you stop the `pub` command, the `sub` will also exit. You can also prevent
this by adding `-k` (`k` for keepalive) to the `sub` command. With both blocking
disabled and keepalive enabled, the commands would look like so:

Remote:

```bash
tail -f -n 0 /tmp/foo.log | ssh pipe.pico.sh pub -b=false foo.log
```

Local:

```bash
ssh pipe.pico.sh sub -k foo.log | grep --line-buffered "ERROR" | xargs -I{} -L1 osascript -e 'display notification "{}" with title "Pipe Notification"'
```

We can combine this with any commands we want and create a pretty robust pub/sub
system. We can even send full command output through a pipe:

```bash
ssh pipe.pico.sh sub htop
```

```bash
htop | ssh pipe.pico.sh pub htop
```

Now what if you wanted to have bi-directional IO? That's where our last command
of Pipe comes in, `pipe`! Pipe allow you to open a bi-directional client on any
topic. It's fully non-blocking and can allow you to do things like have fully
interactive chat over pipe. For example, running the following in two terminals:

```bash
ssh pipe.pico.sh pipe chat
```

Will allow you to type and read messages as if you were sitting at the same
terminal!

# CLI

We have a bunch of demo examples on the main [pipe](https://pipe.pico.sh)
website so be sure to check those out.

```bash
~$ ssh pipe.pico.sh help
Command: ssh pipe.pico.sh <help | ls | pub | sub | pipe> <topic> [-h | args...]

The simplest authenticated pubsub system.  Send messages through
user-defined topics.  Topics are private to the authenticated
ssh user.  The default pubsub model is multicast with bidirectional
blocking, meaning a publisher ("pub") will send its message to all
subscribers ("sub").  Further, both "pub" and "sub" will wait for
at least one event to be sent or received. Pipe ("pipe") allows
for bidirectional messages to be sent between any clients connected
to a pipe.

Think of these different commands in terms of the direction the
data is being sent:

- pub => writes to client
- sub => reads from client
- pipe => read and write between clients
```

## Subs

```bash
~$ ssh pipe.pico.sh sub -h
Usage: sub <topic> [args...]
Args:
  -a string
        Comma separated list of pico usernames or ssh-key fingerprints to allow access to a topic
  -c    Don't send status messages
  -k    Keep the subscription alive even after the publisher has died
  -p    Subscribe to a public topic
```

## Pubs

```bash
~$ ssh pipe.pico.sh pub -h
Usage: pub <topic> [args...]
Args:
  -a string
        Comma separated list of pico usernames or ssh-key fingerprints to allow access to a topic
  -b    Block writes until a subscriber is available (default true)
  -c    Don't send status messages
  -e    Send an empty message to subs
  -p    Publish message to public topic
  -t duration
        Timeout as a Go duration to block for a subscriber to be available. Valid time units are 'ns', 'us' (or 'Âµs'), 'ms', 's', 'm', 'h'. Default is 30 days. (default 720h0m0s)
```

## Pipes

```bash
~$ ssh pipe.pico.sh pipe -h
Usage: pipe <topic> [args...]
Args:
  -a string
        Comma separated list of pico usernames or ssh-key fingerprints to allow access to a topic
  -c    Don't send status messages
  -p    Pipe to a public topic
  -r    Replay messages to the client that sent it
```

# Web Interface

Now what if you don't have a terminal available? Not a problem! Pipe has a web
component that works side by side, `pipe-web`. For example, let's start a
notification `sub` through the terminal like so (note: `p` for public, so anyone
can send us a notification):

```bash
ssh pipe.pico.sh sub -p -k notifications | xargs -I{} -L1 osascript -e 'display notification "{}" with title "Pipe Notification"'
```

We can send a `POST` to the `pipe-web` service to send a message onto that topic
like so:

```bash
echo "Hello world" | curl -X POST https://pipe.pico.sh/topic/notifications --data-binary @-
```

And a notification will pop up! Now it's important to note that this is risky,
anyone can use a "`p`ublic" topic (on both the terminal or `pipe-web`). We can
make this less risky by starting the notifications topic with an access list
set. And if we want `pipe-web` to access it, we need to provide "pico" to the
access list setting:

```bash
ssh pipe.pico.sh sub -a pico -p -k notifications | xargs -I{} -L1 osascript -e 'display notification "{}" with title "Pipe Notification"'
```

Now, only `pipe-web` and yourself are able to access this public topic.

`pipe-web` comes with a few caveats, namely all topics need to be public for it
to work. You can set an access list on the topic, but `pipe-web` is an
unauthenticated service. Therefore, anyone can send a post request to the
process.

## WebSockets

Just like you can use `pipe-web` for `GET` and `POST` requests, you can also use
WebSockets to get a bi-directional stream to a topic. This functionality is most
closely related to the `pipe` CLI option. You can try this functionality
[here](https://antonio-ws-pipe-term.pgs.sh/). The API is simple and documented
below. This example terminal also accepts query params like the following:

| name        | type     | data type | description                         |
| ----------- | -------- | --------- | ----------------------------------- |
| topic       | optional | boolean   | The topic to connect to             |
| replay      | optional | boolean   | Whether or not to replay message    |
| binary      | optional | string    | Connect the client as binary        |
| autoConnect | optional | string    | Autoconnect once the page is loaded |

## Open a pipe

<details>
 <summary><code>GET</code> <code><b>/socket/:topic</b></code> open a pipe to a topic (with upgrade)</summary>

### Parameters

> | name  | type     | data type | description                |
> | ----- | -------- | --------- | -------------------------- |
> | topic | required | string    | topic name to subscribe to |

### Query Parameters

> | name   | type     | data type | description                                   |
> | ------ | -------- | --------- | --------------------------------------------- |
> | status | optional | boolean   | Receive pipe status messages                  |
> | replay | optional | boolean   | Whether or not to replay message              |
> | binary | optional | string    | Connect the client as binary                  |
> | access | optional | string    | Comma separated list of permissible accessors |

### Responses

> | http code | content-type | response            |
> | --------- | ------------ | ------------------- |
> | `101`     | `N/A`        | Switching Protocols |

### Example websocat

> ```bash
> websocat wss://pipe.pico.sh/socket/test
> ```

</details>

## Subscribe to a topic

<details>
 <summary><code>GET</code> <code><b>/topic/:topic</b></code> subscribe to a topic</summary>

### Parameters

> | name  | type     | data type | description                |
> | ----- | -------- | --------- | -------------------------- |
> | topic | required | string    | topic name to subscribe to |

### Query Parameters

> | name    | type     | data type | description                                         |
> | ------- | -------- | --------- | --------------------------------------------------- |
> | persist | optional | boolean   | Persist the subscription after the publisher closes |
> | access  | optional | string    | Comma separated list of permissible accessors       |
> | mime    | optional | string    | Content type to return to the client                |

### Responses

> | http code | content-type                                         | response                                        |
> | --------- | ---------------------------------------------------- | ----------------------------------------------- |
> | `200`     | `text/plain;charset=UTF-8` or `mime` query parameter | Subscription data. Will hang until a pub occurs |

### Example cURL

> ```bash
> curl -vvv https://pipe.pico.sh/topic/test?persist=true
> ```

</details>

## Publish to a topic

<details>
 <summary><code>POST</code> <code><b>/topic/:topic</b></code> publish to a topic</summary>

### Parameters

> | name  | type     | data type | description                |
> | ----- | -------- | --------- | -------------------------- |
> | topic | required | string    | topic name to subscribe to |

### Query Parameters

> | name   | type     | data type | description                                   |
> | ------ | -------- | --------- | --------------------------------------------- |
> | access | optional | string    | Comma separated list of permissible accessors |

### Responses

> | http code | content-type | response            |
> | --------- | ------------ | ------------------- |
> | `200`     |              | No content returned |

### Example cURL

> ```bash
> curl -vvvv https://pipe.pico.sh/topic/test -d "hello"
> ```

</details>

## Subscribe to a topic

<details>
 <summary><code>GET</code> <code><b>/pubsub/:topic</b></code> subscribe to a topic</summary>

### Parameters

> | name  | type     | data type | description                |
> | ----- | -------- | --------- | -------------------------- |
> | topic | required | string    | topic name to subscribe to |

### Query Parameters

> | name    | type     | data type | description                                         |
> | ------- | -------- | --------- | --------------------------------------------------- |
> | persist | optional | boolean   | Persist the subscription after the publisher closes |
> | access  | optional | string    | Comma separated list of permissible accessors       |
> | mime    | optional | string    | Content type to return to the client                |

### Responses

> | http code | content-type                                         | response                                        |
> | --------- | ---------------------------------------------------- | ----------------------------------------------- |
> | `200`     | `text/plain;charset=UTF-8` or `mime` query parameter | Subscription data. Will hang until a pub occurs |

### Example cURL

> ```bash
> curl -vvv https://pipe.pico.sh/pubsub/test?persist=true
> ```

</details>

## Publish to a topic

<details>
 <summary><code>POST</code> <code><b>/pubsub/:topic</b></code> publish to a topic</summary>

### Parameters

> | name  | type     | data type | description                |
> | ----- | -------- | --------- | -------------------------- |
> | topic | required | string    | topic name to subscribe to |

### Query Parameters

> | name   | type     | data type | description                                   |
> | ------ | -------- | --------- | --------------------------------------------- |
> | access | optional | string    | Comma separated list of permissible accessors |

### Responses

> | http code | content-type | response            |
> | --------- | ------------ | ------------------- |
> | `200`     |              | No content returned |

### Example cURL

> ```bash
> curl -vvvv https://pipe.pico.sh/pubsub/test -d "hello"
> ```

</details>

# `pipemgr`

[pipemgr](https://github.com/picosh/pipemgr) is a docker image that will listen
for logs from other running containers and pipe their logs through a topic
called `container-drain`.

```yaml
services:
  pipemgr:
    build:
      context: .
    image: ghcr.io/picosh/pipemgr:latest
    restart: always
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - $HOME/.ssh/id_ed25519:/key:ro
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
      interval: 2s
      timeout: 5s
      retries: 5
      start_period: 1s
  httpbin:
    image: kennethreitz/httpbin
    command: gunicorn -b 0.0.0.0:3000 httpbin:app -k gevent --access-logfile -
    ports:
      - 3000:3000
    labels:
      pipemgr.enable: true
      # filter log lines with:
      # pipemgr.filter: "GET.+(404)"
    depends_on:
      pipemgr:
        condition: service_healthy
```

# Caveats

You must always pipe **something** into pub or else it will block indefinitely
until the process is killed. However, you can provide a flag to send an empty
message: `pub topic -e`.

# Inspiration

A special thanks to [patchbay.pub](https://patchbay.pub) for our inspiration.

# Latest posts

2024-10-06 [pipe: our pubsub ssh service](https://blog.pico.sh/ann-022-pubsub)

<hr />
<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
