package main

import (
   "context"
   "fmt"
   
   "github.com/benn1087/topshot-sales/topshot"

   "github.com/onflow/flow-go-sdk/client"
   "google.golang.org/grpc"
)
func handleErr(err error) {
   if err != nil {
      panic(err)
   }
}

func main() {
    flowClient, err := client.New("access.mainnet.nodes.onflow.org:9000", grpc.WithInsecure())
    handleErr(err)
    err = flowClient.Ping(context.Background())
    handleErr(err)
	latestBlock, err := flowClient.GetLatestBlock(context.Background(), false)
	handleErr(err)
	fmt.Println("current height: ", latestBlock.Height)

	blockEvents, err := flowClient.GetEventsForHeightRange(context.Background(), client.EventRangeQuery{
		Type:        "A.c1e4f4f4c4257510.TopShotMarketV3.MomentPurchased",
		StartHeight: latestBlock.Height-30,
		EndHeight:   latestBlock.Height,
	 })
	 handleErr(err)

	 for _, blockEvent := range blockEvents {
		for _, purchaseEvent := range blockEvent.Events {
		  fmt.Println(purchaseEvent)
		  fmt.Println(purchaseEvent.Value)
		}
	 }


	 

	for _, blockEvent := range blockEvents {
   		if len(blockEvent.Events) != 0 {
      		for _, purchaseEvent := range blockEvent.Events {
         		e := topshot.MomentPurchasedEvent(purchaseEvent.Value)
         		fmt.Println(e)
				 saleMoment, err := topshot.GetSaleMomentFromOwnerAtBlock(flowClient, blockEvent.Height-1, *e.Seller(), e.Id())
				 handleErr(err)
				 fmt.Println(saleMoment)
				 fmt.Printf("transactionID: %s, block height: %d\n",
					 purchaseEvent.TransactionID.String(), blockEvent.Height)
				 fmt.Println()
      		}
   		}
	}
	 
}
