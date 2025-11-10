package golazy

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestNewWithLoader(t *testing.T) {
	loader := func(ctx any) (string, error) {
		return "value", nil
	}
	wl := newWithLoader(loader, false, 0)

	if wl.loader == nil {
		t.Error("loader should not be nil")
	}
	if wl.withTTL != false {
		t.Error("withTTL should be false")
	}
	if len(wl.values) != 0 {
		t.Error("values map should be empty")
	}
}

func TestNewWithLoaderPreloaded(t *testing.T) {
	loader := func(ctx any) (string, error) {
		return "new_value", nil
	}
	ctx := "test_ctx"
	wl := newWithLoaderPreloaded(loader, "preloaded", ctx, false, 0)

	if v, ok := wl.values[ctx]; !ok || v != "preloaded" {
		t.Error("preloaded value not set correctly")
	}
}

func TestValue_CachedValue(t *testing.T) {
	callCount := 0
	loader := func(ctx any) (string, error) {
		callCount++
		return "value", nil
	}
	wl := newWithLoader(loader, false, 0)

	v1, err := wl.Value("ctx1")
	if err != nil || v1 != "value" || callCount != 1 {
		t.Error("first call should invoke loader")
	}

	v2, err := wl.Value("ctx1")
	if err != nil || v2 != "value" || callCount != 1 {
		t.Error("second call should use cache")
	}
}

func TestValue_LoaderError(t *testing.T) {
	loader := func(ctx any) (string, error) {
		return "", errors.New("loader error")
	}
	wl := newWithLoader(loader, false, 0)

	_, err := wl.Value("ctx1")
	if err == nil || err.Error() != "loader error" {
		t.Error("error should be propagated")
	}
}

func TestValue_TTLExpiration(t *testing.T) {
	callCount := 0
	loader := func(ctx any) (string, error) {
		callCount++
		return "value", nil
	}
	wl := newWithLoader(loader, true, 100*time.Millisecond)

	wl.Value("ctx1")
	if callCount != 1 {
		t.Error("loader should be called once")
	}

	wl.Value("ctx1")
	if callCount != 1 {
		t.Error("cached value should be used")
	}

	time.Sleep(150 * time.Millisecond)
	wl.Value("ctx1")
	if callCount != 2 {
		t.Error("expired cache should trigger reload")
	}
}

func TestValue_NoLoaderNil(t *testing.T) {
	wl := newWithLoader[string](nil, false, 0)
	v, err := wl.Value("ctx1")
	if err != nil || v != "" {
		t.Error("should return nil value without error when loader is nil")
	}
}

func TestClear(t *testing.T) {
	loader := func(ctx any) (string, error) {
		return "value", nil
	}
	wl := newWithLoader(loader, false, 0)
	wl.Value("ctx1")

	if _, ok := wl.values["ctx1"]; !ok {
		t.Error("value should be cached")
	}

	wl.Clear("ctx1")
	if _, ok := wl.values["ctx1"]; ok {
		t.Error("value should be cleared")
	}
}

func TestClearAll(t *testing.T) {
	loader := func(ctx any) (string, error) {
		return "value", nil
	}
	wl := newWithLoader(loader, false, 0)
	wl.Value("ctx1")
	wl.Value("ctx2")

	wl.ClearAll()
	if len(wl.values) != 0 {
		t.Error("all values should be cleared")
	}
}

func TestConcurrency(t *testing.T) {
	callCount := 0
	loader := func(ctx any) (string, error) {
		callCount++
		time.Sleep(10 * time.Millisecond)
		return "value", nil
	}
	wl := newWithLoader(loader, false, 0)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			wl.Value("ctx1")
		}()
	}
	wg.Wait()

	if callCount != 1 {
		t.Errorf("loader should be called once, but was called %d times", callCount)
	}
}
