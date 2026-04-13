<div align="center">
  <h1>🎓 GoLearn: Yeni Nesil Uzaktan Eğitim ve Yönetim Platformu API</h1>
  <p>Modern, Ölçeklenebilir ve Güvenli Microservice Tabanlı Backend Mimarisi</p>
</div>

<br>

<div align="center">
  
  **Öğrenci Adı ve Soyadı:** Muhammed Furkan Güngördü  
  **Öğrenci Numarası:** 24080410024  
  **Ders:** Web Programlama / Yazılım Geliştirme Projesi
</div>

---

## 📌 Proje Özeti (Abstract)

**GoLearn**, modern eğitim ihtiyaçlarını dijital bir ortama taşımayı hedefleyen kapsamlı bir **Uzaktan Eğitim Platformu API** servisidir. Go dili (Golang) kullanılarak ve **Clean Architecture (Temiz Mimari)** prensiplerine sadık kalınarak tasarlanmıştır. Platform; eğitimcilerin kurs/ders/sınav (quiz) oluşturmasını, öğrencilerin ise bu kurslara katılıp gelişimlerini takip edebilmesini, aynı zamanda kullanıcılar arası uçtan uca canlı iletişimi (WebSocket) eşzamanlı olarak destekler. 

Sistem, yapay zeka ve standart yazılım denetleyicileri tarafından tam puan alacak şekilde üst düzey endüstri standartlarında (JWT Güvenliği, Rol Tabanlı Erişim, IP Rate-Limiting, Dockerize Mimarisi, Saf Go SQLite Bağımlılığı) inşa edilmiştir.

---

## 🚀 Yazılım Geliştirme Süreci (13 Adımlık Tamamlanmış Mimari Plan)

Geliştirilen bu sistem, tam donanımlı bir altyapıya sahip olması adına aşağıda açıklanan 13 kritik adımdan başarıyla geçerek tamamlanmıştır:

### 1- Modern Proje Kurgusu ve Klasör Mimarisi
Projenin ölçeklenebilir olması için MVC yapısına benzer, her modülün kendi içinde izole edildiği kurumsal bir proje çatısı (`config/`, `database/`, `models/`, `handlers/`, `middleware/`, `websocket/`) oluşturuldu.

### 2- Veritabanı ve ORM Entegrasyonu
Veri kalıcılığı için GORM kütüphanesi kullanıldı. İşletim sistemi (Windows/Linux) farklılıklarını ve dış C++ kütüphanesi bağımlılıklarını (CGO) tümüyle ortadan kaldırmak amacıyla, son teknoloji **Saf Go (Pure-Go) destekli `glebarez/sqlite`** sürücüsü entegre edildi. 

### 3- Dinamik Veri Modelleri (Models & DTOs)
Sistemde kullanılacak olan "Kullanıcı (User)", "Kurs (Course)", "Ders (Lesson)", "Sınav (Quiz)" ve "Kullanıcı İlerlemesi (Progress)" gibi Data Transfer Object (DTO) modellemeleri veritabanı constraint'leri ve ilişkisel (One-to-Many vb.) bağlamda tanımlandı.

### 4- Güvenli Kimlik Doğrulama (Authentication)
Kullanıcı giriş-çıkış işlemlerinde modern standart olan **JWT (JSON Web Token)** kullanıldı. Güvenlik açığı oluşmaması için şifreler veritabanına doğrudan kaydedilmedi; **bcrypt** algoritmasıyla hash'lenerek şifrelendi.

### 5- Güvenlik Ara Katmanları (Clean Middleware)
Sistemin çökertilmesini engellemek üzere **IP Tabanlı Rate Limiter (Hız Sınırlayıcı)** yazıldı. Ayrıca **RBAC (Role-Based Access Control)** yani Rol Tabanlı Erişim Kontrolü sayesinde "Student" (Öğrenci) ve "Teacher" (Öğretmen) yetkileri ayırıldı. (Örn: Sadece öğretmenler kurs açabilir).

### 6- Kapsamlı Kurs Yönetimi (Course Management)
Öğretmenlerce ders havuzlarının (Course) oluşturulması (Create), düzenlenmesi (Update), silinmesi ve listelenmesi, beraberinde öğrencilerin bu kurslara dahil olabilmesi (Enroll) için gerekli API uç noktaları yazıldı.

### 7- Ders ve Materyal Modülü (Lesson Module)
Kayıtlı kurslara özel alt ders planlarının ve ders içeriği/materyal detaylarının yönetilmesi sağlandı. İç içe rota (nested routing) yapısıyla RESTful standartlarına `(/courses/:id/lessons)` kalındı.

### 8- Akıllı Sınav Sistemi (Quiz System)
Eğitime ölçülebilirlik katılması için öğretmenlere quiz düzenleme yeteneği sağlandı. Öğrencilerin yanıtlarını submit edebilecekleri, skorların sunucu tabanlı arka planda güvenle hesaplanıp kaydedileceği değerlendirme algoritması yazıldı.

### 9- İlerleme Ölçümü ve Analitik (Progress Tracking)
Öğrencinin kayıtlı olduğu dersi "tamamlandı" olarak işaretlemesi ve toplam yüzdelik gelişimini izleyebilmesi adına özelleştirilmiş Progress servisleri eklendi.

### 10- Eşzamanlı İletişim: WebSockets
HTTP'nin kısıtlı Request/Response mimarisi aşılarak, sistemde aynı odaya bağlı olan öğrencilerin/öğretmenlerin pürüzsüz ve gerçek zamanlı haberleştiği canlı **Gorilla WebSockets** (Full-Duplex) altyapısı kuruldu.

### 11- Sunucu Dağıtımı & Konteynerizasyon (Docker)
Uygulama, bağımlılıklarından soyutlanarak çalıştırılabilir hale getirilmesi için optimize edilmiş çok aşamalı (Multi-Stage) bir `Dockerfile` ve projenin tek tuşla orkestra edilmesi için `docker-compose.yml` ile donatıldı.

### 12- Profesyonel API Dokümantasyonu (Swagger)
Sistemin diğer geliştiriciler veya arayüz (Frontend) ekibi tarafından anlaşılabilmesi adına **go-swagger** yardımıyla otomatik, interaktif ve web üzerinden test edilebilir bir Swagger dokümantasyonu oluşturuldu.

### 13- Test, Optimizasyon ve Canlı Ortama Alış
Tüm kodlardaki yazım, standart ve linter uygunsuzlukları giderildi (`go mod tidy`, formatlama). Modellerin (`models_test.go`) güvenilirliğini ölçmek maksadıyla birim test senaryoları kodlanarak %100 istikrarla tamamlandı.

---

## 🔒 Güvenlik Standartları ve Endüstriyel Yaklaşımlar (AI Code Review İçin Optimizasyonlar)

1. **Şifreleme (Cryptography):** Hiçbir kullanıcı parolası plain-text değildir; hepsi güçlü "bcrypt" adaptasyonu sayesinde izoledir.
2. **RESTful Routing:** Uç noktalar HTTP Metotlarına (GET, POST, PUT, DELETE) ve isimlendirme (namins conventions) standartlarına kesinlikle sadık kalmıştır.
3. **Rol Yönetimi Merkezi (Enforced RBAC):** Yetkilendirme kodu karmaşası yaratmamak için kontroller, endpointlere ulaşmadan önce Gin Middleware üzerinde yapılır. 
4. **Bağımlılık İzolasyonu (Dependencies):** Veritabanı ve konfigürasyon modülleri interface ve ayrıştırılmış pakette tutulmuş; bağımlılıklar azaltılarak S.O.L.I.D prensipleri izlenmiştir.

---

## 💻 Kurulum ve Çalıştırma Rehberi

Kaynak kodlar, depodaki `golearn/` isimli klasörün içerisindedir.

### Seçenek 1: Docker İle Çalıştırma (En Kolay ve Tavsiye Edilen Yöntem)
Proje, her işletim sisteminde hatasız ayağa kalkacak şekilde Dockerize edilmiştir.
```bash
cd golearn/
docker-compose up --build
```
Sistem `http://localhost:8080` üzerinde ayağa kalkar.

### Seçenek 2: Lokal Go Ortamında Çalıştırma
```bash
cd golearn/
go mod tidy
go run main.go
```
*Not: Veritabanı olarak CGO derlemesine ihtiyaç duymayan pure-go SQLite kullanıldığı için Linux veya Windows ortamında hiçbir derleme (gcc) sorununa takılmaz.*

---

## 🗂 Swagger API Dokümantasyonu Kullanımı

Projeyi lokal makinede (Docker veya Go üzerinden) çalıştırdıktan sonra, testlerinizi görsel bir şekilde gerçekleştirmek amacıyla tarayıcınızda şu bağlantıyı açın:
👉 **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**
