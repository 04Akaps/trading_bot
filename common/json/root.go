package json

import (
	"github.com/bytedance/sonic"
)

type handler struct {
	marshal   func(val interface{}) ([]byte, error)
	unmarshal func(buf []byte, val interface{}) error
}

var JsonHandler handler

func init() {
	JsonHandler = handler{
		marshal:   sonic.Marshal,
		unmarshal: sonic.Unmarshal,
	}
}

func (h handler) Marshal(v interface{}) ([]byte, error) {
	bytes, err := h.marshal(v)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (h handler) Unmarshal(buffer []byte, v interface{}) error {
	err := h.unmarshal(buffer, v)

	if err != nil {
		return err
	}

	return nil
}

func (h handler) Handle(buf interface{}, v interface{}) error {

	bytes, err := h.marshal(buf)

	if err != nil {
		return err
	}

	err = h.unmarshal(bytes, v)

	if err != nil {
		return err
	}

	return nil
}
