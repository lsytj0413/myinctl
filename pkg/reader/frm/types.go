package frm

import (
	"io"
)

// https://dev.mysql.com/doc/internals/en/frm-file-format.html
// https://github.com/mysql/mysql-server/blob/7ed30a748964c009d4909cb8b4b22036ebdef239/sql/table.cc#L7749
// https://gist.github.com/mnpenner/92b261dc5d677fd63f694a7fa370ab50#file-frm_reader-py-L1457
// http://docs.dbsake.net/en/stable/appendix/frm_format.html#id11

type Reader interface {
	Read(reader io.ReadSeeker) (*TableDefinition, error)
	ReadFile(file string) (*TableDefinition, error)
}

type TableDefinition struct {
}

type Header struct {
	FrmVersion          string
	LegacyDatabaseType  int8
	IOSize              int16
	Length              int32
	TmpKeyLength        int16
	RevLength           int16
	MaxRows             int32
	MinRows             int32
	UsePackFields       int8
	KeyInfoLength       int16
	CreateOptions       int16
	FrmFileVersion      string
	AvgRowLength        int32
	DefaultTableCharset int8
	RowType             int8
	KeyLength           int32
	MysqlVersionID      int32
	ExtraSize           int32
	ExtraRecBufLength   int16
	DefaultPartDBType   int8
	KeyBlockSize        int16
}

const (
	MagicNumber = 0xfe01
)

// The enum value of table category, see *enum enum_table_category*
const (
	TableUnknownCategory     = 0
	TableCategoryTemporary   = 1
	TableCategoryUser        = 2
	TableCategorySystem      = 3
	TableCategoryInformation = 4
	TableCategoryLog         = 5
	TableCategoryPerformance = 6
	TableCategoryRPLInfo     = 7
	TableCategoryGtid        = 8
)

// legacyDBType
const (
	DBTypeUnknown           = 0
	DBTypeDiabIsam          = 1
	DBTypeHash              = 2
	DBTypeMisam             = 3
	DBTypePisam             = 4
	DBTypeRmsIsam           = 5
	DBTypeHeap              = 6
	DBTypeIsam              = 7
	DBTypeMrgIsam           = 8
	DBTypeMyisam            = 9
	DBTypeMrgMyisam         = 10
	DBTypeBerkeleyDB        = 11
	DBTypeInnodb            = 12
	DBTypeGemini            = 13
	DBTypeNdbCluster        = 14
	DBTypeExampleDB         = 15
	DBTypeArchiveDB         = 16
	DBTypeCsvDB             = 17
	DBTypeFederatedDB       = 18
	DBTypeBlackholeDB       = 19
	DBTypePartitionDB       = 20
	DBTypeBinlog            = 21
	DBTypeSolid             = 22
	DBTypePbxt              = 23
	DBTypeTableFunction     = 24
	DBTypeMemcache          = 25
	DBTypeFalcon            = 26
	DBTypeMaria             = 27
	DBTypePerformanceSchema = 28
	DBTypeFirstDynamic      = 42
	DBTypeDefault           = 127
)

// rowType
const (
	RowTypeNotUsed    = -1
	RowTypeDefault    = 0
	RowTypeFixed      = 1
	RowTypeDynamic    = 2
	RowTypeCompressed = 3
	RowTypeRedundant  = 4
	RowTypeCompact    = 5
	RowTypePage       = 6 // unused, reserved for future versions.
)

// StatsAutoRecalc
const (
	StatsAutoRecalcDefault = 0
	StatsAutoRecalcOn      = 1
	StatsAutoRecalcOff     = 2
)

// keyAlgorithm
const (
	KeyAlgUndef    = 0
	KeyAlgBtree    = 1 // default
	KeyAlgRtree    = 2
	KeyAlgHash     = 3
	KeyAlgFulltext = 4
)

// keyType is the field type
const (
	KeyTypeEnd        = 0
	KeyTypeText       = 1
	KeyTypeBinary     = 2
	KeyTypeShortInt   = 3
	KeyTypeLongInt    = 4
	KeyTypeFloat      = 5
	KeyTypeDouble     = 6
	KeyTypeNum        = 7
	KeyTypeUshortInt  = 8
	KeyTypeUlongInt   = 9
	KeyTypeLongLong   = 10
	KeyTypeUlongLong  = 11
	KeyTypeInt24      = 12
	KeyTypeUint24     = 13
	KeyTypeInt8       = 14
	KeyTypeVartext1   = 15 // 0-255 bytes, length packed 1 byte
	KeyTypeVarbinary1 = 16
	KeyTypeVartext2   = 17 // 0-65535 bytes, length packed 2 byte
	KeyTypeVarbinary2 = 18
	KeyTypeBit        = 19
)

// filed Type
const (
	TypeDecimal    = 0
	TypeTiny       = 1
	TypeShort      = 2
	TypeLong       = 3
	TypeFloat      = 4
	TypeDouble     = 5
	TypeNull       = 6
	TypeTimestamp  = 7
	TypeLongLong   = 8
	TypeInt24      = 9
	TypeData       = 10
	TypeTime       = 11
	TypeDatatime   = 12
	TypeYear       = 13
	TypeNewDate    = 14
	TypeVarchar    = 15
	TypeBit        = 16
	TypeTimestamp2 = 17
	TypeTime2      = 18
	TypeJSON       = 245
	TypeNewDecimal = 246
	TypeEnum       = 247
	TypeSet        = 248
	TypeTinyBlob   = 249
	TypeMediumBlob = 250
	TypeLongBlob   = 251
	TypeBlob       = 252
	TypeVarString  = 253
	TypeString     = 254
	TypeGeometry   = 255
)
