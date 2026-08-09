package main

import (
	"context"
	"testing"
	"time"

	termbox "github.com/nsf/termbox-go"
	"github.com/pierrre/assert"
)

func TestStartEventLoopForwardEvent(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pollCh := make(chan termbox.Event)
	poll := func() termbox.Event {
		return <-pollCh
	}

	evQueue, wait := startEventLoop(ctx, poll)

	ev := termbox.Event{Type: termbox.EventKey, Key: termbox.KeyEsc}
	go func() { pollCh <- ev }()

	select {
	case got := <-evQueue:
		assert.Equal(t, got, ev)
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for event")
	}

	cancel()
	go func() { pollCh <- termbox.Event{} }()

	done := make(chan struct{})
	go func() {
		wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for goroutine to exit")
	}
}

func TestStartEventLoopExitOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	pollCh := make(chan termbox.Event)
	poll := func() termbox.Event {
		return <-pollCh
	}

	_, wait := startEventLoop(ctx, poll)

	cancel()
	go func() { pollCh <- termbox.Event{Type: termbox.EventInterrupt} }()

	done := make(chan struct{})
	go func() {
		wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for goroutine to exit")
	}
}
