package main

import (
	"bufio"
	"cards/go_proto_code/addressbookpb"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	ab := readAddressBook()
	reader := bufio.NewReader(os.Stdin)
	for {
		printMenu()
		switch option, _ := reader.ReadString('\n'); strings.Replace(option, "\n", "", -1) {
		case "1":
			var insertedId int32
			insertedId, ab = addAddress(reader, ab)
			fmt.Printf("Address with id '%v' successfully inserted", insertedId)
		case "2":
			fmt.Println("Please insert the id of the address you want to update.")
			id := inputId(reader)
			addressIndex := getAddressIndexByStringId(id, getIds(ab))
			if addressIndex != -1 {
				ab = updateAddress(addressIndex, id, reader, ab)
				fmt.Println("Address successfully updated in the address book.")
			} else {
				fmt.Println("The address is not present in the address book.")
			}
		case "3":
			abj := toJson(ab)
			fmt.Println(abj)
		case "4":
			ab = writeEmptyBook()
			fmt.Println("All addresses successfully removed from the address book.")
		case "5":
			fmt.Println("Please insert the id of the address you want to remove.")
			id := inputId(reader)
			addressIndex := getAddressIndexByStringId(id, getIds(ab))
			if addressIndex != -1 {
				ab = removeAddress(ab, addressIndex)
				fmt.Println("Address successfully removed from the address book.")
			} else {
				fmt.Println("The address is not present in the address book.")
			}
		case "6":
			fmt.Println("Good bye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid input, please input another option.")
		}
	}
}

func printMenu() {
	fmt.Println("Input one of the following options:")
	fmt.Println("  1) Insert an address.")
	fmt.Println("  2) Update an address.")
	fmt.Println("  3) See all available addresses.")
	fmt.Println("  4) Remove all available addresses.")
	fmt.Println("  5) Remove an address.")
	fmt.Println("  6) Close the program.")
}

func inputId(reader *bufio.Reader) string {
	id, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading the id: %s", err)
	}
	return strings.Replace(id, "\n", "", -1)
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

func getMaxId(ab *addressbookpb.AddressBook) int32 {
	ids := getIds(ab)
	maxId := int32(0)
	for _, id := range ids {
		if id > maxId {
			maxId = id
		}
	}
	return maxId
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

func addAddress(reader *bufio.Reader, ab *addressbookpb.AddressBook) (int32, *addressbookpb.AddressBook) {
	p := getAddressData(reader, ab, 0)
	ab.People = append(ab.People, p)
	writeBook(ab)
	return p.GetId(), ab
}

func updateAddress(addressIndex int32, existentId string, reader *bufio.Reader, ab *addressbookpb.AddressBook) *addressbookpb.AddressBook {
	idn, err := strconv.ParseInt(existentId, 10, 32)
	if err != nil {
		log.Fatalf("Error parsing the id: %s", err)
	}
	p := getAddressData(reader, ab, int32(idn))
	ab.People = append(ab.People[:addressIndex], ab.People[addressIndex+1:]...)
	ab.People = append(ab.People, p)
	writeBook(ab)
	return ab
}

func getAddressData(reader *bufio.Reader, ab *addressbookpb.AddressBook, existentId int32) *addressbookpb.Person {
	name := readName(reader)
	email := readEmail(reader)
	id := getMaxId(ab) + 1
	if existentId != 0 {
		id = existentId
	}
	phones := readPhones(reader)
	return buildAddress(id, name, email, phones)
}

func readName(reader *bufio.Reader) string {
	for {
		fmt.Println("Please input the name of the person you want to insert.")
		name, _ := reader.ReadString('\n')
		name = strings.Replace(name, "\n", "", -1)
		if name != "" {
			return name
		} else {
			fmt.Println("Invalid name.")
		}
	}
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func readEmail(reader *bufio.Reader) string {
	for {
		fmt.Println("Please input the email of the person you want to insert.")
		email, _ := reader.ReadString('\n')
		email = strings.Replace(email, "\n", "", -1)
		if isEmailValid(email) {
			return email
		} else {
			fmt.Println("Invalid email.")
		}
	}
}

func readPhones(reader *bufio.Reader) []*addressbookpb.Person_PhoneNumber {
	phones := []*addressbookpb.Person_PhoneNumber{}
	for {
		phone := new(addressbookpb.Person_PhoneNumber)
		number := readPhoneNumber(reader)
		pt := readPhoneType(reader)
		phone.Number = number
		phone.Type = pt
		phones = append(phones, phone)
		if !shouldContinue(reader) {
			return phones
		}
	}
}

func shouldContinue(reader *bufio.Reader) bool {
	for {
		fmt.Println("Do you want to input a new number? (yes/no)")
		switch input, _ := reader.ReadString('\n'); strings.Replace(input, "\n", "", -1) {
		case "yes":
			return true
		case "no":
			return false
		default:
			fmt.Println("Invalid input.")
		}
	}
}

func isNumberValid(number string) bool {
	numberRegex := regexp.MustCompile(`^[0-9]{7}$`)
	return numberRegex.MatchString(number)
}

func readPhoneNumber(reader *bufio.Reader) string {
	for {
		fmt.Println("Please input the phone number of the person you want to insert.")
		number, _ := reader.ReadString('\n')
		number = strings.Replace(number, "\n", "", -1)
		if isNumberValid(number) {
			return number
		} else {
			fmt.Println("Invalid phone number.")
		}
	}
}

func printPhoneTypes() {
	fmt.Println("Please input the phone type of the person you want to insert.")
	fmt.Println("  0) MOBILE.")
	fmt.Println("  1) HOME.")
	fmt.Println("  2) WORK.")
}

func readPhoneType(reader *bufio.Reader) addressbookpb.Person_PhoneType {
	for {
		printPhoneTypes()
		switch pt, _ := reader.ReadString('\n'); strings.Replace(pt, "\n", "", -1) {
		case "0":
			return addressbookpb.Person_PhoneType(0)
		case "1":
			return addressbookpb.Person_PhoneType(1)
		case "2":
			return addressbookpb.Person_PhoneType(2)
		default:
			fmt.Println("Invalid phone type.")
		}
	}
}

func buildAddress(id int32, name string, email string, phones []*addressbookpb.Person_PhoneNumber) *addressbookpb.Person {
	p := new(addressbookpb.Person)
	p.Id = id
	p.Name = name
	p.Email = email
	p.LastUpdated = timestamppb.New(time.Now())
	p.Phones = phones
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
	writeBook(ab)
	return ab
}

func getAddressIndexByStringId(id string, ids []int32) int32 {
	idn, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		log.Fatalf("Error parsing the id: %s", err)
	}
	return getAddressIndexById(int32(idn), ids)
}

func getAddressIndexById(id int32, ids []int32) int32 {
	for index, iid := range ids {
		if iid == id {
			return int32(index)
		}
	}
	return -1
}
