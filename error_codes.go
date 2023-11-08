package error_codes

const (
	REQUIRED             = 1000
	NUMERIC              = 1001
	ALPHANUMERIC         = 1002
	APLHASPACEDOTSTRIP   = 1003
	EQUALTOPASSWORD      = 1004
	NOTAVAILABLE         = 1005
	INCORRECTEMAILFORMAT = 1006
	INVALIDFILETYPE      = 1007
	OVERLIMIT            = 1008
	CHARACTERNOTALLOWED  = 1009
	DUPLICATE            = 1010
	EMAILNOTSENT         = 1011
	GENERATEOTPFAILED    = 1012
	PERMISSIONDENIED     = 1013
	NULL                 = 0

	MIN3 = 2003
	MIN5 = 2005
	MIN6 = 2006
	MIN8 = 2008
	MIN9 = 2009

	MAX1   = 3001
	MAX4   = 3004
	MAX6   = 3006
	MAX8   = 3008
	MAX20  = 3020
	MAX50  = 3050
	MAX64  = 3064
	MAX100 = 3100
	MAX150 = 3150
	MAX250 = 3250
	MAX255 = 3255
	MAX350 = 3250
	MAX500 = 3500

	BADREQUEST          = 400
	UNAUTHORIZE         = 401
	FORBIDDEN           = 403
	NOTFOUND            = 404
	EXPIRED             = 419
	TOOEARLY            = 425
	INTERNALSERVERERROR = 500
)

func InitiateDataError() []map[string]interface{} {
	return []map[string]interface{}{}
}

func DataError(code int, field string) map[string]interface{} {
	return map[string]interface{}{
		"code":  code,
		"field": field,
	}

}
