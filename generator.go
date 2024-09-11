package snowflake_id

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var randomIdBits uint8 = 12
var nodeIdBits uint8 = 10
var timestampBits uint8 = 41
var nodeIdShift = randomIdBits
var timestampShift = nodeIdShift + nodeIdBits

type ID int64

type Generator struct {
	nodeId int64
}

func NewGenerator(nodeId int64) (*Generator, error) {
	if nodeId < 0 || nodeId >= (1<<nodeIdBits) {
		return nil, errors.New(fmt.Sprintf("Invalid node id: nodeId should limit in %d bits", nodeIdBits))
	}
	return &Generator{
		nodeId: nodeId,
	}, nil
}

func (g *Generator) NewId() ID {
	// timestamp should limit in 41 bits
	nowMilli := time.Now().UnixMilli() & ((1 << timestampBits) - 1)
	// random number should limit in 12 bits
	randomLimit := int64(1 << randomIdBits)
	rnd := rand.Int63n(randomLimit)
	return g.new(nowMilli, rnd)
}

func (g *Generator) new(nowTimestamp, randomNum int64) ID {
	var rs int64
	rs = nowTimestamp << timestampShift
	rs += g.nodeId << nodeIdShift
	rs += randomNum
	return ID(rs)
}

func (id ID) Int64() int64 {
	return int64(id)
}
func (id ID) Base2() string {
	return strconv.FormatInt(id.Int64(), 2)
}
