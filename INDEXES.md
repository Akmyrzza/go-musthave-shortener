# Полезные ссылки с курса Yandex

# Спринт 1

## Пакет net/http. Работа с HTTP

### Структура проекта

[go.dev | Go Modules Reference. Vendoring](https://pkg.go.dev/net/http) — о директории vendor на официальном сайте Go.

[Dave Cheney's Blog | Use internal packages to reduce your public API surface](https://dave.cheney.net/2019/10/06/use-internal-packages-to-reduce-your-public-api-surface) — зачем нужна директория internal.

[GitHub | Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ru.md) — рекомендации по структуре проектов от Go-сообщества.

### Создание HTTP-сервера

[go.dev/net/http](https://pkg.go.dev/net/http) — документация пакета net/http.

[go.dev | Writing Web Applications](https://go.dev/doc/articles/wiki/) — о создании веб-приложений.

[GitHub | Build web application with golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/ru/preface.md) — о веб-приложениях от авторитетного Go-разработчика @astaxie, автора библиотеки beego.

[IANA | HTTP Status Code Registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) — список кодов статуса.

### Тестирование HTTP-приложения

[go.dev/testing](https://pkg.go.dev/testing) — документация пакета testing.

[go.dev/net/http/httptest](https://pkg.go.dev/net/http/httptest) — документация пакета httptest.

[go.dev | Testing flags](https://pkg.go.dev/cmd/go#hdr-Testing_flags) — о команде go test и флагах.

[GitHub | Testify — Thou Shalt Write Tests](https://github.com/stretchr/testify) — библиотека testify.

[Ilija Eftimov | Testing in Go: go test](https://ieftimov.com/posts/testing-in-go-go-test/) — о тестировании в Go.

[YouTube | Unit testing HTTP servers](https://www.youtube.com/watch?v=hVFEV-ieeew) — выпуск “justforfunc: Programming in Go” о тестировании сервера.

### Использование HTTP-клиента

[go.dev/net/http/cookiejar](https://pkg.go.dev/net/http/cookiejar) — документация пакета cookiejar.

[go.dev | Type Client](https://pkg.go.dev/net/http#Client) — описание HTTP-клиента.

[Golang By Example | Cookies in Go](https://golangbyexample.com/cookies-golang/) — о куках в Go.

### Выбор HTTP-библиотеки

[go.dev | resty](https://pkg.go.dev/github.com/go-resty/resty/v2) — документация пакета resty.

[go.dev | chi](https://pkg.go.dev/github.com/go-chi/chi/v5) — документация роутера chi.

[go.dev | middleware](https://pkg.go.dev/github.com/go-chi/chi/middleware) — список доступных middleware для роутера chi.

[GitHub | HTTP clients](https://github.com/avelino/awesome-go#http-clients) — список клиентских пакетов на Awesome Go.

[GitHub | Routers](https://github.com/avelino/awesome-go#routers) — список серверных пакетов на Awesome Go.

[JSONPlaceholder](https://jsonplaceholder.typicode.com) — бесплатный тестовый API.

## Пакет flag. Чтение аргументов командной строки

### Аргументы командной строки

[go.dev/flag](https://pkg.go.dev/flag) — документация пакета flag.

[GitHub | Cobra](https://github.com/spf13/cobra) — о библиотеке Cobra.

[GitHub | pflag](https://github.com/spf13/pflag) — о пакете pflag.

[GNU.org | Program Argument Syntax Conventions](https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html) — рекомендации GNU.

## Пакет os. Получение переменных окружения

### Переменные окружения

[go.dev/os](https://pkg.go.dev/os) — документация пакета os.

[Losst | Переменные окружения в Linux](https://losst.pro/peremennye-okruzheniya-v-linux) — подробная статья о переменных окружения.

[GitHub | caarlos0/env](https://github.com/caarlos0/env) — пакет caarlos0/env.

[GitHub | kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) — пакет kelseyhightower/envconfig.

# Спринт 2

## Пакет log. Логирование в приложении

### Стандартные и сторонние пакеты для логирования

[go.dev/log](https://pkg.go.dev/log) — документация пакета log.

[go.dev/go.uber.org/zap](https://pkg.go.dev/go.uber.org/zap) — документация пакета zap.

[GitHub | zap](https://github.com/uber-go/zap) — пакет zap.

[GitHub | logrus](https://github.com/Sirupsen/logrus) — пакет logrus.

[GitHub | Awesome Go | Logging](https://github.com/avelino/awesome-go#logging) — библиотеки для логирования.

## Пакет encoding. Сериализация и десериализация данных.

### Основы REST API

[Хабр | Архитектура REST](https://habr.com/ru/articles/38730/) — статья о стиле REST.

[Хабр | Hypermedia](https://habr.com/ru/companies/aligntechnology/articles/281206/) — то, без чего ваше API не совсем REST — подробнее о гиперссылках в сообщениях.

[REST API Tutorial](https://restapitutorial.com/resources.html) — лучшие практики в разработке RESTful-сервисов.

### Структурные теги

[go.dev/reflect](https://pkg.go.dev/reflect) — документация пакета reflect.

[go.dev/encoding/json](https://pkg.go.dev/encoding/json) — документация пакета json, который использует структурные теги для определения параметров JSON-преобразования.

[GitHub | Gorm](https://github.com/jinzhu/gorm) — ORM-библиотека, которая использует структурные теги для указания SQL-связей.

[GitHub | Govalidator](https://github.com/asaskevich/govalidator) — пакет для валидации значений полей на основе описания структурных тегов.

### Стандартные сериализаторы: JSON

[go.dev/encoding/json](https://pkg.go.dev/encoding/json) — документация пакета encoding/json.

[Developer 2.0 | JSON in Go](https://developer20.com/json/) — статья о работе с JSON в Go.

### Стандартные сериализаторы: XML и gob

[go.dev/encoding/xml](https://pkg.go.dev/encoding/xml) — документация пакета encoding/xml.

[go.dev/encoding/xml#Token](https://pkg.go.dev/encoding/xml#Token) — описание XML-токена.

[go.dev/encoding/gob](https://pkg.go.dev/encoding/gob) — документация пакета encoding/gob.

[XMLBeans | Understanding XML Tokens](https://xmlbeans.apache.org/docs/3.0.0/guide/conUnderstandingXMLTokens.html) — статья о том, что такое XML-токен на примере Java.

### Сторонние сериализаторы

[go-yaml](https://pkg.go.dev/gopkg.in/yaml.v3) — пакет для работы с YAML-форматом.

[yaml.org](https://yaml.org/spec/1.2.2/) — спецификация формата YAML.

[toml.io](https://toml.io/en/) — спецификация формата TOML.

[go-toml](https://github.com/pelletier/go-toml) — пакет для работы с TOML-форматом.

[easyjson](https://github.com/mailru/easyjson) — пакет для генерации кода работы с JSON-форматом без использования рефлексии.

[MessagePack](https://msgpack.org) — описание и примеры использования бинарного формата MessagePack.

[msgp](https://github.com/tinylib/msgp) — реализация формата MessagePack для Go.

[Protocol Buffers](https://protobuf.dev) — описание формата Protocol Buffers.

[proto3](https://protobuf.dev/programming-guides/proto3/) — спецификация формата proto3.

[gRPC](https://grpc.io) — описание фреймворка gRPC для реализации RPC-протокола общения между приложениями с форматом данных Protocol Buffers.

[protobuf-go](https://github.com/protocolbuffers/protobuf-go) — пакет для работы с Protocol Buffers для Go.

## Пакет compress. Сжатие данных

### Оптимизация передачи данных

[Хабр | Интерфейсные типы io.Reader и io.Writer](https://habr.com/ru/articles/306914/) — статья об интерфейсных типах.

[Echo | Gzip Middleware](https://echo.labstack.com/docs/middleware/gzip) — пример реализации gzip в HTTP-сервере.

[GitHub | Zstandard](https://facebook.github.io/zstd/) — алгоритм сжатия Zstandard.

[GitHub | NYTimes/gziphandler](https://pkg.go.dev/github.com/NYTimes/gziphandler) — пакет gziphandler.

[GitHub | mholt/archiver](https://github.com/mholt/archiver) — пакет mholt/archiver для работы с архивами.

[GitHub | Squash Compression Benchmark](https://quixdb.github.io/squash-benchmark/) — Squash Compression Benchmark.

[MDN Web Docs | Accept-Encoding](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding) — спецификация заголовка Accept-Encoding.

[MDN Web Docs | Content-Encoding](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding) — спецификация заголовка Content-Encoding.

## Пакет os. Операции с файлами

### Чтение и запись в файл

[go.dev/os](https://pkg.go.dev/os) — документация пакета os.

[go.dev/bufio](https://pkg.go.dev/bufio) — документация пакета bufio.

[The Linux Foundation | Classic SysAdmin: Understanding Linux File Permissions](https://www.linuxfoundation.org/blog/blog/classic-sysadmin-understanding-linux-file-permissions) — модификаторы доступа в Linux.

# Спринт 3

## Пакеты time, context. Отмена операций и управление временем выполнения

### Время: даты, интервалы, таймеры

[go.dev/time](https://pkg.go.dev/time) — документация пакета time.

### Контекст: отмена операций

[go.dev/context](https://pkg.go.dev/context) — документация пакета context.

[go.dev/github.com/jmoiron/sqlx#DB.SelectContext](https://pkg.go.dev/github.com/jmoiron/sqlx#DB.SelectContext) — контекст в запросах к БД на примере пакета sqlx.

[go.dev/net/http#NewRequestWithContext](https://pkg.go.dev/net/http#NewRequestWithContext) — контекст в HTTP-запросах на примере пакета net/http.

[go.dev/runtime/trace](https://pkg.go.dev/runtime/trace) — трассировка на примере пакета runtime/trace.

[DigitalOcean | How To Use Contexts in Go](https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go) — об использовании контекста в Go.

[Ильдар Карымов | Разбираемся с контекстами в Go](https://blog.ildarkarymov.ru/posts/context-guide/) — конспект видео про контексты.

[Хабр | Разбираемся с пакетом Context в Golang](https://habr.com/ru/companies/nixys/articles/461723/) — о лучших практиках и подводных камнях при использовании контекста.

## Пакет database/sql. Взаимодействие с базами данных SQL

### Пакет gomock. Имитация баз данных

[go.dev/github.com/golang/mock/gomock](https://pkg.go.dev/github.com/golang/mock/gomock) — документация пакета gomock.

[GitHub | Testing with GoMock: A Tutorial](https://gist.github.com/thiagozs/4276432d12c2e5b152ea15b3f8b0012e) — руководство по тестированию с

### Абстрактный интерфейс и SQL-драйверы

[go.dev/database/sql](https://pkg.go.dev/database/sql) — документация пакета database/sql.

[GitHub | SQLDrivers](https://github.com/golang/go/wiki/SQLDrivers) — список доступных драйверов.

[Хабр | Go и MySQL: настраиваем пул соединений](https://habr.com/ru/companies/oleg-bunin/articles/583558/) — статья о MySQL.

[GolangBot | MySQL Tutorial: Connecting to MySQL and Creating a DB using Go](https://golangbot.com/connect-create-db-mysql/) — руководство по работе с MySQL.

[Go database/sql tutorial](https://go-database-sql.org/) — популярный самоучитель по работе с БД в Go.

[Soham Kamani | Using an SQL Database in Go](https://www.sohamkamani.com/golang/sql-database/) — о работе с БД SQL в Go.

[GoLinuxCloud | Golang SQLite3 Tutorial](https://www.golinuxcloud.com/golang-sqlite3/) — руководство по работе с SQLite с примерами.

[Ardan Labs | Working with SQLite using Go and Python](https://www.ardanlabs.com/blog/2020/11/working-with-sqlite-using-go-python.html) — о работе с SQLite.

[YouTube | Go Northwest | David Crawshaw SQLite and Go](https://www.youtube.com/watch?v=RqubKSF3wig&ab_channel=GoNorthwest) — видео о Go и SQLite.

### Запросы к базе данных

[SQL-запросы быстро.](https://habr.com/ru/articles/480838/)

[Querying for data - The Go Programming Language.](https://go.dev/doc/database/querying)

[Scanners and Valuers with Go.](https://husobee.github.io/golang/database/2015/06/12/scanner-valuer.html)


### Запись в базу данных

[Памятка/шпаргалка по SQL.](https://habr.com/ru/articles/564390/)

[Пакет SQLx.](https://github.com/jmoiron/sqlx)

[GORM.](https://gorm.io/docs/)


## Пакет errors. Обработка ошибок

### Интроспекция ошибок

[go.dev/errors](https://pkg.go.dev/errors) — документация пакета errors.

[bitfieldconsulting | Error wrapping in Go](https://bitfieldconsulting.com/golang/wrapping-errors) — упаковка ошибок в Go.

# Спринт 4

## Пакеты hash, crypto. Безопасность информации

### Хэширование и шифрование

[go.dev/math/rand](https://pkg.go.dev/math/rand) — документация пакета math/rand.

[go.dev/crypto/sha256](https://pkg.go.dev/crypto/sha256) — документация пакета crypto/sha256.

[go.dev/crypto/md5](https://pkg.go.dev/crypto/md5) — документация пакета crypto/md5.

[go.dev/crypto/cipher](https://pkg.go.dev/crypto/cipher) — документация пакета crypto/cipher.

[Cryptanalytic Attacks on Pseudorandom Number Generators](https://www.schneier.com/wp-content/uploads/2017/10/paper-prngs.pdf) — научная статья о видах атак.

[Educative | What is the AES algorithm?](https://www.educative.io/answers/what-is-the-aes-algorithm) — про алгоритм AES.

[Practical Cryptography for Developers | Cipher Block Modes](https://cryptobook.nakov.com/symmetric-key-ciphers/cipher-block-modes) — про алгоритм GCM.

[Practical Cryptography for Developers | HMAC and Key Derivation](https://cryptobook.nakov.com/mac-and-key-derivation/hmac-and-key-derivation) — про алгоритм HMAC.

### Авторизация: JSON Web Token

[go.dev/github.com/dgrijalva/jwt-go](https://pkg.go.dev/github.com/dgrijalva/jwt-go) — документация пакета jwt-go.

[JSON Web Tokens](https://jwt.io/introduction) — официальная страница JWT.

[RSA](https://ru.wikipedia.org/wiki/RSA) — про алгоритм RSA.

[ECDSA](https://ru.wikipedia.org/wiki/ECDSA) — про алгоритм ECDSA.
