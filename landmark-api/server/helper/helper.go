package helper

import "google.golang.org/protobuf/types/known/wrapperspb"

func ConvertToStringPtr(value *wrapperspb.StringValue) *string {
	val := value.GetValue()
	return &val
}
