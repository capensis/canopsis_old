module git.canopsis.net/canopsis/canopsis-community/community/go-engines-community

// Needs to be synced with Makefile:DOCKER_GOLANG_VERSION and go.mod from Pro
// Only two integers are allowed here (https://golang.org/ref/mod#go-mod-file-go)
go 1.17

// Note: External libs under GPL, AGPL or other viral licenses are not allowed here.
// Canopsis Pro contains Canopsis Community, and Canopsis Pro can't contain viral
// licenses, since it's a proprietary product.
//
// Please always maintain this file with the rules described here:
// https://git.canopsis.net/canopsis/canopsis-pro/-/issues/590

require (
	// No GPL or AGPL libs allowed below!

	github.com/ajg/form v1.5.1
	github.com/alecthomas/participle v0.7.1
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/beevik/etree v1.1.0
	github.com/bsm/redislock v0.7.1
	github.com/casbin/casbin/v2 v2.36.1
	github.com/cucumber/godog v0.12.0
	github.com/dlclark/regexp2 v1.4.0
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ldap/ldap/v3 v3.4.1
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/go-redis/redis/v8 v8.11.3
	github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.6
	github.com/google/uuid v1.3.0
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.1
	github.com/gorilla/websocket v1.4.2
	github.com/json-iterator/go v1.1.11
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/pelletier/go-toml v1.9.3
	github.com/rs/zerolog v1.23.0
	github.com/russellhaering/gosaml2 v0.6.0
	github.com/russellhaering/goxmldsig v1.1.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/streadway/amqp v1.0.0
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.3.1
	github.com/swaggo/swag v1.6.7
	github.com/teambition/rrule-go v1.7.1
	github.com/tidwall/pretty v1.1.0 // indirect
	github.com/valyala/fastjson v1.6.3
	go.mongodb.org/mongo-driver v1.7.1
	golang.org/x/net v0.0.0-20211013171255-e13a2654a71e // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20211013075003-97ac67df715c // indirect
	golang.org/x/text v0.3.7
	golang.org/x/tools v0.1.5 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

require github.com/kylelemons/godebug v1.1.0

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/jackc/pgx/v4 v4.13.0
	github.com/klauspost/compress v1.13.6 // indirect
)

require (
	github.com/go-testfixtures/testfixtures/v3 v3.6.1
	github.com/golang-migrate/migrate/v4 v4.15.1
	github.com/jackc/pgconn v1.10.0
	github.com/jackc/pgproto3/v2 v2.1.1
)

require (
	github.com/Azure/go-ntlmssp v0.0.0-20200615164410-66371956d46c // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/cucumber/gherkin-go/v19 v19.0.3 // indirect
	github.com/cucumber/messages-go/v16 v16.0.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-asn1-ber/asn1-ber v1.5.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181017120253-0766667cb4d1 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.0 // indirect
	github.com/hashicorp/go-memdb v1.3.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgerrcode v0.0.0-20201024163028-a0d42d470451 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.8.1 // indirect
	github.com/jackc/puddle v1.1.3 // indirect
	github.com/jonboulle/clockwork v0.2.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.0.0-20201208211235-fe770d50d911 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/smartystreets/assertions v0.0.0-20180927180507-b2de0cb4f26d // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/ugorji/go/codec v1.1.13 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.uber.org/atomic v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

// No effect on the real canopsis-community repo, but necessary when it's part of the canopsis-pro monorepo
replace git.canopsis.net/canopsis/canopsis-community/community/go-engines-community => ./
