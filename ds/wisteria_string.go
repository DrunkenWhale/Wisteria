package ds

type WisteriaString struct {
	data string
}

func (ws *WisteriaString) ToBytes() []byte {
	return []byte(ws.data)
}

func NewWisteriaStringFromBytes(bytes []byte) string {
	return string(bytes)
}

// convert golang data type

func NewWisteriaString(str string) *WisteriaString {
	return &WisteriaString{
		data: str,
	}
}

func (ws *WisteriaString) Get() string {
	return ws.data
}
