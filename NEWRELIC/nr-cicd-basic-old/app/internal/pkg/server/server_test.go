package server

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	addr := "localhost:9999"
	s := New(addr)
	assert.NotNil(t, s)
	hc := http.DefaultClient
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		t.Log("Starting server...")
		err := s.ListenAndServe()
		require.NoError(t, err)
	}()

	go func() {
		defer func() {
			t.Log("Shutting down...")
			err := s.Shutdown()
			assert.NoError(t, err)
			wg.Done()
		}()

		t.Log("testing")
		timeOut := time.After(10 * time.Second)
		tick := time.Tick(100 * time.Millisecond)
		for {
			select {
			case <-timeOut:
				assert.Fail(t, "timed out")
				return
			case <-tick:
				url := fmt.Sprintf("http://%s/test", addr)
				req, err := http.NewRequest("GET", url, nil)
				require.NoError(t, err)
				resp, errDo := hc.Do(req)
				if errDo == nil {
					assert.Equal(t, 200, resp.StatusCode)
					return
				}
			}
		}
	}()
	wg.Wait()
}
