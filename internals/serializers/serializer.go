package serializers

type ISerializer[T any] interface {
	Deserialize([]byte, *T) error
	Serialize(T) []byte
}
