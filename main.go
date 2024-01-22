package main

import (
	"github.com/picosh/pdocs"
)

func main() {
	pager := pdocs.Pager("./posts")
	sitemap := []*pdocs.Sitemap{
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
			Href: "/getting-started",
			Page: pager("getting-started.md"),
			Tag: "Help",
		},
		{
			Text: "How it works",
			Href: "/how-it-works",
			Page: pager("how-it-works.md"),
			Tag: "Help",
		},
		{
			Text: "File uploads",
			Href: "/file-uploads",
			Page: pager("file-uploads.md"),
			Tag: "Help",
		},
		{
			Text: "Custom domains",
			Href: "/custom-domains",
			Page: pager("custom-domains.md"),
			Tag: "Help",
		},
		{
			Text: "Pages",
			Href: "/pgs",
			Page: pager("pgs.md"),
			Tag: "Services",
		},
		{
			Text: "Prose",
			Href: "/prose",
			Page: pager("prose.md"),
			Tag: "Services",
		},
		{
			Text: "Pastes",
			Href: "/pastes",
			Page: pager("pastes.md"),
			Tag: "Services",
		},
		{
			Text: "Images",
			Href: "/imgs",
			Page: pager("imgs.md"),
			Tag: "Services",
		},
		{
			Text: "Feeds",
			Href: "/feeds",
			Page: pager("feeds.md"),
			Tag: "Services",
		},
		{
			Text: "Tunnels",
			Href: "/tuns",
			Page: pager("tuns.md"),
			Tag: "Services",
		},
		{
			Text: "IRC",
			Href: "/irc",
			Page: pager("irc.md"),
			Children: []*pdocs.Sitemap{
				{Text: "Web"},
				{Text: "Senpai"},
				{Text: "Connect to Libera"},
			},
			Tag: "Community",
		},
		{
			Text: "FAQ",
			Href: "/faq",
			Page: pager("faq.md"),
			Tag: "Help",
		},
		{
			Text: "Plain text lists",
			Href: "/plain-text-lists",
			Page: pager("plain-text-lists.md"),
			Tag: "Spec",
		},
		{
			Text: "About us",
			Href: "/about",
			Page: pager("about.md"),
			Tag: "About",
		},
		{
			Text: "Abuse",
			Href: "/abuse",
			Page: pager("abuse.md"),
			Tag: "About",
		},
		{
			Text: "Operations",
			Href: "/ops",
			Page: pager("ops.md"),
			Tag: "About",
		},
		{
			Text: "Privacy policy",
			Href: "/privacy",
			Page: pager("privacy.md"),
			Tag: "About",
		},
	}

	config := &pdocs.DocConfig{
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
