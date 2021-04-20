package internal

import (
	"testing"
	"unsafe"

	. "github.com/onsi/gomega"

	"github.com/lsytj0413/myinctl/pkg/helper"
)

func TestFrmFileHeader(t *testing.T) {
	t.Run("check struct size", func(t *testing.T) {
		g := NewWithT(t)

		g.Expect(unsafe.Sizeof(FrmFileHeader{})).To(BeEquivalentTo(4096 + 8)) // 8 bytes for align
	})

	t.Run("check data size", func(t *testing.T) {
		g := NewWithT(t)

		v := &FrmFileHeader{}
		g.Expect(helper.DataSize(v)).To(Equal(4096))
	})
}

func TestFrmKeyInfoSectionHeader(t *testing.T) {
	t.Run("check struct size", func(t *testing.T) {
		g := NewWithT(t)

		g.Expect(unsafe.Sizeof(FrmKeyInfoSectionHeader{})).To(BeEquivalentTo(6))
	})

	t.Run("check data size", func(t *testing.T) {
		g := NewWithT(t)

		v := &FrmKeyInfoSectionHeader{}
		g.Expect(helper.DataSize(v)).To(Equal(6))
	})
}

func TestKeyMetadata(t *testing.T) {
	t.Run("check struct size", func(t *testing.T) {
		g := NewWithT(t)

		g.Expect(unsafe.Sizeof(KeyMetadata{})).To(BeEquivalentTo(8))
	})

	t.Run("check data size", func(t *testing.T) {
		g := NewWithT(t)

		v := &KeyMetadata{}
		g.Expect(helper.DataSize(v)).To(Equal(8))
	})
}

func TestKeyParts(t *testing.T) {
	t.Run("check struct size", func(t *testing.T) {
		g := NewWithT(t)

		g.Expect(unsafe.Sizeof(KeyParts{})).To(BeEquivalentTo(10)) // 1 bytes align
	})

	t.Run("check data size", func(t *testing.T) {
		g := NewWithT(t)

		v := &KeyParts{}
		g.Expect(helper.DataSize(v)).To(Equal(9))
	})
}

func TestColumnMetadata(t *testing.T) {
	t.Run("check struct size", func(t *testing.T) {
		g := NewWithT(t)

		g.Expect(unsafe.Sizeof(ColumnMetadata{})).To(BeEquivalentTo(30))
	})

	t.Run("check data size", func(t *testing.T) {
		g := NewWithT(t)

		v := &ColumnMetadata{}
		g.Expect(helper.DataSize(v)).To(Equal(30))
	})
}

func TestColumnField(t *testing.T) {
	t.Run("check struct size", func(t *testing.T) {
		g := NewWithT(t)

		g.Expect(unsafe.Sizeof(ColumnField{})).To(BeEquivalentTo(20))
	})

	t.Run("check data size", func(t *testing.T) {
		g := NewWithT(t)

		v := &ColumnField{}
		g.Expect(helper.DataSize(v)).To(Equal(17))
	})
}
