package meta

import "sync"

type Status string

const (
    Uploaded   Status = "uploaded"
    Processing Status = "processing"
    Ready      Status = "ready"
    Failed     Status = "failed"
)

type ImageMeta struct {
    ID     string
    Status Status
}

var (
    mu   sync.RWMutex
    data = map[string]*ImageMeta{}
)

func Create(id string) {
    mu.Lock()
    defer mu.Unlock()
    data[id] = &ImageMeta{ID: id, Status: Uploaded}
}

func Update(id string, s Status) {
    mu.Lock()
    defer mu.Unlock()
    if img, ok := data[id]; ok {
        img.Status = s
    }
}

func Get(id string) (*ImageMeta, bool) {
    mu.RLock()
    defer mu.RUnlock()
    img, ok := data[id]
    return img, ok
}

func Delete(id string) {
    mu.Lock()
    defer mu.Unlock()
    delete(data, id)
}
