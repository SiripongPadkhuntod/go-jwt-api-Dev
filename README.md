# Go JWT API (Learning Project)

โปรเจกต์นี้เป็นตัวอย่าง **RESTful API ด้วยภาษา Go (Golang)**  
ออกแบบมาเพื่อใช้เรียนรู้พื้นฐานของ Backend Development โดยเน้นเรื่อง **JWT Authentication**

เหมาะสำหรับ **มือใหม่ที่เริ่มต้นเขียน API ด้วย Go**

---

## 🧠 Features

- ✅ JWT Authentication (Access Token)
- ✅ Register / Login
- ✅ Protected Routes ด้วย Middleware
- ✅ Gin Framework
- ✅ โครงสร้างโค้ดอ่านง่าย แยกเป็นสัดส่วน
- ⏳ พร้อมต่อยอด Database / Docker / Swagger

---

## 🗂 Project Structure

```
go-jwt-api-Dev/
│
├── main.go                # Entry point
├── routes/                # API routes
│   ├── auth_routes.go
│   └── item_routes.go
│
├── handlers/              # Controller / Handler logic
│   ├── auth_handler.go
│   └── item_handler.go
│
├── middleware/            # JWT Middleware
│   └── auth_middleware.go
│
├── utils/                 # Helper functions
│
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Prerequisites

- Go 1.20+ → https://go.dev/dl/
- Git → https://git-scm.com/

ตรวจสอบเวอร์ชัน Go:
```
go version
```

---

## 🚀 Installation & Setup

### 1️⃣ Clone Repository

```
git clone https://github.com/SiripongPadkhuntod/go-jwt-api-Dev.git
cd go-jwt-api-Dev
```

---

### 2️⃣ Install Dependencies

```
go mod tidy
```

---

### 3️⃣ Environment Variables

สร้างไฟล์ `.env`

```
JWT_SECRET=supersecretkey
PORT=8080
```

> ⚠️ ห้าม commit `.env` ขึ้น GitHub

---

### 4️⃣ Run Project

```
go run main.go
```

---

## 📡 API Endpoints

### Register
```
POST /auth/register
```

### Login
```
POST /auth/login
```

### Protected Route
```
GET /items
Authorization: Bearer <JWT_TOKEN>
```

---

## 🛡 JWT Flow

1. Login
2. Server สร้าง JWT
3. Client เก็บ Token
4. ส่ง Token ใน Header
5. Middleware ตรวจสอบ
6. ผ่าน → เข้าถึง API ได้

---

## 📦 Libraries

- Gin
- golang-jwt/jwt

---

## 🔮 Ideas สำหรับพัฒนาต่อ

- Refresh Token
- Role-based Access
- Database
- Unit Test
- Docker
- Swagger

---

## 👨‍💻 Author

Siripong Padkhuntod (Thyme)

⭐ หาก repo นี้ช่วยให้คุณเรียนรู้ อย่าลืมกด Star
