// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: orderpb/order.proto

package orderpb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ItemWithQuantity struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ItemWithQuantity) Reset() {
	*x = ItemWithQuantity{}
	mi := &file_orderpb_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemWithQuantity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemWithQuantity) ProtoMessage() {}

func (x *ItemWithQuantity) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemWithQuantity.ProtoReflect.Descriptor instead.
func (*ItemWithQuantity) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{0}
}

func (x *ItemWithQuantity) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ItemWithQuantity) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Quantity      int32                  `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	PriceID       string                 `protobuf:"bytes,4,opt,name=PriceID,proto3" json:"PriceID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_orderpb_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Item) GetPriceID() string {
	if x != nil {
		return x.PriceID
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CustomerID    string                 `protobuf:"bytes,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Items         []*Item                `protobuf:"bytes,4,rep,name=Items,proto3" json:"Items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_orderpb_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Order) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerID    string                 `protobuf:"bytes,1,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Items         []*ItemWithQuantity    `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_orderpb_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderRequest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *CreateOrderRequest) GetItems() []*ItemWithQuantity {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetOrdersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	CustomerID    string                 `protobuf:"bytes,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrdersRequest) Reset() {
	*x = GetOrdersRequest{}
	mi := &file_orderpb_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersRequest) ProtoMessage() {}

func (x *GetOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetOrdersRequest) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrdersRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *GetOrdersRequest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

var File_orderpb_order_proto protoreflect.FileDescriptor

var file_orderpb_order_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x10, 0x49,
	0x74, 0x65, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x60, 0x0a, 0x04, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x44, 0x22, 0x74, 0x0a,
	0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23,
	0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x65, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x4c, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x32, 0xc1, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x36, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3d, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x68, 0x79, 0x78, 0x7a,
	0x6d, 0x6b, 0x68, 0x2f, 0x67, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_orderpb_order_proto_rawDescOnce sync.Once
	file_orderpb_order_proto_rawDescData []byte
)

func file_orderpb_order_proto_rawDescGZIP() []byte {
	file_orderpb_order_proto_rawDescOnce.Do(func() {
		file_orderpb_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_orderpb_order_proto_rawDesc), len(file_orderpb_order_proto_rawDesc)))
	})
	return file_orderpb_order_proto_rawDescData
}

var file_orderpb_order_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_orderpb_order_proto_goTypes = []any{
	(*ItemWithQuantity)(nil),   // 0: orderpb.ItemWithQuantity
	(*Item)(nil),               // 1: orderpb.Item
	(*Order)(nil),              // 2: orderpb.Order
	(*CreateOrderRequest)(nil), // 3: orderpb.CreateOrderRequest
	(*GetOrdersRequest)(nil),   // 4: orderpb.GetOrdersRequest
	(*empty.Empty)(nil),        // 5: google.protobuf.Empty
}
var file_orderpb_order_proto_depIdxs = []int32{
	1, // 0: orderpb.Order.Items:type_name -> orderpb.Item
	0, // 1: orderpb.CreateOrderRequest.Items:type_name -> orderpb.ItemWithQuantity
	3, // 2: orderpb.OrderService.CreateOrder:input_type -> orderpb.CreateOrderRequest
	4, // 3: orderpb.OrderService.GetOrders:input_type -> orderpb.GetOrdersRequest
	2, // 4: orderpb.OrderService.UpdateOrder:input_type -> orderpb.Order
	5, // 5: orderpb.OrderService.CreateOrder:output_type -> google.protobuf.Empty
	2, // 6: orderpb.OrderService.GetOrders:output_type -> orderpb.Order
	5, // 7: orderpb.OrderService.UpdateOrder:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_orderpb_order_proto_init() }
func file_orderpb_order_proto_init() {
	if File_orderpb_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_orderpb_order_proto_rawDesc), len(file_orderpb_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderpb_order_proto_goTypes,
		DependencyIndexes: file_orderpb_order_proto_depIdxs,
		MessageInfos:      file_orderpb_order_proto_msgTypes,
	}.Build()
	File_orderpb_order_proto = out.File
	file_orderpb_order_proto_goTypes = nil
	file_orderpb_order_proto_depIdxs = nil
}
