package module

type Item map[string]interface{}

func (item Item) Valid() bool {
	return item != nil
}
