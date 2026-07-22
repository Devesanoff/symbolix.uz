#  Symbolix.uz

> O'zbek tili uchun yengil va tezkor matn konverteri (Kirill <-> Lotin) hamda in-memory PDF to DOCX aylantiruvchi veb-xizmat.

[![Live Demo](https://img.shields.io/badge/Website-symbolix.uz-blue?style=for-the-badge&logo=render)](https://symbolix.uz)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](LICENSE)

---

##  Imkoniyatlar (Features)

- **📝 Text Translator:** Matnlarni bir zumda Kirill alifbosidan Lotinga va aksincha o'girish (O'zbek tili qoidalariga moslashtirilgan).
- **📄 PDF to DOCX:** PDF fayllarni Word (`.docx`) formatiga tezkor va xavfsiz aylantirish.
- **⚡ In-Memory Processing:** Fayllar server xotirasida qayta ishlanadi va diskda saqlanmaydi (Maxfiylik кафолати).
- ** Ultra Fast:** Go (Golang) tilida yozilgan yuqori unumdorlikka ega backend architecture.

---

## Texnologiyalar (Tech Stack)

- **Backend:** Go (Golang)
- **Frontend:** HTML5, CSS3, Vanilla JavaScript


---

## 📁 Loyiha strukturasi

```text
├── cmd/
│   └── server/       # Ilovani ishga tushirish (Main entrypoint)
├── internal/         # Server logikasi va handlerlar
├── pkg/
│   └── translator/   # Kirill-Lotin tarjima algoritmi
├── web/              # Frontend statik fayllar (HTML, CSS, JS)
├── go.mod
└── README.md