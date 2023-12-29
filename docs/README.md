# Gothrix CMS/CMF

## В разработке!

### Цели
- единая точка входа
- слой промежуточного ПО
- компоненты для вывода данных с возможностью повторного применения
- простая "DB first" ОРМ
- расширяемость через модули
- СЕО ориентированность (SSR в комплекте)
- аутентификация/авторизация в комплекте
- кэш
- очереди задач
- электронная почта
- логирование
- готовый раздел администрирования

### Стек технологий
- [`golang`](https://go.dev/) - компилируемый многопоточный язык программирования
- [`echo`](https://echo.labstack.com/) - каркас, роутинг, слой промежуточного ПО
- [`sqlc`](https://sqlc.dev/) - простая "DB first" ОРМ, генерация типобезопасного кода из SQL
- [`templ`](https://templ.guide/) - шаблонизатор, создание HTML с помощью Go
- [`htmx`](https://htmx.org/) - динамический HTML, возможность реализовать SPA+SSR на Go(обертка над AJAX)
- [`alpinejs`](https://alpinejs.dev/) - реактивный HTML, легкий JS фреймворк
- [`purecss`](https://purecss.io/) - легкий CSS каркас
- [`mysql`](https://www.mysql.com/) - реляционная база данных
- [`redis`](https://redis.io/) - резидентная база данных (для сессий, кэша, очередей и т.п.)


### Структура проекта

```
cmd/                        Основные приложения для текущего проекта
    web/                    Web приложение
        main.php      
    console/                Приложение консольных команд
        main.php      
components/                 Переиспользуемые компоненты сайта
config/                     Файлы конфигурации
    config.php              Основной конфиг
docs/                       Документация
internal/                   Внутренний код приложения и библиотек
    middleware/ 
        session/   
    tasks/                  Задания для очереди задач 
    hooks/                  Хуки(события) используемые для этого какого то модуля   
    common_components/      Общие компоненты (шаблоны и представления)
modules/                    Модули
    module_name/            Исходный код модуля
        components/         Компоненты (шаблоны и представления) конкретного модуля    
        handlers/           HTTP обработчики (действия, промежуточное ПО и т. д.)
        helpers/            Вспомогательные структуры (фабрики, слушатели и т. д.)
            db/             Sql запросы для sqlc
        models/             Структуры модели предметной области (сущности, репозитории и т. д.)
        module_name.go      Основной файл модуля
        services/           Сервисный слой моудля используемый в обработчиках
        route.go            Файл с маршрутизацией модуля
static/                     

```

```mermaid
graph LR
    task{{Task}} --> queue
    queue --> services
    web{{Web}} --> router
    api{{API}} --> router
    hooks <--> services
    subgraph Module
        router --> handler
        handler --> services
        services --> models[models]
        handler -- renders --> components
    end
    models --> mysql[(MySQL)]
    external_service{{External Service}} --> hooks
    handler --> JSON[\JSON\]
    components[\components\]  --> common_components[\Common components\]
```

### Потенциальные модули
> **Модуль** - это блок, отвечающий за определенную функциональность, содержит API доступа к данным, бизнес-логику и события. Каждый модуль имеет метод инициализации.

#### Сайт
-   `site` - модуль сайта который использует остальные модули, если надо, для реализации сайта или свой функционал

#### Админка
-   `user` - пользователи, аутентификация
-   `admin` - настройки и т.п., зависит от user
-   `rbac` - авторизация по схеме роль-модуль-действие, зависит от user и admin


#### Контент
-   `content` - генерация сущностей page+category, если торговля - то сущности по факту это продукт + категория
-   `media` - работа с файлами, от него может зависеть модуль content
-   `seo` - метаданные страниц, от него может зависеть модуль content
-   `property` - зависит от content - заведение свойств для контентной сущности
-   `ffilter` - фильтр по свойствам, зависит от content и property, если есть price или stock, то + они
-   `comment` - комментарии, зависит от content и user


#### Торговля

-   `customer` - клиент, зависит от user
-   `wish` - зависит от content и customer
-   `compare` - зависит от content и customer
-   `price` - цена, зависит от content
-   `stock` - запас, зависит от price и content
-   `сart` - корзина, зависит от customer и price (если есть stock, то можно учитывать количество на складе)
-   `order` - заказ, зависит от customer и сart
-   `sale` - скидки у заказов, зависит от order
-   `payment` - платежные системы у заказов, зависит от order
-   `delivery` - доставки у заказов, зависит от order

```mermaid
graph LR
    subgraph Admin
        admin --> user
        rbac --> user
        rbac --> admin
    end
    subgraph Content
        content --> media
        content --> seo
        property --> content
        ffilter --> content
        ffilter --> property
        comment --> content
        comment --> user
    end
    subgraph Сommerce
        customer --> user
        wish --> customer
        wish --> content
        compare --> customer
        compare --> content
        price --> content
        stock --> content
        stock --> price
        сart --> customer
        сart --> price
        сart --> stock
        order --> customer
        order --> сart
        sale --> order
        payment --> order
        delivery --> order
    end
```