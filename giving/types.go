package giving

import (
	"time"

	"github.com/Budmin/pcgo"
)

type Batch struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		CommitedAt    time.Time `json:"commited_at"`
		Description   string    `json:"description"`
		TotalCents    int       `json:"total_cents"`
		TotalCurrency string    `json:"total_currency"`
		Status        string    `json:"status"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type BatchGroup struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		Description   string    `json:"description"`
		Commited      bool      `json:"commited"`
		TotalCents    int       `json:"total_cents"`
		TotalCurrency string    `json:"total_currency"`
		Status        string    `json:"status"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Campus struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name    string      `json:"name"`
		Address interface{} `json:"address"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Designation struct {
	Type string `json:"Designation"`
	ID   string `json:"id"`

	Attributes struct {
		AmountCents     int    `json:"amount_cents"`
		AmmountCurrency string `json:"ammount_currency"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type DesignationRefund struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		AmountCents    int    `json:"amount_cents"`
		AmountCurrency string `json:"amount_currency"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Donation struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
		PaymentMethodSub    string    `json:"payment_method_sub"`
		PaymentLast4        string    `json:"payment_last4"`
		PaymentBrand        string    `json:"payment_brand"`
		PaymentCheckNumber  int       `json:"payment_check_number"`
		PaymentCheckDatedAt string    `json:"payment_check_dated_at"`
		FeeCents            int       `json:"fee_cents"`
		PaymentMethod       string    `json:"payment_method"`
		ReceivedAt          time.Time `json:"received_at"`
		AmountCents         int       `json:"amount_cents"`
		PaymentStatus       string    `json:"payment_status"`
		CompletedAt         time.Time `json:"completed_at"`
		AmountCurrency      string    `json:"amount_currency"`
		FeeCurrency         string    `json:"fee_currency"`
		Refunded            bool      `json:"refunded"`
		Refundable          bool      `json:"refundable"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Fund struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Name        string    `json:"name"`
		LedgerCode  string    `json:"ledger_code"`
		Description string    `json:"description"`
		Visibility  string    `json:"visibility"`
		Default     bool      `json:"default"`
		Color       string    `json:"color"`
		Deletable   bool      `json:"deletable"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Lable struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Slug string `json:"slug"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Note struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Body string `json:"body"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Organization struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name string `json:"name"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PaymentMethod struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		MethodType    string    `json:"method_type"`
		MethodSubtype string    `json:"method_subtype"`
		Last4         string    `json:"last4"`
		Brand         string    `json:"brand"`
		Expiration    string    `json:"expiration"`
		Verified      bool      `json:"verified"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PaymentSource struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Person struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Permissions    string        `json:"permissions"`
		EmailAddresses []interface{} `json:"email_addresses"`
		Addresses      []interface{} `json:"addresses"`
		PhoneNumbers   []interface{} `json:"phone_numbers"`
		FirstName      string        `json:"first_name"`
		LastName       string        `json:"last_name"`
		DonorNumber    int           `json:"donor_number"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Pledge struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt                   time.Time `json:"created_at"`
		UpdatedAt                   time.Time `json:"updated_at"`
		AmountCents                 int       `json:"amount_cents"`
		AmountCurrency              string    `json:"amount_currency"`
		JointGiverAmountCents       int       `json:"joint_giver_amount_cents"`
		DonatedTotalCents           int       `json:"donated_total_cents"`
		JointGiverDonatedTotalCents int       `json:"joint_giver_donated_total_cents"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type PledgeCampaign struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt                          time.Time `json:"created_at"`
		UpdatedAt                          time.Time `json:"updated_at"`
		Name                               string    `json:"name"`
		Description                        string    `json:"description"`
		StartsAt                           time.Time `json:"starts_at"`
		EndsAt                             time.Time `json:"ends_at"`
		GoalCents                          int       `json:"goal_cents"`
		GoalCurrency                       string    `json:"goal_currency"`
		ShowGoalInChurchCenter             bool      `json:"show_goal_in_church_center"`
		ReceivedTotalFromPledgesCents      int       `json:"received_total_from_pledges_cents"`
		ReceivedTotalOutsideOfPledgesCents int       `json:"received_total_outside_of_pledges_cents"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type RecurringDonation struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt              time.Time `json:"created_at"`
		UpdatedAt              time.Time `json:"updated_at"`
		ReleaseHoldAt          time.Time `json:"release_hold_at"`
		AmountCents            int       `json:"amount_cents"`
		Status                 string    `json:"status"`
		LastDonationReceivedAt time.Time `json:"last_donation_received_at"`
		NextOccurrence         time.Time `json:"next_occurrence"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type RecurringDonationDesignation struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		AmountCents    int    `json:"amount_cents"`
		AmountCurrency string `json:"amount_currency"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Refund struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		AmountCents    int       `json:"amount_cents"`
		AmountCurrency string    `json:"amount_currency"`
		FeeCents       int       `json:"fee_cents"`
		RefundedAt     time.Time `json:"refunded_at"`
		FeeCurrency    string    `json:"fee_currency"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}
