package main

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	result, err := SayHello("zeta", false)

	if err != nil {
		t.Fatal("error occured")
	}

	if result != "Hello zeta" {
		t.Fatal("result must be 'Hello zeta'")
	}

}

// func TestSayHello2(t *testing.T) {
// 	result, err := SayHello("zeta", false)
// 	assert.NoError(t, err, "error occured")
// 	assert.Equal(t, "Hello zeta", result, "result must be 'Hello zeta'")
// }

// func TestSayHello3(t *testing.T) {
// 	t.Run("no error check", func(t *testing.T) {
// 		result, err := SayHello("zeta", false)
// 		assert.NoError(t, err, "error occured")
// 		assert.Equal(t, "Hello zeta", result, "result must be 'Hello zeta'")
// 	})
// 	t.Run("when error check", func(t *testing.T) {
// 		result, err := SayHello("zeta", false)
// 		assert.Error(t, err, "error occured")
// 		assert.Empty(t, result, "still got result")
// 	})
// }
