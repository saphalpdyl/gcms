package serializers

type ISerializer[T any] interface {
	Serialize(string) (T, error)
	Deserialize(T) string
}
