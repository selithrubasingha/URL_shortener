# URL Shortener (Go)

A lightweight, high-performance URL shortener built with Go.

### Features

* **Fast:** Minimal overhead using Go's standard library/fast router.
* **Simple:** Straightforward API for shortening and redirecting.
* **Persistent:** [Mention your database, e.g., Redis or PostgreSQL].

### Getting Started

1. **Clone the repo:**
```bash
git clone https://github.com/selithrubasingha/URL_shortener.git
cd URL_shortener

```


2. **Run it:**
```bash
go run main.go

```



### API Usage

* **Shorten a URL:**
`POST /shorten` with body `{"url": "[https://your-long-url.com](https://your-long-url.com)"}`
* **Redirect:**
`GET /{shortCode}`

---

*Built by Selith Rubasingha.*

---
