module test

go 1.16

// The version doesn't matter here, as we're replacing it with the currently checked out code anyway.
require github.com/mzz2017/quic-go v0.21.0

replace github.com/mzz2017/quic-go => ../../
