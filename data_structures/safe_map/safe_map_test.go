package safemap_test

import (
	"fmt"
	"sync"
	"testing"

	safemap "github.com/juanfer2/go-concurrency/data_structures/safe_map"
	"github.com/stretchr/testify/assert"
)

func TestSafeMap(t *testing.T) {
	var wg sync.WaitGroup
	m := safemap.New[string, int]()
	assert := assert.New(t)

	for i := 0; i < 1000; i++ {
		keyName := fmt.Sprintf("key_%d", i)
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			//t.Error(m)
			m.Add(keyName, i)
			//t.Error(m.Get(keyName))
			value, err := m.Get(keyName)

			if err != nil {
				t.Error(err)
			}

			if value != i {
				t.Errorf("%s should be %d", keyName, i)
			}

			assert.Equal(value, i, "The value should be the same.")
		}(i)
	}

	wg.Wait()
}

func TestDeleteSafeMap(t *testing.T) {
	var wg sync.WaitGroup
	m := safemap.New[string, int]()
	assert := assert.New(t)
	wg.Add(1)

	go func() {
		defer wg.Done()
		keyName := "key"
		m.Add(keyName, 2)
		m.Add("key_2", 2)
		//t.Error(m.Get(keyName))
		err := m.Delete(keyName)
		if err != nil {
			t.Error(err)
		}

		assert.Equalf(m.HasKey(keyName), false, "The key %s doesn't exist", "keyName")
	}()

	wg.Wait()
}
