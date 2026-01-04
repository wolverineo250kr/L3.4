package imageproc

import "path/filepath"

var (
    OriginalsDir = "storage/originals"
    ProcessedDir = "storage/processed"
)

func OriginalPath(id string) string {
    return filepath.Join(OriginalsDir, id)
}

func LargePath(id string) string {
    return filepath.Join(ProcessedDir, id+"_large.jpg")
}

func ThumbPath(id string) string {
    return filepath.Join(ProcessedDir, id+"_thumb.jpg")
}
