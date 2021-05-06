package innodb

// https://github.com/alibaba/innodb-java-reader/blob/master/innodb-java-reader/src/main/java/com/alibaba/innodb/java/reader/page/FilHeader.java
// https://github.com/jeremycole/innodb_ruby/blob/master/bin/innodb_space
// https://zhuanlan.zhihu.com/p/103582178

// FileHeader defines the header of every page.
// NOTE: it always start at the 0x00 in page.
// Size: 38 bytes
type FileHeader struct {
	// Checksum is the checksum value of current page.
	Checksum uint32

	// PageNumber is the number of page
	// UNIQUE.
	PageNumber uint32

	// PrevPage Number.
	PrevPage uint32

	// NextPage Number.
	NextPage uint32

	// LastLSN is the last Log Sequence Number which modified this page and save to disk.
	LastLSN uint64

	// PageType is the type of this page, available value:
	//    FILE_PAGE_TYPE_ALLOCATED: 0x0000, unused
	//    FILE_PAGE_UNDO_LOG:       0x0002
	//    FILE_PAGE_INODE:          0x0003
	//    FILE_PAGE_IBUF_FREE_LIST: 0x0004
	//    FILE_PAGE_IBUF_BITMAP:    0x0005
	//    FILE_PAGE_TYPE_SYS:       0x0006
	//    FILE_PAGE_TYPE_TRX_SYS:   0x0007
	//    FILE_PAGE_TYPE_FSP_HDR:   0x0008
	//    FILE_PAGE_TYPE_XDES:      0x0009
	//    FILE_PAGE_TYPE_BLOB:      0x000A
	//    FILE_PAGE_TYPE_COMPRESSED_BLOB: 0x0011
	//    FILE_PAGE_TYPE_COMPRESSED_BLOB2: 0x0012
	//    FILE_PAGE_UNKNOWN: 0x0013, I_S_PAGE_TYPE_UNKNOWN
	//    FILE_PAGE_RTREE: 0x45BE
	//    FILE_PAGE_INDEX:          0x45BF
	// since mysql 8
	//    FILE_PAGE_SDI: 0x45BD
	//    FILE_PAGE_LOB_INDEX: 0x0016
	//    FILE_PAGE_LOB_DATA: 0x0017
	//    FILE_PAGE_LOB_FIRST: 0x0018
	PageType uint16

	// FlushLSN is the last flush lsn, used in system-table-space
	FlushLSN uint64

	// SpaceID is the table space id
	SpaceID uint32
}

// FileTrailer defines the trailer of every page.
// NOTE: it always start at the 0x00 in page.
// Size: 8 bytes
type FileTrailer struct {
	// Checksum is the checksum value of current page.
	// It should equal to header.Chechsum
	Checksum uint32

	// LastLSN is the last Log Sequence Number which modified this page and save to disk.
	LastLSN uint32
}

// PageHeader defines the INDEX type page's data header
// Size: 36 bytes
type PageHeader struct {
	// DirSlots is the item count of page dir
	DirSlots uint16

	// HeapTop is the min address of unused space, all the address greater than this value is Free Space
	HeapTop uint16

	// HeapCount is the count of record in this page, include the delete-marked record.
	// If the first bit is 1, the record is COMPACT format, otherwise is REDUNDANT
	HeapCount uint16

	// Free is the address of first delete-marked record.
	Free uint16

	// Garbage is the bytes of all delete-marked record.
	Garbage uint16

	// LastInsert is the address of last insert record.
	LastInsert uint16

	// Direction is the insert direction.
	Direction uint16

	// DirectionCount is the count of number insert with current direction.
	DirectionCount uint16

	// RecordCount is the count of record, exclude the MIN/MAX and delete-marked record.
	RecordCount uint16

	// MaxTrxID is the last transaction id, used in secondary-index.
	MaxTrxID uint64

	// Level is the level in the tree, leaf-node is 0.
	Level uint16

	// IndexID is the key-index id this page belong to.
	IndexID uint64

	// Change this 2 field to PageSegmentHeader struct
	// // SegmentLeaf is the leaf-segment header, used in root node.
	// SegmentLeaf [10]byte

	// // SegmentTop is the non-leaf-segment header, used in root node.
	// SegmentTop [10]byte
}

// PageSegmentHeader defines the header of index segment, only used in root node page.
// Size: 20 bytes
type PageSegmentHeader struct {
	// LeafPageInodeSpace is the space of leaf page.
	LeafPageInodeSpace uint32

	// LeafPageInodePageNumber is the page number of leaf page.
	LeafPageInodePageNumber uint32

	// NonLeafPageInodeOffset is the offset of leaf page.
	LeafPageInodeOffset uint16

	// NonLeafPageInodeSpace is the space of non-leaf page.
	NonLeafPageInodeSpace uint32

	// NonLeafPageInodePageNumber is the page number of non-leaf page.
	NonLeafPageInodePageNumber uint32

	// NonLeafPageInodeOffset is the offset of non-leaf page.
	NonLeafPageInodeOffset uint16
}

// Record Format (Compact)
// | var byte | the variable bytes of variable length field's length |
// | var byte | null field bitmap |
// | 5 byte | Record Header |
// | var byte | column data, if record type = min, the bytes is "infimum", if record type = max, the bytes is "supremum", if record type = normal, the bytes is "primary,transaction_id,roll_pointer,column,column" |

// RecordHeader defines the header of each record.
type RecordHeader struct {
	// Mask1 is the first byte of header.
	// | 1bit | reserved
	// | 1bit | reserved
	// | 1bit | delete_mask
	// | 1bit | min_rec_mask, every level non-leaf min record will set this bit
	// | 4bit | n_owned, the number this record owned(in page directory)
	Mask1 uint8

	// Mask2 is the second 2 bytes of header.
	// | 13bit | heap_no, this record address in the heap
	// | 3bit  | record_type, current record type, 0=normal, 1=non-leaf,2=min,3=max
	Mask2 uint16

	// NextRecord is the relative position
	NextRecord int16
}

func (r *RecordHeader) IsDeleted() bool {
	return ((r.Mask1 >> 5) & 0x01) == 1
}

func (r *RecordHeader) OwnedRecordNumber() uint8 {
	return r.Mask1 >> 4
}

func (r *RecordHeader) RecordType() uint16 {
	return uint16((r.Mask2) & 0x07)
}

func (r *RecordHeader) HeapNumber() uint16 {
	return (r.Mask2 & 0xfff8) >> 13
}

// PageDirectorySlot is the slot of page directory.
type PageDirectorySlot struct {
	// Position is the address of current slot(max record) primary key start at page address 0.
	Position uint16
}
