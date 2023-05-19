package proxy_database

import "time"

type Cache struct {
	database   Database
	cachedData map[string]CachedData
}

func NewCache(database Database) *Cache {
	return &Cache{database: database, cachedData: make(map[string]CachedData)}
}

func (c Cache) Set(input Input) {
	c.database.Set(input)
}

func (c Cache) Get(id string) string {
	value, ok := c.cachedData[id]

	foundButCacheExpired := ok && time.Now().After(value.ExpiredAt)

	if !ok || foundButCacheExpired {
		content := c.database.Get(id)

		c.cachedData[id] = CachedData{
			ExpiredAt: time.Now().Add(3 * time.Second),
			Content:   content,
		}

		return content
	}

	return value.Content
}

// data structure

type CachedData struct {
	ExpiredAt time.Time
	Content   string
}
