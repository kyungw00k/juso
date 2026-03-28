package i18n

import "testing"

func TestT_ReturnsMessage(t *testing.T) {
	current = en
	got := T(MsgRootShort)
	want := "Korean postal code lookup CLI"
	if got != want {
		t.Errorf("T(MsgRootShort) = %q, want %q", got, want)
	}
}

func TestT_UnknownKey(t *testing.T) {
	current = en
	got := T("UnknownKey")
	if got != "UnknownKey" {
		t.Errorf("T(unknown) = %q, want key string back", got)
	}
}

func TestTf_WithArgs(t *testing.T) {
	current = en
	got := Tf(MsgCacheEntries, 42)
	want := "Cache entries: 42"
	if got != want {
		t.Errorf("Tf(MsgCacheEntries, 42) = %q, want %q", got, want)
	}
}
