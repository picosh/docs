---
title: Analytics
description: how to get analytics for your blog and sites
---

> Read our
> [privacy policy on storing user analytics](https://pico.sh/privacy#analytics).

For [pico+](/plus) users we provide the ability to view user analytics. This
feature is opt-in so you need to first enable analytics:

```bash
ssh pico.sh
# -> Settings
# -> Enable analtics [OK]
```

# Notes

- We do **not** delete your data when toggling analytics, it's non-destructive
- We keep analytics for up to 1-year

# What metrics are captured?

- Unique visitors
- Top Referrers
- Top URLs

# Blog

```bash
ssh prose.sh stats
```

```bash
Top URLs
╭──────────────────────────────────────────────────────────┬───────────────────╮
│URL                                                       │Count              │
├──────────────────────────────────────────────────────────┼───────────────────┤
│/why-structured-concurrency                               │31                 │
│/                                                         │23                 │
│/what-is-the-smol-web                                     │9                  │
│/zig-devlog-ep1                                           │7                  │
│/what-is-starfx                                           │7                  │
│/typescript-terrible-for-library-developers               │6                  │
│/front-end-complexity                                     │6                  │
│/consciousness-is-wet                                     │5                  │
│/homelab                                                  │5                  │
│/neovimcraft-configs-search                               │5                  │
╰──────────────────────────────────────────────────────────┴───────────────────╯
Top Referers
╭──────────────────────────────────────────────┬───────────────────────────────╮
│URL                                           │Count                          │
├──────────────────────────────────────────────┼───────────────────────────────┤
│bower.sh                                      │11                             │
│erock.io                                      │3                              │
│neovimcraft.com                               │2                              │
│www.google.com                                │2                              │
│github.com                                    │2                              │
│kagi.com                                      │2                              │
│www.linkedin.com                              │1                              │
│metager.de                                    │1                              │
│www.codeproject.com                           │1                              │
│www.google.co.uk                              │1                              │
╰──────────────────────────────────────────────┴───────────────────────────────╯
Unique Visitors this Month
╭─────────────────────────────────────────┬────────────────────────────────────╮
│Date                                     │Unique Visitors                     │
├─────────────────────────────────────────┼────────────────────────────────────┤
│2024-09-01T00:00:00Z                     │64                                  │
│2024-09-02T00:00:00Z                     │6                                   │
╰─────────────────────────────────────────┴────────────────────────────────────╯
```

```bash
ssh prose.sh {post-slug}
```

```bash
ssh prose.sh stats why-structured-concurrency

Top URLs
╭──────────────────────────────────────────────────┬───────────────────────────╮
│URL                                               │Count                      │
├──────────────────────────────────────────────────┼───────────────────────────┤
│/why-structured-concurrency                       │31                         │
╰──────────────────────────────────────────────────┴───────────────────────────╯
Top Referers
╭─────────────────────────────────────────────┬────────────────────────────────╮
│URL                                          │Count                           │
├─────────────────────────────────────────────┼────────────────────────────────┤
│kagi.com                                     │2                               │
│erock.io                                     │1                               │
│github.com                                   │1                               │
│bower.sh                                     │1                               │
│www.google.co.uk                             │1                               │
│www.inoreader.com                            │1                               │
│www.linkedin.com                             │1                               │
╰─────────────────────────────────────────────┴────────────────────────────────╯
Unique Visitors this Month
╭─────────────────────────────────────────┬────────────────────────────────────╮
│Date                                     │Unique Visitors                     │
├─────────────────────────────────────────┼────────────────────────────────────┤
│2024-09-01T00:00:00Z                     │28                                  │
│2024-09-02T00:00:00Z                     │4                                   │
╰─────────────────────────────────────────┴────────────────────────────────────╯
```

# Pages

```bash
ssh pgs.sh stats
```

```bash
╭───────────────────────┬────────────────────────┬──────────────────────┬──────────────────────────╮
│Used (GB)              │Quota (GB)              │Used (%)              │Projects (#)              │
├───────────────────────┼────────────────────────┼──────────────────────┼──────────────────────────┤
│0.3919                 │20.0000                 │1.9595                │29                        │
╰───────────────────────┴────────────────────────┴──────────────────────┴──────────────────────────╯

Visitor Analytics this Month

Top URLs
╭──────────────────────────┬──────────────────────────────────────────────────┬────────────────────╮
│Project                   │URL                                               │Count               │
├──────────────────────────┼──────────────────────────────────────────────────┼────────────────────┤
│neovimcraft               │/                                                 │142                 │
│neovimcraft               │/plugin/siduck76/NvChad/                          │15                  │
│neovimcraft               │/plugin/tadmccorkle/markdown.nvim/                │14                  │
│neovimcraft               │/plugin/NvChad/nvim-colorizer.lua/                │13                  │
│starfx-prod               │/                                                 │13                  │
│neovimcraft               │/plugin/folke/noice.nvim/                         │12                  │
│neovimcraft               │/plugin/catppuccin/nvim/                          │9                   │
│neovimcraft               │/plugin/hrsh7th/nvim-cmp/index.html               │9                   │
╰──────────────────────────┴──────────────────────────────────────────────────┴────────────────────╯
Top Referers
╭──────────────────────────────────────────────────────┬───────────────────────────────────────────╮
│URL                                                   │Count                                      │
├──────────────────────────────────────────────────────┼───────────────────────────────────────────┤
│www.google.com                                        │185                                        │
│neovimcraft.com                                       │82                                         │
│duckduckgo.com                                        │44                                         │
│yandex.ru                                             │15                                         │
│search.brave.com                                      │14                                         │
│bower.sh                                              │14                                         │
│github.com                                            │12                                         │
│www.bing.com                                          │9                                          │
│lobste.rs                                             │6                                          │
│starfx.bower.sh                                       │6                                          │
╰──────────────────────────────────────────────────────┴───────────────────────────────────────────╯
Unique Visitors by Site
╭──────────────────────────────────────────────────────┬───────────────────────────────────────────╮
│URL                                                   │Count                                      │
├──────────────────────────────────────────────────────┼───────────────────────────────────────────┤
│neovimcraft                                           │902                                        │
│starfx-prod                                           │30                                         │
│git-starfx                                            │13                                         │
│docs-local                                            │5                                          │
│git                                                   │2                                          │
│git-neovimcraft                                       │2                                          │
╰──────────────────────────────────────────────────────┴───────────────────────────────────────────╯
```

```bash
ssh pgs.sh stats {project-name}
```

```bash
ssh pgs.sh stats neovimcraft

Top URLs
╭─────────────────────────────────────────────────────────────────┬────────────────────────────────╮
│URL                                                              │Count                           │
├─────────────────────────────────────────────────────────────────┼────────────────────────────────┤
│/                                                                │142                             │
│/plugin/siduck76/NvChad/                                         │15                              │
│/plugin/tadmccorkle/markdown.nvim/                               │14                              │
│/plugin/NvChad/nvim-colorizer.lua/                               │13                              │
│/plugin/folke/noice.nvim/                                        │12                              │
│/plugin/hrsh7th/nvim-cmp/index.html                              │9                               │
│/plugin/catppuccin/nvim/                                         │9                               │
│/plugin/yetone/avante.nvim/                                      │9                               │
│/plugin/iamcco/markdown-preview.nvim/                            │8                               │
│/plugin/folke/tokyonight.nvim/                                   │8                               │
╰─────────────────────────────────────────────────────────────────┴────────────────────────────────╯
Top Referers
╭──────────────────────────────────────────────────────┬───────────────────────────────────────────╮
│URL                                                   │Count                                      │
├──────────────────────────────────────────────────────┼───────────────────────────────────────────┤
│www.google.com                                        │183                                        │
│neovimcraft.com                                       │80                                         │
│duckduckgo.com                                        │44                                         │
│yandex.ru                                             │15                                         │
│search.brave.com                                      │14                                         │
│www.bing.com                                          │9                                          │
│github.com                                            │9                                          │
│www.ecosia.org                                        │3                                          │
│www.reddit.com                                        │3                                          │
│cn.bing.com                                           │3                                          │
╰──────────────────────────────────────────────────────┴───────────────────────────────────────────╯
Unique Visitors this Month
╭───────────────────────────────────────────────────┬──────────────────────────────────────────────╮
│Date                                               │Unique Visitors                               │
├───────────────────────────────────────────────────┼──────────────────────────────────────────────┤
│2024-09-01T00:00:00Z                               │428                                           │
│2024-09-02T00:00:00Z                               │46                                            │
╰───────────────────────────────────────────────────┴──────────────────────────────────────────────╯
```
