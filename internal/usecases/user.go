package usecases

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	utils "github.com/Wisata-Kuliner/oslo/internal/utils"
)

// const projectID = "<your ID>"

// //
// const authKeyID = "your ID"
const path = "https://api.apple-cloudkit.com"
const version = "1"

// const container = "your container"
// const environment = "development"
const database = "public"

// const privPEM = `-----BEGIN EC PRIVATE KEY-----
// your privste key
// -----END EC PRIVATE KEY-----`

// const pubPEM = `-----BEGIN PUBLIC KEY-----
// your public key
// -----END PUBLIC KEY-----`

func GetUser() ([]byte, error) {
	// t0 := time.Now().UTC()
	// t1 := time.Now().UTC()

	authKeyID := os.Getenv("CLOUDKIT_API_KEY")

	// fmt.Printf("\nprivate key:\n")
	// https://golang.org/pkg/crypto/x509/#example_ParsePKIXPublicKey
	privPEM := strings.Replace(os.Getenv("CLOUDKIT_EC_KEY"), `\n`, "\n", -1)
	// fmt.Println()
	privBlock, _ := pem.Decode([]byte(privPEM))
	if privBlock == nil {
		panic("failed to parse PEM block containing the public key")
	}

	private_key, err := x509.ParseECPrivateKey(privBlock.Bytes)
	if err != nil {
		panic("failed to parse PEM block containing the public key")
	}

	pubPEM := strings.Replace(os.Getenv("CLOUDKIT_PUBLIC_KEY"), `\n`, "\n", -1)

	pubBlock, _ := pem.Decode([]byte(pubPEM))
	if pubBlock == nil {
		panic("failed to parse PEM block containing the public key")
	}

	var public_key *ecdsa.PublicKey
	public_k, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		panic("failed to parse PEM block containing the public key")
	}
	switch public_k1 := public_k.(type) {
	case *ecdsa.PublicKey:
		public_key = public_k1
	default:
		//return false
	}
	//////////
	// Config
	//////////
	requestPath := "/database/" +
		version + "/" +
		os.Getenv("CLOUDKIT_CONTAINER_ID") + "/" +
		os.Getenv("CLOUDKIT_ENV") + "/" +
		database + "/" +
		"records/query"

	requestBody := `{"query": {"recordType": "Users"}}`

	f := "2006-01-02T15:04:05Z"
	requestDate := time.Now().UTC().Format(f)

	h := sha256.New()
	h.Write([]byte(requestBody))
	b := h.Sum(nil)
	hashedBody := base64.StdEncoding.EncodeToString(b)

	rawPayload := requestDate + ":" + hashedBody + ":" + requestPath

	signedSignature, err := utils.SignMessage(private_key, []byte(rawPayload))
	if err != nil {
		fmt.Printf("SignMessage  error: %s\n", err.Error())
		return nil, err
	}

	verify := utils.VerifyMessage(public_key, []byte(rawPayload), signedSignature)
	fmt.Printf("signature verification result: %t\n", verify)

	requestSignature := base64.StdEncoding.EncodeToString(signedSignature)
	url := path + requestPath

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestBody)))

	if err != nil {
		fmt.Printf("NewRequest  error: %s\n", err.Error())
		return nil, err
	}

	req.Header.Add("content-type", "text/plain")
	req.Header.Add("X-Apple-CloudKit-Request-KeyID", authKeyID)

	req.Header.Add("X-Apple-CloudKit-Request-ISO8601Date", requestDate)

	req.Header.Add("X-Apple-CloudKit-Request-SignatureV1", requestSignature)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("\nresp.err:%v\n", err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("\nioutil.ReadAll.err:%v\n", err.Error())
		return nil, err
	}
	// fmt.Printf("\nrbody:%s\n", rbody)

	// curl := "curl -X POST -H \"content-type: text/plain\"" + " " +
	// 	"-H X-Apple-CloudKit-Request-KeyID:" + authKeyID + " " +
	// 	"-H X-Apple-CloudKit-Request-ISO8601Date:" + requestDate + " " +
	// 	"-H X-Apple-CloudKit-Request-SignatureV1:" + base64.StdEncoding.EncodeToString(signedSignature) + " " +
	// 	" -d " + "'" + requestBody + "'" + " " +
	// 	url

	// fmt.Printf("\n%s\n", curl)

	return rbody, nil
}
