package services

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	shell "github.com/ipfs/go-ipfs-api"
)

var sh = shell.NewShell("http://localhost:5001") 

func UploadToIPFS(text string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("text cannot be empty")
	}

	reader := strings.NewReader(text + "\n")  
	cid, err := sh.Add(reader)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}	

	return cid, nil
}

func FetchFromIPFS(cid string) (string, error) {
	url := fmt.Sprintf("http://localhost:8080/ipfs/%s", cid)
	res, err := http.Get(url);
	if err != nil {
		return "", err
	}

	defer res.Body.Close();

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
