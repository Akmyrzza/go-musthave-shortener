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
