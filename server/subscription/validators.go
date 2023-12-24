package subscriptions

import "github.com/AlexJMcLean/subscriptions/users"

type SubscriptionModelValidator struct {
	Subscription struct {
		User users.UserModel
	}
}

func NewSubscriptionValidator() SubscriptionModelValidator {
	subscriptionModelValidator := SubscriptionModelValidator{}
	return subscriptionModelValidator
}


