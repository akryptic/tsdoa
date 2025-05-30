package db

import (
	"tsdoa/internal/constants"

	"github.com/dgraph-io/badger/v4"
)

type DBMgr struct {
	DB *badger.DB
}

var (
	DataDBManager = &DBMgr{
		DB: nil,
	}
	IndexDBManager = &DBMgr{
		DB: nil,
	}
)

func Init(path string) error {
	// Data DB setup
	dataDBOps := badger.DefaultOptions(path + "/data").WithLogger(nil)
	dataDB, err := badger.Open(dataDBOps)
	if err != nil {
		return err
	}
	DataDBManager.DB = dataDB

	// Index DB setup
	indexDBOps := badger.DefaultOptions(path + "/index").WithLogger(nil)
	indexDB, err := badger.Open(indexDBOps)
	if err != nil {
		DataDBManager.DB.Close()
		return err
	}
	IndexDBManager.DB = indexDB

	return nil
}

func Close() {
	if DataDBManager.DB != nil {
		DataDBManager.DB.Close()
	}

	if IndexDBManager.DB != nil {
		IndexDBManager.DB.Close()
	}
}

func (dbm *DBMgr) ensureReady() error {
	if dbm.DB == nil {
		return constants.ErrDBNotReady
	}
	return nil
}

func (dbm *DBMgr) Set(key string, value []byte) error {

	if err := dbm.ensureReady(); err != nil {
		return err
	}

	return dbm.DB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
}

func (dbm *DBMgr) GetByKey(key string) ([]byte, error) {

	if err := dbm.ensureReady(); err != nil {
		return nil, err
	}

	var value []byte
	err := dbm.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			value = make([]byte, len(val))
			copy(value, val)
			return nil
		})
	})
	return value, err
}

func (dbm *DBMgr) GetPaginatedWithPrefix(prefix string, offset, limit int) ([][]byte, bool, error) {

	if err := dbm.ensureReady(); err != nil {
		return nil, false, err
	}

	var res [][]byte
	hasMore := false
	pref := []byte(prefix)

	err := dbm.DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true
		opts.Prefix = pref

		it := txn.NewIterator(opts)
		defer it.Close()

		count := 0
		index := 0

		for it.Seek(pref); it.ValidForPrefix(pref); it.Next() {

			if index < offset {
				index++
				continue
			}

			if limit != -1 && count >= limit {
				hasMore = true
				break
			}

			item := it.Item()

			err := item.Value(func(val []byte) error {
				valCopy := make([]byte, len(val))
				copy(valCopy, val)
				res = append(res, valCopy)
				return nil
			})

			if err != nil {
				return err
			}
			count++
		}

		return nil
	})

	if err != nil {
		return nil, false, err
	}

	return res, hasMore, nil
}

func (dbm *DBMgr) DeleteByKey(key string) error {

	if err := dbm.ensureReady(); err != nil {
		return err
	}

	return dbm.DB.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

func (dbm *DBMgr) DeleteAllWithPrefix(prefix string) error {

	if err := dbm.ensureReady(); err != nil {
		return err
	}

	pref := []byte(prefix)
	return dbm.DB.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = pref

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(pref); it.ValidForPrefix(pref); it.Next() {
			item := it.Item()
			if err := txn.Delete(item.Key()); err != nil {
				return err
			}
		}

		return nil
	})
}

func (dbm *DBMgr) WipeOut() error {

	if err := dbm.ensureReady(); err != nil {
		return err
	}

	return dbm.DB.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			if err := txn.Delete(item.Key()); err != nil {
				return err
			}
		}

		return nil
	})
}
