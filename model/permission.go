package model

const (
	PERM_HOTELS_READ   = "hotels:read"
	PERM_HOTELS_WRITE  = "hotels:write"
	PERM_REVIEWS_READ  = "reviews:read"
	PERM_REVIEWS_WRITE = "reviews:write"
	PERM_ORDERS_READ   = "orders:read"
	PERM_ORDERS_WRITE  = "orders:write"
)

var (
	CustomerPermissions = []string{
		PERM_HOTELS_READ,
		PERM_REVIEWS_READ,
		PERM_REVIEWS_WRITE,
		PERM_ORDERS_READ,
		PERM_ORDERS_WRITE,
	}

	PartnerPermissions = []string{
		PERM_HOTELS_WRITE,
		PERM_HOTELS_READ,
		PERM_REVIEWS_READ,
	}
)

type Permissions []string

func (p Permissions) Include(code string) bool {
	for i := range p {
		if code == p[i] {
			return true
		}
	}
	return false
}
