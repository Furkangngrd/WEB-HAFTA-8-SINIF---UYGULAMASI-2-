# GoLearn Education API

GoLearn Uzaktan Eğitim Platformu API, Go, Gin, SQLite, GORM ve JWT kimlik doğrulama ile geliştirilmiştir. Temiz mimari (clean architecture), RBAC (Rol Tabanlı Erişim Kontrolü), Sınav (Quiz) sistemleri, Gelişim takibi ve WebSockets işlevlerini destekler.

### Öğrenci Bilgileri
- **Ad Soyad:** Muhammed Furkan Güngördü
- **Numara:** 24080410024

---

## 🚀 Geliştirme Süreci (13 Adımlık Plan)

Bu proje baştan sona aşağıdaki 13 adım eksiksiz takip edilerek kurulmuştur:

1. **Proje Kurulumu ve Modeller** - Go modüllerinin yüklenmesi, klasör yapısı (config, models vb.)
2. **Veritabanı Entegrasyonu** - Saf Go (Pure-Go) destekli `glebarez/sqlite` ve GORM kullanılarak db bağlantısı sağlandı.
3. **Roller ve Veri Modelleri** - DTO'lar, Öğrenci/Öğretmen gibi user rolleri sisteme işlendi.
4. **Kimlik Doğrulama (Auth JWT)** - `Register` ve `Login` işlemleri ve güvenli Token algoritması.
5. **Ara Katman (Middleware)** - Request Limitleri (RateLimiter), Auth/Roll doğrulama mekanizmaları.
6. **Kurs Yönetimi (Courses)** - Öğretmenlerin ders açabileceği, öğrencilerin kaydolabileceği uç noktalar.
7. **Ders Yönetimi (Lessons)** - Kursa bağlı video veya metin materyalleri yönetimi.
8. **Sınav Sistemi (Quizzes)** - Öğretmenlerin sınav hazırlama ve öğrencilerin çözümlerini gönderme özelliği.
9. **Gelişim Takibi (Progress)** - Öğrencilerin izlediği dersleri ilerlemesi ile kaydetme yeteneği.
10. **Canlı İletişim (WebSockets)** - Sistem odaları ve kullanıcılar arası canlı etkileşim entegresi.
11. **Konteyner (Docker)** - `Dockerfile` ve `docker-compose.yml` ayarları yapıldı.
12. **API Dokümantasyonu** - Bütün endpoint'ler için `Swagger` kuruldu ve kod yorumları ile üretildi.
13. **Testler ve İyileştirmeler** - Gerekli `_test.go` dosyalarının yazılması, versiyon kontrol (Git) ve son pürüzlerin temizlenmesi.

---

## 🛠 Kullanılan Teknolojiler

- **Dil:** Go 1.21+
- **Web Framework:** Gin
- **Veritabanı ve ORM:** SQLite (`github.com/glebarez/sqlite`) & GORM
- **Güvenlik ve İletişim:** JWT (JSON Web Tokens), Gorilla WebSockets
- **Dokümantasyon:** Swagger (go-swagger)

## 📌 Hızlı Başlangıç

(Kaynak kodlar `golearn/` klasörü içerisindedir.)

### Ön Gereksinimler
- Go 1.21+ (CGO gerekmez)
- İsteğe Bağlı: Docker Desktop

### 1- Docker İle Çalıştırma (Önerilen)
`golearn` klasörüne girip terminali açarak şu komutu uygulayın:
```bash
cd golearn
docker-compose up --build
```
Sunucu direkt olarak `http://localhost:8080/swagger/index.html` üzerinden aktif olacaktır.

### 2- Lokal (Docker Olmadan) Çalıştırma
Projeyi kendi cihazınızda derlemek için:
```bash
cd golearn
go mod tidy
go run main.go
```
API arayüzü ve dökümantasyon şu linkte başlayacaktır: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
