package sub_str

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestIndexOf(t *testing.T) {

	url := "https://finance.yahoo.com/quote/AAPL/"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	str := string(body)
	substrIndex := strings.Index(str, "price")
	fmt.Println(substrIndex)
}
