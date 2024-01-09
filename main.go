package main

import (
	"github.com/picosh/pdocs"
)

func main() {
	pager := pdocs.Pager("./posts")
	sitemap := []*pdocs.Sitemap{
		{
			Text: "Marketing",
			Href: "/",
			Page: pager("marketing.md"),
		},
		{
			Text: "Getting Started",
			Href: "/getting-started",
			Page: pager("getting-started.md"),
		},
		{
			Text: "How it Works",
			Href: "/how-it-works",
			Page: pager("how-it-works.md"),
		},
		{
			Text: "File Uploads",
			Href: "/file-uploads",
			Page: pager("file-uploads.md"),
		},
		{
			Text: "Custom Domains",
			Href: "/custom-domains",
			Page: pager("custom-domains.md"),
		},
		{
			Text: "Pages",
			Href: "/pgs",
			Page: pager("pgs.md"),
		},
		{
			Text: "Prose",
			Href: "/prose",
			Page: pager("prose.md"),
		},
		{
			Text: "Pastes",
			Href: "/pastes",
			Page: pager("pastes.md"),
		},
		{
			Text: "Images",
			Href: "/imgs",
			Page: pager("imgs.md"),
		},
		{
			Text: "Feeds",
			Href: "/feeds",
			Page: pager("feeds.md"),
		},
		{
			Text: "Lists",
			Href: "/lists",
			Page: pager("lists.md"),
		},
		{
			Text: "Tunnels",
			Href: "/tuns",
			Page: pager("tuns.md"),
		},
		{
			Text: "Community",
			Href: "/comms",
			Page: pager("comms.md"),
		},
		{
			Text: "FAQ",
			Href: "/faq",
			Page: pager("faq.md"),
		},
		{
			Text: "Plain Text Lists",
			Href: "/plain-text-lists",
			Page: pager("plain-text-lists.md"),
		},
		{
			Text: "Operations",
			Href: "/ops",
			Page: pager("ops.md"),
		},
		{
			Text: "Privacy Policy",
			Href: "/privacy",
			Page: pager("privacy.md"),
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
