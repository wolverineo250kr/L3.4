package main

import (
    "context"
    "image-processor/internal/imageproc"
    "image-processor/internal/kafka"
    "image-processor/internal/meta"
)

func main() {
    r := kafka.NewConsumer()
    for {
        msg, _ := r.ReadMessage(context.Background())
        id := string(msg.Value)

        meta.Update(id, meta.Processing)

        if err := imageproc.Process(id); err != nil {
            meta.Update(id, meta.Failed)
            continue
        }

        meta.Update(id, meta.Ready)
    }
}
