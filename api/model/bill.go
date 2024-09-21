package model

type Bill struct {
	ID        int `json:"id"`
	Price     int `json:"price"`
	ProjectID int `json:"project_id"`
	SrcUser   int `json:"src_user"`
	DstUser   int `json:"dst_user"`
}
