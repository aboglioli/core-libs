package errors

type Metadata struct {
	metadata map[string]interface{}
}

func NewMetadata() Metadata {
	return Metadata{
		metadata: make(map[string]interface{}),
	}
}

func WithMetadata(key string, value interface{}) Metadata {
	return NewMetadata().And(key, value)
}

func (m Metadata) And(key string, value interface{}) Metadata {
	m.metadata[key] = value
	return m
}

func (m Metadata) Merge(other Metadata) Metadata {
	for k, v := range other.metadata {
		m.metadata[k] = v
	}

	return m
}

func (m Metadata) Values() map[string]interface{} {
	return m.metadata
}
