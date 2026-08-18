package htmlutil

import "testing"

func TestEscapeAndUnescape(t *testing.T) {
	src := `<a href="x&y">'z'</a>`
	esc := Escape(src)
	want := `&lt;a href=&#34;x&amp;y&#34;&gt;&#39;z&#39;&lt;/a&gt;`
	if esc != want {
		t.Errorf("Escape = %q, want %q", esc, want)
	}
	if got := Unescape(esc); got != src {
		t.Errorf("Unescape = %q, want %q", got, src)
	}
}

func TestUnescapeNamedEntities(t *testing.T) {
	if got := Unescape("&amp;&lt;&gt;&quot;&#39;"); got != `&<>"'` {
		t.Errorf("Unescape core = %q", got)
	}
	if got := Unescape("&nbsp;"); got != " " {
		t.Errorf("Unescape nbsp = %q", got)
	}
	if got := Unescape("&copy;"); got != "©" {
		t.Errorf("Unescape copy = %q", got)
	}
}

func TestCleanHtmlTag(t *testing.T) {
	if got := CleanHtmlTag("<b>hi</b>"); got != "hi" {
		t.Errorf("CleanHtmlTag = %q", got)
	}
	if got := CleanHtmlTag(`<p class="x">hello <br/> world</p>`); got != "hello  world" {
		t.Errorf("CleanHtmlTag = %q", got)
	}
	if got := CleanHtmlTag("no tags here"); got != "no tags here" {
		t.Errorf("CleanHtmlTag(plain) = %q", got)
	}
}

func TestRemoveHtmlAttr(t *testing.T) {
	in := `<a href="/x" onclick="evil" TITLE="t">link</a>`
	if got := RemoveHtmlAttr(in, "onclick"); got != `<a href="/x" TITLE="t">link</a>` {
		t.Errorf("RemoveHtmlAttr = %q", got)
	}
	// case-sensitive: TITLE is not removed by RemoveHtmlAttr
	if got := RemoveHtmlAttr(in, "title"); got != in {
		t.Errorf("RemoveHtmlAttr(title) should be case-sensitive: %q", got)
	}
	// the insensitive variant removes it
	gotI := RemoveHtmlAttrI(in, "TITLE")
	if gotI != `<a href="/x" onclick="evil">link</a>` {
		t.Errorf("RemoveHtmlAttrI = %q", gotI)
	}
}

func TestRemoveHtmlAttrUnquoted(t *testing.T) {
	in := `<img src=a.png onclick=evil>`
	if got := RemoveHtmlAttr(in, "onclick"); got != "<img src=a.png>" {
		t.Errorf("RemoveHtmlAttr(unquoted) = %q", got)
	}
}

func TestFilterRemovesScripts(t *testing.T) {
	in := `<div>ok</div><script>alert(1)</script><p>after</p>`
	if got := Filter(in); got != "<div>ok</div><p>after</p>" {
		t.Errorf("Filter(script) = %q", got)
	}
}

func TestFilterRemovesEventHandlers(t *testing.T) {
	in := `<a href="/ok" onclick="alert(1)">go</a>`
	if got := Filter(in); got != `<a href="/ok">go</a>` {
		t.Errorf("Filter(on*) = %q", got)
	}
}

func TestFilterRemovesJavascriptProtocol(t *testing.T) {
	in := `<a href="javascript:alert(1)">x</a>`
	got := Filter(in)
	if got != `<a href="#">x</a>` && got != `<a href="">x</a>` {
		t.Errorf("Filter(javascript:) = %q", got)
	}
}

func TestToHex(t *testing.T) {
	if got := ToHex("a"); got != "a" {
		t.Errorf("ToHex(ascii) = %q", got)
	}
	if got := ToHex("你"); got != "&#x4F60;" {
		t.Errorf("ToHex(cjk) = %q", got)
	}
	// control chars below 0x80 are left untouched (Hutool semantics)
	if got := ToHex("\t"); got != "\t" {
		t.Errorf("ToHex(tab) = %q", got)
	}
}

func TestUnwrapHtml(t *testing.T) {
	if got := UnwrapHtml("<p>hello</p>"); got != "hello" {
		t.Errorf("UnwrapHtml = %q", got)
	}
	if got := UnwrapHtml(`<div class="x"><b>bold</b></div>`); got != "<b>bold</b>" {
		t.Errorf("UnwrapHtml nested = %q", got)
	}
	if got := UnwrapHtml("plain"); got != "plain" {
		t.Errorf("UnwrapHtml(plain) = %q", got)
	}
	if got := UnwrapHtml("<p>one</p><p>two</p>"); got != "one</p><p>two" {
		t.Errorf("UnwrapHtml multi = %q", got)
	}
}

func TestEncodeDecodeUnicode(t *testing.T) {
	src := "你好go"
	// Hutool encodeUnicode semantics: only non-ASCII runes encoded,
	// lower-case 4-digit hex. The expected value is the literal
	// backslash-u-4-f-6-0 sequence.
	if got := EncodeUnicode(src); got != "\\u4f60\\u597dgo" {
		t.Errorf("EncodeUnicode = %q", got)
	}
	// decode round-trips
	if got := DecodeUnicode("\\u4f60\\u597dgo"); got != "你好go" {
		t.Errorf("DecodeUnicode = %q", got)
	}
	// uppercase hex also decodes
	if got := DecodeUnicode("\\u4F60"); got != "你" {
		t.Errorf("DecodeUnicode upper = %q", got)
	}
	if got := DecodeUnicode(`plain`); got != "plain" {
		t.Errorf("DecodeUnicode passthrough = %q", got)
	}
}