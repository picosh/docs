package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/picosh/pdocs"
)

// staticHash hashes the static assets, so pages' cache-busting query only
// changes with them and unchanged pages rebuild identically.
func staticHash() (string, error) {
	files, err := filepath.Glob("./static/*")
	if err != nil {
		return "", err
	}

	h := sha256.New()
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return "", err
		}
		h.Write([]byte(f))
		h.Write(data)
	}

	return hex.EncodeToString(h.Sum(nil))[:8], nil
}

type LogoData struct {
	Src string
}

func main() {
	pager := pdocs.Pager("./posts")
	sitemap := &pdocs.Sitemap{
		Children: []*pdocs.Sitemap{
			{
				Text: "Home",
				Href: "/",
				Page: pager("home.md"),
			},
			{
				Text: "Sitemap",
				Href: "/sitemap",
				Page: pager("sitemap.md"),
			},

			{
				Text: "Getting started",
				Children: []*pdocs.Sitemap{
					{
						Text: "Create account",
						Href: "/getting-started",
						Page: pager("getting-started.md"),
						Data: LogoData{
							Src: "/logo.svg",
						},
					},
					{
						Text: "Invites",
						Href: "/invites",
						Page: pager("invites.md"),
					},
					{
						Text: "File uploads",
						Href: "/file-uploads",
						Page: pager("file-uploads.md"),
					},
					{
						Text: "UI",
						Href: "/ui",
						Page: pager("ui.md"),
					},
				},
			},

			{
				Text: "Services",
				Children: []*pdocs.Sitemap{
					{
						Text: "Pages",
						Href: "/pgs",
						Page: pager("pgs.md"),
						Data: LogoData{
							Src: "/logo-pgs.svg",
						},
					},
					{
						Text: "Tuns",
						Href: "/tuns",
						Page: pager("tuns.md"),
						Data: LogoData{
							Src: "/logo-tuns.svg",
						},
					},
					{
						Text: "Pipe",
						Href: "/pipe",
						Page: pager("pipe.md"),
						Data: LogoData{
							Src: "/logo-pipe.svg",
						},
					},
					{
						Text: "Prose",
						Href: "/prose",
						Page: pager("prose.md"),
						Data: LogoData{
							Src: "/logo-prose.svg",
						},
					},
					{
						Text: "Patchbin",
						Href: "/patchbin",
						Page: pager("patchbin.md"),
						Data: LogoData{
							Src: "/logo-patchbin.svg",
						},
					},
					{
						Text: "Feeds",
						Href: "/feeds",
						Page: pager("feeds.md"),
					},
					{
						Text: "Pastes",
						Href: "/pastes",
						Page: pager("pastes.md"),
					},
					{
						Text: "IRC Bouncer",
						Href: "/bouncer",
						Page: pager("bouncer.md"),
					},
				},
			},

			{
				Text: "Advanced",
				Children: []*pdocs.Sitemap{
					{
						Text: "Custom domains",
						Href: "/custom-domains",
						Page: pager("custom-domains.md"),
					},
					{
						Text: "Access control",
						Href: "/access-control",
						Page: pager("access-control.md"),
					},
					{
						Text: "Images",
						Href: "/images",
						Page: pager("images.md"),
					},
					{
						Text: "API Tokens",
						Href: "/api-token",
						Page: pager("api-token.md"),
					},
					{
						Text: "Analytics",
						Href: "/analytics",
						Page: pager("analytics.md"),
					},
					{
						Text: "Logs",
						Href: "/logs",
						Page: pager("logs.md"),
					},
					{
						Text: "Access Logs",
						Href: "/access-logs",
						Page: pager("access-logs.md"),
					},
					{
						Text: "Host Keys",
						Href: "/host-keys",
						Page: pager("host-keys.md"),
					},
				},
			},

			{
				Text: "Help",
				Children: []*pdocs.Sitemap{
					{
						Text: "FAQ",
						Href: "/faq",
						Page: pager("faq.md"),
					},
					{
						Text: "Regions",
						Href: "/regions",
						Page: pager("regions.md"),
					},
					{
						Text: "IRC",
						Href: "/irc",
						Page: pager("irc.md"),
					},
				},
			},

			{
				Text: "X",
				Children: []*pdocs.Sitemap{
					{
						Text: "Styles",
						Href: "/styles",
						Page: pager("styles.md"),
					},
				},
			},

			{
				Text: "About",
				Children: []*pdocs.Sitemap{
					{
						Text: "About us",
						Href: "/about",
						Page: pager("about.md"),
					},
					{
						Text: "Contact us",
						Href: "/contact",
						Page: pager("contact.md"),
					},
					{
						Text: "Abuse",
						Href: "/abuse",
						Page: pager("abuse.md"),
					},
					{
						Text: "Operations",
						Href: "/ops",
						Page: pager("ops.md"),
					},
					{
						Text: "Privacy policy",
						Href: "/privacy",
						Page: pager("privacy.md"),
					},
					{
						Text: "pico+",
						Href: "/plus",
						Page: pager("plus.md"),
					},
				},
			},
		},
	}

	cacheID, err := staticHash()
	if err != nil {
		panic(err)
	}

	logger := slog.Default()
	config := &pdocs.DocConfig{
		Logger:   logger,
		Sitemap:  sitemap,
		Out:      "./public",
		Tmpl:     "./tmpl",
		PageTmpl: "post.page.tmpl",
		CacheId:  cacheID,
	}

	err = config.GenSite()
	if err != nil {
		panic(err)
	}
}
