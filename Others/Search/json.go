package Search

type Search struct {
	Code int      `json:"code"`
	Msg string    `json:"msg"`
	Page int      `json:"page"`
	Pagecount int `json:"pagecount"`
	Limit string  `json:"limit"`
	Total int     `json:"total"`
	List []List   `json:"list"`
	Class []Class `json:"class"`
}
type List struct {
	VodID int `json:"vod_id"`
	VodName string `json:"vod_name"`
	TypeID int `json:"type_id"`
	TypeName string `json:"type_name"`
	VodEn string `json:"vod_en"`
	VodTime string `json:"vod_time"`
	VodRemarks string `json:"vod_remarks"`
	VodPlayFrom string `json:"vod_play_from"`
}
type Class struct {
	TypeID int `json:"type_id"`
	TypeName string `json:"type_name"`
}