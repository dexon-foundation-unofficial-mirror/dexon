package common

import (
	"io"

	"golang.org/x/crypto/sha3"

	"github.com/dexon-foundation/dexon/common"
	"github.com/dexon-foundation/dexon/core/vm/sqlvm/schema"
	"github.com/dexon-foundation/dexon/rlp"
)

// Index related operations of Storage.

// indexBase contain the base internal paramters/methods for all index objects.
type indexBase struct {
	storage  *Storage
	contract common.Address
	pathHash common.Hash
}

// ListEntryLoad load an entry in the list from storage.
// Entry size should <= 32 and entries are packed into a slot if possible.
func (ib *indexBase) ListEntryLoad(
	headerSizeInSlot uint64,
	entrySizeInByte uint64,
	i uint64,
) []byte {
	entryNumInSlot := common.HashLength / entrySizeInByte
	slotShift := headerSizeInSlot + i/entryNumInSlot
	slot := ib.storage.ShiftHashUint64(ib.pathHash, slotShift)
	start := entrySizeInByte * (i % entryNumInSlot)
	end := start + entrySizeInByte
	data := ib.storage.GetState(ib.contract, slot)
	return data[start:end]
}

// ListEntryStore store an entry of the list into storage.
// if len(data) exceeds entrySize, it will be cropped from the right (tail).
func (ib *indexBase) ListEntryStore(
	headerSizeInSlot uint64,
	entrySizeInByte uint64,
	i uint64,
	data []byte) {
	entryNumInSlot := common.HashLength / entrySizeInByte
	slotShift := headerSizeInSlot + i/entryNumInSlot
	slot := ib.storage.ShiftHashUint64(ib.pathHash, slotShift)
	start := entrySizeInByte * (i % entryNumInSlot)
	end := start + entrySizeInByte
	d := ib.storage.GetState(ib.contract, slot)
	copy(d[start:end], data)
	ib.storage.SetState(ib.contract, slot, d)
}

// IndexValues contain addresses to all possible values of an index.
type IndexValues struct {
	indexBase

	// Header.
	Length uint64
	// 3 unused uint64 fields here.
	// Contents.
	ValueHashes []common.Hash
}

// LoadIndexValues load IndexValues struct of a given index.
func (s *Storage) LoadIndexValues(
	contract common.Address,
	tableRef schema.TableRef,
	indexRef schema.IndexRef,
	onlyHeader bool,
) *IndexValues {
	ret := &IndexValues{}
	ret.storage = s
	ret.contract = contract
	ret.pathHash = s.GetIndexValuesPathHash(tableRef, indexRef)
	ret.LoadHeader()
	if !onlyHeader {
		// Load all ValueHashes.
		ret.LoadValues()
	}
	return ret
}

// LoadHeader load header data from Storage.
func (iv *IndexValues) LoadHeader() {
	data := iv.storage.GetState(iv.contract, iv.pathHash)
	iv.Length = bytesToUint64(data[:8])
}

// StoreHeader store header data into Storage.
func (iv *IndexValues) StoreHeader() {
	var header common.Hash
	copy(header[:8], uint64ToBytes(iv.Length))
	iv.storage.SetState(iv.contract, iv.pathHash, header)
}

// LoadValues load all value addresses in the list.
// This function assumes header is already loaded.
func (iv *IndexValues) LoadValues() {
	// NOTICE: empty slice and nil slice have different meaning.
	iv.ValueHashes = make([]common.Hash, iv.Length)
	slot := iv.pathHash
	for i := uint64(0); i < iv.Length; i++ {
		slot = iv.storage.ShiftHashUint64(slot, 1)
		iv.ValueHashes[i] = iv.storage.GetState(iv.contract, slot)
	}
}

// Load the i-th PV address from Storage. (i starts from 0)
func (iv *IndexValues) Load(i uint64) common.Hash {
	if iv.ValueHashes != nil {
		// All values have been loaded, return from cache.
		return iv.ValueHashes[i]
	}
	slot := iv.storage.ShiftHashUint64(iv.pathHash, i+1)
	return iv.storage.GetState(iv.contract, slot)
}

// Store the i-th PV address into Storage. (i starts from 0)
func (iv *IndexValues) Store(i uint64, addr common.Hash) {
	slot := iv.storage.ShiftHashUint64(iv.pathHash, i+1)
	iv.storage.SetState(iv.contract, slot, addr)
	// FIXME(yenlin): update iv.Length here if necessary?

	if iv.ValueHashes != nil {
		// Update cache if loaded.
		if size := uint64(len(iv.ValueHashes)); size <= i {
			iv.ValueHashes = append(iv.ValueHashes,
				make([]common.Hash, i-size+1)...)
		}
		iv.ValueHashes[i] = addr
	}
}

// IndexEntry contain row ids of a given value in an index.
type IndexEntry struct {
	indexBase

	// Header.
	Length              uint64
	IndexToValuesOffset uint64
	ForeignKeyRefCount  uint64
	// 1 unused uint64 field here.
	// Contents.
	RowIDs []uint64
}

// LoadIndexEntry load IndexEntry struct of a given value key on an index.
func (s *Storage) LoadIndexEntry(
	contract common.Address,
	tableRef schema.TableRef,
	indexRef schema.IndexRef,
	onlyHeader bool,
	values ...[]byte,
) *IndexEntry {
	ret := &IndexEntry{}
	ret.storage = s
	ret.contract = contract
	ret.pathHash = s.GetIndexEntryPathHash(tableRef, indexRef, values...)
	ret.LoadHeader()
	if !onlyHeader {
		ret.LoadRowIDs()
	}
	return ret
}

// LoadHeader load header data from Storage.
func (ie *IndexEntry) LoadHeader() {
	data := ie.storage.GetState(ie.contract, ie.pathHash)
	ie.Length = bytesToUint64(data[:8])
	ie.IndexToValuesOffset = bytesToUint64(data[8:16])
	ie.ForeignKeyRefCount = bytesToUint64(data[16:24])
}

// StoreHeader store header data into Storage.
func (ie *IndexEntry) StoreHeader() {
	var header common.Hash
	copy(header[:8], uint64ToBytes(ie.Length))
	copy(header[8:16], uint64ToBytes(ie.IndexToValuesOffset))
	copy(header[16:24], uint64ToBytes(ie.ForeignKeyRefCount))
	ie.storage.SetState(ie.contract, ie.pathHash, header)
}

// LoadRowIDs load all row ids in the list.
// This function assumes header is already loaded.
func (ie *IndexEntry) LoadRowIDs() {
	// NOTICE: empty slice and nil slice have different meaning.
	ie.RowIDs = make([]uint64, 0, ie.Length)
	remain := ie.Length
	slot := ie.pathHash
	// To reduce GetState call, parse by ourselves instead of ListEntryLoad.
	for remain > 0 {
		bound := remain
		if bound > 4 {
			bound = 4
		}
		slot = ie.storage.ShiftHashUint64(slot, 1)
		data := ie.storage.GetState(ie.contract, slot).Bytes()
		for i := uint64(0); i < bound; i++ {
			ie.RowIDs = append(ie.RowIDs, bytesToUint64(data[:8]))
			data = data[8:]
		}
		remain -= bound
	}
}

// Load the i-th row id from Storage. (i starts from 0)
func (ie *IndexEntry) Load(i uint64) uint64 {
	if ie.RowIDs != nil {
		// All ids have been loaded, return from cache.
		return ie.RowIDs[i]
	}
	data := ie.ListEntryLoad(1, 8, i)
	return bytesToUint64(data)
}

// Store the i-th row id into Storage. (i starts from 0)
func (ie *IndexEntry) Store(i uint64, id uint64) {
	ie.ListEntryStore(1, 8, i, uint64ToBytes(id))
	// FIXME(yenlin): update ie.Length here if necessary?

	if ie.RowIDs != nil {
		// Update cache if loaded.
		if size := uint64(len(ie.RowIDs)); size <= i {
			ie.RowIDs = append(ie.RowIDs, make([]uint64, i-size+1)...)
		}
		ie.RowIDs[i] = id
	}
}

// indexRowIDRev is the reverse index from a rowID to the offsets of IndexEntry
// row id lists.
type indexRowIDRev struct {
	indexBase
}

// newIndexRowIDRev construct a indexRowIDRev instance.
func (s *Storage) newIndexRowIDRev(
	contract common.Address,
	tableRef schema.TableRef,
	rowID uint64,
) *indexRowIDRev {
	ret := &indexRowIDRev{}
	ret.storage = s
	ret.contract = contract
	ret.pathHash = s.GetIndexRowIDRevPathHash(tableRef, rowID)
	return ret
}

func (ir *indexRowIDRev) Load(indexRef schema.IndexRef) uint64 {
	data := ir.ListEntryLoad(0, 8, uint64(indexRef))
	return bytesToUint64(data)
}

func (ir *indexRowIDRev) Store(indexRef schema.IndexRef, offset uint64) {
	ir.ListEntryStore(0, 8, uint64(indexRef), uint64ToBytes(offset))
}

func (s *Storage) calPVLoc(payload []byte) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	hw.Write(payload)
	// Length of common.Hash is 256bit,
	// so it can properly match the size of hw.Sum
	hw.Sum(h[:0])
	return
}

func (s *Storage) encodePV(
	tableRef schema.TableRef,
	indexRef schema.IndexRef,
	values ...[]byte,
) []byte {
	// FIXME(yenlin): discuss if the payload is ok.
	payload := make([][]byte, 0, len(values)+2)
	payload = append(payload, tableRefToBytes(tableRef))
	payload = append(payload, indexRefToBytes(indexRef))
	payload = append(payload, values...)
	ret, err := rlp.EncodeToBytes(payload)
	if err != nil {
		// TODO(yenlin): normalize error.
		panic(err)
	}
	return ret
}

func (s *Storage) decodePV(reader io.Reader) [][]byte {
	var ret [][]byte
	err := rlp.Decode(reader, &ret)
	if err != nil {
		// TODO(yenlin): normalize error.
		panic(err)
	}
	return ret[2:]
}

// LoadPV load rlp encoded possible value from Storage.
func (s *Storage) LoadPV(
	contract common.Address,
	loc common.Hash,
) [][]byte {
	reader := NewStorageReader(s, contract, loc)
	return s.decodePV(reader)
}

// StorePV store rlp encoded possible value into Storage and return its
// location.
func (s *Storage) StorePV(
	contract common.Address,
	tableRef schema.TableRef,
	indexRef schema.IndexRef,
	values ...[]byte,
) common.Hash {
	payload := s.encodePV(tableRef, indexRef, values...)
	// FIXME(yenlin): Is hash(rlp([table, index, values...]) collision-free?
	loc := s.calPVLoc(payload)
	writer := NewStorageWriter(s, contract, loc)
	_, err := writer.Write(payload)
	if err != nil {
		// TODO(yenlin): normalize error.
		panic(err)
	}
	return loc
}

// ErasePV erase one possible value record on Storage.
func (s *Storage) ErasePV(
	contract common.Address,
	tableRef schema.TableRef,
	indexRef schema.IndexRef,
	values ...[]byte) {
	// TODO(yenlin): is there some better way to estimate the length...?
	payload := s.encodePV(tableRef, indexRef, values...)
	loc := s.calPVLoc(payload)
	writer := NewStorageWriter(s, contract, loc)
	_, err := writer.Write(make([]byte, len(payload)))
	if err != nil {
		// TODO(yenlin): normalize error.
		panic(err)
	}
}

// InsertIndex insert one rowid with a possible value into an index.
func (s *Storage) InsertIndex(
	contract common.Address,
	tableRef schema.TableRef,
	indexRef schema.IndexRef,
	pk uint64,
	values ...[]byte,
) {
	entry := s.LoadIndexEntry(contract, tableRef, indexRef, true, values...)
	if entry.Length == 0 {
		// Append to IndexValues.
		loc := s.StorePV(contract, tableRef, indexRef, values...)
		indexValues := s.LoadIndexValues(contract,
			tableRef, indexRef, true)
		indexValues.Store(indexValues.Length, loc)
		indexValues.Length++
		entry.IndexToValuesOffset = indexValues.Length
		indexValues.StoreHeader()
	}
	entry.Store(entry.Length, pk)
	entry.Length++
	entry.StoreHeader()
	rowRev := s.newIndexRowIDRev(contract, tableRef, pk)
	rowRev.Store(indexRef, entry.Length)
}

// DeleteIndex delete one rowid from an index.
func (s *Storage) DeleteIndex(
	contract common.Address,
	tableRef schema.TableRef,
	indexRef schema.IndexRef,
	pk uint64,
	values ...[]byte,
) {
	rowRev := s.newIndexRowIDRev(contract, tableRef, pk)
	entry := s.LoadIndexEntry(contract, tableRef, indexRef, true, values...)
	idx := rowRev.Load(indexRef) - 1
	rowRev.Store(indexRef, 0)
	if idx != entry.Length-1 {
		lastPK := entry.Load(entry.Length - 1)
		entry.Store(idx, lastPK)
		lastRowRev := s.newIndexRowIDRev(contract, tableRef, pk)
		lastRowRev.Store(indexRef, idx+1)
	}
	entry.Length--
	entry.Store(entry.Length, 0)
	entry.StoreHeader()
	if entry.Length == 0 {
		idx := entry.IndexToValuesOffset - 1
		indexValues := s.LoadIndexValues(contract, tableRef, indexRef, true)
		if idx != indexValues.Length-1 {
			lastValueLoc := indexValues.Load(indexValues.Length - 1)
			lastValue := s.LoadPV(contract, lastValueLoc)
			lastEntry := s.LoadIndexEntry(contract, tableRef, indexRef, true,
				lastValue...)
			lastEntry.IndexToValuesOffset = idx + 1
			lastEntry.StoreHeader()
		}
		indexValues.Length--
		indexValues.Store(indexValues.Length, common.Hash{})
		s.ErasePV(contract, tableRef, indexRef, values...)
		entry.IndexToValuesOffset = 0
	}
	entry.StoreHeader()
}

// CheckIndex checks the validity of the index entry on the possible value.
func (s *Storage) CheckIndex(
	contract common.Address,
	tableRef schema.TableRef,
	indexRef schema.IndexRef,
	unique bool,
	values ...[]byte,
) bool {
	entry := s.LoadIndexEntry(contract, tableRef, indexRef, true, values...)
	if unique && entry.Length > 1 {
		return false
	}
	if entry.ForeignKeyRefCount > 0 && entry.Length == 0 {
		return false
	}
	return true
}
