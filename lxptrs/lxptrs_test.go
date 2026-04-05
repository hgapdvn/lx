package lxptrs_test

import (
	"testing"

	"github.com/hgapdvn/lx/lxptrs"
)

func TestRef(t *testing.T) {

	t.Run("String", func(t *testing.T) {
		v := "ref check"
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Bool", func(t *testing.T) {
		v := true
		p := lxptrs.Ref(true)
		if *p != true {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Int", func(t *testing.T) {
		v := 30
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Int64", func(t *testing.T) {
		v := int64(98)
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Int32", func(t *testing.T) {
		v := int32(12)
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Int16", func(t *testing.T) {
		v := int16(-55)
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Int8", func(t *testing.T) {
		v := int8(12)
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Unit", func(t *testing.T) {
		v := uint(12)
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Unit64", func(t *testing.T) {
		v := uint64(1999)
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Unit32", func(t *testing.T) {
		v := uint32(19921)
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Unit16", func(t *testing.T) {
		v := uint16(1000)
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Unit8", func(t *testing.T) {
		v := uint8(1)
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})
}

func TestDeref(t *testing.T) {

	t.Run("String", func(t *testing.T) {
		v := "my string"
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Bool", func(t *testing.T) {
		v := true
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Int", func(t *testing.T) {
		v := 10
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Int64", func(t *testing.T) {
		v := int64(10)
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Int32", func(t *testing.T) {
		v := int32(10)
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Int16", func(t *testing.T) {
		v := int16(10)
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Int8", func(t *testing.T) {
		v := int8(10)
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("UInt", func(t *testing.T) {
		v := uint(10)
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("UInt64", func(t *testing.T) {
		v := uint64(10)
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("UInt32", func(t *testing.T) {
		v := uint32(10)
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("UInt16", func(t *testing.T) {
		v := uint16(10)
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("UInt8", func(t *testing.T) {
		v := uint8(10)
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})
}
