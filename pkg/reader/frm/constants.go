package frm

// DatabaseType is the database type definition.
// DatabaseType is the engine of current table, see enum legacy_db_type
type DatabaseType int

const (
	// DatabaseTypeUnknown is the unknown database type
	DatabaseTypeUnknown DatabaseType = 0

	// DatabaseTypeDiabIsam is the diab_isam database type
	DatabaseTypeDiabIsam DatabaseType = 1

	// DatabaseTypeHash is the hash database type
	DatabaseTypeHash DatabaseType = 2

	// DatabaseTypeMisam is the misam database type
	DatabaseTypeMisam DatabaseType = 3

	// DatabaseTypePisam is the pisam database type
	DatabaseTypePisam DatabaseType = 4

	// DatabaseTypeRmsIsam is the rms_isam database type
	DatabaseTypeRmsIsam DatabaseType = 5

	// DatabaseTypeHeap is the heap database type
	DatabaseTypeHeap DatabaseType = 6

	// DatabaseTypeIsam is the isam database type
	DatabaseTypeIsam DatabaseType = 7

	// DatabaseTypeMrgIsam is the mrg_isam database type
	DatabaseTypeMrgIsam DatabaseType = 8

	// DatabaseTypeMyIsam is the MyIsam database type
	DatabaseTypeMyIsam DatabaseType = 9

	// DatabaseTypeMrgMyIsam is the mrg_myisam database type
	DatabaseTypeMrgMyIsam DatabaseType = 10

	// DatabaseTypeBerkeleyDB is the berkeley_db database type
	DatabaseTypeBerkeleyDB DatabaseType = 11

	// DatabaseTypeInnoDB is the innodb database type
	DatabaseTypeInnoDB DatabaseType = 12

	// DatabaseTypeGemini is the gemini database type
	DatabaseTypeGemini DatabaseType = 13

	// DatabaseTypeNdbCluster is the ndb_cluster database type
	DatabaseTypeNdbCluster DatabaseType = 14

	// DatabaseTypeExampleDB is the example_db database type
	DatabaseTypeExampleDB DatabaseType = 15

	// DatabaseTypeArchiveDB is the archive_db database type
	DatabaseTypeArchiveDB DatabaseType = 16

	// DatabaseTypeCSVDB is the csv_db database type
	DatabaseTypeCSVDB DatabaseType = 17

	// DatabaseTypeFederatedDB is the federeated_db database type
	DatabaseTypeFederatedDB DatabaseType = 18

	// DatabaseTypeBlackholeDB is the blackhold_db database type
	DatabaseTypeBlackholeDB DatabaseType = 19

	// DatabaseTypePartitionDB is the partition_db database type
	DatabaseTypePartitionDB DatabaseType = 20

	// DatabaseTypeBinlog is the binlog database type
	DatabaseTypeBinlog DatabaseType = 21

	// DatabaseTypeSolid is the solid database type
	DatabaseTypeSolid DatabaseType = 22

	// DatabaseTypePbxt is the pbxt database type
	DatabaseTypePbxt DatabaseType = 23

	// DatabaseTypeTableFunction is the table_function database type
	DatabaseTypeTableFunction DatabaseType = 24

	// DatabaseTypeMemcache is the memcache database type
	DatabaseTypeMemcache DatabaseType = 25

	// DatabaseTypeFalcon is the falcon database type
	DatabaseTypeFalcon DatabaseType = 26

	// DatabaseTypeMaria is the maria database type
	DatabaseTypeMaria DatabaseType = 27

	// DatabaseTypePerformanceSchema is the performance_schema database type
	DatabaseTypePerformanceSchema DatabaseType = 28

	// DatabaseTypeFirstDynamic is the first_dynamic database type
	DatabaseTypeFirstDynamic DatabaseType = 42

	// DatabaseTypeDefault is the default database type
	DatabaseTypeDefault DatabaseType = 127
)

// RowType is the row format definition.
// See enum row_type
type RowType int

const (
	// RowTypeNotUsed is the not used row format.
	RowTypeNotUsed RowType = -1

	// RowTypeDefault is the default row format.
	RowTypeDefault RowType = 0

	// RowTypeFixed is the fixed row format.
	RowTypeFixed RowType = 1

	// RowTypeDynamic is the dynamic row format.
	RowTypeDynamic RowType = 2

	// RowTypeCompressed is the compressed row format.
	RowTypeCompressed RowType = 3

	// RowTypeRedundant is the redundant row format.
	RowTypeRedundant RowType = 4

	// RowTypeCompact is the compact row format.
	RowTypeCompact RowType = 5

	// RowTypePage is UNUSED, reserved for future versions.
	RowTypePage RowType = 6
)

// DataType is the column DATA_TYPE value.
// See enum enum_field_types
// See https://dev.mysql.com/doc/refman/5.7/en/data-types.html
type DataType int

const (
	// DataTypeDecimal types: DECIMAL
	// NOTE: used in OLDER MySQL version.
	DataTypeDecimal DataType = 0

	// DataTypeTiny types: TINYINT, BOOL, BOOLEAN
	// Length: 1 bytes.
	DataTypeTiny DataType = 1

	// DataTypeShort types: SMALLINT
	// Length: 2 bytes.
	DataTypeShort DataType = 2

	// DataTypeLong types: INT, INTEGER
	// Length: 4 bytes.
	DataTypeLong DataType = 3

	// DataTypeFloat types: FLOAT, FLOAT(p) when p < 25
	// Length: 4 bytes.
	DataTypeFloat DataType = 4

	// DataTypeDouble types: DOUBLE, DOUBLE PRECISION, REAL, FLOAT(p) when p >= 25
	// Length 8 bytes.
	DataTypeDouble DataType = 5

	// DataTypeNull is the null type
	DataTypeNull DataType = 6

	// DataTypeTimestamp is the TIMESTAMP type, without ms support.
	DataTypeTimestamp = 7

	// DataTypeLongLong types: BIGINT
	// Length: 8 bytes.
	DataTypeLongLong = 8

	// DataTypeInt24 is the MEDIUMINT type, 3 bytes.
	// TYPES: MEDIUMINT
	DataTypeInt24 = 9

	// DataTypeDate is the DATE type, used before MySQL 5.0
	DataTypeDate = 10

	// DataTypeTime is the TIME type, without ms support.
	DataTypeTime = 11

	// DataTypeDateTime is the DATETIME type
	DataTypeDateTime = 12

	// DataTypeYear types: YEAR
	DataTypeYear = 13

	// DataTypeNewDate types: DATE
	// NOTE: used after MySQL 5.0
	DataTypeNewDate = 14

	// DataTypeVarchar types: VARCHAR, VARBINARY
	DataTypeVarchar = 15

	// DataTypeBit is the BIT type
	// TYPES: BIT
	DataTypeBit = 16

	// DataTypeTimestamp2 types: TIMESTAMP
	// NOTE: with ms support.
	DataTypeTimestamp2 = 17

	// DataTypeDatetime2 types: DATETIME
	// NOTE: with ms support
	DataTypeDatetime2 = 18

	// DataTypeTime2 types: TIME
	// NOTE: with ms support.
	DataTypeTime2 = 19

	// DataTypeJSON types: JSON
	DataTypeJSON = 245

	// DataTypeNewDecimal types: DECIMAL, DEC, FIXED, NUMERIC
	DataTypeNewDecimal = 246

	// DataTypeEnum types: ENUM
	DataTypeEnum = 247

	// DataTypeSet types: SET
	DataTypeSet = 248

	// DataTypeTinyBlob types: TINYBLOB, TINYTEXT
	// Length: < (2^8) bytes
	DataTypeTinyBlob = 249

	// DataTypeMediumBlob types: MEDIUMBLOB, MEDIUMTEXT
	// Length: < (2^24) bytes
	DataTypeMediumBlob = 250

	// DataTypeLongBlob types: LONGBLOB, LONGTEXT
	// Length: < (2^32) bytes
	DataTypeLongBlob = 251

	// DataTypeBlob types: BLOB, TEXT
	// Length: < (2^16) bytes
	DataTypeBlob = 252

	// DataTypeVarString is the VARCHAR type
	DataTypeVarString = 253

	// DataTypeString types: CHAR, BINARY
	DataTypeString = 254

	// DataTypeGeometry types: GEOMETRY, POINT, LINESTRING, POLYGON, MULTIPOINT, MULTILINESTRING, MULTIPOLYGON, GEOMETRYCOLLECTION
	DataTypeGeometry = 255
)

// KeyType is the column TYPE value.
// See enum keytype
type KeyType int

const (
	// KeyTypePrimary is the PRIMARY KEY
	KeyTypePrimary KeyType = 0

	// KeyTypeUnique is the UNIQUE INDEX
	KeyTypeUnique KeyType = 1

	// KeyTypeMultiple is the unknown index
	KeyTypeMultiple KeyType = 2

	// KeyTypeFulltext is the FULLTEXT
	KeyTypeFulltext KeyType = 3

	// KeyTypeSpatial is the SPATIAL
	KeyTypeSpatial KeyType = 4

	// KeyTypeForeign is the FOREIGN
	KeyTypeForeign KeyType = 5
)

const (
	// FileTypeMagicNumber is the magic number of frm file format.
	FileTypeMagicNumber = 0x01fe
)
