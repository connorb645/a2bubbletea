// Package a2ui holds the A2UI v1.0 messages and surfaces used by the renderer.
package a2ui

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Version is the A2UI specification version this package implements.
const Version = "v1.0"

// Message is one envelope in an agent-to-renderer stream. Exactly one of the
// payload fields is set.
type Message struct {
	Version          string            `json:"version"`
	CreateSurface    *CreateSurface    `json:"createSurface,omitempty"`
	UpdateComponents *UpdateComponents `json:"updateComponents,omitempty"`
}

// DecodeMessages reads a stream of newline-delimited A2UI envelopes. Blank
// lines are ignored.
func DecodeMessages(reader io.Reader) ([]Message, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 64*1024), 4*1024*1024)

	var messages []Message
	for line := 1; scanner.Scan(); line++ {
		raw := strings.TrimSpace(scanner.Text())
		if raw == "" {
			continue
		}

		var message Message
		if err := json.Unmarshal([]byte(raw), &message); err != nil {
			return nil, fmt.Errorf("line %d: %w", line, err)
		}
		messages = append(messages, message)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read messages: %w", err)
	}
	return messages, nil
}

func (message Message) payloadCount() int {
	count := 0
	if message.CreateSurface != nil {
		count++
	}
	if message.UpdateComponents != nil {
		count++
	}
	return count
}
