package topshot

import (
	"fmt"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
)
// pub event MomentMinted(momentID: UInt64, playID: UInt32, setID: UInt32, serialNumber: UInt32)
type MomentMintedEvent cadence.Event

func (evt MomentMintedEvent) Id() uint64 {
	return uint64(evt.Fields[0].(cadence.UInt64))
}

func (evt MomentMintedEvent) PlayId() uint32 {
	return uint32(evt.Fields[1].(cadence.UInt32))
}

func (evt MomentMintedEvent) SetId() uint32 {
	return uint32(evt.Fields[2].(cadence.UInt32))
}

func (evt MomentMintedEvent) serialNumber() uint32 {
	return uint32(evt.Fields[3].(cadence.UInt32))
}


func (evt MomentMintedEvent) String() string {
	return fmt.Sprintf("moment mint4ed: momentid: %d, playid: %d, setid: %d, serialnum: %d",
		evt.Id(), evt.PlayId(), evt.SetID(), evt.serialNumber())
}
