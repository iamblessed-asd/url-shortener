package storage

type MemoryStorage struct {
	data map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]string),
	}
}

func (m *MemoryStorage) Save(code, url string) {
	m.data[code] = url
}

func (m *MemoryStorage) Load(code string) (string, bool) {
	url, ok := m.data[code]
	return url, ok
}
