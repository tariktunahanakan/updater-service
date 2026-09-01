# Notlar

## Bilerek yapmadığım 2 şey

**1. Deployment'a probe ve resource limit eklemedim.**
readiness/liveness probe ve cpu/memory limitleri prod için şart, farkındayım.
Önce en kritik olanlara odaklandım; bunlar da eklenmeli.

**2. main.go'daki hardcoded token'ı koddan silmedim.**
Deployment tarafında token'ı Secret'a taşıdım, artık plaintext değil.
Ama main.go içinde token hâlâ sabit olarak duruyor; onu os.Getenv ile
okuyacak şekilde değiştirmedim. Bu kodu değiştirmem ayrı bir iş, önce
deploy tarafındaki security açığını kapattım. Kalıcı çözüm: value'yu
env'den okumak.

## Bu repo 60 repoya çoğalsa ne kırılır?

- **Secret her repoda kopyalanır.** Şu an Jenkinsfile'da şifre plaintext.
  Bunu 60 repoya kopyala-yapıştır yaparsak 60 yerde şifre sızar, birini
  değiştirmemiz gerekince hepsini tek tek düzeltmemiz gerekir. Merkezi bir
  secret yönetimi şart olur.
- **latest ve elle yazılan Dockerfile.** Her repo golang:latest ve kendi
  Dockerfile'ını taşırsa, hangi sürümün nerede olduğunu kimse bilemez.
  Ortak, pinlenmiş bir base image ve paylaşılan CI şablonu gerekir.
- **Aynı hatalar kopyalanır.** Buradaki selector bug'ı gibi bir hata bir
  şablondan 60 repoya yayılırsa, tek bir yanlış tüm servisleri aynı anda
  çökertebilir.
- **Flaky testler CI'yı kilitler.** "Bazen geçen" testler 60 repoda sürekli
  kırmızı CI demek; determinism en baştan zorunlu olmalı.

- **Manuel release ve kopyalanan pipeline tutmaz.** Şu an release'i elle
  tetikliyoruz ve her repo kendi Jenkinsfile'ını taşıyor. 60 repoda bu
  imkânsız; ortak bir reusable pipeline (Jenkins Shared Library gibi)
  olmalı ki her servis aynı standart CI'ı kullansın, tek yerden güncellensin.

- **Görünürlük** kaybolur. 60 servisin
  hangisi ayakta, hangisi hangi sürümde, prod git'le eşleşiyor mu takibi zorlaşır.
  Bunun için GitOps (ArgoCD/Flux) ile deploy'u git'e bağlamak — hangi sürümün
  nerede olduğu git'ten görünür, drift otomatik yakalanır — ve servis
  sağlığı için merkezi observability (metrics/log toplama) gerekir.




## AI kullanımı

DevOpsum ve Go günlük işimin parçası değil. Bu case'teki Go tarafını
(flaky testin sebebi, reproducible build ayarları) ve çözümleri bir AI
asistanıyla (Claude) anlayarak ilerledim. Her değişikliğin neden gerektiğini
kavrayarak yaptım, körü körüne kopyalamadım. Kararlar bana ait; AI'ı
"bu neden böyle" diye sorduğum bir yardımcı gibi kullandım.