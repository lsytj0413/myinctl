package charset

// Collation defines the collation, this will affect the sort of charset.
// Rules:
//  1. suffix _ci: case insensitive
//  2. suffix _cs: case sensitive
//  3. suffix _bin: order by binary
// NOTE:
//   SHOW COLLATION;
type Collation struct {
	// Name is the collation name
	Name string

	// Charset is the charset name witch the collation belong to
	Charset string

	// ID is the collation id
	ID int

	// IsDefault will be true if the collation is the default collation in charset
	// There will only one default collation in charset
	IsDefault bool

	// Sortlen is the sort bytes length of collation
	Sortlen int
}

// Charset defines the charset.
// NOTE:
//   show character set;
type Charset struct {
	// Name is the name of charset
	Name string

	// Description is the desc of charset
	Description string

	// DefaultCollation is the name of default collation
	DefaultCollation string

	// MaxLength is the max bytes length of charset
	MaxLength int
}

// Factory is the interface for retrieve charset & collation
type Factory interface {
	// FindCharset will return the Charset from charset name
	FindCharset(charsetName string) (Charset, error)

	// FindCollation will return the Collation from collation name
	FindCollation(collationName string) (Collation, error)

	// FindCollationsFromCharset will return the array of Collation belong the charset name
	FindCollationsFromCharset(charsetName string) []Collation

	// FindCollationFromID will return the Collation from collation id
	FindCollationFromID(id int) (Collation, error)
}
