package snap

type (
	SnapTokenReq struct {
		PrivateKey string `json:"privateKey"`
		TimeStamp  string `json:"timeStamp"`
	}

	AsymmetricSignatureSnap struct {
		PrivateKey string `json:"privateKey,omitempty"`
		PublicKey  string `json:"publicKey,omitempty"`
		TimeStamp  string `json:"timeStamp,omitempty"`
		ClientKey  string `json:"clientKey"`
	}

	SymetricSignatureSnap struct {
		Url          string `json:"url"`
		Method       string `json:"method"`
		AccessToken  string `json:"accessToken"`
		TimeStamp    string `json:"timeStamp,omitempty"`
		ClientSecret string `json:"clientSecret"`
		Body         string `json:"body"`
	}
)
