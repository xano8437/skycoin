package coin

import (
    "github.com/skycoin/skycoin/src/lib/encoder"
)



type UxManager struct {
    UxMap map[SHA256]int
    UXArray []Ux
}

type (UxManager *self) AppendUx(ux Ux) {

    _, exists := self.UxMap[hash]
    if exists {
        log.Panic()
    }
    self.UXArray = append(self.UXArray, ux)
    //TODO: check this element does not exist!
    UxMap[ux.Hash()] = ux 
}

type (UxManager *self) RemoveUx(hash SHA256) {
    //TODO: check element exists
    idx, exists := self.UxMap[hash]
    if !exists {
        log.Panic()
    }
    delete(self.UxMap, hash)
    append(self.UXArray[:i], self.UXArray[i+1:]...)
}

type (UxManager *self) GetUx(hash SHA256) Ux {
    idx, exists := self.UxMap[hash]
    if !exists {
        log.Panic()
    }
    return self.UxArray[idx]
}

/*
	Unspent Outputs
*/

//needs a nonce
//think through replay atacks

/*

- hash must only depend on factors known to sender
-- hash cannot depend on block executed
-- hash cannot depend on sequence number
-- hash may depend on nonce

- hash must depend only on factors known to sender
-- needed to minimize divergence during block chain forks
- it should be difficult to create outputs with duplicate ids

- Uxhash cannot depend on time or block it was created
- time is still needed for
*/

/*
	For each transaction, keep track of
	- order created
	- order spent (for rollbacks)
*/

type UxOut struct {
    Head UxHead
    Body UxBody //hashed part
    //Meta UxMeta
}

//not hashed, metdata
type UxHead struct {
    Time  uint64 //time of block it was created in
    BkSeq uint64 //block it was created in
    SpSeq uint64 //block it was spent in
}

//part that is hashed
type UxBody struct {
    SrcTransaction SHA256
    Address        Address //address of receiver
    Coins          uint64  //number of coins
    Hours          uint64  //coin hours
}

//type UxMeta struct {
//}

func (self UxOut) Hash() SHA256 {
    b1 := encoder.Serialize(self.Body)
    return SumSHA256(b1)
}

/*
func (self UxOut) HashTotal() *SHA256 {
	b1 := encoder.Serialize(self.Head)
	b2 := encoder.Serialize(self.Body)
	b3 = append(b1, b2...)
	return SumSHA256(b3)
}
*/

/*
	Make indepedent of block rate?
	Then need creation time of output
	Creation time of transaction cant be hashed
*/

//calculate coinhour balance of output
func (self *UxOut) CoinHours(t uint64) uint64 {
    if t < self.Head.Time {
        return 0
    }

    v1 := self.Body.Hours             //starting coinshour
    ch := (t - self.Head.Time) / 3600 //number of hours, one hour every 240 block
    v2 := ch * self.Body.Coins        //accumulated coin-hours
    return v1 + v2                    //starting+earned
}
