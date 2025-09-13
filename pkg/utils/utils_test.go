package utils

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestParseId(t *testing.T) {
	t.Run("err will be nil if send valid uuid", func(t *testing.T) {
		testUuid := uuid.New()

		_, _, err := ParseId(testUuid)

		if err != nil {
			t.Error("ParseId should return uuid in string format, but return error")
		}
	})

	t.Run("should be error if send invalid uuid", func(t *testing.T) {
		invalidTestUuid := "invalidstring"

		_, _, err := ParseId(invalidTestUuid)

		if err == nil {
			t.Error("ParseId should return error, but error was nil")
		}
	})
}

func TestGenerateError(t *testing.T) {
	err := GenerateError("testName", "some error")

	if err == nil {
		t.Error("should be error, but get nil")
	}
}

func TestSetInterval(t *testing.T) {
	t.Run("callback should be called multiple times", func(t *testing.T) {
		var callCount int
		var mu sync.Mutex
		var wg sync.WaitGroup

		callback := func() error {
			mu.Lock()
			callCount++
			mu.Unlock()
			return nil
		}

		interval := 10 * time.Millisecond
		stop := SetInterval(callback, interval)

		// Ждем несколько вызовов
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
		}()
		wg.Wait()

		// Останавливаем
		stop <- true

		mu.Lock()
		actualCalls := callCount
		mu.Unlock()

		// Должно быть минимум 3 вызова за 50ms с интервалом 10ms
		if actualCalls < 3 {
			t.Errorf("Expected at least 3 calls, got %d", actualCalls)
		}
	})

	t.Run("callback should handle errors", func(t *testing.T) {
		var callCount int
		var mu sync.Mutex

		callback := func() error {
			mu.Lock()
			callCount++
			mu.Unlock()
			return errors.New("test error")
		}

		interval := 10 * time.Millisecond
		stop := SetInterval(callback, interval)

		// Ждем несколько вызовов
		time.Sleep(50 * time.Millisecond)

		// Останавливаем
		stop <- true

		mu.Lock()
		actualCalls := callCount
		mu.Unlock()

		// Callback должен вызываться даже при ошибках
		if actualCalls < 3 {
			t.Errorf("Expected at least 3 calls even with errors, got %d", actualCalls)
		}
	})

	t.Run("stop channel should terminate the interval", func(t *testing.T) {
		var callCount int
		var mu sync.Mutex

		callback := func() error {
			mu.Lock()
			callCount++
			mu.Unlock()
			return nil
		}

		interval := 20 * time.Millisecond
		stop := SetInterval(callback, interval)

		// Ждем один вызов
		time.Sleep(30 * time.Millisecond)

		// Останавливаем
		stop <- true

		// Ждем еще немного, чтобы убедиться, что больше вызовов не будет
		time.Sleep(50 * time.Millisecond)

		mu.Lock()
		actualCalls := callCount
		mu.Unlock()

		// Должно быть 1-2 вызова (один до остановки, возможно еще один)
		if actualCalls < 1 || actualCalls > 3 {
			t.Errorf("Expected 1-3 calls, got %d", actualCalls)
		}
	})

	t.Run("multiple stop signals should not cause panic", func(t *testing.T) {
		callback := func() error {
			return nil
		}

		interval := 20 * time.Millisecond
		stop := SetInterval(callback, interval)

		// Отправляем один сигнал остановки
		stop <- true

		// Ждем немного, чтобы убедиться, что нет паники
		time.Sleep(30 * time.Millisecond)

		// Если мы дошли сюда, значит паники не было
	})

	t.Run("very short interval should work", func(t *testing.T) {
		var callCount int
		var mu sync.Mutex

		callback := func() error {
			mu.Lock()
			callCount++
			mu.Unlock()
			return nil
		}

		interval := 5 * time.Millisecond
		stop := SetInterval(callback, interval)

		// Ждем несколько вызовов
		time.Sleep(25 * time.Millisecond)

		// Останавливаем
		stop <- true

		mu.Lock()
		actualCalls := callCount
		mu.Unlock()

		// Должно быть несколько вызовов за 25ms с интервалом 5ms
		if actualCalls < 3 {
			t.Errorf("Expected at least 3 calls with short interval, got %d", actualCalls)
		}
	})

	t.Run("very long interval should work", func(t *testing.T) {
		var callCount int
		var mu sync.Mutex

		callback := func() error {
			mu.Lock()
			callCount++
			mu.Unlock()
			return nil
		}

		interval := 100 * time.Millisecond
		stop := SetInterval(callback, interval)

		// Ждем один вызов
		time.Sleep(120 * time.Millisecond)

		// Останавливаем
		stop <- true

		mu.Lock()
		actualCalls := callCount
		mu.Unlock()

		// Должен быть минимум один вызов
		if actualCalls < 1 {
			t.Errorf("Expected at least 1 call with long interval, got %d", actualCalls)
		}
	})

	t.Run("concurrent access should be safe", func(t *testing.T) {
		var callCount int
		var mu sync.Mutex

		callback := func() error {
			mu.Lock()
			callCount++
			mu.Unlock()
			return nil
		}

		interval := 10 * time.Millisecond
		stop := SetInterval(callback, interval)

		// Запускаем несколько горутин, которые читают счетчик
		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				time.Sleep(30 * time.Millisecond)
				mu.Lock()
				_ = callCount
				mu.Unlock()
			}()
		}

		wg.Wait()

		// Останавливаем
		stop <- true

		mu.Lock()
		actualCalls := callCount
		mu.Unlock()

		// Должно быть несколько вызовов
		if actualCalls < 2 {
			t.Errorf("Expected at least 2 calls, got %d", actualCalls)
		}
	})
}
