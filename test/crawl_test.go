package test

import (
	"io"
	"net/http"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipleEndpoints(t *testing.T) {
	assert := assert.New(t)
	endpoints := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/users",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/comments/46",
		"https://jsonplaceholder.typicode.com/photos/22",
	}

	result := make(chan string, 5)
	for _, v := range endpoints {
		go func(v string) {
			request, _ := http.NewRequest(http.MethodGet, v, nil)
			client := http.Client{}
			response, _ := client.Do(request)
			defer response.Body.Close() // @dev prevent FD leak

			raw, _ := io.ReadAll(response.Body)
			data := string(raw)
			result <- data
		}(v)
	}

	var final []string
	for i := 0; i < len(endpoints); i++ {
		final = append(final, <-result)
	}

	t.Logf("final: %s", final)
	assert.Equal(len(endpoints), 5)
}
