
## [v2.7.2](https://github.com/ArtalkJS/Artalk/compare/v2.7.1...v2.7.2) (2023-12-20)

### Features

* patch `locale` config from `zh-cn` to `zh-CN` ([#678](https://github.com/ArtalkJS/Artalk/issues/678))

### Bug Fixes

* **sidebar:** sidebar language switch and config save issue ([#708](https://github.com/ArtalkJS/Artalk/issues/708))
* **ui:** add `referrerpolicy` attribute for iframe ([#687](https://github.com/ArtalkJS/Artalk/issues/687)) ([#707](https://github.com/ArtalkJS/Artalk/issues/707))
* **ui/dark-mode:** potential memory leak issue in auto mode ([#688](https://github.com/ArtalkJS/Artalk/issues/688))
* **ui/pagination:** auto switch page by url hashtag issue ([#693](https://github.com/ArtalkJS/Artalk/issues/693))
* **ui/sidebar:** refactor settings and fix save issue ([#677](https://github.com/ArtalkJS/Artalk/issues/677)) ([#706](https://github.com/ArtalkJS/Artalk/issues/706))

### Code Refactoring

* **ui:** move window references into function scope ([#675](https://github.com/ArtalkJS/Artalk/issues/675))
* **ui/api:** losing coupling between `Api` and `User`
* **ui/user:** user class no longer globally singleton

### Documentation

* **i18n:** add translation for `open` button
* **refactor:** organize the pkgs and fix some issues ([#702](https://github.com/ArtalkJS/Artalk/issues/702))


## [v2.7.1](https://github.com/ArtalkJS/Artalk/compare/v2.7.0...v2.7.1) (2023-12-17)

### Bug Fixes

* **build:** generate source maps to ease debugging
* **cout-widget:** `loadCountWidget` issue
* **ui:** static method export issue in JS module ([#669](https://github.com/ArtalkJS/Artalk/issues/669))

### Documentation

* **i18n:** add translation of update notice


## [v2.7.0](https://github.com/ArtalkJS/Artalk/compare/v2.6.4...v2.7.0) (2023-12-17)

### Features

* **ui:** support `scrollRelativeTo` config option
* **ui/test:** add end-to-end (e2e) testing using Playwright

### Bug Fixes

* **api/img-upload:** `public_path` config enables the use of full url ([#659](https://github.com/ArtalkJS/Artalk/issues/659)) ([#664](https://github.com/ArtalkJS/Artalk/issues/664))
* **lifecycle:** create multi-instances at the same time ([#660](https://github.com/ArtalkJS/Artalk/issues/660)) ([#663](https://github.com/ArtalkJS/Artalk/issues/663))
* **notify_pusher:** panic when admin ids array is empty ([#634](https://github.com/ArtalkJS/Artalk/issues/634))
* **style/list:** replace `float` to `text-align` in footer part ([#619](https://github.com/ArtalkJS/Artalk/issues/619))
* **ui:** scroll to the center of the viewport issue
* **ui/conf:** sanitize `noComment` conf option for security ([#624](https://github.com/ArtalkJS/Artalk/issues/624))
* **ui/dark-mode:** `setDarkMode` cannot save to instance config ([#661](https://github.com/ArtalkJS/Artalk/issues/661))
* **ui/editor:** fix position of the comment box when replying ([#643](https://github.com/ArtalkJS/Artalk/issues/643)) ([#648](https://github.com/ArtalkJS/Artalk/issues/648))

### Performance Improvements

* **cache:** implement GC feature for simple_cache pkg ([#656](https://github.com/ArtalkJS/Artalk/issues/656))

### Code Refactoring

* **editor:** refactor editor plugin manager ([#609](https://github.com/ArtalkJS/Artalk/issues/609))
* **editor:** refactor plug kit and events ([#613](https://github.com/ArtalkJS/Artalk/issues/613))
* **ui:** reduce coupling with `ContextApi`
* **ui:** separate standalone admin-only-elem checker
* **ui/api:** loose coupling between `Api` and `Context`
* **ui/conf:** loose coupling between config and list fetch
* **ui/event:** refactor core event manager ([#611](https://github.com/ArtalkJS/Artalk/issues/611))
* **ui/layer:** better layer implements and independence ([#662](https://github.com/ArtalkJS/Artalk/issues/662))
* **ui/list:** separate list into standalone components ([#618](https://github.com/ArtalkJS/Artalk/issues/618))
* **ui/marked:** separate markdown related codes
* **ui/marked:** loose coupling with `marked` func
* **ui/plugin:** further divide functionality into plugins ([#615](https://github.com/ArtalkJS/Artalk/issues/615))
* **ui/stat:** losing coupling of `CountWidget`
* **ui/types:** better `ArtalkType` import and export ([#620](https://github.com/ArtalkJS/Artalk/issues/620))

### Documentation

* revise and add more examples
* provide clearer and more detailed instructions in CONTRIBUTING.md
* **conf:** disable `frontend.uaBadge` config option by default
* **deploy:** add deployment guide for render.com ([#649](https://github.com/ArtalkJS/Artalk/issues/649))
* **import:** update the examples in import-framework.md ([#665](https://github.com/ArtalkJS/Artalk/issues/665))

### BREAKING CHANGE


The following top-level functions exported by the 'artalk' npm package have been deprecated: `Artalk.update`, `Artalk.reload`, and `Artalk.destroy`. These methods now require invocation on an instance created by either `Artalk.init` or `new Artalk`. Please utilize instance-level methods instead, such as `artalkInstance.update`. For more information, refer to  [the documentation](https://artalk.js.org/guide/frontend/import-framework.html). The update was implemented to enable the concurrent creation of multiple instances, adapting to situations where Vue components are simultaneously referenced across various pages. Initially, only a singular instance was permitted to mitigate memory leak concerns. However, this proved limiting for scenarios involving the caching of multiple component instances through 'keep-alive.' To better suit intricate SPA application needs, the choice was made to permit the creation of multiple independent instances. It's crucial to remember to manually invoke the `artalk.destroy` method when releasing components to avoid memory leaks (Issue [#660](https://github.com/ArtalkJS/Artalk/issues/660)).


## [v2.6.4](https://github.com/ArtalkJS/Artalk/compare/v2.6.3...v2.6.4) (2023-10-12)

### Features

* **go:** upgrade golang to v1.21.3 and some deps

### Bug Fixes

* **editor:** fix duplicate event binding when update conf ([#605](https://github.com/ArtalkJS/Artalk/issues/605))
* **email:** sender not initialized correctly
* **style:** fix copyright text overlay on send button ([#607](https://github.com/ArtalkJS/Artalk/issues/607))

### Code Refactoring

* **anti_spam:** refactor anti_spam pkg and add unit test ([#589](https://github.com/ArtalkJS/Artalk/issues/589))


## [v2.6.3](https://github.com/ArtalkJS/Artalk/compare/v2.6.2...v2.6.3) (2023-09-20)

### Features

* **config:** add `db.ssl` to enable db ssl mode ([#587](https://github.com/ArtalkJS/Artalk/issues/587))

### Bug Fixes

* **notify_tpl:** fix template render result in notify case

### Performance Improvements

* **config:** add error message for timezone setting ([#586](https://github.com/ArtalkJS/Artalk/issues/586))

### Code Refactoring

* **renderer:** abstract func of template renderer for multi-cases ([#585](https://github.com/ArtalkJS/Artalk/issues/585))

### Documentation

* new landing page


## [v2.6.2](https://github.com/ArtalkJS/Artalk/compare/v2.6.1...v2.6.2) (2023-09-12)

### Bug Fixes

* **cmd:** remove app instance creation when mini boot ([#577](https://github.com/ArtalkJS/Artalk/issues/577))

### Performance Improvements

* **style:** modify `z-index` less than 5 ([#578](https://github.com/ArtalkJS/Artalk/issues/578))


## [v2.6.1](https://github.com/ArtalkJS/Artalk/compare/v2.6.0...v2.6.1) (2023-09-11)

### Bug Fixes

* **captcha:** image captcha base64 response issue ([#575](https://github.com/ArtalkJS/Artalk/issues/575))

### Documentation

* **config:** add info about env variables ([#566](https://github.com/ArtalkJS/Artalk/issues/566))
* **extras:** add deploy guide for railway.app ([#567](https://github.com/ArtalkJS/Artalk/issues/567))


## [v2.6.0](https://github.com/ArtalkJS/Artalk/compare/v2.5.5...v2.6.0) (2023-09-01)

### Features

* **config:** support load env variables as config ([#564](https://github.com/ArtalkJS/Artalk/issues/564))
* **go:** upgrade golang to v1.21.0 and some deps ([#543](https://github.com/ArtalkJS/Artalk/issues/543))
* **hook:** add hook pkg
* **simple_cache:** a simple thread-safe cache with sync.Map
* **utils:** add common random utils

### Bug Fixes

* **cache:** child comment ids store after comment updated
* **cmd:** fix parse root cmd global flags
* **config:** fix ip region patch
* **db:** fix dsn strings for mysql connections ([#541](https://github.com/ArtalkJS/Artalk/issues/541))
* **limiter:** limiter log ignore options method req
* **limiter:** fix limiter always mode
* **log:** add missing log for web server listen func ([#556](https://github.com/ArtalkJS/Artalk/issues/556))
* **ui:** fix plug cannot disable issue ([#563](https://github.com/ArtalkJS/Artalk/issues/563))

### Performance Improvements

* **cmd:** support command mini boot mode

### Code Refactoring

* **CI:** replace using `docker/metadata-action` for the docker push ([#545](https://github.com/ArtalkJS/Artalk/issues/545))
* **CI:** replace renovate with github dependabot ([#546](https://github.com/ArtalkJS/Artalk/issues/546))
* **anti_spam:** refactor anti_spam pkg
* **artransfer:** refactor artransfer pkg
* **cache:** remove cached fields in struct to save memory ([#555](https://github.com/ArtalkJS/Artalk/issues/555))
* **cache:** cache code tidy-up
* **cache:** refactor cache pkg
* **cache:** refactor cache pkg
* **cache:** refactor cache pkg
* **captcha:** refactor captcha pkg
* **cloud:** refactor cloud pkg
* **cmd:** refactor cmd pkg
* **conf:** better config.NewFromFile with error handle
* **config:** refactor config pkg
* **core:** refactor core pkg
* **core:** better AppService func with error handle
* **cors:** refactor cors middleware
* **dao:** refactor dao pkg
* **dao_cache:** add dao_cache pkg
* **db:** refactor db pkg
* **email:** refactor email pkg
* **i18n:** refactor i18n pkg
* **ip_region:** refactor ip_region pkg
* **limiter:** add limiter pkg
* **log:** create common log pkg
* **main:** refactor main.go
* **notify:** add notify_pusher pkg
* **server/common:** refactor server/common pkg
* **server/handler:** refactor server/handler pkg
* **server/middleware:** refactor server/middleware pkg
* **test:** create common test pkg

### Documentation

* **extras:** add deploy guide for fly.io ([#520](https://github.com/ArtalkJS/Artalk/issues/520))


## [v2.5.5](https://github.com/ArtalkJS/Artalk/compare/v2.5.4...v2.5.5) (2023-05-10)

### Features

* **build:** add embedded fe lite and i18n scripts ([#493](https://github.com/ArtalkJS/Artalk/issues/493)) ([#505](https://github.com/ArtalkJS/Artalk/issues/505))
* **go:** upgrade golang to v1.20.4 and some deps ([#506](https://github.com/ArtalkJS/Artalk/issues/506))
* **gravatar:** add `gravatar.params` config option ([#508](https://github.com/ArtalkJS/Artalk/issues/508))

### Bug Fixes

* **api:** comment IP data loss after editing ([#504](https://github.com/ArtalkJS/Artalk/issues/504))
* **editor:** prevent editor reset on comment failure ([#507](https://github.com/ArtalkJS/Artalk/issues/507))

### BREAKING CHANGE


The `gravatar.default` config option has been removed. Please use `gravatar.params` instead. The default value for `gravatar.params` is now `"d=mp&s=240"`.


## [v2.5.4](https://github.com/ArtalkJS/Artalk/compare/v2.5.3...v2.5.4) (2023-04-09)

### Features

* **go:** upgrade golang to v1.20.3 and some deps ([#485](https://github.com/ArtalkJS/Artalk/issues/485))

### Bug Fixes

* **CI:** upgrade and pin node.js and pnpm version ([#484](https://github.com/ArtalkJS/Artalk/issues/484))

### Documentation

* update swagger docs and some deps


## [v2.5.3](https://github.com/ArtalkJS/Artalk/compare/v2.5.2...v2.5.3) (2023-04-06)

### Bug Fixes

* **api/comment_add:** SIGSEGV crash in comment_add ([#478](https://github.com/ArtalkJS/Artalk/issues/478)) ([#481](https://github.com/ArtalkJS/Artalk/issues/481))

### Performance Improvements

* **ip_region:** add db file existence check and error message ([#482](https://github.com/ArtalkJS/Artalk/issues/482))


## [v2.5.2](https://github.com/ArtalkJS/Artalk/compare/v2.5.1...v2.5.2) (2023-03-19)

### Bug Fixes

* **ui:** destroy instance when calling init func twice

### Performance Improvements

* **notify:** shorten the length of notify key ([#466](https://github.com/ArtalkJS/Artalk/issues/466))
* **style:** add `!important` for white-space css in code block ([#467](https://github.com/ArtalkJS/Artalk/issues/467))


## [v2.5.1](https://github.com/ArtalkJS/Artalk/compare/v2.5.0...v2.5.1) (2023-03-16)

### Features

* **go:** upgrade golang to v1.20.2

### Bug Fixes

* **ui/count-widget:** context api undefined issue ([#464](https://github.com/ArtalkJS/Artalk/issues/464))
* **ui/i18n:** duplicate packaging built-in language in the external script
* **ui/paginator:** showErr func call issue
* **ui/sort-dropdown:** dropdown menu disappears after call reload func ([#461](https://github.com/ArtalkJS/Artalk/issues/461))

### Code Refactoring

* **anti_spam/aliyun:** accessing aliyun green text api without sdk ([#459](https://github.com/ArtalkJS/Artalk/issues/459))
* **ui/user:** modify user to standalone module ([#463](https://github.com/ArtalkJS/Artalk/issues/463))


## [v2.5.0](https://github.com/ArtalkJS/Artalk/compare/v2.4.4...v2.5.0) (2023-03-10)

### Features

* migrate from `echo` to `go-fiber`
* upgrade go to v1.20.1
* display IP region of comment ([#418](https://github.com/ArtalkJS/Artalk/issues/418)) ([#447](https://github.com/ArtalkJS/Artalk/issues/447))
* docker ci add support for building arm64 wheels
* more functions to handle artalk lifecycle ([#426](https://github.com/ArtalkJS/Artalk/issues/426))
* **ui:** add some static methods
* **ui/height_limit:** support scrollable height limit area ([#451](https://github.com/ArtalkJS/Artalk/issues/451))
* **ui/sidebar:** add dark mode support ([#450](https://github.com/ArtalkJS/Artalk/issues/450))
* **captcha:** add support for reCAPTCHA and hCaptcha ([#456](https://github.com/ArtalkJS/Artalk/issues/456))
* **captcha:** support turnstile captcha by cloudflare ([#453](https://github.com/ArtalkJS/Artalk/issues/453))
* **i18n:** add i18n support for backend ([#343](https://github.com/ArtalkJS/Artalk/issues/343))
* **i18n:** translations for backend ([#344](https://github.com/ArtalkJS/Artalk/issues/344))
* **i18n:** add i18n support for sidebar ([#353](https://github.com/ArtalkJS/Artalk/issues/353))
* **i18n:** add zh-TW i18n translation for sidebar and app

### Bug Fixes

* **ui:** hash goto function check condition issue
* **ui/conf:** avoid some conf overrides frontend from the backend ([#449](https://github.com/ArtalkJS/Artalk/issues/449))
* **ui/editor:** disable img upload cannot hide its btn
* **ui/i18n:** subscribe event priority issue
* **ui/sidebar:** array type of preference initial data issue
* **ui/sidebar:** array type config option cannot be shown
* **ui/sidebar:** boolean type setting option save issue ([#431](https://github.com/ArtalkJS/Artalk/issues/431)) ([#444](https://github.com/ArtalkJS/Artalk/issues/444))
* **ui/sidebar:** setting item save follow type of template
* **lint:** add tsc check before vite compile ([#440](https://github.com/ArtalkJS/Artalk/issues/440))
* **email:** duplicate sending with multiple admins using same email addrs ([#375](https://github.com/ArtalkJS/Artalk/issues/375))
* **email:** email queue initialization issue ([#374](https://github.com/ArtalkJS/Artalk/issues/374))
* **email:** failback to `email.mail_tpl` if `admin_notify.email.mail_tpl` is empty
* `timeAgo` function does not display the now
* add `.npmignore` to fix NPM publish inclusion issue
* sidebar navigation sorting ([#361](https://github.com/ArtalkJS/Artalk/issues/361))

### Performance Improvements

* improve some css styles
* add graceful shutdown
* **conf/i18n:** detect and change locale when config file contains chinese
* **ui/list:** remove useless function call
* **ui/sidebar:** improve sidebar i18n

### Code Refactoring

* bump to monorepo
* renamed from artalk-go to artalk
* http origin checker
* abstract email service
* project package structure
* remove version two-way check ([#452](https://github.com/ArtalkJS/Artalk/issues/452))
* build scripts and CI tests
* replace pkger with go:embed
* launch with vscode debugger
* **CI:** one-key site creating with artalk integrated
* **CI:** improve build and release workflows ([#358](https://github.com/ArtalkJS/Artalk/issues/358))
* **captcha:** abstract captcha service ([#455](https://github.com/ArtalkJS/Artalk/issues/455))
* **comment:** separate comment ui renders from single file ([#427](https://github.com/ArtalkJS/Artalk/issues/427))
* **style:** convert to use Sass as a style interpreter ([#439](https://github.com/ArtalkJS/Artalk/issues/439))
* **ui:** automatic dependency injection ([#429](https://github.com/ArtalkJS/Artalk/issues/429))
* **ui/checker:** simplify checker lifecycle function param table ([#428](https://github.com/ArtalkJS/Artalk/issues/428))
* **ui/dark-mode:** separate dark mode logic into its own module ([#430](https://github.com/ArtalkJS/Artalk/issues/430))
* **ui/editor:** modify editor ui to standalone module ([#441](https://github.com/ArtalkJS/Artalk/issues/441))
* **ui/editor:** change functions of editor to standalone modules ([#443](https://github.com/ArtalkJS/Artalk/issues/443))
* **ui/height-limit:** modify height limit function to standalone module ([#435](https://github.com/ArtalkJS/Artalk/issues/435))
* **ui/i18n:** improve i18n function to standalone module ([#434](https://github.com/ArtalkJS/Artalk/issues/434))
* **ui/list:** modify list pagination to standalone module ([#437](https://github.com/ArtalkJS/Artalk/issues/437))
* **page/fetch:** remove goquery dependency when extracting page data ([#442](https://github.com/ArtalkJS/Artalk/issues/442))
* **anti_spam/qcloud-tms:** implement qcloud tms api without sdk ([#438](https://github.com/ArtalkJS/Artalk/issues/438))

### Documentation

* migrate ArtalkJS/Docs to monorepo docs
* add `CODE_OF_CONDUCT.md`
* add simplified README for artalk npm package
* add translation section to `CONTRIBUTING.md`
* fix wrong config value ([#371](https://github.com/ArtalkJS/Artalk/issues/371))
* fix broken links ([#364](https://github.com/ArtalkJS/Artalk/issues/364))
* add open api ([#360](https://github.com/ArtalkJS/Artalk/issues/360))
* add `Project Structure` section to `CONTRIBUTING.md`
* refine and add frontend api docs
* update setup-example-site.sh script usage
* init artalk with new frontend api
* **deploy:** add `restart=always` for docker to auto restart ([#425](https://github.com/ArtalkJS/Artalk/issues/425))
* **extras:** add deploy guide for vuepress ([#436](https://github.com/ArtalkJS/Artalk/issues/436))
