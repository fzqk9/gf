// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.


package gmap

import (
	"sync"
    "gitee.com/johng/gf/g/util/gconv"
    "time"
)

type IntStringMap struct {
	mu sync.RWMutex
	m  map[int]string
}

func NewIntStringMap() *IntStringMap {
	return &IntStringMap{
        m: make(map[int]string),
    }
}

// 给定回调函数对原始内容进行遍历，回调函数返回true表示继续遍历，否则停止遍历
func (this *IntStringMap) Iterator(f func (k int, v string) bool) {
    this.mu.RLock()
    for k, v := range this.m {
        if !f(k, v) {
            break
        }
    }
    this.mu.RUnlock()
}

// 哈希表克隆
func (this *IntStringMap) Clone() *map[int]string {
	m := make(map[int]string)
	this.mu.RLock()
	for k, v := range this.m {
		m[k] = v
	}
    this.mu.RUnlock()
	return &m
}

// 设置键值对
func (this *IntStringMap) Set(key int, val string) {
	this.mu.Lock()
	this.m[key] = val
	this.mu.Unlock()
}

// 批量设置键值对
func (this *IntStringMap) BatchSet(m map[int]string) {
	this.mu.Lock()
	for k, v := range m {
		this.m[k] = v
	}
	this.mu.Unlock()
}

// 获取键值
func (this *IntStringMap) Get(key int) string {
	this.mu.RLock()
	val, _ := this.m[key]
	this.mu.RUnlock()
	return val
}

func (this *IntStringMap) GetBool(key int) bool {
    return gconv.Bool(this.Get(key))
}

func (this *IntStringMap) GetInt(key int) int {
    return gconv.Int(this.Get(key))
}

func (this *IntStringMap) GetInt8(key int) int8 {
    return gconv.Int8(this.Get(key))
}

func (this *IntStringMap) GetInt16(key int) int16 {
    return gconv.Int16(this.Get(key))
}

func (this *IntStringMap) GetInt32(key int) int32 {
    return gconv.Int32(this.Get(key))
}

func (this *IntStringMap) GetInt64(key int) int64 {
    return gconv.Int64(this.Get(key))
}

func (this *IntStringMap) GetUint (key int) uint {
    return gconv.Uint(this.Get(key))
}

func (this *IntStringMap) GetUint8 (key int) uint8 {
    return gconv.Uint8(this.Get(key))
}

func (this *IntStringMap) GetUint16 (key int) uint16 {
    return gconv.Uint16(this.Get(key))
}

func (this *IntStringMap) GetUint32 (key int) uint32 {
    return gconv.Uint32(this.Get(key))
}

func (this *IntStringMap) GetUint64 (key int) uint64 {
    return gconv.Uint64(this.Get(key))
}

func (this *IntStringMap) GetFloat32 (key int) float32 {
    return gconv.Float32(this.Get(key))
}

func (this *IntStringMap) GetFloat64 (key int) float64 {
    return gconv.Float64(this.Get(key))
}

func (this *IntStringMap) GetString (key int) string {
    return gconv.String(this.Get(key))
}

func (this *IntStringMap) GetTime (key int, format...string) time.Time {
    return gconv.Time(this.Get(key), format...)
}

func (this *IntStringMap) GetTimeDuration (key int) time.Duration {
    return gconv.TimeDuration(this.Get(key))
}

// 获取键值，如果键值不存在则写入默认值
func (this *IntStringMap) GetWithDefault(key int, value string) string {
    this.mu.Lock()
    val, ok := this.m[key]
    if !ok {
        this.m[key] = value
        val         = value
    }
    this.mu.Unlock()
    return val
}

// 删除键值对
func (this *IntStringMap) Remove(key int) {
    this.mu.Lock()
    delete(this.m, key)
    this.mu.Unlock()
}

// 批量删除键值对
func (this *IntStringMap) BatchRemove(keys []int) {
    this.mu.Lock()
    for _, key := range keys {
        delete(this.m, key)
    }
    this.mu.Unlock()
}

// 返回对应的键值，并删除该键值
func (this *IntStringMap) GetAndRemove(key int) (string) {
    this.mu.Lock()
    val, exists := this.m[key]
    if exists {
        delete(this.m, key)
    }
    this.mu.Unlock()
    return val
}

// 返回键列表
func (this *IntStringMap) Keys() []int {
    this.mu.RLock()
    keys := make([]int, 0)
    for key, _ := range this.m {
        keys = append(keys, key)
    }
    this.mu.RUnlock()
    return keys
}

// 返回值列表(注意是随机排序)
func (this *IntStringMap) Values() []string {
    this.mu.RLock()
    vals := make([]string, 0)
    for _, val := range this.m {
        vals = append(vals, val)
    }
    this.mu.RUnlock()
    return vals
}

// 是否存在某个键
func (this *IntStringMap) Contains(key int) bool {
    this.mu.RLock()
    _, exists := this.m[key]
    this.mu.RUnlock()
    return exists
}

// 哈希表大小
func (this *IntStringMap) Size() int {
    this.mu.RLock()
    length := len(this.m)
    this.mu.RUnlock()
    return length
}

// 哈希表是否为空
func (this *IntStringMap) IsEmpty() bool {
    this.mu.RLock()
    empty := (len(this.m) == 0)
    this.mu.RUnlock()
    return empty
}

// 清空哈希表
func (this *IntStringMap) Clear() {
    this.mu.Lock()
    this.m = make(map[int]string)
    this.mu.Unlock()
}

// 使用自定义方法执行加锁修改操作
func (this *IntStringMap) LockFunc(f func(m map[int]string)) {
    this.mu.Lock()
    f(this.m)
    this.mu.Unlock()
}

// 使用自定义方法执行加锁读取操作
func (this *IntStringMap) RLockFunc(f func(m map[int]string)) {
    this.mu.RLock()
    f(this.m)
    this.mu.RUnlock()
}
