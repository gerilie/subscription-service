// Package ping provides a simple ping service.
package ping

//	@Summary	Ping service
//	@Tags		health
//	@ID			ping
//	@Produce	json
//	@Success	200	{object}	pingResp	"Ping success"
//	@Failure	500	{string}	string		"Internal server error"
//	@Router		/ping [get].
