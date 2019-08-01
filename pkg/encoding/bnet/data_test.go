package bnet_test

import "encoding/binary"

type TTUint8 uint8
type TTBool8 bool
type TTUint16 uint16
type TTUint32 uint32
type TTBool32 bool
type TTUint64 uint64
type TTString string
type TTArray [4]uint8
type TTUint8Slice []uint8
type TTStringList []string

type nestedStruct struct {
	U8     uint8
	U8TT   TTUint8
	B8F    bool    `bnet:"size-uint8"`
	B8FTT  TTBool8 `bnet:"size-uint8"`
	B8T    bool    `bnet:"size-uint8"`
	B8TTT  TTBool8 `bnet:"size-uint8"`
	U16    uint16
	U16BE  uint16 `bnet:"bigendian"`
	U16TT  TTUint16
	U32    uint32
	U32BE  uint32 `bnet:"bigendian"`
	U32TT  TTUint32
	B32F   bool     `bnet:"size-uint32"`
	B32FTT TTBool32 `bnet:"size-uint32"`
	B32T   bool     `bnet:"size-uint32"`
	B32TTT TTBool32 `bnet:"size-uint32"`
	U64    uint64
	U64BE  uint64 `bnet:"bigendian"`
	U64TT  TTUint64

	U8A   [4]uint8
	U8ATT TTArray

	U8SliceCount   uint32       `bnet:"save-U8SliceCount"`
	U8Slice        []uint8      `bnet:"len-U8SliceCount"`
	U8SliceCountTT TTUint32     `bnet:"save-U8SliceCountTT"`
	U8SliceTT      TTUint8Slice `bnet:"len-U8SliceCountTT"`

	StrA      string
	StrB      string
	StrEmpty  string
	StrTT     TTString
	StrList   []string
	StrListTT TTStringList
}

type supportedTypes struct {
	U8     uint8
	U8TT   TTUint8
	B8F    bool    `bnet:"size-uint8"`
	B8FTT  TTBool8 `bnet:"size-uint8"`
	B8T    bool    `bnet:"size-uint8"`
	B8TTT  TTBool8 `bnet:"size-uint8"`
	U16    uint16
	U16BE  uint16 `bnet:"bigendian"`
	U16TT  TTUint16
	U32    uint32
	U32BE  uint32 `bnet:"bigendian"`
	U32TT  TTUint32
	B32F   bool     `bnet:"size-uint32"`
	B32FTT TTBool32 `bnet:"size-uint32"`
	B32T   bool     `bnet:"size-uint32"`
	B32TTT TTBool32 `bnet:"size-uint32"`
	U64    uint64
	U64BE  uint64 `bnet:"bigendian"`
	U64TT  TTUint64

	U8A   [4]uint8
	U8ATT TTArray

	U8SliceCount   uint32       `bnet:"save-U8SliceCount"`
	U8Slice        []uint8      `bnet:"len-U8SliceCount"`
	U8SliceCountTT TTUint32     `bnet:"save-U8SliceCountTT"`
	U8SliceTT      TTUint8Slice `bnet:"len-U8SliceCountTT"`

	StrA      string
	StrB      string
	StrEmpty  string
	StrTT     TTString
	StrList   []string
	StrListTT TTStringList

	Nested    nestedStruct
	NestedPtr *nestedStruct
}

var supportedTypesTests = struct {
	encoded []byte
	decoded supportedTypes
}{
	encoded: []byte{
		0x02,       // U8
		0x02,       // U8TT
		0x00,       // B8F
		0x00,       // B8FTT
		0x01,       // B8T
		0x01,       // B8FTT
		0x02, 0x00, // U16
		0x02, 0x00, // U16BE
		0x02, 0x00, // U16TT
		0x02, 0x00, 0x00, 0x00, // U32
		0x02, 0x00, 0x00, 0x00, // U32BE
		0x02, 0x00, 0x00, 0x00, // U32TT
		0x00, 0x00, 0x00, 0x00, // U32F
		0x00, 0x00, 0x00, 0x00, // U32FTT
		0x01, 0x00, 0x00, 0x00, // U32T
		0x01, 0x00, 0x00, 0x00, // U32TT
		0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // U64
		0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // U64BE
		0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // U64TT
		0x02, 0x02, 0x02, 0x02, // U8A
		0x02, 0x02, 0x02, 0x02, // U8ATT
		0x04, 0x00, 0x00, 0x00, // U8SliceCount
		0x01, 0x02, 0x03, 0x04, // U8Slice
		0x04, 0x00, 0x00, 0x00, // U8SliceCountTT
		0x01, 0x02, 0x03, 0x04, // U8SliceTT
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, // StrA
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, //StrB
		0x00,                               // StrEmpty
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, // StrTT
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x00, // StrList
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x00, // StrListTT

		// Nested
		0x02,       // U8
		0x02,       // U8TT
		0x00,       // B8F
		0x00,       // B8FTT
		0x01,       // B8T
		0x01,       // B8FTT
		0x02, 0x00, // U16
		0x02, 0x00, // U16BE
		0x02, 0x00, // U16TT
		0x02, 0x00, 0x00, 0x00, // U32
		0x02, 0x00, 0x00, 0x00, // U32BE
		0x02, 0x00, 0x00, 0x00, // U32TT
		0x00, 0x00, 0x00, 0x00, // U32F
		0x00, 0x00, 0x00, 0x00, // U32FTT
		0x01, 0x00, 0x00, 0x00, // U32T
		0x01, 0x00, 0x00, 0x00, // U32TT
		0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // U64
		0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // U64BE
		0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // U64TT
		0x02, 0x02, 0x02, 0x02, // U8A
		0x02, 0x02, 0x02, 0x02, // U8ATT
		0x04, 0x00, 0x00, 0x00, // U8SliceCount
		0x01, 0x02, 0x03, 0x04, // U8Slice
		0x04, 0x00, 0x00, 0x00, // U8SliceCountTT
		0x01, 0x02, 0x03, 0x04, // U8SliceTT
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, // StrA
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, //StrB
		0x00,                               // StrEmpty
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, // StrTT
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x00, // StrList
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x00, // StrListTT

		// NestedPtr
		0x02,       // U8
		0x02,       // U8TT
		0x00,       // B8F
		0x00,       // B8FTT
		0x01,       // B8T
		0x01,       // B8FTT
		0x02, 0x00, // U16
		0x02, 0x00, // U16BE
		0x02, 0x00, // U16TT
		0x02, 0x00, 0x00, 0x00, // U32
		0x02, 0x00, 0x00, 0x00, // U32BE
		0x02, 0x00, 0x00, 0x00, // U32TT
		0x00, 0x00, 0x00, 0x00, // U32F
		0x00, 0x00, 0x00, 0x00, // U32FTT
		0x01, 0x00, 0x00, 0x00, // U32T
		0x01, 0x00, 0x00, 0x00, // U32TT
		0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // U64
		0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // U64BE
		0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // U64TT
		0x02, 0x02, 0x02, 0x02, // U8A
		0x02, 0x02, 0x02, 0x02, // U8ATT
		0x04, 0x00, 0x00, 0x00, // U8SliceCount
		0x01, 0x02, 0x03, 0x04, // U8Slice
		0x04, 0x00, 0x00, 0x00, // U8SliceCountTT
		0x01, 0x02, 0x03, 0x04, // U8SliceTT
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, // StrA
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, //StrB
		0x00,                               // StrEmpty
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, // StrTT
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x00, // StrList
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x00, // StrListTT

	},
	decoded: supportedTypes{
		U8:     0x02,
		U8TT:   0x02,
		B8F:    false,
		B8FTT:  false,
		B8T:    true,
		B8TTT:  true,
		U16:    binary.LittleEndian.Uint16([]byte{0x02, 0x00}),
		U16BE:  binary.BigEndian.Uint16([]byte{0x02, 0x00}),
		U16TT:  TTUint16(binary.LittleEndian.Uint16([]byte{0x02, 0x00})),
		U32:    binary.LittleEndian.Uint32([]byte{0x02, 0x00, 0x00, 0x00}),
		U32BE:  binary.BigEndian.Uint32([]byte{0x02, 0x00, 0x00, 0x00}),
		U32TT:  TTUint32(binary.LittleEndian.Uint32([]byte{0x02, 0x00, 0x00, 0x00})),
		B32F:   false,
		B32FTT: false,
		B32T:   true,
		B32TTT: true,
		U64:    binary.LittleEndian.Uint64([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
		U64BE:  binary.BigEndian.Uint64([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
		U64TT:  TTUint64(binary.LittleEndian.Uint64([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})),
		U8A:    [4]uint8{0x02, 0x02, 0x02, 0x02},
		U8ATT:  TTArray{0x02, 0x02, 0x02, 0x02},

		U8SliceCount:   4,
		U8Slice:        []uint8{0x01, 0x02, 0x03, 0x04},
		U8SliceCountTT: 4,
		U8SliceTT:      []uint8{0x01, 0x02, 0x03, 0x04},

		StrA:      "hello",
		StrB:      "hello",
		StrEmpty:  "",
		StrTT:     "hello",
		StrList:   []string{"hello", "hello"},
		StrListTT: []string{"hello", "hello"},

		Nested: nestedStruct{
			U8:     0x02,
			U8TT:   0x02,
			B8F:    false,
			B8FTT:  false,
			B8T:    true,
			B8TTT:  true,
			U16:    binary.LittleEndian.Uint16([]byte{0x02, 0x00}),
			U16BE:  binary.BigEndian.Uint16([]byte{0x02, 0x00}),
			U16TT:  TTUint16(binary.LittleEndian.Uint16([]byte{0x02, 0x00})),
			U32:    binary.LittleEndian.Uint32([]byte{0x02, 0x00, 0x00, 0x00}),
			U32BE:  binary.BigEndian.Uint32([]byte{0x02, 0x00, 0x00, 0x00}),
			U32TT:  TTUint32(binary.LittleEndian.Uint32([]byte{0x02, 0x00, 0x00, 0x00})),
			B32F:   false,
			B32FTT: false,
			B32T:   true,
			B32TTT: true,
			U64:    binary.LittleEndian.Uint64([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
			U64BE:  binary.BigEndian.Uint64([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
			U64TT:  TTUint64(binary.LittleEndian.Uint64([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})),
			U8A:    [4]uint8{0x02, 0x02, 0x02, 0x02},
			U8ATT:  TTArray{0x02, 0x02, 0x02, 0x02},

			U8SliceCount:   4,
			U8Slice:        []uint8{0x01, 0x02, 0x03, 0x04},
			U8SliceCountTT: 4,
			U8SliceTT:      []uint8{0x01, 0x02, 0x03, 0x04},

			StrA:      "hello",
			StrB:      "hello",
			StrEmpty:  "",
			StrTT:     "hello",
			StrList:   []string{"hello", "hello"},
			StrListTT: []string{"hello", "hello"},
		},

		NestedPtr: &nestedStruct{
			U8:     0x02,
			U8TT:   0x02,
			B8F:    false,
			B8FTT:  false,
			B8T:    true,
			B8TTT:  true,
			U16:    binary.LittleEndian.Uint16([]byte{0x02, 0x00}),
			U16BE:  binary.BigEndian.Uint16([]byte{0x02, 0x00}),
			U16TT:  TTUint16(binary.LittleEndian.Uint16([]byte{0x02, 0x00})),
			U32:    binary.LittleEndian.Uint32([]byte{0x02, 0x00, 0x00, 0x00}),
			U32BE:  binary.BigEndian.Uint32([]byte{0x02, 0x00, 0x00, 0x00}),
			U32TT:  TTUint32(binary.LittleEndian.Uint32([]byte{0x02, 0x00, 0x00, 0x00})),
			B32F:   false,
			B32FTT: false,
			B32T:   true,
			B32TTT: true,
			U64:    binary.LittleEndian.Uint64([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
			U64BE:  binary.BigEndian.Uint64([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
			U64TT:  TTUint64(binary.LittleEndian.Uint64([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})),
			U8A:    [4]uint8{0x02, 0x02, 0x02, 0x02},
			U8ATT:  TTArray{0x02, 0x02, 0x02, 0x02},

			U8SliceCount:   4,
			U8Slice:        []uint8{0x01, 0x02, 0x03, 0x04},
			U8SliceCountTT: 4,
			U8SliceTT:      []uint8{0x01, 0x02, 0x03, 0x04},

			StrA:      "hello",
			StrB:      "hello",
			StrEmpty:  "",
			StrTT:     "hello",
			StrList:   []string{"hello", "hello"},
			StrListTT: []string{"hello", "hello"},
		},
	},
}
