// MyType is an example type.
type MyType struct {
	// Id is a unique ID for this MyType.
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name is the name of this MyType.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}
