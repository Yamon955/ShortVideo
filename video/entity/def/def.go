package def

const (
	DefaultPieceSize = 1024 * 1024 // 默认分片大小 1MB

	MinIOBucketName = "svtest"

	MachineID = "machine-id"
)

// VideoInfo 上传的视频信息
type VideoInfo struct {
	UID      uint64   `json:"uid"`
	VID      string   `json:"vid"`
	PlayURL  string   `json:"play_url"`
	CoverURL string   `json:"cover_url"`
	Tags     []string `json:"tags"`
}
