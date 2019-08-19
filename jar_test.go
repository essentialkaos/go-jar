package jar

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type JARSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&JARSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *JARSuite) TestParsing(c *C) {
	m, err := ReadFile("testdata/wagon-file-3.3.2.jar")

	c.Assert(err, IsNil)
	c.Assert(m, NotNil)

	c.Assert(
		m["Implementation-URL"], Equals,
		"https://maven.apache.org/wagon/wagon-providers/wagon-file",
	)
}

func (s *JARSuite) TestParsingErrors(c *C) {
	_, err := ReadFile("testdata/broken.jar")
	c.Assert(err, NotNil)

	_, err = ReadFile("testdata/jsr250-api-1.0-broken-manifest.jar")
	c.Assert(err, NotNil)

	_, err = ReadFile("testdata/jsr250-api-1.0-no-manifest.jar")
	c.Assert(err, NotNil)
}
