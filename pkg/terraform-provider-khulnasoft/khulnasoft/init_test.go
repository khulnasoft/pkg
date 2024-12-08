package khulnasoft

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
)

func init() {
	log.Println("setup suite")
	var (
		present, verifyTLS                                       bool
		username, password, khulnasoftURL, verifyTLSString, caCertPath string
		err                                                      error
		caCertByte                                               []byte
	)

	username, present = os.LookupEnv("KHULNASOFT_USER")
	if !present {
		panic("KHULNASOFT_USER env is missing, please set it")
	}

	password, present = os.LookupEnv("KHULNASOFT_PASSWORD")
	if !present {
		panic("KHULNASOFT_PASSWORD env is missing, please set it")
	}

	khulnasoftURL, present = os.LookupEnv("KHULNASOFT_URL")
	if !present {
		panic("KHULNASOFT_URL env is missing, please set it")
	}

	verifyTLSString, present = os.LookupEnv("KHULNASOFT_TLS_VERIFY")
	if !present {
		verifyTLSString = "true"
	}

	caCertPath, present = os.LookupEnv("KHULNASOFT_CA_CERT_PATH")
	if present {
		if caCertPath != "" {
			caCertByte, err = os.ReadFile(caCertPath)
			if err != nil {
				panic("Unable to read CA certificates")
			}
		}
		panic("KHULNASOFT_CA_CERT_PATH env is missing, please set it")
	}

	verifyTLS, _ = strconv.ParseBool(verifyTLSString)

	khulnasoftClient := client.NewClient(khulnasoftURL, username, password, verifyTLS, caCertByte)
	token, url, err := khulnasoftClient.GetAuthToken()

	if err != nil {
		panic(fmt.Errorf("failed to receive token, error: %s", err))
	}

	err = os.Setenv("TESTING_AUTH_TOKEN", token)
	if err != nil {
		panic("Failed to set AUTH_TOKEN env")
	}

	err = os.Setenv("TESTING_URL", url)
	if err != nil {
		panic("Failed to set TESTING_URL env")
	}
	log.Println("Finished to set token")

}
