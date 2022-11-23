package gochan

import (
	"encoding/json"
	"net/http"
	pathpkg "path"
	"sync"
	"time"
)

type Session struct {
	APIURL         string        // 4chan API URL. Default: "a.4cdn.org"
	ImageURL       string        // 4chan image URL. Default: "i.4cdn.org"
	StaticURL      string        // 4chan static URL. Default: "s.4cdn.org"
	SSL            bool          // Specifies whether SSL is used (default: true)
	UpdateCooldown time.Duration // Thread update cooldown (default: 15s, minimum: 10s)

	cooldown      <-chan time.Time
	cooldownMutex sync.Mutex
}

// NewSession creates a new session.
func NewSession() *Session {
	sess := &Session{
		APIURL:         API_URL,
		ImageURL:       IMAGE_URL,
		StaticURL:      STATIC_URL,
		SSL:            true,
		UpdateCooldown: 15 * time.Second,
	}

	return sess
}

func (sess *Session) urlPrefix() string {
	if sess.SSL {
		return "https://"
	} else {
		return "http://"
	}
}

func (sess *Session) getReq(base, path string, beforeRequest func(*http.Request) error) (*http.Response, error) {
	url := sess.urlPrefix() + pathpkg.Join(base, path)
	sess.cooldownMutex.Lock()

	// https://github.com/4chan/4chan-API#api-rules
	if sess.cooldown != nil {
		<-sess.cooldown
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if beforeRequest != nil {
		if err := beforeRequest(req); err != nil {
			return nil, err
		}
	}

	resp, err := http.DefaultClient.Do(req)
	sess.cooldown = time.After(1 * time.Second)
	sess.cooldownMutex.Unlock()
	return resp, err
}

func (sess *Session) getUnmarshal(base, path string, dest interface{}, modify func(*http.Request) error) error {
	resp, err := sess.getReq(base, path, modify)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(dest)
}
