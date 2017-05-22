package bbucket

import (
//	"github.com/jinxmcg/quickbites"
//	"unsafe"
)

const (
	DEFAULT_BUCKET_SIZE = 248
	DEFAULT_NO_VALUES = 4000000 * DEFAULT_BUCKET_SIZE
)

type BucketIndex struct {	
	len int
	cap int
}

type bucket struct {
	bucketId int 
	content []byte //the bucket
	maxValue int //max bucket bits
	nextBucket *bucket
}

func New() *BucketIndex {
	return &BucketIndex{0,DEFAULT_NO_VALUES}
}

func (b *BucketIndex) Len() int {
	return b.len
}

func (b *BucketIndex) Cap() int {
	return b.cap
}
func (b *BucketIndex) SetBit() bool {
	return true
}
