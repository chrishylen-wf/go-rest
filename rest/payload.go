package rest

import "fmt"

// Payload is the unmarshalled request body.
type Payload map[string]interface{}

// Get returns the value with the given key as an interface{}. If the key doesn't
// exist, nil is returned with an error.
func (p Payload) Get(key string) (interface{}, error) {
	if value, ok := p[key]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("No value with key '%s'", key)
}

// GetInt returns the value with the given key as an int. If the key doesn't
// exist or is not an int, the zero value is returned with an error.
func (p Payload) GetInt(key string) (int, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(int); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not an int", key)
}

// GetInt8 returns the value with the given key as an int8. If the key doesn't
// exist or is not an int8, the zero value is returned with an error.
func (p Payload) GetInt8(key string) (int8, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(int8); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not an int8", key)
}

// GetInt16 returns the value with the given key as an int16. If the key doesn't
// exist or is not an int16, the zero value is returned with an error.
func (p Payload) GetInt16(key string) (int16, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(int16); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not an int16", key)
}

// GetInt32 returns the value with the given key as an int32. If the key doesn't
// exist or is not an int32, the zero value is returned with an error.
func (p Payload) GetInt32(key string) (int32, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(int32); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not an int32", key)
}

// GetInt64 returns the value with the given key as an int64. If the key doesn't
// exist or is not an int64, the zero value is returned with an error.
func (p Payload) GetInt64(key string) (int64, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(int64); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not an int64", key)
}

// GetUint returns the value with the given key as a uint. If the key doesn't
// exist or is not a uint, the zero value is returned with an error.
func (p Payload) GetUint(key string) (uint, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(uint); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not a uint", key)
}

// GetUint8 returns the value with the given key as a uint8. If the key doesn't
// exist or is not a uint8, the zero value is returned with an error.
func (p Payload) GetUint8(key string) (uint8, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(uint8); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not a uint8", key)
}

// GetUint16 returns the value with the given key as a uint16. If the key doesn't
// exist or is not a uint16, the zero value is returned with an error.
func (p Payload) GetUint16(key string) (uint16, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(uint16); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not a uint16", key)
}

// GetUint32 returns the value with the given key as a uint32. If the key doesn't
// exist or is not a uint32, the zero value is returned with an error.
func (p Payload) GetUint32(key string) (uint32, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(uint32); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not a uint32", key)
}

// GetUint64 returns the value with the given key as a uint64. If the key doesn't
// exist or is not a uint64, the zero value is returned with an error.
func (p Payload) GetUint64(key string) (uint64, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(uint64); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not a uint64", key)
}

// GetFloat32 returns the value with the given key as a Float32. If the key doesn't
// exist or is not a Float32, the zero value is returned with an error.
func (p Payload) GetFloat32(key string) (float32, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(float32); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not a float32", key)
}

// GetFloat64 returns the value with the given key as a Float64. If the key doesn't
// exist or is not a Float32, the zero value is returned with an error.
func (p Payload) GetFloat64(key string) (float64, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(float64); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not a float64", key)
}

// GetByte returns the value with the given key as a byte. If the key doesn't
// exist or is not a byte, the zero value is returned with an error.
func (p Payload) GetByte(key string) (byte, error) {
	value, err := p.Get(key)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(byte); ok {
		return value, nil
	}
	return 0, fmt.Errorf("Value with key '%s' not a byte", key)
}

// GetString returns the value with the given key as a string. If the key doesn't
// exist or is not a string, the zero value is returned with an error.
func (p Payload) GetString(key string) (string, error) {
	value, err := p.Get(key)
	if err != nil {
		return "", err
	}
	if value, ok := value.(string); ok {
		return value, nil
	}
	return "", fmt.Errorf("Value with key '%s' not a string", key)
}

// GetBool returns the value with the given key as a bool. If the key doesn't
// exist or is not a bool, false is returned with an error.
func (p Payload) GetBool(key string) (bool, error) {
	value, err := p.Get(key)
	if err != nil {
		return false, err
	}
	if value, ok := value.(bool); ok {
		return value, nil
	}
	return false, fmt.Errorf("Value with key '%s' not a bool", key)
}

// GetArray returns the value with the given key as an []interface{}. If the value
// doesn't exist or is not an []interface{}, nil is returned with an error.
func (p Payload) GetArray(key string) ([]interface{}, error) {
	value, err := p.Get(key)
	if err != nil {
		return nil, err
	}
	if value, ok := value.([]interface{}); ok {
		return value, nil
	}
	return nil, fmt.Errorf("Value with key '%s' not an array", key)
}

// GetMap returns the value with the given key as a map[string]interface{}. If the
// key doesn't exist or is not a map[string]interface{}, nil is returned with an
// error.
func (p Payload) GetMap(key string) (map[string]interface{}, error) {
	value, err := p.Get(key)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(map[string]interface{}); ok {
		return value, nil
	}
	return nil, fmt.Errorf("Value with key '%s' not a map", key)
}
