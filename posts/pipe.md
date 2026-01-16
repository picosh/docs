---
title: pipe
description: PubSub without the infrastructure.
keywords: [pico, pipe]
toc: 2
---

> pipe.pico.sh is a free service

Stream data between machines using the SSH key you already have -- no API keys, no SDKs, no setup.

```bash
echo "hello world" | ssh pipe.pico.sh pub mykey
# separate terminal
ssh pipe.pico.sh sub mykey
hello world
```

That's it. Private and authenticated by default.

[pipe](https://pipe.pico.sh) brings the simplicity of \*nix pipes to the network. If you know how to use `|`, you already know how to use pipe.

# Features

| Feature                        | What it means for you                               |
| ------------------------------ | --------------------------------------------------- |
| **Works on any machine**       | No installation, no dependencies, just SSH          |
| **SSH key is your credential** | No API keys to manage or rotate                     |
| **Private by default**         | Topics are only accessible to you unless you opt-in |
| **Public topics**              | Share with specific users or make fully public      |
| **Primitives that compose**    | Read (`sub`), write (`pub`), or both (`pipe`)       |
| **Multicast built-in**         | Many publishers, many subscribers, just works       |

# Examples

## A basic API

Pipe some data into our ssh app and we will send it to anyone listening.

```bash
ssh pipe.pico.sh sub mykey
```

```bash
echo "hello world!" | ssh pipe.pico.sh pub mykey
```

## Simple desktop notifications

Want to quickly receive a notification when a job is done? It can be as simple as:

```bash
ssh pipe.pico.sh sub notify; notify-send "job done!"
```

```bash
./longjob.sh; ssh pipe.pico.sh pub notify -e
```

## File sharing

Sometimes you need data exfiltration and all you have is SSH:

```bash
cat doc.md | ssh pipe.pico.sh pub thedoc
```

```bash
ssh pipe.pico.sh sub thedoc > ./important.md
```

## Pipe command output

Send command output through our `pipe` command. The `pipe` command is just like `pub` except it is non-blocking and also acts like a `sub`. So a client that can read and write to the topic.

```bash
ssh pipe.pico.sh sub htop
```

```bash
htop | ssh pipe.pico.sh pipe htop
```

## Chat

Use our `pipe` command to have a chat with someone.

```bash
ssh pipe.pico.sh pipe mychan -p
```

Now anyone with a `pico` account can subscribe to this topic using the same command and start typing!

## Pipe reverse shell

If you squint hard enough you can give users interactive access to your shell.

```bash
mkfifo /tmp/f; cat /tmp/f | /bin/sh -i 2>&1 | ssh pipe.pico.sh pipe myshell > /tmp/f
```

```bash
ssh pipe.pico.sh pipe myshell
```

## Simple CI/CD

Having an authenticated, zero-install event system is handy for deploying apps automatically.

```bash
while true; do ssh pipe.pico.sh sub deploy-app; docker compose pull && docker compose up -d; done
```

```bash
docker buildx build --push -t myapp .; ssh pipe.pico.sh pub deploy-app -e
```

## Send a public message

```bash
echo "hello world!" | ssh pipe.pico.sh pub mychan -p
```

Now anyone with a `pico` account can subscribe to this topic:

```bash
ssh pipe.pico.sh sub mychan -p
```

# Pubsub interactions

## Multiple subs

Have many subscribers, they will all receive the message.

```bash
ssh pipe.pico.sh sub foobar
```

```bash
ssh pipe.pico.sh sub foobar
```

```bash
while true; do echo "foobar1"; sleep 1; done | ssh pipe.pico.sh pub foobar
```

## Multiple pubs

Have many publishers send messages to subscribers.

```bash
while true; do echo "foobar1"; sleep 1; done | ssh pipe.pico.sh pub foobar
```

```bash
while true; do echo "foobar2"; sleep 1; done | ssh pipe.pico.sh pub foobar
```

```bash
ssh pipe.pico.sh sub foobar
```

## Multiple pubs and subs

Have many publishers send messages to many subscribers.

```bash
ssh pipe.pico.sh sub foobar
```

```bash
ssh pipe.pico.sh sub foobar
```

```bash
while true; do echo "foobar1"; sleep 1; done | ssh pipe.pico.sh pub foobar
```

```bash
while true; do echo "foobar2"; sleep 1; done | ssh pipe.pico.sh pub foobar
```

# Topic names

All topics are converted into the following format: `{user}/{topic}`. Depending on how the `pub` is created, it changes how to access the topic. These rules are created in an effort to make is ergonomic for multiple use cases.

- `pub {topic}` -> owner can access with `sub {topic}` **or** `sub {owner}/{topic}`
- `pub -a {other} {topic}` -> `other` pico user must access with `sub {owner}/{topic}`
- `pub -p {topic}` -> anyone can access with `sub -p {topic}`

# Pipe monitors

Pipe monitors are health checks on your topics. Think of it as a status monitor for any topic. The simplest version of this can be used to check uptime of any site:

```bash
ssh pipe.pico.sh monitor pico-uptime 1h
```

Then you would setup a cron to ping your site and on success send a pub to pipe:

```bash
*/10 * * * * curl -fsS https://pico.sh && ssh pipe.pico.sh pub -e pico-uptime
# -or-
0/10 * * * * my_script.sh && ssh pipe.pico.sh pub -e pico-uptime
```

Now you can check the status of the monitors:

```bash
ssh pipe.pico.sh status
```

You can then subscribe to an RSS that will send alerts when a topic has not been hit at least once within the time period specied:

```bash
https://pipe.pico.sh/rss/<token>
```

And you can delete monitors:

```bash
ssh pipe.pico.sh monitor -d pico-uptime
```

You can read our original RFC to learn more about use cases: [rfc-008 pipe monitors](https://blog.pico.sh/rfc-008-ping).

# Web Interface

Now what if you don't have a terminal available? Not a problem! Pipe has a web component that works side by side, `pipe-web`. For example, let's start a notification `sub` through the terminal like so (note: `p` for public, so anyone can send us a notification):

```bash
ssh pipe.pico.sh sub -p -k notifications | xargs -I{} -L1 osascript -e 'display notification "{}" with title "Pipe Notification"'
```

We can send a `POST` to the `pipe-web` service to send a message onto that topic like so:

```bash
echo "Hello world" | curl -X POST https://pipe.pico.sh/topic/notifications --data-binary @-
```

And a notification will pop up! Now it's important to note that this is risky, anyone can use a "`p`ublic" topic (on both the terminal or `pipe-web`). We can make this less risky by starting the notifications topic with an access list set. And if we want `pipe-web` to access it, we need to provide "pico" to the access list setting:

```bash
ssh pipe.pico.sh sub -a pico -p -k notifications | xargs -I{} -L1 osascript -e 'display notification "{}" with title "Pipe Notification"'
```

Now, only `pipe-web` and yourself are able to access this public topic.

`pipe-web` comes with a few caveats, namely all topics need to be public for it to work. You can set an access list on the topic, but `pipe-web` is an unauthenticated service. Therefore, anyone can send a post request to the process.

## WebSockets

Just like you can use `pipe-web` for `GET` and `POST` requests, you can also use WebSockets to get a bi-directional stream to a topic. This functionality is most closely related to the `pipe` CLI option. You can try this functionality [here](https://antonio-ws-pipe-term.pgs.sh/). The API is simple and documented below. This example terminal also accepts query params like the following:

| name        | type     | data type | description                         |
| ----------- | -------- | --------- | ----------------------------------- |
| topic       | optional | boolean   | The topic to connect to             |
| replay      | optional | boolean   | Whether or not to replay message    |
| binary      | optional | string    | Connect the client as binary        |
| autoConnect | optional | string    | Autoconnect once the page is loaded |

## Open a pipe

<details>
 <summary><code>GET</code> <code><b>/pipe/:topic</b></code> open a pipe to a topic (with upgrade)</summary>

### Parameters

| name  | type     | data type | description                |
| ----- | -------- | --------- | -------------------------- |
| topic | required | string    | topic name to subscribe to |

### Query Parameters

| name   | type     | data type | description                                   |
| ------ | -------- | --------- | --------------------------------------------- |
| status | optional | boolean   | Receive pipe status messages                  |
| replay | optional | boolean   | Whether or not to replay message              |
| binary | optional | string    | Connect the client as binary                  |
| access | optional | string    | Comma separated list of permissible accessors |

### Responses

| http code | content-type | response            |
| --------- | ------------ | ------------------- |
| `101`     | `N/A`        | Switching Protocols |

### Example websocat

```bash
websocat wss://pipe.pico.sh/pipe/test
```

</details>

## Subscribe to a topic

<details>
 <summary><code>GET</code> <code><b>/topic/:topic</b></code> subscribe to a topic</summary>

### Parameters

| name  | type     | data type | description                |
| ----- | -------- | --------- | -------------------------- |
| topic | required | string    | topic name to subscribe to |

### Query Parameters

| name    | type     | data type | description                                                 |
| ------- | -------- | --------- | ----------------------------------------------------------- |
| persist | optional | boolean   | Persist the subscription after the publisher closes         |
| block   | optional | boolean   | Block writes until a subscriber is available (default true) |
| access  | optional | string    | Comma separated list of permissible accessors               |
| mime    | optional | string    | Content type to return to the client                        |

### Responses

| http code | content-type                                         | response                                        |
| --------- | ---------------------------------------------------- | ----------------------------------------------- |
| `200`     | `text/plain;charset=UTF-8` or `mime` query parameter | Subscription data. Will hang until a pub occurs |

### Example cURL

```bash
curl -vvv https://pipe.pico.sh/topic/test?persist=true
```

</details>

## Publish to a topic

<details>
 <summary><code>POST</code> <code><b>/topic/:topic</b></code> publish to a topic</summary>

### Parameters

| name  | type     | data type | description                |
| ----- | -------- | --------- | -------------------------- |
| topic | required | string    | topic name to subscribe to |

### Query Parameters

| name   | type     | data type | description                                   |
| ------ | -------- | --------- | --------------------------------------------- |
| access | optional | string    | Comma separated list of permissible accessors |

### Responses

| http code | content-type | response            |
| --------- | ------------ | ------------------- |
| `200`     |              | No content returned |

### Example cURL

```bash
curl -vvvv https://pipe.pico.sh/topic/test -d "hello"
```

</details>

## Subscribe to a topic (pubsub)

<details>
 <summary><code>GET</code> <code><b>/pubsub/:topic</b></code> subscribe to a topic</summary>

### Parameters

| name  | type     | data type | description                |
| ----- | -------- | --------- | -------------------------- |
| topic | required | string    | topic name to subscribe to |

### Query Parameters

| name    | type     | data type | description                                                 |
| ------- | -------- | --------- | ----------------------------------------------------------- |
| persist | optional | boolean   | Persist the subscription after the publisher closes         |
| block   | optional | boolean   | Block writes until a subscriber is available (default true) |
| access  | optional | string    | Comma separated list of permissible accessors               |
| mime    | optional | string    | Content type to return to the client                        |

### Responses

| http code | content-type                                         | response                                        |
| --------- | ---------------------------------------------------- | ----------------------------------------------- |
| `200`     | `text/plain;charset=UTF-8` or `mime` query parameter | Subscription data. Will hang until a pub occurs |

### Example cURL

```bash
curl -vvv https://pipe.pico.sh/pubsub/test?persist=true
```

</details>

## Publish to a topic (pubsub)

<details>
 <summary><code>POST</code> <code><b>/pubsub/:topic</b></code> publish to a topic</summary>

### Parameters

| name  | type     | data type | description                |
| ----- | -------- | --------- | -------------------------- |
| topic | required | string    | topic name to subscribe to |

### Query Parameters

| name   | type     | data type | description                                   |
| ------ | -------- | --------- | --------------------------------------------- |
| access | optional | string    | Comma separated list of permissible accessors |

### Responses

| http code | content-type | response            |
| --------- | ------------ | ------------------- |
| `200`     |              | No content returned |

### Example cURL

```bash
curl -vvvv https://pipe.pico.sh/pubsub/test -d "hello"
```

</details>

# `pipemgr`

[pipemgr](https://github.com/picosh/pipemgr) is a docker image that will listen for logs from other running containers and pipe their logs through a topic called `container-drain`.

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

You must always pipe **something** into pub or else it will block indefinitely until the process is killed. However, you can provide a flag to send an empty message: `pub topic -e`.

# Inspiration

A special thanks to [patchbay.pub](https://patchbay.pub) for our inspiration.

# Latest posts

2024-10-06 [pipe: our pubsub ssh service](https://blog.pico.sh/ann-022-pubsub)

<hr />
<div class="flex flex-col items-center justify-center">
  <p>Create an account using only your SSH key.</p>
  <a href="/getting-started" class="btn-link">Get Started</a>
</div>
