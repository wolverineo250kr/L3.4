package imageproc

import (
    "github.com/disintegration/imaging"
)

func Process(id string) error {
    img, err := imaging.Open(OriginalPath(id))
    if err != nil {
        return err
    }

    large := imaging.Resize(img, 1200, 0, imaging.Lanczos)
    thumb := imaging.Thumbnail(img, 300, 300, imaging.Lanczos)

    if err := imaging.Save(large, LargePath(id)); err != nil {
        return err
    }
    if err := imaging.Save(thumb, ThumbPath(id)); err != nil {
        return err
    }
    return nil
}
