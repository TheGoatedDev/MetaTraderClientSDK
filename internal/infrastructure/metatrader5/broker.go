package metatrader5

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/TheGoatedDev/MetaTraderClientSDK/internal/domain/broker"
	broker_utils "github.com/TheGoatedDev/MetaTraderClientSDK/internal/shared/broker"
)

func Search(company string) ([]broker.Company, error) {
	if len(company) < 4 {
		return nil, fmt.Errorf("specify at least 4 symbols")
	}
	signature := broker_utils.GenerateSignature(company, "mt5")
	signatureText := hex.EncodeToString(signature)

	data := url.Values{}
	data.Set("company", company)
	data.Set("code", "mt5")
	data.Set("signature", signatureText)
	data.Set("ver", "2")

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://updates.metaquotes.net/public/mt5/network", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("User-Agent", "MetaTrader 5 Terminal/5.4420 (Windows NT 10.0.22621; x64)")
	req.Header.Add("Cookie", broker_utils.GetCookies())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	split := strings.Split(string(body), "\n")

	return jsonToCompanies(split[len(split)-1]).Result, nil
}

func jsonToCompanies(jsonStr string) broker.Companies {
	var companies broker.Companies

	err := json.Unmarshal([]byte(jsonStr), &companies)
	if err != nil {
		// Handle error
		return broker.Companies{}
	}
	return companies
}
