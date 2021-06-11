package main

import (
	"bufio"
	"cards/go_proto_code/addressbookpb"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	ab := readAddressBook()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Input one of the following options:")
		fmt.Println("  1) Insert/Update an address.")
		fmt.Println("  2) See all available addresses.")
		fmt.Println("  3) Remove all available addresses.")
		fmt.Println("  4) Remove an address.")
		fmt.Println("  5) Close the program.")
		switch option, _ := reader.ReadString('\n'); strings.Replace(option, "\n", "", -1) {
		case "1":
			fmt.Println("Insert bleh")
		case "2":
			abj := toJson(ab)
			fmt.Println(abj)
		case "3":
			ab = writeEmptyBook()
			fmt.Println("All addresses successfully removed from the address book.")
		case "4":
			fmt.Println("Please insert the id of the address you want to remove.")
			id, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading the id: %s", err)
			}
			id = strings.Replace(id, "\n", "", -1)
			addressIndex := getAddressIndexById(id, getIds(ab))
			if addressIndex != -1 {
				ab = removeAddress(ab, addressIndex)
				fmt.Println("Address successfully removed from the address book.")
			} else {
				fmt.Println("The address is not present in the address book.")
			}
		case "5":
			fmt.Println("Good bye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid input, please input another option.")
		}
	}
}

func readAddressBook() *addressbookpb.AddressBook {
	ab := new(addressbookpb.AddressBook)
	ba, err := ioutil.ReadFile("./code_for_protos/addressbook.bin")
	if err != nil {
		log.Fatalf("Error reading the addressbook.bin file: %s", err)
	}
	proto.Unmarshal(ba, ab)
	return ab
}

func getIds(ab *addressbookpb.AddressBook) []int32 {
	ids := []int32{}
	for _, p := range ab.GetPeople() {
		ids = append(ids, p.GetId())
	}
	return ids
}

func toJson(ab *addressbookpb.AddressBook) string {
	jsonM := jsonpb.Marshaler{}
	abj, err := jsonM.MarshalToString(ab)
	if err != nil {
		log.Fatalf("Error when marshaling the address book to JSON format: %s", err)
	}
	return abj
}

func getAddressData()

func buildAddress() *addressbookpb.Person {
	p := new(addressbookpb.Person)
	p.Id = asd
	p.Name = asd
	p.Email = asd
	p.LastUpdated = asd
	p.Phones = asd
	return p
}

func insertAddress(ab *addressbookpb.AddressBook, p *addressbookpb.Person) *addressbookpb.AddressBook {
	ab.People = append(ab.People, p)
	return ab
}

func writeBook(ab *addressbookpb.AddressBook) {
	bc, err := proto.Marshal(ab)
	if err != nil {
		log.Fatalf("Error while marshaling the proto message: %s", err)
	}
	ioutil.WriteFile("./code_for_protos/addressbook.bin", bc, 0644)
}

func writeEmptyBook() *addressbookpb.AddressBook {
	ab := new(addressbookpb.AddressBook)
	writeBook(ab)
	return ab
}

func removeAddress(ab *addressbookpb.AddressBook, index int32) *addressbookpb.AddressBook {
	ab.People = append(ab.People[:index], ab.People[index+1:]...)
	return ab
}

func getAddressIndexById(id string, ids []int32) int32 {
	idn, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		log.Fatalf("Error parsing the id: %s", err)
	}
	for index, iid := range ids {
		if iid == int32(idn) {
			return int32(index)
		}
	}
	return -1
}
