# API Endpoint Report

## Overview
This report lists every API route in the **BackEnd** project, providing:
- **HTTP method & path**
- **Roles** that can access the endpoint (derived from middleware)
- Whether **URL parameters** are required
- **Exact JSON schema** for the request body (if any)
- **Exact JSON schema** for the response body

The schemas are taken directly from `api_schema.json` and formatted for readability.

---

## Auth

### POST `/auth/login`
- **Roles:** Public (no auth)
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {
    "email": { "type": "string", "format": "email" },
    "password": { "type": "string" }
  },
  "required": ["email", "password"]
}
```

**Response Schema:**
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": {
    "token": { "type": "string" },
    "user": {
      "type": "object",
      "properties": {
        "id": { "type": "integer" },
        "username": { "type": "string" },
        "email": { "type": "string", "format": "email" },
        "role": { "type": "string" }
      },
      "required": ["id", "username", "email", "role"]
    }
  },
  "required": ["token", "user"]
}
```

---

### POST `/auth/register`
- **Roles:** Public (no auth)
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {
    "username": { "type": "string" },
    "email": { "type": "string", "format": "email" },
    "password": { "type": "string" },
    "role": { "type": "string" },
    "fullname": { "type": "string" }
  },
  "required": ["username", "email", "password", "fullname"]
}
```

**Response Schema:**
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "user": { "$ref": "#/definitions/UserResponse" }
  },
  "required": ["message", "user"]
}
```

---

### POST `/auth/refresh`
- **Roles:** Public (no auth)
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "token": { "type": "string" }
  },
  "required": ["message", "token"]
}
```

---

### POST `/auth/logout`
- **Roles:** Public (no auth)
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": { "message": { "type": "string" } },
  "required": ["message"]
}
```

---

## Blog

### POST `/api/v1/blogs/` (and POST `/admin/blogs`)
- **Roles:** SuperAdmin, KoorMedcrev
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {
    "title": { "type": "string" },
    "slug": { "type": "string" },
    "content": { "type": "string" },
    "kategori": { "type": "array", "items": { "type": "string" } },
    "published_at": { "type": "string", "format": "date-time" },
    "status": { "type": "string" },
    "id_work": { "type": "integer" },
    "id_pengurus": { "type": "integer" },
    "files": { "type": "array", "items": { "type": "string", "format": "binary" } }
  },
  "required": ["title", "slug", "content", "kategori", "status", "id_work", "id_pengurus"]
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "data": { "$ref": "#/definitions/BlogResponse" }
  },
  "required": ["message", "data"]
}
```

---

### GET `/blogs`
- **Roles:** Public
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {
    "page": { "type": "integer", "default": 1 },
    "limit": { "type": "integer", "default": 10 },
    "kategori": { "type": "array", "items": { "type": "string" } }
  }
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "data": { "type": "array", "items": { "$ref": "#/definitions/BlogThumbnail" } },
    "totalPage": { "type": "integer" },
    "currentPage": { "type": "integer" }
  },
  "required": ["message", "data", "totalPage", "currentPage"]
}
```

---

### GET `/blogs/:id`
- **Roles:** Public
- **URL Params:** **Yes** (`id`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "blog": { "$ref": "#/definitions/BlogResponse" }
  },
  "required": ["message", "blog"]
}
```

---

### GET `/admin/blogs`
- **Roles:** SuperAdmin, KoorMedcrev
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "data": { "type": "array", "items": { "$ref": "#/definitions/BlogResponse" } }
  },
  "required": ["message", "data"]
}
```

---

### GET `/admin/blogs/:id`
- **Roles:** SuperAdmin, KoorMedcrev
- **URL Params:** **Yes** (`id`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "blog": { "$ref": "#/definitions/BlogResponse" }
  },
  "required": ["message", "blog"]
}
```

---

### PUT `/admin/blogs/:id`
- **Roles:** SuperAdmin, KoorMedcrev
- **URL Params:** **Yes** (`id`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {
    "title": { "type": "string" },
    "slug": { "type": "string" },
    "content": { "type": "string" },
    "kategori": { "type": "array", "items": { "type": "string" } },
    "status": { "type": "string" },
    "published_at": { "type": "string", "format": "date-time" },
    "files": { "type": "array", "items": { "type": "string", "format": "binary" } }
  }
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "data": { "$ref": "#/definitions/BlogResponse" }
  },
  "required": ["message", "data"]
}
```

---

### DELETE `/admin/blogs/:id`
- **Roles:** SuperAdmin, KoorMedcrev
- **URL Params:** **Yes** (`id`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" } },
  "required": ["message"]
}
```

---

## Gallery

### POST `/admin/gallery`
- **Roles:** SuperAdmin, KoorMedcrev
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {
    "gallery_name": { "type": "string" },
    "gallery_type": { "type": "string" },
    "description": { "type": "string" },
    "event_date": { "type": "string", "format": "date" },
    "files": { "type": "array", "items": { "type": "string", "format": "binary" } }
  },
  "required": ["gallery_name", "gallery_type", "description", "event_date"]
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "data": { "$ref": "#/definitions/GalleryResponse" }
  },
  "required": ["message", "data"]
}
```

---

### GET `/gallery`
- **Roles:** Public
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "data": { "type": "array", "items": { "$ref": "#/definitions/GalleryResponse" } }
  },
  "required": ["message", "data"]
}
```

---

### DELETE `/admin/gallery/:id`
- **Roles:** SuperAdmin, KoorMedcrev
- **URL Params:** **Yes** (`id`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" } },
  "required": ["message"]
}
```

---

## Pengurus

### POST `/pengurus` (and POST `/admin/pengurus`)
- **Roles:** PENGURUS, KOOR, BPH, ADMIN
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": { "...": { "type": "string" } }
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" }, "data": { "$ref": "#/definitions/PengurusResponse" } },
  "required": ["message", "data"]
}
```

---

### GET `/pengurus/division/:division`
- **Roles:** Public
- **URL Params:** **Yes** (`division`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" }, "data": { "type": "array", "items": { "$ref": "#/definitions/PengurusPublicResponse" } } },
  "required": ["message", "data"]
}
```

---

### PUT `/pengurus/me`
- **Roles:** PENGURUS, KOOR, BPH, ADMIN
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {
    "full_name": { "type": "string" },
    "email": { "type": "string", "format": "email" }
  }
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" }, "data": { "$ref": "#/definitions/PengurusResponse" } },
  "required": ["message", "data"]
}
```

---

### PUT `/admin/pengurus/:id`
- **Roles:** KOOR, BPH, ADMIN
- **URL Params:** **Yes** (`id`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": { "...": { "type": "string" } }
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" }, "data": { "$ref": "#/definitions/PengurusResponse" } },
  "required": ["message", "data"]
}
```

---

### DELETE `/admin/pengurus/delete/:id`
- **Roles:** KOOR, BPH, ADMIN
- **URL Params:** **Yes** (`id`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" } },
  "required": ["message"]
}
```

---

## Upload

### DELETE `/upload/file`
- **Roles:** Any authenticated user
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" } },
  "required": ["message"]
}
```

---

### GET `/upload/files`
- **Roles:** Any authenticated user
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "files": { "type": "array", "items": { "type": "string" } } }
}
```

---

## User

### POST `/user` (and POST `/admin/user/super-admin`)
- **Roles:** ADMIN, KOOR, BPH
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {
    "username": { "type": "string" },
    "email": { "type": "string", "format": "email" },
    "password": { "type": "string" },
    "role": { "type": "string" },
    "fullname": { "type": "string" }
  },
  "required": ["username", "email", "password", "fullname"]
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" }, "data": { "$ref": "#/definitions/UserResponse" } },
  "required": ["message", "data"]
}
```

---

### GET `/user/me`
- **Roles:** ADMIN, KOOR, BPH, PENGURUS
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "data": { "$ref": "#/definitions/UserResponse" }
  },
  "required": ["message", "data"]
}
```

---

### PUT `/user/:id`
- **Roles:** ADMIN, KOOR, BPH
- **URL Params:** **Yes** (`id`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {
    "username": { "type": "string" },
    "email": { "type": "string", "format": "email" },
    "fullname": { "type": "string" }
  }
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "data": { "$ref": "#/definitions/UserResponse" }
  },
  "required": ["message", "data"]
}
```

---

### DELETE `/user/:id`
- **Roles:** ADMIN, KOOR, BPH
- **URL Params:** **Yes** (`id`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" } },
  "required": ["message"]
}
```

---

## Work

### POST `/admin/works`
- **Roles:** SuperAdmin, Koordinator
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {
    "pengurus_id": { "type": "integer" },
    "title": { "type": "string" },
    "tagline": { "type": "string" },
    "description": { "type": "string" },
    "slug": { "type": "string" },
    "project_type": { "type": "string" },
    "technologies": { "type": "array", "items": { "type": "string" } },
    "project_date": { "type": "string", "format": "date" },
    "status": { "type": "string" },
    "division": { "type": "string" },
    "files": { "type": "array", "items": { "type": "string", "format": "binary" } }
  },
  "required": ["pengurus_id", "title", "tagline", "description", "slug", "project_type", "technologies", "project_date", "status"]
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" }, "data": { "$ref": "#/definitions/WorkResponseClient" } },
  "required": ["message", "data"]
}
```

---

### GET `/works`
- **Roles:** Public
- **URL Params:** No

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "data": { "type": "array", "items": { "$ref": "#/definitions/WorkResponseClient" } }
  },
  "required": ["message", "data"]
}
```

---

### GET `/works/:projecttype`
- **Roles:** Public
- **URL Params:** **Yes** (`projecttype`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": {
    "message": { "type": "string" },
    "data": { "type": "array", "items": { "$ref": "#/definitions/WorkResponseClient" } }
  },
  "required": ["message", "data"]
}
```

---

### PUT `/admin/works/:id`
- **Roles:** SuperAdmin, Koordinator
- **URL Params:** **Yes** (`id`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {
    "pengurus_id": { "type": "integer" },
    "title": { "type": "string" },
    "tagline": { "type": "string" },
    "description": { "type": "string" },
    "slug": { "type": "string" },
    "project_type": { "type": "string" },
    "project_date": { "type": "string", "format": "date" },
    "status": { "type": "string" },
    "technologies": { "type": "array", "items": { "type": "string" } },
    "existingID_image": { "type": "array", "items": { "type": "integer" } }
  }
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" }, "data": { "$ref": "#/definitions/WorkUpdateResponse" } },
  "required": ["message", "data"]
}
```

---

### DELETE `/admin/works/:id`
- **Roles:** SuperAdmin, Koordinator
- **URL Params:** **Yes** (`id`)

**Request Schema:**
```json
{
  "type": "object",
  "properties": {}
}
```

**Response Schema:**
```json
{
  "type": "object",
  "properties": { "message": { "type": "string" } },
  "required": ["message"]
}
```

---

## Definitions (Reusable Schemas)

(Only listed for reference – they are referenced by `$ref` above.)

- **UserResponse**, **BlogResponse**, **BlogThumbnail**, **BlogGalleryResponse**, **GalleryResponse**, **PengurusResponse**, **PengurusPublicResponse**, **WorkResponseClient**, **WorkUpdateResponse**, **WorkGalleryResponse** – see `api_schema.json` for full definitions.

---

*This markdown file was generated automatically to aid developers, testers, and documentation writers.*
