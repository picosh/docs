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
				Children: []*pdocs.Sitemap{
					{
						Text: "Create account",
						Href: "/getting-started",
						Page: pager("getting-started.md"),
						Children: []*pdocs.Sitemap{
							pdocs.AnchorTagSitemap("Create your account with Public-Key Cryptography"),
						},
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
						Children: []*pdocs.Sitemap{
							pdocs.AnchorTagSitemap("How do I upload files"),
							pdocs.AnchorTagSitemap("How do I update files"),
							pdocs.AnchorTagSitemap("How do I delete files"),
							pdocs.AnchorTagSitemap("How do I download files"),
						},
					},
					{
						Text: "UI",
						Href: "/ui",
						Page: pager("ui.md"),
						Children: []*pdocs.Sitemap{
							pdocs.AnchorTagSitemap("SSH TUI"),
							pdocs.AnchorTagSitemap("Web UI"),
							pdocs.AnchorTagSitemap("SSH Config"),
						},
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
						Children: []*pdocs.Sitemap{
							pdocs.AnchorTagSitemap("CLI Reference"),
							pdocs.AnchorTagSitemap("File denylist"),
							pdocs.AnchorTagSitemap("Access Control List"),
							pdocs.AnchorTagSitemap("Pretty URLs"),
							pdocs.AnchorTagSitemap("Custom Domains"),
							pdocs.AnchorTagSitemap("Custom Redirects and rewrites"),
							pdocs.AnchorTagSitemap("Custom Headers"),
							pdocs.AnchorTagSitemap("Single-Page Applications"),
							pdocs.AnchorTagSitemap("Reserved Username Project"),
							pdocs.AnchorTagSitemap("Content Security Policy"),
							pdocs.AnchorTagSitemap("Does pages have a CDN or muilti-region support"),
						},
					},
					{
						Text: "Tuns",
						Href: "/tuns",
						Page: pager("tuns.md"),
					},
					{
						Text: "Prose",
						Href: "/prose",
						Page: pager("prose.md"),
						Children: []*pdocs.Sitemap{
							pdocs.AnchorTagSitemap("How are blogs structured"),
							pdocs.AnchorTagSitemap("How can I customize my blog"),
							pdocs.AnchorTagSitemap("How can I customize a blog post"),
							pdocs.AnchorTagSitemap("Unlisted posts"),
							pdocs.AnchorTagSitemap("How can I add a footer to all of my posts"),
							pdocs.AnchorTagSitemap("How can I change the theme of my blog"),
							pdocs.AnchorTagSitemap("How can I change the default 404 page"),
						},
					},
					{
						Text: "Pastes",
						Href: "/pastes",
						Page: pager("pastes.md"),
						Children: []*pdocs.Sitemap{
							pdocs.AnchorTagSitemap("Pipe Support"),
							pdocs.AnchorTagSitemap("How do I set expiration date"),
							pdocs.AnchorTagSitemap("How do I unlist a paste"),
						},
					},
					{
						Text: "RSS-To-Email",
						Href: "/feeds",
						Page: pager("feeds.md"),
						Children: []*pdocs.Sitemap{
							pdocs.AnchorTagSitemap("Digest interval options"),
							pdocs.AnchorTagSitemap("Inline content"),
							pdocs.AnchorTagSitemap("Can I create multiple email digests?"),
							pdocs.AnchorTagSitemap("Can I fetch Reddit RSS feeds"),
						},
					},
					{
						Text: "Docker Registry",
						Href: "/imgs",
						Page: pager("imgs.md"),
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
						Children: []*pdocs.Sitemap{
							pdocs.AnchorTagSitemap("prose.sh"),
							pdocs.AnchorTagSitemap("pgs.sh"),
							pdocs.AnchorTagSitemap("My DNS does not support CNAME flattening"),
						},
					},
					{
						Text: "Images",
						Href: "/images",
						Page: pager("images.md"),
						Children: []*pdocs.Sitemap{
							pdocs.AnchorTagSitemap("What file types are supported"),
							pdocs.AnchorTagSitemap("Image manipulation"),
						},
					},
					{
						Text:     "API Tokens",
						Href:     "/api-token",
						Page:     pager("api-token.md"),
						Children: []*pdocs.Sitemap{},
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
						Children: []*pdocs.Sitemap{
							pdocs.AnchorTagSitemap("Why can't I login to pico"),
							pdocs.AnchorTagSitemap("How do I force the correct pico SSH key"),
							pdocs.AnchorTagSitemap("How do I generate an SSH key"),
							pdocs.AnchorTagSitemap("How can I setup my ssh-agent"),
							pdocs.AnchorTagSitemap("How can I use git to sync my content"),
							pdocs.AnchorTagSitemap("Can I create multiple pico accounts"),
							pdocs.AnchorTagSitemap("Can I associate multiple SSH keypairs to a single account"),
							pdocs.AnchorTagSitemap("Are there any bandwidth limitations"),
							pdocs.AnchorTagSitemap("How can I download a copy of all of my content"),
							pdocs.AnchorTagSitemap("How can I delete my content"),
							pdocs.AnchorTagSitemap("How can I delete my account with my content"),
							pdocs.AnchorTagSitemap("I lost my SSH private key how can I recover my account"),
						},
					},
					{
						Text: "IRC",
						Href: "/irc",
						Page: pager("irc.md"),
						Children: []*pdocs.Sitemap{
							pdocs.AnchorTagSitemap("Realtime chat"),
							pdocs.AnchorTagSitemap("Our public bouncer"),
						},
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
