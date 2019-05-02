package app

/*import (
	"net/http"
	"strings"

	"github.schibsted.io/Yapo/payment-schedule/pkg/domain"
)
*/
// CreateSubscriptionInteractor Define function for user case
/*type CreateSubscriptionInteractor interface {
	CreateSubscription(accountID int, email, adminEmail, name string, billingType domain.DocType, products []int) (interface{}, error)
}
*/
// CreateSubscriptionHandler  implements the handler interface and responds to /subscribe
// requests, then executes the CreateSubscription usecase and returns the response
// that the use case give to us.
/*
type CreateSubscriptionHandler struct {
	CreateSubscriptionInteractor CreateSubscriptionInteractor
	Logger                       domain.Logger
}

type createSubscriptionRequestInput struct {
	SubscriberID int            `json:"subscriberID"`
	Products     []int          `json:"products"`
	BillingType  domain.DocType `json:"billingType"`
	Email        string         `json:"email"`
	AdminEmail   string         `json:"adminEmail"`
	Name         string         `json:"name"`
}

type createSubscriptionRequestOutput struct {
	Subscription interface{} `json:"subscription"`
}

type createSubscriptionRequestError goutils.GenericError

func (*CreateSubscriptionHandler) Input() HandlerInput {
	return &createSubscriptionRequestInput{}
}

func (c *CreateSubscriptionHandler) Execute(ig InputGetter) *goutils.Response {
	input, response := ig()
	if response != nil {
		return response
	}

	in := input.(*createSubscriptionRequestInput)
	if !c.validate(in) {
		return &goutils.Response{
			Code: http.StatusBadRequest,
			Body: createSubscriptionRequestError{
				ErrorMessage: "MISSING_PARAMETERS",
			},
		}
	}

	newSubscription, err := c.CreateSubscriptionInteractor.CreateSubscription(
		in.SubscriberID,
		strings.ToLower(in.Email),
		in.AdminEmail,
		in.Name,
		in.BillingType,
		in.Products,
	)

	if err != nil {
		return &goutils.Response{
			Code: http.StatusBadRequest,
			Body: createSubscriptionRequestError{
				err.Error(),
			},
		}
	}

	return &goutils.Response{
		Code: http.StatusOK,
		Body: createSubscriptionRequestOutput{
			Subscription: newSubscription,
		},
	}
}

func (c *CreateSubscriptionHandler) validate(in *createSubscriptionRequestInput) bool {
	if in.SubscriberID == 0 {
		c.Logger.Debug("Missing SubscriberID")
		return false
	}

	if in.Name == "" {
		c.Logger.Debug("Missing Subscriber Name")
		return false
	}

	if len(in.Products) == 0 {
		c.Logger.Debug("Missing Products %#v", in.Products)
		return false
	}

	if in.BillingType == "" || in.Email == "" || in.AdminEmail == "" {
		c.Logger.Debug("Missing Billing (%#v) or Email ((%#v) or AdminEmail %#v", in.BillingType, in.Email, in.AdminEmail)
		return false
	}

	return true
}
*/
