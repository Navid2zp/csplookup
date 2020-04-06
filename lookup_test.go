package csplookup

import (
	"fmt"
	"testing"
	"time"
)

func TestTestResponseTime(t *testing.T) {
	result, err := TestResponseTime()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Name Lookup: ", result.NameLookup)
	fmt.Println("Connect: ", result.Connect)
	fmt.Println("TLS Handshake: ", result.TLSHandshake)
	fmt.Println("First Byte: ", result.FirstByte)
	fmt.Println("Full Response: ", result.FullResponse)
	fmt.Println("Body Size (byte): ", result.BodySize)

	jsonData, err := result.ToJSON()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonData))

	xmlData, err := result.ToXML()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(xmlData))
}

func TestLookup(t *testing.T) {
	client := NewClient("YOUR_API_KEY")

	response, err := client.LookupWithTL("4.2.2.4", time.Millisecond*2)
	if err != nil {
		if err != TimeLimitReached {
			t.Error(err)
		} else {
			fmt.Println("it took too long!")
		}
	}

	response, err = client.Lookup("4.2.2.4")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(response)
}
