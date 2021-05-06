package innodb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/lsytj0413/myinctl/pkg/helper"
)

func Test(t *testing.T) {
	f, err := ioutil.ReadFile(helper.JoinWithProjectAbsPath("./test/user_accounts.ibd"))
	if err != nil {
		panic(err)
	}

	f = f[48*1024:]

	r := bytes.NewReader(f)

	fileHeader := &FileHeader{}
	err = binary.Read(r, binary.BigEndian, fileHeader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("FileHeader: %+v\n", fileHeader)

	pageHeader := &PageHeader{}
	err = binary.Read(r, binary.BigEndian, pageHeader)
	if err != nil {
		panic(err)
	}
	pageHeader.HeapCount = pageHeader.HeapCount & 0x7FFF
	fmt.Printf("PageHeader: %+v\n", pageHeader)

	segHeader := &PageSegmentHeader{}
	err = binary.Read(r, binary.BigEndian, segHeader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("SegHeader: %+v\n", segHeader)

	infimumHeader := &RecordHeader{}
	err = binary.Read(r, binary.BigEndian, infimumHeader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("InfimumHeader: %+v\n", infimumHeader)

	data := make([]byte, 8)
	err = binary.Read(r, binary.BigEndian, &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("InfimumData: %+v\n", string(data))

	suprmumHeader := &RecordHeader{}
	err = binary.Read(r, binary.BigEndian, suprmumHeader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("SuprmumHeader: %+v\n", suprmumHeader)

	data = make([]byte, 8)
	err = binary.Read(r, binary.BigEndian, &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("SuprmumData: %+v\n", string(data))

	startPos := 38 + 56 + 5 + int(infimumHeader.NextRecord)
	for startPos != 0 {
		recordHeader := &RecordHeader{}
		startPos -= 5
		abs, err := r.Seek(int64(startPos), 0)
		if err != nil {
			panic(err)
		}
		_ = abs

		err = binary.Read(r, binary.BigEndian, recordHeader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("RecordHeader: %+v\n", recordHeader)

		if recordHeader.NextRecord == 0 {
			break
		}

		startPos = startPos + int(recordHeader.NextRecord) + 5
	}
}
