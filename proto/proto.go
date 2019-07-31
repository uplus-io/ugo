/*
 * Copyright (c) 2019 uplus.io
 */

package proto

import (
	ggproto "github.com/golang/protobuf/proto"
)

const PROTO_VERSION = 1

type ProtoMessage ggproto.Message

func Marshal(message ProtoMessage) ([]byte, error) {
	return ggproto.Marshal(message)
}

func Unmarshal(buf []byte, message ProtoMessage) (error) {
	return ggproto.Unmarshal(buf, message)
}
