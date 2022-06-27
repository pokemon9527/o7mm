package model

import "time"

// import (
//   "sha/sha256"
//   "bytes"
//   "strconv"
// )

type Block struct {
	Timestamp     int64  // 当前时间戳，区块创建时间
	Data          []byte // 区块实际存储数据
	PrevBlockHash []byte // 前一个区块哈希值
	Hash          []byte // 当前区块哈希值
}

// SetHash
func (this *Block) SetHash() {
	// TODO: 补充算法
	this.Hash = []byte("xxx")
}

// SetHash_bitcoin
// func (b *Block)SetHash_bitcoin(){
//   timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)
//
// 	b.Hash = hash[:]
// }

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}
