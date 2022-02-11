package govmq

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAuthOnRegisterM3ModifierBuilder(t *testing.T) {
	b := &AuthOnRegisterM3ModifierBuilder{}
	b.WithMaxMessageSize(65535).
		WithMaxOfflineMessages(10000).
		WithMountpoint("test")
	r := NewOKResponse(b.Build())
	fmt.Println(r)
	mm := b.Build().MaxMessageSize
	fmt.Println("&mm")
	fmt.Println(*mm)
}

func TestAuthOnRegisterM5ModifierBuilder(t *testing.T) {
	b := AuthOnRegisterM5ResponseBuilder{}
	b.WithMaxMessageSize(65535).
		WithMaxOfflineMessages(10000).
		WithMountpoint("test")
	r := NewOKResponse(b.Build())
	s, _ := json.Marshal(r)
	fmt.Println(r)
	fmt.Println(string(s))
}

// Add a flag in builder that can identify the builder
// Put if in
