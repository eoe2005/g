package gconfcenter

import (
	"errors"
	"sync"
)

type confCache struct {
	cache sync.Map
}

func (d *confCache) Set(key string, value interface{}, expireSeconds int) (err error) {
	d.cache.Store(key, value)
	return nil
}

//EntryCount 获取实体数量
func (d *confCache) EntryCount() (entryCount int64) {
	count := int64(0)
	d.cache.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

//Get 获取缓存
func (d *confCache) Get(key string) (value interface{}, err error) {
	v, ok := d.cache.Load(key)
	if !ok {
		return nil, errors.New("load default cache fail")
	}
	return v.([]byte), nil
}

//Range 遍历缓存
func (d *confCache) Range(f func(key, value interface{}) bool) {
	d.cache.Range(f)
}

//Del 删除缓存
func (d *confCache) Del(key string) (affected bool) {
	d.cache.Delete(key)
	return true
}

//Clear 清除所有缓存
func (d *confCache) Clear() {
	d.cache = sync.Map{}
}
