package k8s

import (
	account "github.com/lastbackend/lastbackend/pkg/account/api/v1"
)


type ComponentAccountsGetter interface {
	Accounts() ComponentAccountInterface
}

// ComponentStatusInterface has methods to work with ComponentStatus resources.
type ComponentAccountInterface interface {
	Create(*account.ComponentAccount) (*account.ComponentAccount, error)
}

// componentStatuses implements ComponentStatusInterface
type componentAccounts struct {
	client *LBClient
}

// newComponentStatuses returns a ComponentStatuses
func newComponentAccounts(c *LBClient) *componentAccounts {
	return &componentAccounts{
		client: c,
	}
}

// Create takes the representation of a componentStatus and creates it.  Returns the server's representation of the componentStatus, and an error, if there is any.
func (c *componentAccounts) Create(componentAccount *account.ComponentAccount) (result *account.ComponentAccount, err error) {
	result = &account.ComponentAccount{}
	err = c.client.Post().
		Resource("componentaccounts").
		Body(componentAccount).
		Do().
		Into(result)
	return
}



