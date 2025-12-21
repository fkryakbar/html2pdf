# HTML to PDF Microservice API

Microservice API untuk mengonversi HTML+CSS menjadi file PDF menggunakan Go dan wkhtmltopdf.

## Features

- ✅ Konversi HTML+CSS ke PDF
- ✅ Full CSS support (Flexbox, Grid, custom fonts)
- ✅ Variable substitution dengan syntax `<<variabel>>`
- ✅ API Key authentication
- ✅ Docker ready

## Quick Start

### Build Docker Image

```bash
docker build -t html2pdf .
```

### Run Container

```bash
docker run -p 8080:8080 -e API_KEY=your-secret-api-key html2pdf
```

## API Documentation

### POST /api/convert

Mengonversi HTML ke PDF.

**Headers:**

```
Content-Type: application/json
X-API-Key: your-api-key
```

**Request Body:**

```json
{
  "html": "<html><head><style>h1{color:red; font-family: Arial;}</style></head><body><h1>Halo <<nama>></h1><p>Tanggal: <<tanggal>></p></body></html>",
  "var": [{ "nama": "John Doe" }, { "tanggal": "21 Desember 2025" }]
}
```

**Response:**

- Content-Type: `application/pdf`
- Content-Disposition: `attachment; filename="document.pdf"`
- Body: PDF file binary

### GET /health

Health check endpoint (no authentication required).

**Response:**

```json
{ "status": "ok" }
```

## Variable Substitution

Gunakan syntax `<<variabel>>` dalam HTML untuk placeholder yang akan diganti dengan nilai dari array `var`.

**Contoh:**

```json
{
  "html": "<p>Nama: <<nama>>, Alamat: <<alamat>></p>",
  "var": [{ "nama": "John" }, { "alamat": "Jakarta" }]
}
```

**Hasil HTML:**

```html
<p>Nama: John, Alamat: Jakarta</p>
```

## Error Responses

| Status Code | Error                 | Description                          |
| ----------- | --------------------- | ------------------------------------ |
| 401         | unauthorized          | Missing or invalid API key           |
| 400         | bad_request           | Invalid JSON or missing HTML content |
| 500         | pdf_generation_failed | Failed to generate PDF               |

## Example Usage

### cURL

```bash
curl -X POST http://localhost:8080/api/convert \
  -H "Content-Type: application/json" \
  -H "X-API-Key: your-secret-api-key" \
  -d '{
    "html": "<html><head><style>body{font-family:Arial;} h1{color:blue;}</style></head><body><h1>Hello <<nama>></h1></body></html>",
    "var": [{"nama": "World"}]
  }' \
  --output document.pdf
```

### JavaScript (fetch)

```javascript
const response = await fetch("http://localhost:8080/api/convert", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    "X-API-Key": "your-secret-api-key",
  },
  body: JSON.stringify({
    html: "<html><body><h1>Hello <<nama>></h1></body></html>",
    var: [{ nama: "World" }],
  }),
});

const blob = await response.blob();
// Download or process the PDF blob
```

## Environment Variables

| Variable | Default                          | Description                |
| -------- | -------------------------------- | -------------------------- |
| API_KEY  | dev-api-key-change-in-production | API key for authentication |
| PORT     | 8080                             | Server port                |

## Development

### Prerequisites

- Go 1.21+
- Docker (for running with wkhtmltopdf)

### Build & Run with Docker

```bash
# Build
docker build -t html2pdf .

# Run
docker run -p 8080:8080 -e API_KEY=my-secret-key html2pdf
```

## License

MIT
