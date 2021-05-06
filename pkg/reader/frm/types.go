package frm

import (
	"io"
)

// nolint
// FrmFileHeader is the frm file header struct (logical).
type FrmFileHeader struct {
	// FrmVer is the frm version.
	FrmVer int

	DatabaseType DatabaseType

	IOSize uint16

	KeyLength uint16

	DefaultValueLength uint16

	MaxRows uint32
	MinRows uint32

	KeyInfoSectionLength uint16

	CreateOptions uint16

	AvgRowLength uint32

	Charset int

	RowType RowType

	StatsSamplePages int
	StatsAutoRecalc  int

	ExtraSize uint32

	DefaultPartDBType uint8

	KeyBlockSize uint16
}

// https://dev.mysql.com/doc/internals/en/frm-file-format.html
// https://github.com/mysql/mysql-server/blob/7ed30a748964c009d4909cb8b4b22036ebdef239/sql/table.cc#L7749
// https://gist.github.com/mnpenner/92b261dc5d677fd63f694a7fa370ab50#file-frm_reader-py-L1457
// http://docs.dbsake.net/en/stable/appendix/frm_format.html#id11

type Reader interface {
	Read(reader io.ReadSeeker) (*TableDefinition, error)
	ReadFile(file string) (*TableDefinition, error)
}

// TableDefinition is the MySQL table def.
type TableDefinition struct {
	// QualifiedName is the full qualified table name, such as database.table
	QualifiedName string

	// Name is the table name
	Name string

	// ColumnNames is the array of column name
	ColumnNames []string

	// Columns is the list of column definition
	Columns []*Column

	// PrimaryKey is the PRIMARY KEY information
	PrimaryKey *KeyMeta

	// SecondaryKeys is the secondary key lists
	SecondaryKeys []*KeyMeta

	// Charset is the table charset
	Charset int

	// Collation is the table charset collation
	Collation int
}

// Column is the MySQL table column def.
// It could be read from table information_schema.columns, such as expression "SELECT * FROM information_schema.columns WHERE table_schema='db' AND table_name='tb'"
// TABLE_CATALOG: def
// TABLE_SCHEMA: myinctl
//   TABLE_NAME: user_accounts
//  COLUMN_NAME: id
// ORDINAL_POSITION: 1
// COLUMN_DEFAULT: NULL
//  IS_NULLABLE: NO
//    DATA_TYPE: bigint
// CHARACTER_MAXIMUM_LENGTH: NULL
// CHARACTER_OCTET_LENGTH: NULL
// NUMERIC_PRECISION: 19
// NUMERIC_SCALE: 0
// DATETIME_PRECISION: NULL
// CHARACTER_SET_NAME: NULL
// COLLATION_NAME: NULL
//  COLUMN_TYPE: bigint(20)
//   COLUMN_KEY: PRI
// 	   EXTRA: auto_increment
//   PRIVILEGES: select,insert,update,references
// COLUMN_COMMENT:
// GENERATION_EXPRESSION:
type Column struct {
	// TableDef is the reference to table
	TableDef *TableDefinition

	// Name is the column name. It's the COLUMN_NAME value.
	Name string

	// Ordinal is the order of current column. It's the ORDINAL_POSITION value.
	Ordinal int

	// FullType is the COLUMN_TYPE of column.
	FullType string

	// DataType is the DATA_TYPE of column.
	DataType DataType

	// Nullable is the IS_NULLABLE of column.
	Nullable bool

	// IsPrimaryKey is the primary key bit of column. If the column is defined with PRIMARY KEY in column this will be true.
	IsPrimaryKey bool

	// Length is the size for integer type, MAX length for varchar and FIXED length for char type.
	Length int

	// Scale is the number of digits to the right of the decimal point, the NUMERIC_SCALE column.
	// D in DECIMAL(M, D)
	Scale int

	// Precision is the maximum number of digits, the NUMERIC_PRECISION column.
	// M in DECIMAL(M, D)
	Precision int

	// Charset is the CHARACTER_SET_NAME column.
	Charset string

	// Collation is the COLLATION_NAME column.
	Collation string

	// MaxBytesPerChar is the bytes of char.
	// EX: If the charset is utf8, the value will be 3.
	MaxBytesPerChar int

	// IsVarLenChar is the charset have variable length per char.
	// EX: If the charset is utf8mb4, CHAR will be treaded as VARCHAR and this field will be true.
	IsVarLenChar bool

	// Attributes hold the attr for ENUM & SET type.
	Attributes map[string]interface{}
}

// KeyMeta defines the key metadata information.
// mysql> select * from information_schema.innodb_sys_indexes limit 1;
// +----------+--------+----------+------+----------+---------+-------+-----------------+
// | INDEX_ID | NAME   | TABLE_ID | TYPE | N_FIELDS | PAGE_NO | SPACE | MERGE_THRESHOLD |
// +----------+--------+----------+------+----------+---------+-------+-----------------+
// |       11 | ID_IND |       11 |    3 |        1 |     270 |     0 |              50 |
// +----------+--------+----------+------+----------+---------+-------+-----------------+
// See: https://dev.mysql.com/doc/refman/5.7/en/create-table.html#create-table-indexes-keys
type KeyMeta struct {
	// Name is the index name.
	Name string

	// Type is the index key type
	Type KeyType

	// NumberOfColumns is the key numbers
	NumberOfColumns int

	// KeyColumns is the key columns
	KeyColumns []*Column

	// KeyColumnNames is the key column names
	KeyColumnNames []string

	// KeyVarLengthColumns is the variable length key columns
	KeyVarLengthColumns []*Column

	// KeyVarLengthColumnNames is the variable length key column names
	KeyVarLengthColumnNames []string

	// KeyVarLength is the variable length of column
	// NOTE: the length of key might be different from column definitions, EX: VARCHAR
	KeyVarLength []int
}

// const (
// 	MagicNumber = 0xfe01
// )

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
