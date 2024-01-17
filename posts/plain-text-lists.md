---
title: Plain Text Lists
description: Our spec for plain text lists
keywords: [pico, lists, spec, plain, text]
---

- Version: **2022.08.05.dev**
- Status: **Draft**
- Author: **Eric Bower**

The goal of this specification is to understand how we render plain text lists.
The overall design of this format is to be easy to parse and render.

The format is line-oriented, and a satisfactory rendering can be achieved with a
single pass of a document, processing each line independently. As per gopher,
links can only be displayed one per line, encouraging neat, list-like structure.

Feedback on any part of this is extremely welcome, please email
[hello@pico.sh](mailto:hello@pico.sh).

The source code for our parser can be found
[here](https://github.com/picosh/pico/blob/85ad4b81370427925328ab24fa568f044fd624ab/shared/listparser.go).

# Parameters

As a subtype of the top-level media type "text", "text/plain" inherits the
"charset" parameter defined in
[RFC 2046](https://datatracker.ietf.org/doc/html/rfc2046#section-4.1). The
default value of "charset" is "UTF-8" for "text" content.

# Line orientation

As mentioned, the text format is line-oriented. Each line of a document has a
single "line type". It is possible to unambiguously determine a line's type
purely by inspecting its first (3) characters. A line's type determines the
manner in which it should be presented to the user. Any details of presentation
or rendering associated with a particular line type are strictly limited in
scope to that individual line.

# File extension

Plain text lists only supports the `.txt` file extension and will ignore all
other file extensions.

# List item

List items are separated by newline characters `\n`. Each list item is on its
own line. A list item does not require any special formatting. A list item can
contain as much text as it wants. We encourage soft wrapping for readability in
your editor of choice. Hard wrapping is not permitted as it will create a new
list item.

Empty lines will be completely removed and not rendered to the end user.

# Hyperlinks

Hyperlinks are denoted by the prefix `=>`. The following text should then be the
hyperlink.

```
=> https://lists.sh
```

Optionally you can supply the hyperlink text immediately following the link.

```
=> https://lists.sh microblog for lists
```

# Nested lists

Users can create nested lists. Tabbing a list will nest it under the list item
directly above it. Both tab character `\t` or whitespace as tabs are permitted.
Do not mix tabs and spaces because the nesting will yield unexpected results.

```
first item
    second item
        third item
last item
```

# Images

List items can be represented as images by prefixing the line with `=<`.

```
=< https://i.imgur.com/iXMNUN5.jpg
```

Optionally you can supply the image alt text immediately following the link.

```
=< https://i.imgur.com/iXMNUN5.jpg I use arch, btw
```

# Headers

List items can be represented as headers. We support two headers currently.
Headers will end the previous list and then create a new one after it. This
allows a single document to contain multiple lists.

```
# Header One
## Header Two
```

# Blockquotes

List items can be represented as blockquotes.

```
> This is a blockquote.
```

# Preformatted

List items can be represented as preformatted text where newline characters are
not considered part of new list items. They can be represented by prefixing the
line with three backticks

````
```
#!/usr/bin/env bash

set -x

echo "this is a preformatted list item!
```
````

You must also close the preformatted text with another backtick on its own line.
The next example with **NOT** work.

````
```
#!/usr/bin/env bash

echo "This will not render properly"```
````

# Variables

Variables allow us to store metadata within our system. Variables are list items
with key value pairs denoted by `=:` followed by the key, a whitespace
character, and then the value.

```
=: publish_at 2022-04-20
```

These variables will not be rendered to the user inside the list.

List of available variables:

```
=: title Hello World
=: description A fine description
=: publish_at 2022-04-20
=: tags feature, announcement
=: list_type none
```

`list_type` value gets sent directly to css property
[list-style-type](https://developer.mozilla.org/en-US/docs/Web/CSS/list-style-type)
