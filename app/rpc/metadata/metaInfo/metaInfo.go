package metaInfo

type BucketInfo struct {
	BucketName string `json:"bucket_name"`
	CreateTime string `json:"create_time"`
	CreateUser string `json:"create_user"`
	Version    bool   `json:"version"`
}

type ObjectInfo struct {
	ObjectName      string   `json:"object_name"`
	ObjectId        string   `json:"object_id"`
	CreateTime      string   `json:"create_time"`
	CreateUser      string   `json:"create_user"`
	IsDir           bool     `json:"is_dir"`
	ChekCode        string   `json:"chek_code"`
	ReplicationAddr []string `json:"replication_addr"`
	Size            int64    `json:"size"`
}

//type ObjectTree struct {
//	ObjectName string `json:"object_name"`
//	CreateTime string `json:"create_time"`
//	Size int64 `json:"size"`
//	ObjectTrees []*ObjectTree `json:"object_trees"`
//}
