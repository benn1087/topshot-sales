package main

import (
	"os" //write csv related
	"io"
	"log"
	//"encoding/csv" //write csv related
	"strings"
    "context"
   "fmt"
    "time"
   
  "github.com/benn1087/topshot-sales/flowhelp"
   
  	"github.com/onflow/cadence"
   "github.com/onflow/flow-go-sdk/client"
   //"google.golang.org/grpc"
)
/*
func handleErr(err error) {
   if err != nil {
      panic(err)
   }
}
*/
//function write to csv
/*
func writeCsv(data []string) {
    file, err := os.Create("result.csv")
    handleErr(err)
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range data {
        err := writer.Write(value)
        handleErr(err)
    }
}
*/


func get_all(startheight, latestheight uint64) (SetcreateMoments []cadence.Event, blockNumbers []uint64, timestamps []time.Time, err error) {
	numQueriedBlocks :=uint64(248)
	flowClient, err := flowhelper.GetFlowClient(startheight)
	//flowClient, err := client.New("access.mainnet.nodes.onflow.org:9000", grpc.WithInsecure())
	for startheight < latestheight {
		m, b, t, err :=get_momentsetcreate(startheight, startheight+numQueriedBlocks,flowClient)
		if err != nil {
			if strings.Contains(err.Error(), "ResourceExhausted") {
				//log.Warn("resource exhausted, decrease number of queried blocks.")
				numQueriedBlocks /= 2
				continue
			}
			//log.Error("getting setcreate moments: ", err)
		}
		SetcreateMoments = append(SetcreateMoments, m...)
		blockNumbers = append(blockNumbers, b...)
		timestamps = append(timestamps, t...)
		// Increase start height ad reset numQueriedBlocks to default
		fmt.Println(startheight, startheight+numQueriedBlocks, numQueriedBlocks)
		startheight += numQueriedBlocks
		numQueriedBlocks = 248
	}
	fmt.Println(blockNumbers)
	return //setcreateMoments, blockNumbers, timestamps, err
}


func get_momentsetcreate(startheight, endheight uint64, flowClient *client.Client) (SetcreateMoments []cadence.Event, blockNumbers []uint64, timestamps []time.Time, err error){
	blockEvents, err := flowClient.GetEventsForHeightRange(context.Background(), client.EventRangeQuery{
		Type:        "A.0b2a3299cc857e29.TopShot.SetCreated",  //Setcreateed",
		StartHeight: startheight,
		EndHeight:   endheight, 
	})
	if err != nil {
		return
	}
	for _, blockEvent := range blockEvents {
		if len(blockEvent.Events) != 0 {
			for _, momentsetcreatedEvent := range blockEvent.Events {
			  SetcreateMoments = append(SetcreateMoments, momentsetcreatedEvent.Value)
			  blockNumbers = append(blockNumbers, blockEvent.Height)
			  timestamps = append(timestamps, blockEvent.BlockTimestamp)
			}
		}	
	}
	return 
}


func main() {
	/*
	var x16,y16,z16,w16 = get_all(23830813, 26554382)

	fmt.Println(w16)

	file16, err := os.Create("myfile16")
	if err != nil {
		log.Fatal(err)
	}

	mw16 := io.MultiWriter(os.Stdout, file16)

	for i := range x16 {
		fmt.Fprintln(mw16, x16[i].Fields[0], ",", x16[i].Fields[1], ",", y16[i],",", z16[i])
	}
	
	

    var x15,y15,z15,w15 = get_all(21291692, 23830812)

	fmt.Println(w15)

	file15, err := os.Create("myfile15")
	if err != nil {
		log.Fatal(err)
	}

	mw15 := io.MultiWriter(os.Stdout, file15)
	for i := range x15 {
		fmt.Fprintln(mw15, x15[i].Fields[0], ",", x15[i].Fields[1], ",",  y15[i],",", z15[i])
	}
	*/

	/*
	var x14,y14,z14,w14 = get_all(19251817, 20812705) //(19050753, 21291691)

	fmt.Println(w14)

	file14, err := os.Create("myfile14")
	if err != nil {
		log.Fatal(err)
	}

	mw14 := io.MultiWriter(os.Stdout, file14)
	for i := range x14 {
		fmt.Fprintln(mw14, x14[i].Fields[0], ",", x14[i].Fields[1], ",",  y14[i],",", z14[i])
	}
	

	var x13,y13,z13,w13 = get_all(18780501, 18780509) //(18587478, 19050752)
	
	
	fmt.Println(w13)

	file13, err := os.Create("myfile13")
	if err != nil {
		log.Fatal(err)
	}

	mw13 := io.MultiWriter(os.Stdout, file13)
	for i := range x13 {
		fmt.Fprintln(mw13, x13[i].Fields[0], ",", x13[i].Fields[1], ",",  y13[i],",", z13[i])
	}
	
	

	var x12,y12,z12,w12 = get_all(17805776, 17805789) //(17544523, 18587477)
	fmt.Println(w12)

	file12, err := os.Create("myfile12")
	if err != nil {
		log.Fatal(err)
	}

	mw12 := io.MultiWriter(os.Stdout, file12)
	for i := range x12 {
		fmt.Fprintln(mw12, x12[i].Fields[0], ",", x12[i].Fields[1], ",",  y12[i],",", z12[i])
	}

	var x11,y11,z11,w11 = get_all(16966860, 17300785) //(16755602, 17544522)
	fmt.Println(w11)

	file11, err := os.Create("myfile11")
	if err != nil {
		log.Fatal(err)
	}

	mw11 := io.MultiWriter(os.Stdout, file11)
	for i := range x11 {
		fmt.Fprintln(mw11, x11[i].Fields[0], ",", x11[i].Fields[1], ",",  y11[i],",", z11[i])
	}
	
	var x10,y10,z10,w10 = get_all(16459660, 16459671) //(15791891, 16755601)
	fmt.Println(w10)

	file10, err := os.Create("myfile10")
	if err != nil {
		log.Fatal(err)
	}

	mw10 := io.MultiWriter(os.Stdout, file10)
	for i := range x10 {
		fmt.Fprintln(mw10, x10[i].Fields[0], ",", x10[i].Fields[1], ",", y10[i],",", z10[i])
	}

	var x09,y09,z09,w09 = get_all(15405460, 15405465) //(14892104, 15791890)
	fmt.Println(w09)

	file09, err := os.Create("myfile09")
	if err != nil {
		log.Fatal(err)
	}

	mw09 := io.MultiWriter(os.Stdout, file09)
	for i := range x09 {
		fmt.Fprintln(mw09, x09[i].Fields[0], ",", x09[i].Fields[1], ",",  y09[i],",", z09[i])
	}

	var x08,y08,z08,w08 = get_all(13990621, 14185879)  //(13950742, 14892103)
	fmt.Println(w08)

	file08, err := os.Create("myfile08")
	if err != nil {
		log.Fatal(err)
	}

	mw08 := io.MultiWriter(os.Stdout, file08)
	for i := range x08 {
		fmt.Fprintln(mw08, x08[i].Fields[0], ",", x08[i].Fields[1], ",",  y08[i],",", z08[i])
	}

	
	var x07,y07,z07,w07 = get_all(13404174, 13950741)
	fmt.Println(w07)

	
	file07, err := os.Create("myfile07")
	if err != nil {
		log.Fatal(err)
	}

	mw07 := io.MultiWriter(os.Stdout, file07)
	fmt.Fprintln(mw07, x07, "end of slice 1", y07,"end of slice 2", z07)
	

	var x06,y06,z06,w06 = get_all(12796820, 12796828) //(12609237, 13404173)
	fmt.Println(w06)

	file06, err := os.Create("myfile06")
	if err != nil {
		log.Fatal(err)
	}

	mw06 := io.MultiWriter(os.Stdout, file06)
	for i := range x06 {
		fmt.Fprintln(mw06, x06[i].Fields[0], ",", x06[i].Fields[1], ",",  y06[i],",", z06[i])
	}

	var x05,y05,z05,w05 = get_all(12433006, 12463968)
	fmt.Println(w05)

	file05, err := os.Create("myfile05")
	if err != nil {
		log.Fatal(err)
	}

	mw05 := io.MultiWriter(os.Stdout, file05)
	for i := range x05 {
		fmt.Fprintln(mw05, x05[i].Fields[0], ",", x05[i].Fields[1], ",",  y05[i],",", z05[i])
	}
	*/
	var x04,y04,z04,w04 = get_all(11798590, 11798598) //(10446366, 10952608) //11798594//(9992020, 12020336)
	fmt.Println(w04)

	file04, err := os.Create("myfile04a")
	if err != nil {
		log.Fatal(err)
	}

	mw04 := io.MultiWriter(os.Stdout, file04)
	for i := range x04 {
		fmt.Fprintln(mw04, x04[i].Fields[0], ",", x04[i].Fields[1], ",",  y04[i],",", z04[i])
	}
	/*
	var x03,y03,z03,w03 = get_all(9737133, 9992019)
	fmt.Println(w03)

	file03, err := os.Create("myfile03")
	if err != nil {
		log.Fatal(err)
	}

	
	mw03 := io.MultiWriter(os.Stdout, file03)
	for i := range x03 {
		fmt.Fprintln(mw03, x03[i].Fields[0], ",", x03[i].Fields[1], ",",  y03[i],",", z03[i])
	}
	
	var x02,y02,z02,w02 = get_all(8973161, 9716725) //(8742959, 9737132)
	fmt.Println(w02)

	file02, err := os.Create("myfile02")
	if err != nil {
		log.Fatal(err)
	}

	mw02 := io.MultiWriter(os.Stdout, file02)
	for i := range x02 {
		fmt.Fprintln(mw02, x02[i].Fields[0], ",", x02[i].Fields[1], ",",   y02[i],",", z02[i])
	}

	var x01,y01,z01,w01 = get_all(7643982, 7951773)  //(7601063, 8742958)
	fmt.Println(w01)

	file01, err := os.Create("myfile01")
	if err != nil {
		log.Fatal(err)
	}

	mw01 := io.MultiWriter(os.Stdout, file01)
	for i := range x01 {
		fmt.Fprintln(mw01, x01[i].Fields[0], ",", x01[i].Fields[1], ",",  y01[i],",", z01[i])
	}
	*/
	
}