package paymail

type (
	// Alias is the users alias.
	Alias string

	//Handle is the users full address i.e. {alias}@{domain.tld}.
	Handle string
)

type (
	// Domain is the current domain (tld).
	Domain interface {
		Verify(handle Handle) bool
	}
)
