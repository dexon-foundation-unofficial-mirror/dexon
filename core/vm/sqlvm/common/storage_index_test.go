package common

import (
	"testing"

	"github.com/dexon-foundation/dexon/common"
	"github.com/dexon-foundation/dexon/core/state"
	"github.com/dexon-foundation/dexon/core/vm/sqlvm/schema"
	"github.com/dexon-foundation/dexon/ethdb"
	"github.com/stretchr/testify/suite"
)

type StorageIndexTestSuite struct {
	suite.Suite
	storage  *Storage
	contract common.Address
}

func (s *StorageIndexTestSuite) SetupTest() {
	db := ethdb.NewMemDatabase()
	state, _ := state.New(common.Hash{}, state.NewDatabase(db))
	s.storage = NewStorage(state)
	s.contract = common.BytesToAddress([]byte("5566"))
}

func (s *StorageIndexTestSuite) TestPV() {
	table := schema.TableRef(0)
	index := schema.IndexRef(0)

	testcases := [][][]byte{
		{[]byte("a"), []byte("b")},
		{[]byte("a"), []byte("b"), []byte("c")},
		{[]byte("hello"), []byte("world")},
	}
	for i := range testcases {
		loc := s.storage.StorePV(s.contract, table, index, testcases[i]...)
		s.Require().Equal(testcases[i], s.storage.LoadPV(s.contract, loc))
	}
}

func (s *StorageIndexTestSuite) TestIndexValues() {
	addrs := []common.Hash{
		common.Hash{1},
		common.Hash{2},
	}
	table := schema.TableRef(0)
	index := schema.IndexRef(0)

	indexValues := s.storage.LoadIndexValues(s.contract, table, index, true)
	var check *IndexValues

	// Test header store / load.
	s.Require().Equal(uint64(0), indexValues.Length)
	indexValues.Length = uint64(len(addrs))
	indexValues.StoreHeader()
	check = s.storage.LoadIndexValues(s.contract, table, index, true)
	s.Require().Equal(indexValues.Length, check.Length)

	// Test record store / load.
	for i := range addrs {
		indexValues.Store(uint64(i), addrs[i])
		s.Require().Equal(addrs[i], indexValues.Load(uint64(i)))
	}

	// Test LoadValues.
	s.Require().Nil(indexValues.ValueHashes)
	indexValues.LoadValues()
	s.Require().Equal(addrs, indexValues.ValueHashes)

	// Check all data again with LoadIndexValues.
	check = s.storage.LoadIndexValues(s.contract, table, index, false)
	s.Require().Equal(indexValues.Length, check.Length)
	s.Require().Equal(addrs, check.ValueHashes)
}

func (s *StorageIndexTestSuite) TestIndexEntry() {
	ids := []uint64{1, 2}
	table := schema.TableRef(0)
	index := schema.IndexRef(0)

	value := [][]byte{[]byte("a"), []byte("b")}

	entry := s.storage.LoadIndexEntry(s.contract, table, index, true, value...)
	var check *IndexEntry
	// Test header store / load.
	s.Require().Equal(uint64(0), entry.Length)
	entry.Length = uint64(len(ids))
	entry.IndexToValuesOffset = uint64(1)
	entry.ForeignKeyRefCount = uint64(2)
	entry.StoreHeader()
	check = s.storage.LoadIndexEntry(s.contract, table, index, true, value...)
	s.Require().Equal(entry.Length, check.Length)
	s.Require().Equal(entry.IndexToValuesOffset, check.IndexToValuesOffset)
	s.Require().Equal(entry.ForeignKeyRefCount, check.ForeignKeyRefCount)

	// Test record store / load.
	for i := range ids {
		entry.Store(uint64(i), ids[i])
		s.Require().Equal(ids[i], entry.Load(uint64(i)))
	}

	// Test LoadRowIDs.
	s.Require().Nil(entry.RowIDs)
	entry.LoadRowIDs()
	s.Require().Equal(ids, entry.RowIDs)

	// Check all data again with LoadIndexEntry.
	check = s.storage.LoadIndexEntry(s.contract, table, index, false, value...)
	s.Require().Equal(entry.Length, check.Length)
	s.Require().Equal(entry.IndexToValuesOffset, check.IndexToValuesOffset)
	s.Require().Equal(entry.ForeignKeyRefCount, check.ForeignKeyRefCount)
	s.Require().Equal(ids, check.RowIDs)
}

func (s *StorageIndexTestSuite) TestIndexRowIDRev() {
	offsets := []uint64{1, 2}
	table := schema.TableRef(0)
	rowid := uint64(0)

	rev := s.storage.newIndexRowIDRev(s.contract, table, rowid)
	for i := range offsets {
		rev.Store(schema.IndexRef(i), offsets[i])
	}
	for i := range offsets {
		s.Require().Equal(offsets[i], rev.Load(schema.IndexRef(i)))
	}
}

func TestStorageIndex(t *testing.T) {
	suite.Run(t, new(StorageIndexTestSuite))
}
