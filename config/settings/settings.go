package settings

type Settings struct {
	ClientId     string `envconfig:"PLAID_CLIENT_ID" default:""`
	Secret       string `envconfig:"PLAID_SECRET" default:""`
	Env          string `envconfig:"PLAID_ENV" default:"sandbox"`
	Products     string `envconfig:"PLAID_PRODUCTS" default:"transactions"`
	CountryCodes string `envconfig:"PLAID_COUNTRY_CODES" default:"CA"`
	RedirectUri  string `envconfig:"PLAID_REDIRECT_URI" default:""`
}
