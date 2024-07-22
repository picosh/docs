---
title: Analytics
description: how to get analytics for your blog and sites
---

For [pico+](/plus) users we provide the ability to view user analytics. This
feature is opt-in so you need to first enable analytics:

```bash
ssh pico.sh
# -> Settings
# -> Enable analtics [OK]
```

# What metrics are captured?

- Unique visitors
- Top Referrers
- Top URLs

# Blog

```bash
ssh prose.sh stats
```

```bash
ssh prose.sh post-slug
```

# Pages

```bash
ssh pgs.sh stats
```

```bash
ssh pgs.sh stats project-name
```
