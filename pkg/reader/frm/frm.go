package frm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"

	"github.com/lsytj0413/myinctl/pkg/reader/frm/internal"
)

type defReader struct {
}

func Readn(r io.Reader, n int) ([]byte, error) {
	data := make([]byte, n)
	v, err := r.Read(data)
	if err != nil {
		return nil, err
	}

	if v != n {
		return nil, errors.Errorf("Readn failed, expect %v, actual %v", n, v)
	}
	return data, nil
}

func Read1(r io.Reader) (byte, error) {
	data, err := Readn(r, 1)
	if err != nil {
		return 0, err
	}

	return data[0], nil
}

func (d *defReader) Read(r io.ReadSeeker) (*TableDefinition, error) {
	frmFileHeader := &internal.FrmFileHeader{}
	err := binary.Read(r, binary.LittleEndian, frmFileHeader)
	if err != nil {
		return nil, err
	}

	if frmFileHeader.MagicNumber != uint16(0x01FE) {
		return nil, errors.Errorf("Read frm file header failed, wrong magic number: %v", frmFileHeader.MagicNumber)
	}

	_, err = r.Seek(int64(frmFileHeader.IOSize), 0)
	if err != nil {
		return nil, err
	}

	// key info section length: frmFileHeader.KeyInfoLength
	keyInfoSectionHeader := &internal.FrmKeyInfoSectionHeader{}
	err = binary.Read(r, binary.LittleEndian, keyInfoSectionHeader)
	if err != nil {
		return nil, err
	}

	keys, keyParts := keyInfoSectionHeader.Count()
	fmt.Println("keys: ", keys)
	fmt.Println("KeyParts: ", keyParts)

	for i := 0; i < keys; i++ {
		fmt.Println("read index key: ", i)
		keyMeta := &internal.KeyMetadata{}
		err = binary.Read(r, binary.LittleEndian, keyMeta)
		if err != nil {
			return nil, err
		}

		fmt.Printf("%+v\n", keyMeta)
		for j := keyMeta.UserDefinedKeyParts; j > 0; j-- {
			fmt.Println("read index key parts: ", j)
			keyParts := &internal.KeyParts{}
			err = binary.Read(r, binary.LittleEndian, keyParts)
			keyParts.FieldNumber = keyParts.FieldNumber & 0x3FFF
			if err != nil {
				return nil, err
			}

			fmt.Printf("%+v\n", keyParts)
		}
	}

	// read key names
	{
		terminator, err := Read1(r)
		if err != nil {
			return nil, err
		}

		for i := 0; i < keys; i++ {
			keyName := ""
			for {
				b, err := Read1(r)
				if err != nil {
					return nil, err
				}
				if b == terminator {
					fmt.Println("index: ", i, "  Name: ", keyName)
					break
				}
				keyName += string(b)
			}
		}
	}

	// read key comment
	{
		// Skip 2 bytes
		_, err = Readn(r, 1)
		if err != nil {
			return nil, err
		}

		for i := 0; i < keys; i++ {
			// if flags & HA_USES_COMMENT
			data, err := Readn(r, 2)
			if err != nil {
				return nil, err
			}

			length := uint16(data[1])<<8 | uint16(data[0])
			fmt.Println("index: ", i, "  CommentLength: ", length)
			data, err = Readn(r, int(length))
			if err != nil {
				return nil, err
			}
			fmt.Println("index: ", i, "  Comment: ", string(data))
		}
	}

	// // read column
	recordOffset := int32(frmFileHeader.IOSize) + int32(frmFileHeader.TmpKeyLength) + int32(frmFileHeader.RecLength)
	columnOffset := ((recordOffset)/int32(frmFileHeader.IOSize)+1)*int32(frmFileHeader.IOSize) + 256 // plus 256, because it will be 256 + fixed 2 bytes(value 01)
	fmt.Println("columnOffset: ", columnOffset)
	_, err = r.Seek(int64(columnOffset), 0)
	if err != nil {
		return nil, err
	}

	columnMeta := &internal.ColumnMetadata{}
	err = binary.Read(r, binary.LittleEndian, columnMeta)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", columnMeta)

	// skip 52 bytes
	_, err = Readn(r, 52)
	if err != nil {
		return nil, err
	}

	columnNameLength := 0
	// read column names
	{
		for i := 0; i < int(columnMeta.NumberOfColumn); i++ {
			n, err := Read1(r)
			if err != nil {
				return nil, err
			}
			fmt.Println("read column name: ", i, " length: ", int(n))

			// the name has 2 bytes padding expect the last column name
			length := int(n) + 2
			if i == int(columnMeta.NumberOfColumn)-1 {
				length -= 2
			}

			data, err := Readn(r, length)
			if err != nil {
				return nil, err
			}
			fmt.Printf("column name: %v\n", string(data[0:length]))
			fmt.Println("")

			columnNameLength += int(n) - 1
		}
	}
	fmt.Println("columnNameLength: ", columnNameLength)

	columnes := make([]*internal.ColumnField, 0, int(columnMeta.NumberOfColumn))
	{
		// read column field
		for i := 0; i < int(columnMeta.NumberOfColumn); i++ {
			v := &internal.ColumnField{}
			err = binary.Read(r, binary.LittleEndian, v)
			if err != nil {
				return nil, err
			}
			fmt.Printf("columnField:%v, %+v\n", i, v)

			columnes = append(columnes, v)
		}
	}

	// read enum
	skipLength := 1 + columnNameLength + int(columnMeta.NumberOfColumn)
	_, err = r.Seek(int64(skipLength), 1)
	if err != nil {
		return nil, err
	}

	{
		fmt.Println("read enum skip: ", skipLength)
		hasEnum := false
		for i := 0; i < int(columnMeta.NumberOfColumn); i++ {
			cf := columnes[i]
			if cf.IntervalNumber == 0 {
				continue
			}

			hasEnum = true
			// is enum field
			splitChar, err := Read1(r)
			if err != nil {
				return nil, err
			}

			values := make([]byte, 0)
			for {
				b, err := Read1(r)
				if err != nil {
					return nil, err
				}

				if b == splitChar {
					break
				}
				values = append(values, b)
			}

			enums := bytes.Split((values), []byte{values[0]})

			enumStr := []string{}
			for _, v := range enums {
				if len(v) > 0 {
					enumStr = append(enumStr, string(v))
				}
			}
			fmt.Printf("field: %v,  enums: %v\n", i, enumStr)

			// put the splitChar back
			_, err = r.Seek(int64(-1), 1)
			if err != nil {
				return nil, err
			}
		}
		if hasEnum {
			// skip last bytes, because the splitChar has been put back
			_, err = Read1(r)
			if err != nil {
				return nil, err
			}
		}
	}

	// read comment
	{
		for i := 0; i < int(columnMeta.NumberOfColumn); i++ {
			cf := columnes[i]
			if cf.CommentLength == 0 {
				continue
			}

			data, err := Readn(r, int(cf.CommentLength))
			if err != nil {
				return nil, err
			}

			fmt.Printf("column: %v, comment: %v\n", i, string(data))
		}
	}

	// read table comment
	{
		_, err = r.Seek(int64(columnOffset-256+46), 0) // between frm file header and column data
		if err != nil {
			return nil, err
		}
		n, err := Read1(r)
		if err != nil {
			return nil, err
		}

		data, err := Readn(r, int(n))
		if err != nil {
			return nil, err
		}
		fmt.Println("table comment: ", string(data))
	}

	// read defaults values
	{
		offset := int(frmFileHeader.IOSize) + int(frmFileHeader.TmpKeyLength)
		_, err = r.Seek(int64(offset)+int64(1), 0)
		if err != nil {
			return nil, err
		}
		data, err := Readn(r, int(frmFileHeader.RecLength))
		if err != nil {
			return nil, err
		}
		fmt.Println("defaults values: ", len(data))
	}

	// read the engine data
	{
		offset := int64(frmFileHeader.IOSize) + int64(frmFileHeader.TmpKeyLength) + int64(frmFileHeader.RecLength)
		_, err = r.Seek(int64(offset)+int64(2), 0)
		if err != nil {
			return nil, err
		}

		engineLenData, err := Readn(r, 2)
		if err != nil {
			return nil, err
		}
		engineLen := int(engineLenData[1])<<8 | int(engineLenData[0])
		engineStr, err := Readn(r, engineLen)
		if err != nil {
			return nil, err
		}
		fmt.Println("engine: ", string(engineStr))

		var partLen uint32
		err = binary.Read(r, binary.LittleEndian, &partLen)
		if err != nil {
			return nil, err
		}
		partStr, err := Readn(r, int(partLen))
		if err != nil {
			return nil, err
		}
		fmt.Println("partStr: ", string(partStr))
	}

	return nil, nil
}

func (d *defReader) ReadFile(file string) (*TableDefinition, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return d.Read(bytes.NewReader(f))
}

func NewReader() Reader {
	return &defReader{}
}
