# DevOps Case — updater servisi

Bu repo, sahadaki cihazlara image manifest'i dönen küçük bir Go servisini içeriyor.
Çalışıyor gibi duruyor ama release süreci güvenilir değil.

**Süre kutusu: 2–3 saat.** Bundan fazlasını harcamanı beklemiyoruz; yetiştiremediğin
maddeyi "yapmadım, sebebi şu" diye yazman yeterli, eksi puan değil.

**Gereken:** Docker ve Go. Başka hiçbir şey kurman gerekmiyor.

---

## Ne istiyoruz

### 1. Release tekrarlanabilir olsun

```
make release
make digest      # -> sha256:...
make release
make digest      # -> aynı çıktı olmalı
```

Şu an bu iki digest tutmuyor. Tutar hale getir.

*Minimum bar:* binary'nin sha256'sı iki build'de aynı olsun.
*Tam çözüm:* image digest'i de aynı olsun.

### 2. Testler güvenilir olsun

```
go test -count=1 ./...
```

Bunu arka arkaya 10 kez koştur, 10'unda da geçsin.
Testi silmek çözüm değil — neyin yanlış olduğunu bul.

### 3. Jenkinsfile

Koşturmana gerek yok, Jenkins kurma. Dosyayı oku, gördüğün **2 sorunu**
ilgili satırın üstüne yorum olarak yaz. Tek cümle yeter.

### 4. `deploy/updater-deployment.yaml`

Bu manifest sahaya çıkmadan önce ne değiştirirsin ve neden?
**En fazla 10 satır** yaz (ayrı bir dosyaya veya PR açıklamasına).
İstersen manifest'i düzelt, istersen sadece yaz.

### 5. İki soru, kısa cevap

`NOTES.md` açıp yaz, her biri birkaç satır:

- **Bilerek yapmadığın 2 şey nedir, neden yapmadın?**
- Bu repo 60 repoya çoğalsa ne kırılır?

---

## Teslim

- Kendi GitHub hesabında bir repo aç, çözümü oraya push et, linki gönder.
- **Commit geçmişi bizim için önemli.** Tek büyük commit yerine adım adım ilerle;
  denemelerin ve geri aldıkların da kalsın.
- Yapay zekâ kullanman serbest, kullandıysan `NOTES.md` içinde nerede kullandığını
  yaz. Sorun değil — teslim sonrası 20 dakikalık görüşmede kodun üstünden birlikte
  geçeceğiz.
