package a2ui

import (
	"strings"
	"testing"
)

func TestStoreAppliesStream(t *testing.T) {
	messages, err := DecodeMessages(strings.NewReader(`
{"version":"v1.0","createSurface":{"surfaceId":"main"}}
{"version":"v1.0","updateComponents":{"surfaceId":"main","components":[{"id":"root","component":"Text","text":"Hello"}]}}
`))
	if err != nil {
		t.Fatalf("DecodeMessages: %v", err)
	}

	store := NewStore()
	for i, message := range messages {
		if err := store.Apply(message); err != nil {
			t.Fatalf("message %d: %v", i+1, err)
		}
	}

	surface, ok := store.Current()
	if !ok {
		t.Fatal("no current surface")
	}
	if surface.ID != "main" || surface.Components["root"].Text != "Hello" {
		t.Fatalf("unexpected surface: %+v", surface)
	}
}

func TestStoreRejectsBadMessages(t *testing.T) {
	store := NewStore()

	tests := []struct {
		name    string
		message Message
	}{
		{name: "wrong version", message: Message{Version: "v0.9"}},
		{name: "no payload", message: Message{Version: Version}},
		{
			name: "two payloads",
			message: Message{
				Version:          Version,
				CreateSurface:    &CreateSurface{SurfaceID: "main"},
				UpdateComponents: &UpdateComponents{SurfaceID: "main"},
			},
		},
		{name: "unknown surface", message: Message{Version: Version, UpdateComponents: &UpdateComponents{SurfaceID: "missing"}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := store.Apply(test.message); err == nil {
				t.Fatal("expected an error")
			}
		})
	}
}
