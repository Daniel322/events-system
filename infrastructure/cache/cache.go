package cache

// TODO: rework to redis in future

type Cache map[string]interface{}

var Instance Cache

func Init() {
	Instance = make(Cache)
}

func (c Cache) Set(key string, value interface{}) error {
	Instance[key] = value

	return nil
}

func (c Cache) Get(key string) (interface{}, bool) {
	val, ok := c[key]

	return val, ok
}

func (c Cache) Remove(key string) {
	delete(c, key)
}
