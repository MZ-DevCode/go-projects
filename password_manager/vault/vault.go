package vault

import(
	"encoding/json"
)

type Account struct{
	Resource string `json:"resource"`
	Login 	 string `json:"login"`
	Password string `json:"password"`
}

type Vault struct{
	Accounts []Account `json:"accounts"`
}

func NewVault() *Vault{
	return &Vault{
		Accounts: []Account{},
	}
}

func (v *Vault) AddAccount(newAcc Account){
	v.Accounts = append(v.Accounts, newAcc)
}

func (v *Vault) ToBytes() ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (v *Vault) FromBytes(data []byte) error {
		err := json.Unmarshal(data, v)
		if err != nil {
			return err
		}
		return nil
}
