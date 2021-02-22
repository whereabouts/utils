package jwt

import "time"

const (
	defaultHeaderAlgorithm   = "HSA256"
	defaultHeaderType        = "JWT"
	defaultPayloadIssuer     = "whereabouts.icu"
	defaultPayloadOwner      = ""
	defaultPayloadPurpose    = "authentication"
	defaultPayloadRecipient  = "browser"
	defaultPayloadExpireTime = time.Minute * 30
	defaultSignature         = "hezebin"
)

type Header struct {
	Algorithm string `json:"algorithm"`
	Type      string `json:"JWT"`
}

type Payload struct {
	Issuer    string                 `json:"issuer"`
	Owner     string                 `json:"owner"`
	Purpose   string                 `json:"purpose"`
	Recipient string                 `json:"recipient"`
	Time      int64                  `json:"time"`
	Expire    int64                  `json:"expire"`
	External  map[string]interface{} `json:"external"`
}

func defaultHeader() *Header {
	return &Header{Algorithm: defaultHeaderAlgorithm, Type: defaultHeaderType}
}

func defaultPayload() *Payload {
	currentTime := time.Now()
	return &Payload{
		Issuer:    defaultPayloadIssuer,
		Owner:     defaultPayloadOwner,
		Purpose:   defaultPayloadPurpose,
		Recipient: defaultPayloadRecipient,
		Time:      currentTime.Unix(),
		Expire:    currentTime.Add(defaultPayloadExpireTime).Unix(),
		External:  nil,
	}
}

var (
	signature = defaultSignature
)

func SetSignature(sign string) {
	signature = sign
}

func getSignature() string {
	return signature
}
