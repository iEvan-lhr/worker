package typ

type WorkerInfo struct {
	Id           int    `json:"id,omitempty" gorm:"primaryKey"`
	WorkerName   string `json:"worker_name" gorm:"column:worker_name"`
	NodeNow      string `json:"node_now" gorm:"column:node_now"`
	NextNode     string `json:"next_node" gorm:"column:next_node"`
	WorkerStatus int    `json:"worker_status" gorm:"column:worker_status"`
}
