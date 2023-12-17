package handlers

// JSONString Tipe
type JSONString string

// MarshalJSON Convert
func (j JSONString) MarshalJSON() ([]byte, error) {
	return []byte(j), nil
}
