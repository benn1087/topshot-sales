/*package topshot

import (
	"fmt"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
)
// pub event MomentPurchased(id: UInt64, price: UFix64, seller: Address?)
type MomentPurchasedEvent cadence.Event

func (evt MomentPurchasedEvent) Id() uint64 {
	return uint64(evt.Fields[0].(cadence.UInt64))
}

func (evt MomentPurchasedEvent) Price() float64 {
	return float64(evt.Fields[1].(cadence.UFix64).ToGoValue().(uint64))/1e8 // ufixed 64 have 8 digits of precision
}

func (evt MomentPurchasedEvent) Seller() *flow.Address {
	optionalAddress := (evt.Fields[2]).(cadence.Optional)
	if cadenceAddress, ok := optionalAddress.Value.(cadence.Address); ok {
		sellerAddress := flow.BytesToAddress(cadenceAddress.Bytes())
		return &sellerAddress
	}
	return nil
}

func (evt MomentPurchasedEvent) String() string {
	return fmt.Sprintf("moment purchased: momentid: %d, price: %f, seller: %s",
		evt.Id(), evt.Price(), evt.Seller())
}
*/


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
