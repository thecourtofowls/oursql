package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMerkleNode(t *testing.T) {
	data := [][]byte{
		[]byte("node1"),
		[]byte("node2"),
		[]byte("node3"),
	}

	// Level 1

	n1 := NewMerkleNode(nil, nil, data[0])
	n2 := NewMerkleNode(nil, nil, data[1])
	n3 := NewMerkleNode(nil, nil, data[2])
	n4 := NewMerkleNode(nil, nil, data[2])

	// Level 2
	n5 := NewMerkleNode(n1, n2, nil)
	n6 := NewMerkleNode(n3, n4, nil)

	// Level 3
	n7 := NewMerkleNode(n5, n6, nil)

	assert.Equal(
		t,
		"64b04b718d8b7c5b6fd17f7ec221945c034cfce3be4118da33244966150c4bd4",
		hex.EncodeToString(n5.Data),
		"Level 1 hash 1 is correct",
	)
	assert.Equal(
		t,
		"08bd0d1426f87a78bfc2f0b13eccdf6f5b58dac6b37a7b9441c1a2fab415d76c",
		hex.EncodeToString(n6.Data),
		"Level 1 hash 2 is correct",
	)
	assert.Equal(
		t,
		"4e3e44e55926330ab6c31892f980f8bfd1a6e910ff1ebc3f778211377f35227e",
		hex.EncodeToString(n7.Data),
		"Root hash is correct",
	)
}

func TestNewMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("node1"),
		[]byte("node2"),
		[]byte("node3"),
	}
	// Level 1
	n1 := NewMerkleNode(nil, nil, data[0])
	n2 := NewMerkleNode(nil, nil, data[1])
	n3 := NewMerkleNode(nil, nil, data[2])
	n4 := NewMerkleNode(nil, nil, data[2])

	// Level 2
	n5 := NewMerkleNode(n1, n2, nil)
	n6 := NewMerkleNode(n3, n4, nil)

	// Level 3
	n7 := NewMerkleNode(n5, n6, nil)

	rootHash := fmt.Sprintf("%x", n7.Data)
	mTree := NewMerkleTree(data)

	assert.Equal(t, rootHash, fmt.Sprintf("%x", mTree.RootNode.Data), "Merkle tree root hash is correct")
}

func TestCases(t *testing.T) {
	data := map[string]string{
		"72ea5be9289faf2c61b5c2e50e5afc580229a3034dbbebcb3d74d55e0f3a251a": "0cff8f020102ff9000010a0000fe09a6ff900005fe02083cff810301010b5472616e73616374696f6e01ff8200010401024944010a00010356696e01ff86000104566f757401ff8a00010454696d65010400000024ff85020101155b5d7472616e73616374696f6e2e5458496e70757401ff860001ff84000040ff83030101075458496e70757401ff84000104010454786964010a000104566f757401040001095369676e6174757265010a0001065075624b6579010a00000025ff89020101165b5d7472616e73616374696f6e2e54584f757470757401ff8a0001ff8800002fff870301010854584f757470757401ff88000102010556616c7565010800010a5075624b657948617368010a000000fe010cff82012030a0312fb11a58509f4282787ffa0b93be1daa805fbac9d08e3be246a02d9ea4010101204e4140ac2342f9098992778efed0bb187468b0dec6128627f6958e3575e1dc7302400f2cf8ad35211ef6fc318649fc10e73c2d90fd7be85836b761f79511437f3cc79885e790293b4e79e6f027383b2bab197beaecf9f67c5e94d20b5a7e49df1cef014013ef8ac7b9fae93bdaf0e9bdb92a451f7062c2230256473b23c0f6aa85dd720ace9e51072431962abb75bd88170438fddbf127b684d945de95bf12a958b2fa4c00010201fef03f011420b4f541f7b1f1a66557cc6dffdfb8d05bfe99530001fe22400114907ee81795c163e82290649290ba2d01b0eabb430001fcb522039a00fe02083cff810301010b5472616e73616374696f6e01ff8200010401024944010a00010356696e01ff86000104566f757401ff8a00010454696d65010400000024ff85020101155b5d7472616e73616374696f6e2e5458496e70757401ff860001ff84000040ff83030101075458496e70757401ff84000104010454786964010a000104566f757401040001095369676e6174757265010a0001065075624b6579010a00000025ff89020101165b5d7472616e73616374696f6e2e54584f757470757401ff8a0001ff8800002fff870301010854584f757470757401ff88000102010556616c7565010800010a5075624b657948617368010a000000fe010cff8201209bffc10782afc040e336be70c45f2eea9fe4ea84a275dbe3b14a100cefb49f4d0101012030a0312fb11a58509f4282787ffa0b93be1daa805fbac9d08e3be246a02d9ea401020140fa0365c8d3456e372c5f589c6c004eae7919c98eb1324215f82d7e98ac7393dcbba2ee58f5c854f412b39117baac7f644425fe8fd0573ea6b258b23d8e17ce38014013ef8ac7b9fae93bdaf0e9bdb92a451f7062c2230256473b23c0f6aa85dd720ace9e51072431962abb75bd88170438fddbf127b684d945de95bf12a958b2fa4c000102014001144461cf6d2ac6218f6b0afc9f2fcd1095f56f082b0001fe1c400114907ee81795c163e82290649290ba2d01b0eabb430001fcb522039e00fe020a3cff810301010b5472616e73616374696f6e01ff8200010401024944010a00010356696e01ff86000104566f757401ff8a00010454696d65010400000024ff85020101155b5d7472616e73616374696f6e2e5458496e70757401ff860001ff84000040ff83030101075458496e70757401ff84000104010454786964010a000104566f757401040001095369676e6174757265010a0001065075624b6579010a00000025ff89020101165b5d7472616e73616374696f6e2e54584f757470757401ff8a0001ff8800002fff870301010854584f757470757401ff88000102010556616c7565010800010a5075624b657948617368010a000000fe010eff820120dc7c99eb7e49047d274deba6eae67e1113a087b6260dee8b0cf66a8e4ed372f3010101209bffc10782afc040e336be70c45f2eea9fe4ea84a275dbe3b14a100cefb49f4d010201403745d9e066ec5b97ed2a593649a016bf155625d0e75732c0d3b69793809703b121988c31ac4c9b7f06b50d705099f0ea46a898007a1a606a3359d22769a7505d014013ef8ac7b9fae93bdaf0e9bdb92a451f7062c2230256473b23c0f6aa85dd720ace9e51072431962abb75bd88170438fddbf127b684d945de95bf12a958b2fa4c00010201fe084001144461cf6d2ac6218f6b0afc9f2fcd1095f56f082b0001fe10400114907ee81795c163e82290649290ba2d01b0eabb430001fcb52203a000fe02083cff810301010b5472616e73616374696f6e01ff8200010401024944010a00010356696e01ff86000104566f757401ff8a00010454696d65010400000024ff85020101155b5d7472616e73616374696f6e2e5458496e70757401ff860001ff84000040ff83030101075458496e70757401ff84000104010454786964010a000104566f757401040001095369676e6174757265010a0001065075624b6579010a00000025ff89020101165b5d7472616e73616374696f6e2e54584f757470757401ff8a0001ff8800002fff870301010854584f757470757401ff88000102010556616c7565010800010a5075624b657948617368010a000000fe010cff820120a6599536dfb2f44241cec8bd0e5cd8a388f1393df014bad922c2050d9f5c9bef010101209bffc10782afc040e336be70c45f2eea9fe4ea84a275dbe3b14a100cefb49f4d0240a911df3ffde2afb73ab711818943aee7519f18c801ada39b5baae41833a02d128bf06699daa27c57c7b16028a3dfda7345b664258a75aa12d61977c84ba125e901405fb8340714b33491dd044f9853f064605b620a0db90d95835ba0343103321c4eef34091cf5762bda749f8e5c198127a0c75680d8728ca778f185664cb57cf15a00010201fef03f011420b4f541f7b1f1a66557cc6dffdfb8d05bfe99530001fef03f01144461cf6d2ac6218f6b0afc9f2fcd1095f56f082b0001fcb52203a400fe01713cff810301010b5472616e73616374696f6e01ff8200010401024944010a00010356696e01ff86000104566f757401ff8a00010454696d65010400000024ff85020101155b5d7472616e73616374696f6e2e5458496e70757401ff860001ff84000040ff83030101075458496e70757401ff84000104010454786964010a000104566f757401040001095369676e6174757265010a0001065075624b6579010a00000025ff89020101165b5d7472616e73616374696f6e2e54584f757470757401ff8a0001ff8800002fff870301010854584f757470757401ff88000102010556616c7565010800010a5075624b657948617368010a00000077ff820120469d402e3017e09074a98b053ec7127975aa93cd5624ac7f8eeb8b54c72fc7800101020102283233303861616163396635636361316438326331343833383364656261323261383162663561313500010101fe24400114907ee81795c163e82290649290ba2d01b0eabb430001fcb52203a600",
	}

	for output, input := range data {
		inputbytes, _ := hex.DecodeString(input)
		var buff bytes.Buffer
		var transactions [][]byte

		buff.Write(inputbytes)
		dec := gob.NewDecoder(&buff)
		dec.Decode(&transactions)

		mTree := NewMerkleTree(transactions)

		rootHash := mTree.RootNode.Data

		assert.Equal(t, output, hex.EncodeToString(rootHash), "Merkle tree root hash is correct")
	}

}
