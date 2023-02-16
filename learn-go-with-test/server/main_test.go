package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("返回佩琪的分数", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "players/Pepper", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("获得 %q，期望 %q", got, want)
		}
	})
}
