// Package internal holds the frm file struct definitions.
// NOTE: all the number is stored as little-endian.
package internal

// FrmFileHeader defines the frm file header struct.
// NOTE: the header always been 64 bytes and padding to 4096 bytes length (4KB).
type FrmFileHeader struct {
	// MagicNumber is the fixed number defined as .frm format file.
	// Value: 0x01fe
	MagicNumber uint16 // 0x00

	// FrmVer is the frm version.
	// RM_VER (which is in include/mysql_version.h) +3 +test(create_info->varchar)
	FrmVer uint8 // 0x02

	// LegacyDBType is the database type (engine).
	LegacyDBType uint8 // 0x03

	// Unknown0 defines the unknown field.
	// UNUSED.
	_ uint8 // 0x04

	// Unknown1 defines the unknown field.
	_ uint8 // 0x05

	// IOSize is the size of io, the size of file header and other section.
	// Value: Always 4096
	IOSize uint16 // 0x06

	// Unknown2 defines the unknown field.
	_ uint16 // 0x08

	// Unknown3 defines the unknown field.
	_ uint32 // 0x0a

	// TmpKeyLength is the key length, if the value is 0xFFFF then the value is stored at 0x2f
	TmpKeyLength uint16 // 0x0e

	// RecLength is the default value length
	RecLength uint16 // 0x10

	// MaxRows is the table MAX_ROWS option.
	MaxRows uint32 // 0x12

	// MinRows is the table MIN_ROWS option.
	MinRows uint32 // 0x16

	// Unknown4 defines the unknown field.
	// Always been 0x0200, meas "use long pack-fields"
	_ uint16 // 0x1a

	// KeyInfoLength is the keyinfo section length.
	KeyInfoLength uint16 // 0x1c

	// CreateOptions is the db_create_options.
	// EX: HA_LONG_BLOB_PTR
	CreateOptions uint16 // 0x1e

	// Unknown5 defines the unknown field.
	// UNUSED.
	_ uint8 // 0x20

	// Version5FrmFile is the mark for 5.0 frm file.
	// Value: always 1 after mysql 5
	Version5FrmFile uint8 // 0x21

	// AvgRowLength is the table AVG_ROW_LENGTH option.
	AvgRowLength uint32 // 0x22

	// DefaultTableCharset is the table DEFAULT CHARACTER SET option.
	DefaultTableCharset uint8 // 0x26

	// Unknown7 defines the unknown field.
	_ uint8 // 0x27

	// RowType is the row types.
	RowType uint8 // 0x28

	// TableCharsetHighByte is the high byte of charset
	TableCharsetHighByte uint8 // 0x29

	// StatsSamplePages is the pages.
	StatsSamplePages uint16 // 0x2a

	// StatsAutoRecalc is the auto recalc flag.
	StatsAutoRecalc uint8 // 0x2c

	// Unknown8 defines the unknown field.
	_ uint16 // 0x2d

	// KeyLength is the key length.
	KeyLength uint32 // 0x2f

	// MysqlVersionID is the mysql version.
	MysqlVersionID uint32 // 0x33

	// ExtraSize of the extra info.
	// EX: CONNECTION=<>
	//     ENGINE=<>
	//     PARTITION BY clause + partitioning flags
	//     WITH PARSER names (MySQL 5.1+)
	//     Table COMMENT
	ExtraSize uint32 // 0x37

	// ExtraRecBufLength is the extra rec buf length.
	ExtraRecBufLength uint16 // 0x3b

	// DefaultPartDBType is the enum legacy_db_type, if table is partitioned.
	DefaultPartDBType uint8 // 0x3d

	// KeyBlockSize is the table KEY_BLOCK_SIZE option.
	KeyBlockSize uint16 // 0x3e

	// Unknown9 defines the unknown field.
	// NOTE: padding
	// _ [4096 - 64]byte // 0x40
}

// FrmKeyInfoSectionHeader defines the frm file key information section header.
// NOTE: it always start at 0x1000 and padding to 6 bytes.
type FrmKeyInfoSectionHeader struct {
	// Data is the header.
	// If the Data[0] == 0x80:
	//   Then:
	//     Key Count is Data[1] << 7 | Data[0] &0x7F
	//     Key Parts Count is Data[2] << 8 | Data[3]
	//   Else:
	//     Key Count is Data[0]
	//     Key Parts Count is Data[1]
	// NOTE:
	//   Key Count defines the index count in this table (include PRIMARY KEY)
	//   Key Parts Count defines the index column count in all index
	Data [4]byte

	// NOTE: padding
	_ [2]byte
}

// Count return the keyCount and keyPartsCount
func (k *FrmKeyInfoSectionHeader) Count() (keyCount int, keyPartsCount int) {
	if k.Data[0] == 0x80 {
		return int(k.Data[1])<<7 | int(k.Data[0])&0x7F, int(k.Data[3])<<7 | int(k.Data[2])
	}

	return int(k.Data[0]), int(k.Data[1])
}

// KeyMetadata defines the metadata of index key.
type KeyMetadata struct {
	// Flags is the key flags, such as HA_USES_COMMENT
	Flags uint16

	// KeyLength is the length of index
	KeyLength uint16

	// UserDefinedKeyParts is the column count cover by the index
	UserDefinedKeyParts uint8

	// Algorithm is the index algorithm, such as HA_KEY_ALG_BTREE
	Algorithm uint8

	// BlockSize is the block size of index, table KEY_BLOCK_SIZE option
	BlockSize uint16
}

// KeyParts defines the struct of index key user defined key parts.
type KeyParts struct {
	// FieldNumber is the field index of current parts.
	// NOTE: this field should be mark with 0x3FFF
	FieldNumber uint16

	// Offset in the mysql internal data struct.
	// NOTE: internal-usage
	Offset uint16

	// KeyPartFlag is the flag of key part.
	KeyPartFlag int8

	// KeyType is the key_type, SEE ha_base_keytype
	KeyType uint16

	// Length is the column index length
	Length uint16
}

// ColumnMetadata defines the metadata of column.
type ColumnMetadata struct {
	// MagicNumber is the fixed number defined as .frm column metadata section.
	// Value: 0x03
	MagicNumber uint16 // 0x00

	// NumberOfColumn is the count of column.
	NumberOfColumn uint16 // 0x02

	// Pos is the length of all screen.
	Pos uint16 // 0x04

	// BytesInColumn is the bytes in all column
	BytesInColumn uint16 // 0x06

	_ [4]byte // 0x08

	// Length is the column length.
	Length uint16 // 0x0c

	// IntervalCount is the number of different SET/ENUM column.
	IntervalCount uint16 // 0x0e

	// IntervalParts is the number of different strings in SET/ENUM column.
	IntervalParts uint16 // 0x10

	// IntLength is the column length.
	IntLength uint16 // 0x12

	_ [6]byte // 0x14

	// NumberOfNullColumn is the number of nullable columns.
	NumberOfNullColumn uint16 // 0x1a

	// CommentLength is the comment length of all column.
	CommentLength uint16 // 0x1c
}

// ColumnField defines the column information of field.
type ColumnField struct {
	_ [2]byte // 0x00

	// Length is the column length.
	Length uint8 // 0x02

	// BytesInColumn is the bytes length.
	BytesInColumn uint16 // 0x03

	_ [2]byte // 0x05

	// Unireg unknown.
	Unireg uint8 // 0x07

	// Flags is the column flag, such as FIELDFLAG_MAYBE_NULL
	Flags uint16 // 0x08

	// UniregType is the type for field, such as NEXT_NUMBER
	UniregType uint8 // 0x0a

	// CharsetLow is the charset number (<<8)
	CharsetLow uint8 // 0x0b

	// IntervalNumber is the number of field.
	IntervalNumber uint8 // 0x0c

	// DataType is the field type, see enum_field_types.
	DataType uint8 // 0x0d

	// Charset is the charset number.
	Charset uint8 // 0x0e

	// CommentLength is the comment string length of field.
	CommentLength uint8 // 0x0f

	_ [1]byte
}
