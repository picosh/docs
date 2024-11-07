package main

import (
	"log/slog"

	"github.com/picosh/pdocs"
)

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
					},
					{
						Text: "How it works",
						Href: "/how-it-works",
						Page: pager("how-it-works.md"),
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
						Href: "https://pipe.pico.sh",
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
						Text: "Pastes",
						Href: "/pastes",
						Page: pager("pastes.md"),
					},
					{
						Text: "RSS-To-Email",
						Href: "/feeds",
						Page: pager("feeds.md"),
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
						Text: "Web Tunnels",
						Href: "/tunnels",
						Page: pager("tunnels.md"),
					},
					{
						Text: "Lab",
						Href: "/lab",
						Page: pager("lab.md"),
					},
					{
						Text: "Plain text lists",
						Href: "/plain-text-lists",
						Page: pager("plain-text-lists.md"),
					},
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

	logger := slog.Default()
	config := &pdocs.DocConfig{
		Logger:   logger,
		Sitemap:  sitemap,
		Out:      "./public",
		Tmpl:     "./tmpl",
		PageTmpl: "post.page.tmpl",
	}

	err := config.GenSite()
	if err != nil {
		panic(err)
	}
}
