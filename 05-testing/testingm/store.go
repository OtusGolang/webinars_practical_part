package testingm

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Store — примитивное key-value хранилище поверх текстового файла.
//
// Для наших целей важно одно: чтобы протестировать Store, нужен "тяжёлый"
// внешний ресурс — каталог на диске. Создавать и удалять его в каждом тесте
// дорого и шумно, поэтому он поднимается один раз на весь пакет — в TestMain.
// В реальной жизни на этом месте обычно БД, docker-контейнер (testcontainers),
// kafka, redis или прогретый HTTP-сервер.
type Store struct {
	mu   sync.RWMutex
	path string
	data map[string]string
}

const storeFileName = "store.tsv"

// Open открывает (или создаёт) хранилище в каталоге dir.
func Open(dir string) (*Store, error) {
	s := &Store{
		path: filepath.Join(dir, storeFileName),
		data: make(map[string]string),
	}

	f, err := os.Open(s.path)
	if errors.Is(err, os.ErrNotExist) {
		return s, nil // пустое хранилище — это нормально
	}
	if err != nil {
		return nil, fmt.Errorf("open store %q: %w", s.path, err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		key, value, ok := strings.Cut(sc.Text(), "\t")
		if !ok {
			continue
		}
		s.data[key] = value
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("read store %q: %w", s.path, err)
	}

	return s, nil
}

// Set сохраняет значение и сразу сбрасывает всё хранилище на диск.
func (s *Store) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	return s.flushLocked()
}

// Get возвращает значение и признак его наличия.
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]

	return value, ok
}

// Len возвращает количество ключей.
func (s *Store) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.data)
}

// Close имитирует освобождение ресурса (закрытие соединения с БД и т.п.).
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = nil

	return nil
}

func (s *Store) flushLocked() error {
	var b strings.Builder
	for key, value := range s.data {
		fmt.Fprintf(&b, "%s\t%s\n", key, value)
	}

	if err := os.WriteFile(s.path, []byte(b.String()), 0o600); err != nil {
		return fmt.Errorf("flush store %q: %w", s.path, err)
	}

	return nil
}
