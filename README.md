В корне проекта:

go mod tidy
mkdir -p storage/originals storage/processed

Запуск сервисов
Запусти HTTP API (приём изображений)
go run cmd/api/main.go


Ожидаешь в консоли:
Listening on :8080

Запустить Worker (обработка изображений)

В другом терминале:
go run cmd/worker/main.go

Использование через браузер

Открой:
http://localhost:8080

Что можно делать:
Выбрать файл (jpg / png / gif)

Нажать Upload

Увидеть статус:
uploaded
processing 

После ready изображение обработано

Файлы появятся тут:
```
storage/
 ├── originals/     # оригинал
 └── processed/     # resized + thumbnail
```
Использование через API (curl)
Загрузка изображения
curl -F "file=@cat.jpg" http://localhost:8080/upload
 
Ответ: 
b1f6e7b2-9c6e-4c8e-bb8d-1d8c5a1c3c33


Это image_id.

Проверка статуса
curl http://localhost:8080/image/b1f6e7b2-9c6e-4c8e-bb8d-1d8c5a1c3c33


Ответ: 
processing


или 
ready

Где лежат результаты

После успешной обработки:

storage/processed/
 ├── {id}_large.jpg   # большая версия
 └── {id}_thumb.jpg   # миниатюра

Удаление изображения (вручную)

(в текущей версии)

rm storage/originals/{id}

rm storage/processed/{id}_*.jpg 
