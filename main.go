package main

import (
	"github.com/picosh/pdocs"
)

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
				Href: "/getting-started",
				Page: pager("getting-started.md"),
				Tag:  "Help",
				Children: []*pdocs.Sitemap{
					pdocs.AnchorTagSitemap("Create your account with Public-Key Cryptography"),
				},
			},
			{
				Text: "How it works",
				Href: "/how-it-works",
				Page: pager("how-it-works.md"),
				Tag:  "Help",
			},
			{
				Text: "File uploads",
				Href: "/file-uploads",
				Page: pager("file-uploads.md"),
				Tag:  "Help",
				Children: []*pdocs.Sitemap{
					pdocs.AnchorTagSitemap("How do I upload files"),
					pdocs.AnchorTagSitemap("How do I update files"),
					pdocs.AnchorTagSitemap("How do I delete files"),
					pdocs.AnchorTagSitemap("How do I download files"),
				},
			},
			{
				Text:   "pico+",
				Href:   "/plus",
				Page:   pager("plus.md"),
				Tag:    "pico+",
				Hidden: true,
			},
			{
				Text:   "Join",
				Href:   "/join",
				Page:   pager("plus-join.md"),
				Tag:    "pico+",
				Hidden: true,
			},
			{
				Text: "Custom domains",
				Href: "/custom-domains",
				Page: pager("custom-domains.md"),
				Tag:  "Help",
				Children: []*pdocs.Sitemap{
					pdocs.AnchorTagSitemap("prose.sh"),
					pdocs.AnchorTagSitemap("pgs.sh"),
					pdocs.AnchorTagSitemap("My DNS does not support CNAME flattening"),
				},
			},
			{
				Text: "pgs.sh",
				Href: "/pgs",
				Page: pager("pgs.md"),
				Tag:  "Services",
				Children: []*pdocs.Sitemap{
					pdocs.AnchorTagSitemap("CLI Reference"),
					pdocs.AnchorTagSitemap("What file types are supported"),
					pdocs.AnchorTagSitemap("user-defined-redirects"),
					pdocs.AnchorTagSitemap("single-page-applications"),
				},
			},
			{
				Text: "tuns.sh",
				Href: "/tuns",
				Page: pager("tuns.md"),
				Tag:  "Services",
			},
			{
				Text: "imgs.sh",
				Href: "/imgs",
				Page: pager("imgs.md"),
				Tag:  "Services",
			},
			{
				Text: "prose.sh",
				Href: "/prose",
				Page: pager("prose.md"),
				Tag:  "Services",
				Children: []*pdocs.Sitemap{
					pdocs.AnchorTagSitemap("Metadata"),
					pdocs.AnchorTagSitemap("How can I add a footer to all of my posts"),
					pdocs.AnchorTagSitemap("How can I customize my blog page"),
					pdocs.AnchorTagSitemap("How can I change the theme of my blog"),
					pdocs.AnchorTagSitemap("How can I change the layout of my blog"),
				},
			},
			{
				Text: "pastes.sh",
				Href: "/pastes",
				Page: pager("pastes.md"),
				Tag:  "Services",
				Children: []*pdocs.Sitemap{
					pdocs.AnchorTagSitemap("Pipe Support"),
					pdocs.AnchorTagSitemap("How do I set expiration date"),
					pdocs.AnchorTagSitemap("How do I unlist a paste"),
				},
			},
			{
				Text: "feeds.sh",
				Href: "/feeds",
				Page: pager("feeds.md"),
				Tag:  "Services",
				Children: []*pdocs.Sitemap{
					pdocs.AnchorTagSitemap("Digest interval options"),
					pdocs.AnchorTagSitemap("Inline content"),
					pdocs.AnchorTagSitemap("Can I create multiple email digests?"),
					pdocs.AnchorTagSitemap("Can I fetch Reddit RSS feeds"),
				},
			},
			{
				Text: "Images",
				Href: "/images",
				Page: pager("images.md"),
				Tag:  "Help",
				Children: []*pdocs.Sitemap{
					pdocs.AnchorTagSitemap("What file types are supported"),
					pdocs.AnchorTagSitemap("Image manipulation"),
				},
			},
			{
				Text: "IRC",
				Href: "/irc",
				Page: pager("irc.md"),
				Children: []*pdocs.Sitemap{
					pdocs.AnchorTagSitemap("Web"),
					pdocs.AnchorTagSitemap("Senpai"),
					pdocs.AnchorTagSitemap("Connect to Libera"),
				},
				Tag: "Community",
			},
			{
				Text: "API Tokens",
				Href: "/api-token",
				Page: pager("api-token.md"),
				Children: []*pdocs.Sitemap{
					pdocs.AnchorTagSitemap("Web"),
					pdocs.AnchorTagSitemap("Senpai"),
					pdocs.AnchorTagSitemap("Connect to Libera"),
				},
				Tag: "Help",
			},
			{
				Text: "FAQ",
				Href: "/faq",
				Page: pager("faq.md"),
				Tag:  "Help",
				Children: []*pdocs.Sitemap{
					pdocs.AnchorTagSitemap("Permission denied when using SSH"),
					pdocs.AnchorTagSitemap("Generating a new SSH key"),
					pdocs.AnchorTagSitemap("Why do I provide my username when using SSH"),
					pdocs.AnchorTagSitemap("How can I use git to sync my content"),
				},
			},
			{
				Text: "Lab",
				Href: "/lab",
				Page: pager("lab.md"),
				Tag:  "Lab",
			},
			{
				Text: "Plain text lists",
				Href: "/plain-text-lists",
				Page: pager("plain-text-lists.md"),
				Tag:  "Spec",
			},
			{
				Text: "About us",
				Href: "/about",
				Page: pager("about.md"),
				Tag:  "About",
			},
			{
				Text: "Abuse",
				Href: "/abuse",
				Page: pager("abuse.md"),
				Tag:  "About",
			},
			{
				Text: "Operations",
				Href: "/ops",
				Page: pager("ops.md"),
				Tag:  "About",
			},
			{
				Text: "Privacy policy",
				Href: "/privacy",
				Page: pager("privacy.md"),
				Tag:  "About",
			},
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
