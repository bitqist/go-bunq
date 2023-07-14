package bunq

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type companyService service

// GetUserPerson retrieves a signle user company
// https://doc.bunq.com/#/user-company/READ_UserCompany
func (u *companyService) GetUserPerson() (*responseUserCompany, error) {
	userID, err := u.client.GetUserID()
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodGet, u.client.formatRequestURL(fmt.Sprintf(endpointUserCompanyGet, userID)), nil)
	if err != nil {
		return nil, errors.Wrap(err, "bunq: could not create request for user-company")
	}

	res, err := u.client.do(r)
	if err != nil {
		return nil, errors.Wrap(err, "bunq: request to user-company failed")
	}

	var resUserCompany responseUserCompany
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&resUserCompany)
	if err != nil {
		return nil, errors.Wrap(err, "bunq: parsing response failed")
	}

	return &resUserCompany, nil
}
