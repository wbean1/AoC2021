package day16

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var operations []string

func Run() {
	binaryString := Input("E20D4100AA9C0199CA6A3D9D6352294D47B3AC6A4335FBE3FDD251003873657600B46F8DC600AE80273CCD2D5028B6600AF802B2959524B727D8A8CC3CCEEF3497188C017A005466DAA6FDB3A96D5944C014C006865D5A7255D79926F5E69200A164C1A65E26C867DDE7D7E4794FE72F3100C0159A42952A7008A6A5C189BCD456442E4A0A46008580273ADB3AD1224E600ACD37E802200084C1083F1540010E8D105A371802D3B845A0090E4BD59DE0E52FFC659A5EBE99AC2B7004A3ECC7E58814492C4E2918023379DA96006EC0008545B84B1B00010F8E915E1E20087D3D0E577B1C9A4C93DD233E2ECF65265D800031D97C8ACCCDDE74A64BD4CC284E401444B05F802B3711695C65BCC010A004067D2E7C4208A803F23B139B9470D7333B71240050A20042236C6A834600C4568F5048801098B90B626B00155271573008A4C7A71662848821001093CB4A009C77874200FCE6E7391049EB509FE3E910421924D3006C40198BB11E2A8803B1AE2A4431007A15C6E8F26009E002A725A5292D294FED5500C7170038C00E602A8CC00D60259D008B140201DC00C401B05400E201608804D45003C00393600B94400970020C00F6002127128C0129CDC7B4F46C91A0084E7C6648DC000DC89D341B23B8D95C802D09453A0069263D8219DF680E339003032A6F30F126780002CC333005E8035400042635C578A8200DC198890AA46F394B29C4016A4960C70017D99D7E8AF309CC014FCFDFB0FE0DA490A6F9D490010567A3780549539ED49167BA47338FAAC1F3005255AEC01200043A3E46C84E200CC4E895114C011C0054A522592912C9C8FDE10005D8164026C70066C200C4618BD074401E8C90E23ACDFE5642700A6672D73F285644B237E8CCCCB77738A0801A3CFED364B823334C46303496C940")
	fmt.Printf("%s\n", binaryString)
	versionSum := sumVersionPackets(binaryString)
	fmt.Printf("part1: versionSum is: %d", versionSum)
	fmt.Println("part2: operations array is:")
	fmt.Println(operations)
}

func parseHeader(s string) (int64, int64) {
	versionNum, err := strconv.ParseInt(s[:3], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	typeID, err := strconv.ParseInt(s[3:6], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return versionNum, typeID
}

func sumVersionPackets(s string) int64 {

	typeMap := map[int64]string{
		0: "sum",
		1: "product",
		2: "min",
		3: "max",
		5: "greaterThan",
		6: "lessThan",
		7: "equalTo",
	}
	fmt.Printf("have string: %s\n", s)
	if !strings.Contains(s, "1") {
		return 0
	}
	if len(s) < 6 {
		return 0
	}
	versionNum, typeID := parseHeader(s)
	s = s[6:] //strip off versionNum, typeID

	if typeID != 4 {
		operations = append(operations, typeMap[typeID])
		lengthTypeID, err := strconv.ParseInt(s[:1], 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		s = s[1:] //strip off legthTypeID
		var lengthBits int
		if lengthTypeID == 0 {
			lengthBits = 15
			length, err := strconv.ParseInt(s[:lengthBits], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			s = s[lengthBits:]      // strip off lengthBits
			contained := s[:length] // contained packets??
			s = s[length:]          // strip off length
			fmt.Printf("Found TypeID: %d, VersionNum: %d, LengthType: 0, Length: %d\n", typeID, versionNum, length)
			operations = append(operations, "(")
			versionNum += sumVersionPackets(contained)
			operations = append(operations, ")")
			return versionNum + sumVersionPackets(s)
		} else {
			lengthBits = 11
			numPackets, err := strconv.ParseInt(s[:lengthBits], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			s = s[lengthBits:] // strip off lengthBits
			operations = append(operations, fmt.Sprintf("%d(", int(numPackets)))
			fmt.Printf("Found TypeID: %d, VersionNum: %d, LengthType: 1, NumPackets %d\n", typeID, versionNum, numPackets)
			return versionNum + sumVersionPackets(s)
		}
	} else {
		dataBlocks := 0
		data := []byte{}
		for {
			if string(s[0]) == "1" {
				data = append(data, s[1:5]...)
				s = s[5:] // strip off 5 bits and continue
				dataBlocks++
			} else {
				data = append(data, s[1:5]...)
				s = s[5:] // strip off 5 bits and stop!
				dataBlocks++
				break
			}
		}
		result, err := strconv.ParseInt(string(data), 2, 64)
		if err != nil {
			log.Fatal(err)
		}
		resultStr := strconv.Itoa(int(result))
		operations = append(operations, resultStr)

		fmt.Printf("Found TypeID: %d, VersionNum: %d, dataBlocks: %d, value: %d\n", typeID, versionNum, dataBlocks, result)
		return versionNum + sumVersionPackets(s)
	}
}

func Input(input string) string {
	myBytes, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	binaryString := ""
	for _, byte := range myBytes {
		binaryString += fmt.Sprintf("%08b", byte)
	}
	return binaryString
}
