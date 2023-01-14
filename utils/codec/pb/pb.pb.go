package pb

type KV struct {
	Key       []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Version   uint64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	ExpiresAt uint64 `protobuf:"varint,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (m *KV) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KV) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KV) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *KV) GetExpiresAt() uint64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	//proto.RegisterType((*KV)(nil), "pb.KV")
}
