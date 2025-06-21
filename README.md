# 📘 EPUB Image Processor

Утилита на Go для пакетной обработки изображений внутри `.epub` книг:

- 🔍 уменьшает изображения (по проценту),
- ⚫ преобразует в чёрно-белые (по флагу),
- ⚡ обрабатывает изображения многопоточно,
- 📦 пересобирает EPUB с тем же именем.

---

## 🚀 Установка

```bash
git clone https://github.com/dekonix/epub-image-processor.git
cd epub-image-processor
go build -o epub-imgproc
```

---

🔧 Использование

```bash
./epub-imgproc -in path/to/book.epub [опции]
```

Примеры:

📉 Уменьшить изображения до 60% и сделать их чёрно-белыми:

```bash
./epub-imgproc -in assets/book.epub -resize 60 -grayscale
```

⚙️ Только изменить размер (без ч/б):

```bash
./epub-imgproc -in assets/book.epub -resize 50
```

🧵 Использовать 8 потоков:

```bash
./epub-imgproc -in assets/book.epub -workers 8
```

📂 Сохранить в конкретный файл:

```bash
./epub-imgproc -in assets/book.epub -out output/optimized.epub
```

---

🛠 Параметры CLI

| Флаг         | Описание                                                        |
| ------------ | --------------------------------------------------------------- |
| `-in`        | Путь к входному EPUB (обязательно)                              |
| `-out`       | Путь к выходному EPUB (по умолчанию: output/имя_файла.epub)     |
| `-resize`    | Масштаб изображений в процентах (по умолчанию: 100)             |
| `-grayscale` | Преобразовать изображения в чёрно-белые                         |
| `-workers`   | Количество потоков для обработки (по умолчанию: число ядер CPU) |

---

📜 Лицензия

MIT © 2025 — dekonix
