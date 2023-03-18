package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	token          string
	testMode       bool
	apiURL         string
}

type Size struct {
	Weight, Length, Width, Height float64
}

type PriceSending struct {
	TariffCode        int     `json:"tariff_code"`
	TariffName        string  `json:"tariff_name"`
	TariffDescription string  `json:"tariff_description"`
	DeliveryMode      int     `json:"delivery_mode"`
	DeliverySum       float64 `json:"delivery_sum"`
	PeriodMin         int     `json:"period_min"`
	PeriodMax         int     `json:"period_max"`
}

func NewClient(token string, testMode bool, apiURL string) *Client {
	return &Client{token: token, testMode: testMode, apiURL: apiURL}
}

func (c *Client) Calculate(addrFrom string, addrTo string, size Size) ([]PriceSending, error) {
	params := url.Values{}
	params.Set("city_from", addrFrom)
	params.Set("city_to", addrTo)
	params.Set("weight", strconv.FormatFloat(size.Weight, 'f', -1, 64))
	params.Set("length", strconv.FormatFloat(size.Length, 'f', -1, 64))
	params.Set("width", strconv.FormatFloat(size.Width, 'f', -1, 64))
	params.Set("height", strconv.FormatFloat(size.Height, 'f', -1, 64))

	req, err := http.NewRequest("GET", c.apiURL+"/calculate?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status: " + resp.Status)
	}

	var result struct {
		TariffCodes []PriceSending `json:"tariff_codes"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.TariffCodes, nil
}
