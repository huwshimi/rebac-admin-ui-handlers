# Changelog

## [0.1.0](https://github.com/canonical/rebac-admin-ui-handlers/compare/v0.0.1...v0.1.0) (2024-07-30)


### Features

* add `/groups` endpoints handler ([3295880](https://github.com/canonical/rebac-admin-ui-handlers/commit/329588054390d6af3062852d6107dbe5c0a453ff))
* add `/resources` endpoints ([4470127](https://github.com/canonical/rebac-admin-ui-handlers/commit/44701277ca219d88d1bca92627a2f44ac87f1a0a))
* add `/roles` endpoint handlers ([070ad57](https://github.com/canonical/rebac-admin-ui-handlers/commit/070ad576c7807c9819f7be1017c37baf99f7d4f5))
* add `/swagger.json` endpoint handler ([41edd8a](https://github.com/canonical/rebac-admin-ui-handlers/commit/41edd8aa118334ec971a5627c76e1ea926e99d68))
* add `Authenticator` interface ([dc45542](https://github.com/canonical/rebac-admin-ui-handlers/commit/dc455428390d1d38705b19f336a62bfebc2ad31b))
* add `NewInvalidRequestError` ([9603f0c](https://github.com/canonical/rebac-admin-ui-handlers/commit/9603f0ceb3193390f5e9e03aa1738ba38007e03f))
* add `NewNotImplementedError` helper function ([7eaf2a6](https://github.com/canonical/rebac-admin-ui-handlers/commit/7eaf2a6f1e9986289ee9964fde224ce075fbf163))
* add `writeResponse` func ([8364b52](https://github.com/canonical/rebac-admin-ui-handlers/commit/8364b52769e13d5904f9a2af2fe15654828dfb3f))
* add Capabilities endpoints handler implementation ([6088299](https://github.com/canonical/rebac-admin-ui-handlers/commit/608829944be0cd6ec8fac4be67dec3efe1d1b8d4))
* add Capabilities interface ([8070618](https://github.com/canonical/rebac-admin-ui-handlers/commit/8070618a9848ba62c6228c903cdcb60906c4bebf))
* add capability inference ([2720f22](https://github.com/canonical/rebac-admin-ui-handlers/commit/2720f22f87475a8cacc173b46726ce79f63a310a))
* add CI linting and unittest run for rebac-admin-backend code ([0435f5c](https://github.com/canonical/rebac-admin-ui-handlers/commit/0435f5c205514be7a55ff91309b7f721b51a044a))
* add coverage for tests in `reback-admin-backend folder` ([5f69042](https://github.com/canonical/rebac-admin-ui-handlers/commit/5f6904251b8d6c2eb701e3521dd14ef6af3cff71))
* add Entitlements endpoints handlers implementations ([52f4b15](https://github.com/canonical/rebac-admin-ui-handlers/commit/52f4b1591834d0b87649d7561b36ca3dec171836))
* add Entitlements service interfaces ([7cbd054](https://github.com/canonical/rebac-admin-ui-handlers/commit/7cbd05401974f8ad7878156005008e8463c7bab8))
* add first handlers' logic + scaffold remaining ([c624cac](https://github.com/canonical/rebac-admin-ui-handlers/commit/c624cac9563fd6bb4a4c5dee913f793fa9f444e7))
* add identities handler implementation for most endpoints ([d4178dc](https://github.com/canonical/rebac-admin-ui-handlers/commit/d4178dcf4a9b2a9d30a90565d0ca3b07a55db629))
* add identitiesServiceBackend interface definition ([7332476](https://github.com/canonical/rebac-admin-ui-handlers/commit/7332476efd7dd0d921740ecaf67277af10a783d7))
* add implementation for `/authentication/*` endpoints ([451a6cb](https://github.com/canonical/rebac-admin-ui-handlers/commit/451a6cb309448be437b816b0219671fbec4e14f9))
* add in-memory example ([afbbbbb](https://github.com/canonical/rebac-admin-ui-handlers/commit/afbbbbb578f31901dac519c69e86915dbfb73bc8))
* add initializers for custom errors + adjust test ([c18938a](https://github.com/canonical/rebac-admin-ui-handlers/commit/c18938aa282fe5a604a029f0548cf1052b3b9a6b))
* add interfaces for IdentityProviders service ([173789e](https://github.com/canonical/rebac-admin-ui-handlers/commit/173789ee5d691ef774cfa2e7b476a1f68e1888a3))
* add library composition/setup types ([e26771d](https://github.com/canonical/rebac-admin-ui-handlers/commit/e26771d93eba124bcecab942c4832c549870dd91))
* add new error creator functions ([41d9736](https://github.com/canonical/rebac-admin-ui-handlers/commit/41d97366105877edf068b42f8ded518c030c8eb9))
* add new error function for authentication errors ([cf024fa](https://github.com/canonical/rebac-admin-ui-handlers/commit/cf024fa52114bc29b3290b0e458e1a9ede3a9ca1))
* add PaginatedResponse generic type ([a5acc87](https://github.com/canonical/rebac-admin-ui-handlers/commit/a5acc876d905863df20ab0202d6455bb3cf5a37e))
* add Patch implementations, improved + missing tests ([4984b3a](https://github.com/canonical/rebac-admin-ui-handlers/commit/4984b3a4223e1c9f15f3d1035b2364c60a047f14))
* add proper error formatting ([63f5217](https://github.com/canonical/rebac-admin-ui-handlers/commit/63f52175775d86e195547825393b2bf68924bb87))
* add Roles related interfaces and error mapper ([7b2e8e3](https://github.com/canonical/rebac-admin-ui-handlers/commit/7b2e8e39195856e1e8cba50429afbe9cc4f3499b))
* add validation for `/authentication` endpoints ([a91602c](https://github.com/canonical/rebac-admin-ui-handlers/commit/a91602cf5c640c067733ec1e3857903313eef940))
* add validation for `/groups` endpoints ([cff7f95](https://github.com/canonical/rebac-admin-ui-handlers/commit/cff7f9570cd29f76597e0fa1e41c57dd4f6292ee))
* add validation for `/roles` endpoints ([8b4d008](https://github.com/canonical/rebac-admin-ui-handlers/commit/8b4d0087d99bae41ac20b905b95df92c5c07253d))
* add validation to `/identities` endpoints ([a41356a](https://github.com/canonical/rebac-admin-ui-handlers/commit/a41356a32f9b06cf37be5ef7ba7371a57342b4c6))
* add ValidationError + adjust tests + godocs ([e0f192e](https://github.com/canonical/rebac-admin-ui-handlers/commit/e0f192ed57fef3860cf5b1b499003944acd018cb))
* extend pull-spec.sh to allow for use of local file path ([56e6795](https://github.com/canonical/rebac-admin-ui-handlers/commit/56e6795089876695016a1770153c5be9580aec0b))
* generate OpenAPI types/handlers ([7d08230](https://github.com/canonical/rebac-admin-ui-handlers/commit/7d08230ac298bda885bd31e90f375b11bc0f4ee8))
* introduce `ErrorResponseMapper` interface and basic delegate implementation ([a5fb259](https://github.com/canonical/rebac-admin-ui-handlers/commit/a5fb2594417e662a5add9c513caa988a1a941189))


### Bug Fixes

* add `/ready` endpoint for testing ([4fedc7b](https://github.com/canonical/rebac-admin-ui-handlers/commit/4fedc7bc6ccbc0a57605880af55aebcd30178cb5))
* add `content-type` header to response ([182580b](https://github.com/canonical/rebac-admin-ui-handlers/commit/182580bcdf9b5c2556698363c25aae945ae9f2db))
* add `mocks` target to Makefile ([04a32db](https://github.com/canonical/rebac-admin-ui-handlers/commit/04a32dbfb460f280ba6357cacf964cd598e7cfa8))
* add `NewUnknownError` method ([98b17fc](https://github.com/canonical/rebac-admin-ui-handlers/commit/98b17fc549824dd58219304ea6ce12f8d85cb94d))
* add `test` target to Makefile ([dee3d62](https://github.com/canonical/rebac-admin-ui-handlers/commit/dee3d627c943822bdffdc0f94ec4ef2a17c21e76))
* add authentication middleware ([8678565](https://github.com/canonical/rebac-admin-ui-handlers/commit/8678565b9d6014b558ce0f142fcf99d4eb4d7983))
* add helper methods to map/write service-level error responses ([03b4a20](https://github.com/canonical/rebac-admin-ui-handlers/commit/03b4a20def6cae024cac1c9a9a8c5e039efd2052))
* add log on write response failure ([f1e3700](https://github.com/canonical/rebac-admin-ui-handlers/commit/f1e3700bfdbf8f7e566e7ed889e8a37a354a961d))
* add preferred pagination parameter ([451586d](https://github.com/canonical/rebac-admin-ui-handlers/commit/451586d7a2df8c97d72d31e0f386901860454c13))
* add public function `ContextWithIdentity` ([b8c1a8f](https://github.com/canonical/rebac-admin-ui-handlers/commit/b8c1a8f9f7580446d9c9b5d98f75fd143f13bc95))
* apply `handlerWithValidation` at composition ([700d618](https://github.com/canonical/rebac-admin-ui-handlers/commit/700d618b9869d1e475b78c7c79bda674bd480ed4))
* apply removed pagination from `/entitlements` response ([6018aee](https://github.com/canonical/rebac-admin-ui-handlers/commit/6018aeeadc57168754047c1137327ce306e2d701))
* apply validation handler at constructor function ([dc58ad8](https://github.com/canonical/rebac-admin-ui-handlers/commit/dc58ad8db1a31d54396b1fe3b8fb434c90ad23c3))
* bail on `nil` identity ([f23fcc3](https://github.com/canonical/rebac-admin-ui-handlers/commit/f23fcc33ddd2a964cb0104a394cf8b4b630e5997))
* bail out if authenticator is nil ([47e0883](https://github.com/canonical/rebac-admin-ui-handlers/commit/47e0883f1a46b781349f4fddee4726ffee40b08a))
* change `UserEntitlements` to `Entitlements` ([59dc095](https://github.com/canonical/rebac-admin-ui-handlers/commit/59dc09589bd9e5a580fc0fc15940bf0ba7a1669b))
* change `UserResources` to `Resources` ([59f7bde](https://github.com/canonical/rebac-admin-ui-handlers/commit/59f7bdec0e308a638a2e134d393a761c363eac67))
* check for `nil` when mapping ([33bc090](https://github.com/canonical/rebac-admin-ui-handlers/commit/33bc0900d41788184219f105324e2f76bcd80fb7))
* correct parameter names ([3dab693](https://github.com/canonical/rebac-admin-ui-handlers/commit/3dab693c55cdc6667418b445837e05faaf7d26df))
* correct wrong request body type ([73307fa](https://github.com/canonical/rebac-admin-ui-handlers/commit/73307fac34e1a3326375b181e3c76e725d02e8b9))
* embed handler struct ([35b49c6](https://github.com/canonical/rebac-admin-ui-handlers/commit/35b49c6988ccc1647e53bfc3d1e5346c4d9d2c05))
* fix lint issues ([3d585ba](https://github.com/canonical/rebac-admin-ui-handlers/commit/3d585bad497c7f37cc5f3b00628c17f54c65e925))
* fix not returning all entities when `entityType` is nil ([3cde9d1](https://github.com/canonical/rebac-admin-ui-handlers/commit/3cde9d14b0bc1123d6246f0b0aae1f7645d4771b))
* fix typo ([b75f9cc](https://github.com/canonical/rebac-admin-ui-handlers/commit/b75f9ccf77b58a3862e8b3fb2579a64f1b7b552a))
* generate mocks so Lint doesn't complain ([bfcf899](https://github.com/canonical/rebac-admin-ui-handlers/commit/bfcf89981726c4e5e349c072e9b53a434d034db6))
* ignore trailing spaces when comparing state files ([a343a7b](https://github.com/canonical/rebac-admin-ui-handlers/commit/a343a7b4283d1b911b4693924fa0df284b2ae15a))
* improve `Authenticator` interface definition ([664f33d](https://github.com/canonical/rebac-admin-ui-handlers/commit/664f33dfc96eee706f77f3878c66f9b2344e1808))
* improve `test.sh` with options ([7fd2818](https://github.com/canonical/rebac-admin-ui-handlers/commit/7fd2818e18970671a46c066b956d09e93b11f836))
* improve authorization failure error message ([d94ff5a](https://github.com/canonical/rebac-admin-ui-handlers/commit/d94ff5a6a783a91504fc8182cca967e04bc0515c))
* improve default page token assignment ([123e4de](https://github.com/canonical/rebac-admin-ui-handlers/commit/123e4de3fbf36a352391bebaf43361f6d06769a5))
* improve error message ([93c7abd](https://github.com/canonical/rebac-admin-ui-handlers/commit/93c7abd53569a230008e1050ed8ecde7e03405c9))
* install `mock/mockgen` in `mocks` target ([e5c68c7](https://github.com/canonical/rebac-admin-ui-handlers/commit/e5c68c7824021eaaa166e3841b5102bc1be2d439))
* issue when linting after merge ([3ff16dc](https://github.com/canonical/rebac-admin-ui-handlers/commit/3ff16dc1fe04de8751525c4692289545143d26b0))
* make `Authenticator` component optional ([0a9fafb](https://github.com/canonical/rebac-admin-ui-handlers/commit/0a9fafbcd643974d27e3560d01a4eb90e317e90d))
* move to repo root ([b4f931f](https://github.com/canonical/rebac-admin-ui-handlers/commit/b4f931fde4e370b9e0cf853e25b4f37ad984c0db))
* persist when `isDirty` is `true` ([6ebc144](https://github.com/canonical/rebac-admin-ui-handlers/commit/6ebc144e770bdb03e879ad65374242d9c313fd62))
* refactor toward removing generics ([141c2ef](https://github.com/canonical/rebac-admin-ui-handlers/commit/141c2efc49b1426bfed6b847915c0636c5edf91a))
* remove `delegateErrorResponseMapper` type ([544943a](https://github.com/canonical/rebac-admin-ui-handlers/commit/544943abdc2969cf53d6c06d4dea1329aa9bb48b))
* remove `Unimplemented` as embedded struct ([0895000](https://github.com/canonical/rebac-admin-ui-handlers/commit/0895000b365e412951427956b66899513d584048))
* remove JSON parsing from handler ([fe92509](https://github.com/canonical/rebac-admin-ui-handlers/commit/fe9250962bb622812d11e6466ad9416d25e3bdd3))
* remove pagination from `/entitlements` response ([7ded910](https://github.com/canonical/rebac-admin-ui-handlers/commit/7ded9101888ac76c0a88aa431698d08f5f4d1647))
* remove pagination from `CapabilitiesService` ([e1c87ed](https://github.com/canonical/rebac-admin-ui-handlers/commit/e1c87edc99b8740929c7a6ae4e3734738a904a50))
* remove unused `skip` field in tests ([b6c3187](https://github.com/canonical/rebac-admin-ui-handlers/commit/b6c3187e336416dfc92de6470bb0a05c0984248a))
* remove unused interfaces ([e077347](https://github.com/canonical/rebac-admin-ui-handlers/commit/e077347835d27533e7f05c539eaec5e735ddbeea))
* remove unused interfaces references ([36bfc1e](https://github.com/canonical/rebac-admin-ui-handlers/commit/36bfc1e885b64c8244892275f8a9356b640acb4f))
* remove unused validator ([153204f](https://github.com/canonical/rebac-admin-ui-handlers/commit/153204fb1338201ec57246a3aa6db91bdc0c58b8))
* rename `getErrorResponse` to `mapErrorResponse` ([2485e15](https://github.com/canonical/rebac-admin-ui-handlers/commit/2485e1586088a07df2dbdbbb3f6678e1d5bd9346))
* rename `internalHandlerError` to `handlerBadRequestError` ([0bd5a21](https://github.com/canonical/rebac-admin-ui-handlers/commit/0bd5a212c43f1595224044269a1914879ea6dcf5))
* reorder writing response headers ([6fdebbc](https://github.com/canonical/rebac-admin-ui-handlers/commit/6fdebbcbbdc5b43e0e933aab2f4a900929fdd2ea))
* replace with idiomatic 200 response ([e36f735](https://github.com/canonical/rebac-admin-ui-handlers/commit/e36f735eda10bb7624ae9a8e4c463f871b58ed77))
* return badrequest response in case of un-decodable request body ([6721722](https://github.com/canonical/rebac-admin-ui-handlers/commit/6721722e9b85ba13e183cae137e953cee5b5b974))
* return error from ctor if authenticator is nil ([c6eb6c8](https://github.com/canonical/rebac-admin-ui-handlers/commit/c6eb6c88c2d2d7ac1c68c6fbe8b3af838e59d0a2))
* revive parts dropped in rebase ([7902560](https://github.com/canonical/rebac-admin-ui-handlers/commit/790256009282b2e10493f7b698956fb6064df0ee))
* run built server ([21e4f2e](https://github.com/canonical/rebac-admin-ui-handlers/commit/21e4f2e0dd9d9ee22fc433c147375bf9d69947bf))
* set `isDirty` to `false` after loading ([c938d0d](https://github.com/canonical/rebac-admin-ui-handlers/commit/c938d0d4e11993d1a123aa850a6a7f7d173e43bb))
* set `next.href` to empty at last page ([767ed79](https://github.com/canonical/rebac-admin-ui-handlers/commit/767ed7976b51606d5121bff96eeb91f8534a314e))
* set provided services on new `handler` instance ([de6a1c9](https://github.com/canonical/rebac-admin-ui-handlers/commit/de6a1c979009f53bf12b4055f91b88b10be71a34))
* support filtering resources by entity type ([b8652d4](https://github.com/canonical/rebac-admin-ui-handlers/commit/b8652d47ab6ddf1a2e9f5b1d39d29720867d2661))
* **swagger:** godoc ([db11b38](https://github.com/canonical/rebac-admin-ui-handlers/commit/db11b385590e9918700d97321c337e75d7abde8d))
* typos ([cd4f699](https://github.com/canonical/rebac-admin-ui-handlers/commit/cd4f6997654109ae56954c495c661a3203890220))
* update `ListEntitlements` return type ([18ddcbe](https://github.com/canonical/rebac-admin-ui-handlers/commit/18ddcbee8e01359acb7523ddd7cdefb987218718))
* use `-o` to write to file instead of shell redirection ([bdfa736](https://github.com/canonical/rebac-admin-ui-handlers/commit/bdfa7364b72f77144eb0d1b0642e9c330d618246))
* use `errorWithStatus` as base of errors ([ad41494](https://github.com/canonical/rebac-admin-ui-handlers/commit/ad414941ea3e89ee5dd200b938ce1be574c88bd7))
* use `test.sh` options when testing ([9a7a07e](https://github.com/canonical/rebac-admin-ui-handlers/commit/9a7a07eff1b1ed6cb89263e646c9a762756fb4f0))
* use `validator` for request body validation ([ad55337](https://github.com/canonical/rebac-admin-ui-handlers/commit/ad55337ed0895ad65e002334449d946218cf4420))
* use correct name for variable ([3d479ce](https://github.com/canonical/rebac-admin-ui-handlers/commit/3d479cec5833a8a227e3e46de3301cb5fb553aef))
* use designated `/ready` endpoint ([30dea67](https://github.com/canonical/rebac-admin-ui-handlers/commit/30dea6794768ba2b58922adf0d5a527d2a795b4f))
