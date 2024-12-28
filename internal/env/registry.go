package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type Registry struct {
	env map[string]string
	mu  sync.RWMutex
}

var (
	instance *Registry
	once     sync.Once
)

func GetRegistry() *Registry {
	once.Do(func() {
		instance = &Registry{
			env: make(map[string]string),
		}
		instance.loadEnv()
	})
	return instance
}

func (r *Registry) loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	for _, e := range os.Environ() {
		pair := splitEnv(e)

		r.env[pair[0]] = pair[1]
	}
}

func splitEnv(env string) [2]string {
	i := 0
	for ; i < len(env); i++ {
		if env[i] == '=' {
			break
		}
	}
	return [2]string{env[:i], env[i+1:]}
}

func (r *Registry) Get(key string) string {

	if value, exists := r.env[key]; exists {
		return value
	}
	log.Printf("Warning: key %s not found in environment", key)
	return ""
}
