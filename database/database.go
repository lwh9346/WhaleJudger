package database

import (
	"log"

	"github.com/xujiajun/nutsdb"
)

//HasKey 返回一个k-v型bucket是否含有某个key
func HasKey(db *nutsdb.DB, bucket, key string) bool {
	var l int
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			e, err1 := tx.RangeScan(bucket, []byte(key), []byte(key))
			l = len(e)
			return err1
		}); err != nil {
		log.Fatal(err)
	}
	return l > 0
}

//SetValue 设置k-v型bucket的键值
func SetValue(db *nutsdb.DB, bucket, key string, value []byte, ttl uint32) {
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			return tx.Put(bucket, []byte(key), value, ttl)
		}); err != nil {
		log.Fatal(err)
	}
}