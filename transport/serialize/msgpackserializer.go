package serialize

import (
	"errors"
	"reflect"

	"github.com/dtegapp/nexus/v3/wamp"
	"github.com/ugorji/go/codec"
)

var mh *codec.MsgpackHandle

func init() {
	mh = new(codec.MsgpackHandle)
	mh.WriteExt = true
	mh.MapType = reflect.TypeOf(map[string]interface{}(nil))
}

// MsgpackRegisterExtension registers a custom type for special serialization.
func MsgpackRegisterExtension(t reflect.Type, ext byte, encode func(reflect.Value) ([]byte, error), decode func(reflect.Value, []byte) error) error { //nolint:lll
	return mh.AddExt(t, ext, encode, decode)
}

// MessagePackSerializer is an implementation of Serializer that handles
// serializing and deserializing msgpack encoded payloads.
type MessagePackSerializer struct{}

// Serialize encodes a Message into a msgpack payload.
func (s *MessagePackSerializer) Serialize(msg wamp.Message) ([]byte, error) {
	var b []byte
	err := codec.NewEncoderBytes(&b, mh).Encode(msgToList(msg))
	return b, err
}

// Deserialize decodes a msgpack payload into a Message.
func (s *MessagePackSerializer) Deserialize(data []byte) (wamp.Message, error) {
	var v []interface{}
	err := codec.NewDecoderBytes(data, mh).Decode(&v)
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, errors.New("invalid message")
	}

	typ, ok := v[0].(int64)
	if !ok {
		return nil, errors.New("unsupported message format")
	}
	return listToMsg(wamp.MessageType(typ), v)
}

// SerializeDataItem encodes any object/structure into a msgpack payload.
func (s *MessagePackSerializer) SerializeDataItem(item interface{}) ([]byte, error) {
	var b []byte
	err := codec.NewEncoderBytes(&b, mh).Encode(item)
	return b, err
}

// DeserializeDataItem decodes a json payload into an object/structure.
func (s *MessagePackSerializer) DeserializeDataItem(data []byte, v interface{}) error {
	return codec.NewDecoderBytes(data, mh).Decode(&v)
}
