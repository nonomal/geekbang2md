package image

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/DuC-cnZj/geekbang2md/waiter"
	"golang.org/x/time/rate"
)

type Manager struct {
	sync.RWMutex
	images  map[string]string
	baseDir string
	waiter  *waiter.Waiter
}

func NewManager(baseDir string) *Manager {
	os.MkdirAll(filepath.Join(baseDir, "mp3"), 0755)

	return &Manager{
		RWMutex: sync.RWMutex{},
		images:  map[string]string{},
		baseDir: baseDir,
		waiter:  waiter.NewWaiter(rate.Every(100*time.Millisecond), 30),
	}
}

func (m *Manager) Download(u string) (string, error) {
	if path := m.Get(u); path != "" {
		return path, nil
	}
	parse, _ := url.Parse(u)
	split := strings.Split(parse.Path, "/")
	name := split[len(split)-1]
	var p string
	if strings.HasSuffix(name, ".mp3") {
		p = filepath.Join(m.baseDir, "mp3", name)
	} else {
		p = filepath.Join(m.baseDir, name)
	}
	stat, err := os.Stat(p)
	if err == nil && stat.Mode().IsRegular() {
		m.Add(u, p)
		return p, nil
	}
	m.waiter.Wait(context.TODO())
	c := &http.Client{}
	res, err := c.Get(u)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	all, _ := io.ReadAll(res.Body)
	if err := os.WriteFile(p, all, 0644); err != nil {
		return "", err
	}
	m.Add(u, p)
	return p, nil
}

func (m *Manager) Has(url string) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.images[url]
	return ok
}

func (m *Manager) Get(url string) string {
	m.RLock()
	defer m.RUnlock()
	path, ok := m.images[url]
	if !ok {
		return ""
	}
	return path
}
func (m *Manager) Add(url, path string) {
	m.Lock()
	defer m.Unlock()
	m.images[url] = path
}