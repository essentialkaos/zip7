package zip7

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	"github.com/essentialkaos/ek/v13/fsutil"

	check "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { check.TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

type Z7Suite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = check.Suite(&Z7Suite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *Z7Suite) TestListing7zNoComp(c *check.C) {
	info, err := List(Props{File: "testdata/test-no-compression.7z"})

	c.Assert(err, check.IsNil)
	c.Assert(info, check.NotNil)

	c.Assert(info.Path, check.Equals, "testdata/test-no-compression.7z")
	c.Assert(info.Type, check.Equals, TYPE_7Z)
	c.Assert(info.Method, check.DeepEquals, []string{"Copy"})
	c.Assert(info.Solid, check.Equals, false)
	c.Assert(info.Blocks, check.Equals, 3)
	c.Assert(info.PhysicalSize, check.Equals, 246)
	c.Assert(info.HeadersSize, check.Equals, 224)

	c.Assert(info.Files, check.HasLen, 6)

	c.Assert(info.Files[0].Path, check.Equals, "test")
	c.Assert(info.Files[0].Folder, check.Equals, "")
	c.Assert(info.Files[0].Size, check.Equals, 0)
	c.Assert(info.Files[0].PackedSize, check.Equals, 0)
	c.Assert(info.Files[0].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[0].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Attributes, check.Equals, "D_ drwxr-xr-x")
	c.Assert(info.Files[0].CRC, check.Equals, 0)
	c.Assert(info.Files[0].Encrypted, check.Equals, false)
	c.Assert(info.Files[0].Method, check.DeepEquals, []string{""})
	c.Assert(info.Files[0].Block, check.Equals, 0)
	c.Assert(info.Files[0].Comment, check.Equals, "")
	c.Assert(info.Files[0].HostOS, check.Equals, "")
	c.Assert(info.Files[0].Version, check.Equals, 0)

	c.Assert(info.Files[1].Path, check.Equals, "test/dir1")
	c.Assert(info.Files[1].Folder, check.Equals, "")
	c.Assert(info.Files[1].Size, check.Equals, 0)
	c.Assert(info.Files[1].PackedSize, check.Equals, 0)
	c.Assert(info.Files[1].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[1].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[1].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[1].Attributes, check.Equals, "D_ drwxr-xr-x")
	c.Assert(info.Files[1].CRC, check.Equals, 0)
	c.Assert(info.Files[1].Encrypted, check.Equals, false)
	c.Assert(info.Files[1].Method, check.DeepEquals, []string{""})
	c.Assert(info.Files[1].Block, check.Equals, 0)
	c.Assert(info.Files[1].Comment, check.Equals, "")
	c.Assert(info.Files[1].HostOS, check.Equals, "")
	c.Assert(info.Files[1].Version, check.Equals, 0)

	c.Assert(info.Files[2].Path, check.Equals, "test/file2.log")
	c.Assert(info.Files[2].Folder, check.Equals, "")
	c.Assert(info.Files[2].Size, check.Equals, 0)
	c.Assert(info.Files[2].PackedSize, check.Equals, 0)
	c.Assert(info.Files[2].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[2].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[2].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[2].Attributes, check.Equals, "A_ -rw-r--r--")
	c.Assert(info.Files[2].CRC, check.Equals, 0)
	c.Assert(info.Files[2].Encrypted, check.Equals, false)
	c.Assert(info.Files[2].Method, check.DeepEquals, []string{""})
	c.Assert(info.Files[2].Block, check.Equals, 0)
	c.Assert(info.Files[2].Comment, check.Equals, "")
	c.Assert(info.Files[2].HostOS, check.Equals, "")
	c.Assert(info.Files[2].Version, check.Equals, 0)

	c.Assert(info.Files[3].Path, check.Equals, "test/dir1/file1.log")
	c.Assert(info.Files[3].Folder, check.Equals, "")
	c.Assert(info.Files[3].Size, check.Equals, 11)
	c.Assert(info.Files[3].PackedSize, check.Equals, 11)
	c.Assert(info.Files[3].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[3].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[3].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[3].Attributes, check.Equals, "A_ -rw-r--r--")
	c.Assert(info.Files[3].CRC, check.Equals, 2073076399)
	c.Assert(info.Files[3].Encrypted, check.Equals, false)
	c.Assert(info.Files[3].Method, check.DeepEquals, []string{"Copy"})
	c.Assert(info.Files[3].Block, check.Equals, 0)
	c.Assert(info.Files[3].Comment, check.Equals, "")
	c.Assert(info.Files[3].HostOS, check.Equals, "")
	c.Assert(info.Files[3].Version, check.Equals, 0)

	c.Assert(info.Files[4].Path, check.Equals, "test/file1.log")
	c.Assert(info.Files[4].Folder, check.Equals, "")
	c.Assert(info.Files[4].Size, check.Equals, 6)
	c.Assert(info.Files[4].PackedSize, check.Equals, 6)
	c.Assert(info.Files[4].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[4].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[4].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[4].Attributes, check.Equals, "A_ -rw-r--r--")
	c.Assert(info.Files[4].CRC, check.Equals, 3157996776)
	c.Assert(info.Files[4].Encrypted, check.Equals, false)
	c.Assert(info.Files[4].Method, check.DeepEquals, []string{"Copy"})
	c.Assert(info.Files[4].Block, check.Equals, 1)
	c.Assert(info.Files[4].Comment, check.Equals, "")
	c.Assert(info.Files[4].HostOS, check.Equals, "")
	c.Assert(info.Files[4].Version, check.Equals, 0)

	c.Assert(info.Files[5].Path, check.Equals, "test/file3.log")
	c.Assert(info.Files[5].Folder, check.Equals, "")
	c.Assert(info.Files[5].Size, check.Equals, 5)
	c.Assert(info.Files[5].PackedSize, check.Equals, 5)
	c.Assert(info.Files[5].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[5].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[5].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[5].Attributes, check.Equals, "A_ -rw-r--r--")
	c.Assert(info.Files[5].CRC, check.Equals, 3168002993)
	c.Assert(info.Files[5].Encrypted, check.Equals, false)
	c.Assert(info.Files[5].Method, check.DeepEquals, []string{"Copy"})
	c.Assert(info.Files[5].Block, check.Equals, 2)
	c.Assert(info.Files[5].Comment, check.Equals, "")
	c.Assert(info.Files[5].HostOS, check.Equals, "")
	c.Assert(info.Files[5].Version, check.Equals, 0)
}

func (s *Z7Suite) TestListing7zMaxComp(c *check.C) {
	info, err := List(Props{File: "testdata/test-max-compression.7z"})

	c.Assert(err, check.IsNil)
	c.Assert(info, check.NotNil)

	c.Assert(info.Path, check.Equals, "testdata/test-max-compression.7z")
	c.Assert(info.Type, check.Equals, TYPE_7Z)
	c.Assert(info.Method, check.DeepEquals, []string{"LZMA2:12"})
	c.Assert(info.Solid, check.Equals, true)
	c.Assert(info.Blocks, check.Equals, 1)
	c.Assert(info.PhysicalSize, check.Equals, 254)
	c.Assert(info.HeadersSize, check.Equals, 228)

	c.Assert(info.Files, check.HasLen, 6)

	c.Assert(info.Files[0].Path, check.Equals, "test")
	c.Assert(info.Files[0].Folder, check.Equals, "")
	c.Assert(info.Files[0].Size, check.Equals, 0)
	c.Assert(info.Files[0].PackedSize, check.Equals, 0)
	c.Assert(info.Files[0].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[0].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Attributes, check.Equals, "D_ drwxr-xr-x")
	c.Assert(info.Files[0].CRC, check.Equals, 0)
	c.Assert(info.Files[0].Encrypted, check.Equals, false)
	c.Assert(info.Files[0].Method, check.DeepEquals, []string{""})
	c.Assert(info.Files[0].Block, check.Equals, 0)
	c.Assert(info.Files[0].Comment, check.Equals, "")
	c.Assert(info.Files[0].HostOS, check.Equals, "")
	c.Assert(info.Files[0].Version, check.Equals, 0)

	c.Assert(info.Files[1].Path, check.Equals, "test/dir1")
	c.Assert(info.Files[1].Folder, check.Equals, "")
	c.Assert(info.Files[1].Size, check.Equals, 0)
	c.Assert(info.Files[1].PackedSize, check.Equals, 0)
	c.Assert(info.Files[1].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[1].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[1].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[1].Attributes, check.Equals, "D_ drwxr-xr-x")
	c.Assert(info.Files[1].CRC, check.Equals, 0)
	c.Assert(info.Files[1].Encrypted, check.Equals, false)
	c.Assert(info.Files[1].Method, check.DeepEquals, []string{""})
	c.Assert(info.Files[1].Block, check.Equals, 0)
	c.Assert(info.Files[1].Comment, check.Equals, "")
	c.Assert(info.Files[1].HostOS, check.Equals, "")
	c.Assert(info.Files[1].Version, check.Equals, 0)

	c.Assert(info.Files[2].Path, check.Equals, "test/file2.log")
	c.Assert(info.Files[2].Folder, check.Equals, "")
	c.Assert(info.Files[2].Size, check.Equals, 0)
	c.Assert(info.Files[2].PackedSize, check.Equals, 0)
	c.Assert(info.Files[2].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[2].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[2].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[2].Attributes, check.Equals, "A_ -rw-r--r--")
	c.Assert(info.Files[2].CRC, check.Equals, 0)
	c.Assert(info.Files[2].Encrypted, check.Equals, false)
	c.Assert(info.Files[2].Method, check.DeepEquals, []string{""})
	c.Assert(info.Files[2].Block, check.Equals, 0)
	c.Assert(info.Files[2].Comment, check.Equals, "")
	c.Assert(info.Files[2].HostOS, check.Equals, "")
	c.Assert(info.Files[2].Version, check.Equals, 0)

	c.Assert(info.Files[3].Path, check.Equals, "test/dir1/file1.log")
	c.Assert(info.Files[3].Folder, check.Equals, "")
	c.Assert(info.Files[3].Size, check.Equals, 11)
	c.Assert(info.Files[3].PackedSize, check.Equals, 26)
	c.Assert(info.Files[3].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[3].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[3].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[3].Attributes, check.Equals, "A_ -rw-r--r--")
	c.Assert(info.Files[3].CRC, check.Equals, 2073076399)
	c.Assert(info.Files[3].Encrypted, check.Equals, false)
	c.Assert(info.Files[3].Method, check.DeepEquals, []string{"LZMA2:12"})
	c.Assert(info.Files[3].Block, check.Equals, 0)
	c.Assert(info.Files[3].Comment, check.Equals, "")
	c.Assert(info.Files[3].HostOS, check.Equals, "")
	c.Assert(info.Files[3].Version, check.Equals, 0)

	c.Assert(info.Files[4].Path, check.Equals, "test/file1.log")
	c.Assert(info.Files[4].Folder, check.Equals, "")
	c.Assert(info.Files[4].Size, check.Equals, 6)
	c.Assert(info.Files[4].PackedSize, check.Equals, 0)
	c.Assert(info.Files[4].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[4].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[4].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[4].Attributes, check.Equals, "A_ -rw-r--r--")
	c.Assert(info.Files[4].CRC, check.Equals, 3157996776)
	c.Assert(info.Files[4].Encrypted, check.Equals, false)
	c.Assert(info.Files[4].Method, check.DeepEquals, []string{"LZMA2:12"})
	c.Assert(info.Files[4].Block, check.Equals, 0)
	c.Assert(info.Files[4].Comment, check.Equals, "")
	c.Assert(info.Files[4].HostOS, check.Equals, "")
	c.Assert(info.Files[4].Version, check.Equals, 0)

	c.Assert(info.Files[5].Path, check.Equals, "test/file3.log")
	c.Assert(info.Files[5].Folder, check.Equals, "")
	c.Assert(info.Files[5].Size, check.Equals, 5)
	c.Assert(info.Files[5].PackedSize, check.Equals, 0)
	c.Assert(info.Files[5].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[5].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[5].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[5].Attributes, check.Equals, "A_ -rw-r--r--")
	c.Assert(info.Files[5].CRC, check.Equals, 3168002993)
	c.Assert(info.Files[5].Encrypted, check.Equals, false)
	c.Assert(info.Files[5].Method, check.DeepEquals, []string{"LZMA2:12"})
	c.Assert(info.Files[5].Block, check.Equals, 0)
	c.Assert(info.Files[5].Comment, check.Equals, "")
	c.Assert(info.Files[5].HostOS, check.Equals, "")
	c.Assert(info.Files[5].Version, check.Equals, 0)
}

func (s *Z7Suite) TestListingZip(c *check.C) {
	info, err := List(Props{File: "testdata/test.zip"})

	c.Assert(err, check.IsNil)
	c.Assert(info, check.NotNil)

	c.Assert(info.Path, check.Equals, "testdata/test.zip")
	c.Assert(info.Type, check.Equals, TYPE_ZIP)
	c.Assert(info.Method, check.DeepEquals, []string{""})
	c.Assert(info.Solid, check.Equals, false)
	c.Assert(info.Blocks, check.Equals, 0)
	c.Assert(info.PhysicalSize, check.Equals, 0)
	c.Assert(info.HeadersSize, check.Equals, 0)

	c.Assert(info.Files, check.HasLen, 6)

	c.Assert(info.Files[0].Path, check.Equals, "test")
	c.Assert(info.Files[0].Folder, check.Equals, "+")
	c.Assert(info.Files[0].Size, check.Equals, 0)
	c.Assert(info.Files[0].PackedSize, check.Equals, 0)
	c.Assert(info.Files[0].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[0].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Attributes, check.Equals, "D_ drwxr-xr-x")
	c.Assert(info.Files[0].CRC, check.Equals, 0)
	c.Assert(info.Files[0].Encrypted, check.Equals, false)
	c.Assert(info.Files[0].Method, check.DeepEquals, []string{"Store"})
	c.Assert(info.Files[0].Block, check.Equals, 0)
	c.Assert(info.Files[0].Comment, check.Equals, "")
	c.Assert(info.Files[0].HostOS, check.Equals, "Unix")
	c.Assert(info.Files[0].Version, check.Equals, 10)

	c.Assert(info.Files[1].Path, check.Equals, "test/file1.log")
	c.Assert(info.Files[1].Folder, check.Equals, "-")
	c.Assert(info.Files[1].Size, check.Equals, 6)
	c.Assert(info.Files[1].PackedSize, check.Equals, 6)
	c.Assert(info.Files[1].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[1].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[1].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[1].Attributes, check.Equals, "_ -rw-r--r--")
	c.Assert(info.Files[1].CRC, check.Equals, 3157996776)
	c.Assert(info.Files[1].Encrypted, check.Equals, false)
	c.Assert(info.Files[1].Method, check.DeepEquals, []string{"Store"})
	c.Assert(info.Files[1].Block, check.Equals, 0)
	c.Assert(info.Files[1].Comment, check.Equals, "")
	c.Assert(info.Files[1].HostOS, check.Equals, "Unix")
	c.Assert(info.Files[1].Version, check.Equals, 10)

	c.Assert(info.Files[2].Path, check.Equals, "test/file2.log")
	c.Assert(info.Files[2].Folder, check.Equals, "-")
	c.Assert(info.Files[2].Size, check.Equals, 0)
	c.Assert(info.Files[2].PackedSize, check.Equals, 0)
	c.Assert(info.Files[2].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[2].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[2].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[2].Attributes, check.Equals, "_ -rw-r--r--")
	c.Assert(info.Files[2].CRC, check.Equals, 0)
	c.Assert(info.Files[2].Encrypted, check.Equals, false)
	c.Assert(info.Files[2].Method, check.DeepEquals, []string{"Store"})
	c.Assert(info.Files[2].Block, check.Equals, 0)
	c.Assert(info.Files[2].Comment, check.Equals, "")
	c.Assert(info.Files[2].HostOS, check.Equals, "Unix")
	c.Assert(info.Files[2].Version, check.Equals, 10)

	c.Assert(info.Files[3].Path, check.Equals, "test/dir1")
	c.Assert(info.Files[3].Folder, check.Equals, "+")
	c.Assert(info.Files[3].Size, check.Equals, 0)
	c.Assert(info.Files[3].PackedSize, check.Equals, 0)
	c.Assert(info.Files[3].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[3].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[3].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[3].Attributes, check.Equals, "D_ drwxr-xr-x")
	c.Assert(info.Files[3].CRC, check.Equals, 0)
	c.Assert(info.Files[3].Encrypted, check.Equals, false)
	c.Assert(info.Files[3].Method, check.DeepEquals, []string{"Store"})
	c.Assert(info.Files[3].Block, check.Equals, 0)
	c.Assert(info.Files[3].Comment, check.Equals, "")
	c.Assert(info.Files[3].HostOS, check.Equals, "Unix")
	c.Assert(info.Files[3].Version, check.Equals, 10)

	c.Assert(info.Files[4].Path, check.Equals, "test/dir1/file1.log")
	c.Assert(info.Files[4].Folder, check.Equals, "-")
	c.Assert(info.Files[4].Size, check.Equals, 11)
	c.Assert(info.Files[4].PackedSize, check.Equals, 11)
	c.Assert(info.Files[4].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[4].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[4].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[4].Attributes, check.Equals, "_ -rw-r--r--")
	c.Assert(info.Files[4].CRC, check.Equals, 2073076399)
	c.Assert(info.Files[4].Encrypted, check.Equals, false)
	c.Assert(info.Files[4].Method, check.DeepEquals, []string{"Store"})
	c.Assert(info.Files[4].Block, check.Equals, 0)
	c.Assert(info.Files[4].Comment, check.Equals, "")
	c.Assert(info.Files[4].HostOS, check.Equals, "Unix")
	c.Assert(info.Files[4].Version, check.Equals, 10)

	c.Assert(info.Files[5].Path, check.Equals, "test/file3.log")
	c.Assert(info.Files[5].Folder, check.Equals, "-")
	c.Assert(info.Files[5].Size, check.Equals, 5)
	c.Assert(info.Files[5].PackedSize, check.Equals, 5)
	c.Assert(info.Files[5].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[5].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[5].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[5].Attributes, check.Equals, "_ -rw-r--r--")
	c.Assert(info.Files[5].CRC, check.Equals, 3168002993)
	c.Assert(info.Files[5].Encrypted, check.Equals, false)
	c.Assert(info.Files[5].Method, check.DeepEquals, []string{"Store"})
	c.Assert(info.Files[5].Block, check.Equals, 0)
	c.Assert(info.Files[5].Comment, check.Equals, "")
	c.Assert(info.Files[5].HostOS, check.Equals, "Unix")
	c.Assert(info.Files[5].Version, check.Equals, 10)
}

func (s *Z7Suite) TestListingGz(c *check.C) {
	info, err := List(Props{File: "testdata/test.tar.gz"})

	c.Assert(err, check.IsNil)
	c.Assert(info, check.NotNil)

	c.Assert(info.Path, check.Equals, "testdata/test.tar.gz")
	c.Assert(info.Type, check.Equals, TYPE_GZIP)
	c.Assert(info.Method, check.DeepEquals, []string{""})
	c.Assert(info.Solid, check.Equals, false)
	c.Assert(info.Blocks, check.Equals, 0)
	c.Assert(info.PhysicalSize, check.Equals, 0)
	c.Assert(info.HeadersSize, check.Equals, 0)

	c.Assert(info.Files, check.HasLen, 1)

	c.Assert(info.Files[0].Path, check.Equals, "test.tar")
	c.Assert(info.Files[0].Folder, check.Equals, "")
	c.Assert(info.Files[0].Size, check.Equals, 10240)
	c.Assert(info.Files[0].PackedSize, check.Equals, 253)
	c.Assert(info.Files[0].Modified.Unix(), check.Not(check.Equals), 0)
	c.Assert(info.Files[0].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Attributes, check.Equals, "")
	c.Assert(info.Files[0].CRC, check.Equals, 285968217)
	c.Assert(info.Files[0].Encrypted, check.Equals, false)
	c.Assert(info.Files[0].Method, check.DeepEquals, []string{""})
	c.Assert(info.Files[0].Block, check.Equals, 0)
	c.Assert(info.Files[0].Comment, check.Equals, "")
	c.Assert(info.Files[0].HostOS, check.Equals, "Unix")
	c.Assert(info.Files[0].Version, check.Equals, 0)
}

func (s *Z7Suite) TestListingBz(c *check.C) {
	info, err := List(Props{File: "testdata/test.tar.bz2"})

	c.Assert(err, check.IsNil)
	c.Assert(info, check.NotNil)

	c.Assert(info.Path, check.Equals, "testdata/test.tar.bz2")
	c.Assert(info.Type, check.Equals, TYPE_BZIP)
	c.Assert(info.Method, check.DeepEquals, []string{""})
	c.Assert(info.Solid, check.Equals, false)
	c.Assert(info.Blocks, check.Equals, 0)
	c.Assert(info.PhysicalSize, check.Equals, 0)
	c.Assert(info.HeadersSize, check.Equals, 0)

	c.Assert(info.Files, check.HasLen, 1)

	c.Assert(info.Files[0].Path, check.Equals, "")
	c.Assert(info.Files[0].Folder, check.Equals, "")
	c.Assert(info.Files[0].Size, check.Equals, 0)
	c.Assert(info.Files[0].PackedSize, check.Equals, 0)
	c.Assert(info.Files[0].Modified.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Attributes, check.Equals, "")
	c.Assert(info.Files[0].CRC, check.Equals, 0)
	c.Assert(info.Files[0].Encrypted, check.Equals, false)
	c.Assert(info.Files[0].Method, check.DeepEquals, []string{""})
	c.Assert(info.Files[0].Block, check.Equals, 0)
	c.Assert(info.Files[0].Comment, check.Equals, "")
	c.Assert(info.Files[0].HostOS, check.Equals, "")
	c.Assert(info.Files[0].Version, check.Equals, 0)
}

func (s *Z7Suite) TestListingXz(c *check.C) {
	info, err := List(Props{File: "testdata/test.tar.xz"})

	c.Assert(err, check.IsNil)
	c.Assert(info, check.NotNil)

	c.Assert(info.Path, check.Equals, "testdata/test.tar.xz")
	c.Assert(info.Type, check.Equals, TYPE_XZ)
	c.Assert(info.Method, check.DeepEquals, []string{"LZMA2:26", "CRC64"})
	c.Assert(info.Solid, check.Equals, false)
	c.Assert(info.Blocks, check.Equals, 0)
	c.Assert(info.PhysicalSize, check.Equals, 0)
	c.Assert(info.HeadersSize, check.Equals, 0)

	c.Assert(info.Files, check.HasLen, 1)

	c.Assert(info.Files[0].Path, check.Equals, "")
	c.Assert(info.Files[0].Folder, check.Equals, "")
	c.Assert(info.Files[0].Size, check.Equals, 10240)
	c.Assert(info.Files[0].PackedSize, check.Equals, 272)
	c.Assert(info.Files[0].Modified.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Created.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Accessed.Unix() < 0, check.Equals, true)
	c.Assert(info.Files[0].Attributes, check.Equals, "")
	c.Assert(info.Files[0].CRC, check.Equals, 0)
	c.Assert(info.Files[0].Encrypted, check.Equals, false)
	c.Assert(info.Files[0].Method, check.DeepEquals, []string{"LZMA2:26", "CRC64"})
	c.Assert(info.Files[0].Block, check.Equals, 0)
	c.Assert(info.Files[0].Comment, check.Equals, "")
	c.Assert(info.Files[0].HostOS, check.Equals, "")
	c.Assert(info.Files[0].Version, check.Equals, 0)
}

func (s *Z7Suite) TestCheck(c *check.C) {
	ok, err := Check(Props{File: "testdata/test-max-compression.7z"})

	c.Assert(ok, check.Equals, true)
	c.Assert(err, check.IsNil)

	ok, err = Check(Props{File: "testdata/test-broken.7z"})

	c.Assert(ok, check.Equals, false)
	c.Assert(err, check.NotNil)

	ok, err = Check(Props{File: "testdata/test-not-exist.7z"})

	c.Assert(ok, check.Equals, false)
	c.Assert(err, check.NotNil)
}

func (s *Z7Suite) TestAdd(c *check.C) {
	resultFile := c.MkDir() + "/test.7z"

	_, err := Add(Props{File: resultFile})

	c.Assert(err, check.NotNil)

	_, err = Add(Props{File: resultFile}, "testdata/test")

	c.Assert(err, check.IsNil)

	c.Assert(fsutil.IsExist(resultFile), check.Equals, true)
	c.Assert(fsutil.IsReadable(resultFile), check.Equals, true)
	c.Assert(fsutil.IsEmpty(resultFile), check.Equals, false)

	ok, err := Check(Props{File: resultFile})

	c.Assert(ok, check.Equals, true)
	c.Assert(err, check.IsNil)

	resultFile = c.MkDir() + "/test1.7z"

	_, err = Add(Props{File: resultFile}, "testdata/test")

	c.Assert(err, check.IsNil)

	c.Assert(fsutil.IsExist(resultFile), check.Equals, true)
	c.Assert(fsutil.IsReadable(resultFile), check.Equals, true)
	c.Assert(fsutil.IsEmpty(resultFile), check.Equals, false)

	ok, err = Check(Props{File: resultFile})

	c.Assert(ok, check.Equals, true)
	c.Assert(err, check.IsNil)

	resultFile = c.MkDir() + "/test2.7z"

	_, err = Add(Props{Dir: "testdata/test", File: resultFile}, "file1.log")

	c.Assert(err, check.IsNil)
}

func (s *Z7Suite) TestExtract(c *check.C) {
	outputDir := c.MkDir()

	_, err := Extract(
		Props{
			File:      "testdata/test-max-compression.7z",
			OutputDir: outputDir,
		},
	)

	c.Assert(err, check.IsNil)

	c.Assert(fsutil.CheckPerms("DR", outputDir+"/test"), check.Equals, true)
	c.Assert(fsutil.CheckPerms("DR", outputDir+"/test/dir1"), check.Equals, true)
	c.Assert(fsutil.CheckPerms("FRS", outputDir+"/test/dir1/file1.log"), check.Equals, true)
	c.Assert(fsutil.CheckPerms("FRS", outputDir+"/test/file1.log"), check.Equals, true)
	c.Assert(fsutil.CheckPerms("FRS", outputDir+"/test/file2.log"), check.Equals, false)
	c.Assert(fsutil.CheckPerms("FRS", outputDir+"/test/file3.log"), check.Equals, true)
}

func (s *Z7Suite) TestDelete(c *check.C) {
	testArchive := c.MkDir() + "/test.7z"

	fsutil.CopyFile("testdata/test-max-compression.7z", testArchive)

	_, err := Delete(Props{File: testArchive}, "test/file2.log", "test/file3.log")

	c.Assert(err, check.IsNil)

	outputDir := c.MkDir()

	_, err = Extract(
		Props{
			File:      testArchive,
			OutputDir: outputDir,
		},
	)

	c.Assert(err, check.IsNil)

	c.Assert(fsutil.CheckPerms("DR", outputDir+"/test"), check.Equals, true)
	c.Assert(fsutil.CheckPerms("DR", outputDir+"/test/dir1"), check.Equals, true)
	c.Assert(fsutil.CheckPerms("FR", outputDir+"/test/dir1/file1.log"), check.Equals, true)
	c.Assert(fsutil.CheckPerms("FR", outputDir+"/test/file1.log"), check.Equals, true)
	c.Assert(fsutil.CheckPerms("FR", outputDir+"/test/file2.log"), check.Equals, false)
	c.Assert(fsutil.CheckPerms("FR", outputDir+"/test/file3.log"), check.Equals, false)

	_, err = Delete(Props{File: testArchive})

	c.Assert(err, check.NotNil)
}

func (s *Z7Suite) TestValidationErrors(c *check.C) {
	p := Props{File: "unknown.7z"}

	_, err := List(p)
	c.Assert(err, check.NotNil)

	_, err = Check(p)
	c.Assert(err, check.NotNil)

	_, err = Extract(p)
	c.Assert(err, check.NotNil)

	_, err = Delete(p, "")
	c.Assert(err, check.NotNil)

	c.Assert(Props{IncludeFile: "unknown"}.Validate(false), check.NotNil)
	c.Assert(Props{ExcludeFile: "unknown"}.Validate(false), check.NotNil)
	c.Assert(Props{OutputDir: "unknown"}.Validate(false), check.NotNil)
}
