# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/shared#PageInfo">PageInfo</a>

# Templates

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ChatTemplateParam">ChatTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailTemplateParam">EmailTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#InAppFeedTemplateParam">InAppFeedTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PushTemplateParam">PushTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#RequestTemplateParam">RequestTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#SMSTemplateParam">SMSTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WebhookTemplateParam">WebhookTemplateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ChatTemplate">ChatTemplate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailTemplate">EmailTemplate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#InAppFeedTemplate">InAppFeedTemplate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PushTemplate">PushTemplate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#RequestTemplate">RequestTemplate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#SMSTemplate">SMSTemplate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WebhookTemplate">WebhookTemplate</a>

# EmailLayouts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayout">EmailLayout</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutUpsertResponse">EmailLayoutUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutValidateResponse">EmailLayoutValidateResponse</a>

Methods:

- <code title="get /v1/email_layouts/{email_layout_key}">client.EmailLayouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, emailLayoutKey <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutGetParams">EmailLayoutGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayout">EmailLayout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/email_layouts">client.EmailLayouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutListParams">EmailLayoutListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayout">EmailLayout</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/email_layouts/{email_layout_key}">client.EmailLayouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, emailLayoutKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutUpsertParams">EmailLayoutUpsertParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutUpsertResponse">EmailLayoutUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/email_layouts/{email_layout_key}/validate">client.EmailLayouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, emailLayoutKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutValidateParams">EmailLayoutValidateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailLayoutValidateResponse">EmailLayoutValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Commits

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Commit">Commit</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitCommitAllResponse">CommitCommitAllResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitPromoteAllResponse">CommitPromoteAllResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitPromoteOneResponse">CommitPromoteOneResponse</a>

Methods:

- <code title="get /v1/commits/{id}">client.Commits.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Commit">Commit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/commits">client.Commits.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitListParams">CommitListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Commit">Commit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/commits">client.Commits.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitService.CommitAll">CommitAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitCommitAllParams">CommitCommitAllParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitCommitAllResponse">CommitCommitAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/commits/promote">client.Commits.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitService.PromoteAll">PromoteAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitPromoteAllParams">CommitPromoteAllParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitPromoteAllResponse">CommitPromoteAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/commits/{id}/promote">client.Commits.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitService.PromoteOne">PromoteOne</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#CommitPromoteOneResponse">CommitPromoteOneResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Partials

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Partial">Partial</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialUpsertResponse">PartialUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialValidateResponse">PartialValidateResponse</a>

Methods:

- <code title="get /v1/partials/{partial_key}">client.Partials.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, partialKey <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialGetParams">PartialGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Partial">Partial</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/partials">client.Partials.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialListParams">PartialListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Partial">Partial</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/partials/{partial_key}">client.Partials.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, partialKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialUpsertParams">PartialUpsertParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialUpsertResponse">PartialUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/partials/{partial_key}/validate">client.Partials.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, partialKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialValidateParams">PartialValidateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PartialValidateResponse">PartialValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Translations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Translation">Translation</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationGetResponse">TranslationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationUpsertResponse">TranslationUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationValidateResponse">TranslationValidateResponse</a>

Methods:

- <code title="get /v1/translations/{locale_code}">client.Translations.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, localeCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationGetParams">TranslationGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationGetResponse">TranslationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/translations">client.Translations.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationListParams">TranslationListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Translation">Translation</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/translations/{locale_code}">client.Translations.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, localeCode <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationUpsertParams">TranslationUpsertParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationUpsertResponse">TranslationUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/translations/{locale_code}/validate">client.Translations.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, localeCode <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationValidateParams">TranslationValidateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#TranslationValidateResponse">TranslationValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workflows

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ConditionParam">ConditionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ConditionGroupUnionParam">ConditionGroupUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#DurationParam">DurationParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#SendWindowParam">SendWindowParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowBatchStepParam">WorkflowBatchStepParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowBranchStepParam">WorkflowBranchStepParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowChannelStepParam">WorkflowChannelStepParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowDelayStepParam">WorkflowDelayStepParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowFetchStepParam">WorkflowFetchStepParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowStepUnionParam">WorkflowStepUnionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowThrottleStepParam">WorkflowThrottleStepParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowTriggerWorkflowStepParam">WorkflowTriggerWorkflowStepParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Condition">Condition</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ConditionGroupUnion">ConditionGroupUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Duration">Duration</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#SendWindow">SendWindow</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Workflow">Workflow</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowBatchStep">WorkflowBatchStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowBranchStep">WorkflowBranchStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowChannelStep">WorkflowChannelStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowDelayStep">WorkflowDelayStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowFetchStep">WorkflowFetchStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowStepUnion">WorkflowStepUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowThrottleStep">WorkflowThrottleStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowTriggerWorkflowStep">WorkflowTriggerWorkflowStep</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowActivateResponse">WorkflowActivateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowRunResponse">WorkflowRunResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowUpsertResponse">WorkflowUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowValidateResponse">WorkflowValidateResponse</a>

Methods:

- <code title="get /v1/workflows/{workflow_key}">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowKey <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowGetParams">WorkflowGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Workflow">Workflow</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workflows">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowListParams">WorkflowListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Workflow">Workflow</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/workflows/{workflow_key}/activate">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowService.Activate">Activate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowActivateParams">WorkflowActivateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowActivateResponse">WorkflowActivateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/workflows/{workflow_key}/run">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowRunParams">WorkflowRunParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowRunResponse">WorkflowRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/workflows/{workflow_key}">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowUpsertParams">WorkflowUpsertParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowUpsertResponse">WorkflowUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/workflows/{workflow_key}/validate">client.Workflows.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowValidateParams">WorkflowValidateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowValidateResponse">WorkflowValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Steps

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowStepPreviewTemplateResponse">WorkflowStepPreviewTemplateResponse</a>

Methods:

- <code title="post /v1/workflows/{workflow_key}/steps/{step_ref}/preview_template">client.Workflows.Steps.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowStepService.PreviewTemplate">PreviewTemplate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowKey <a href="https://pkg.go.dev/builtin#string">string</a>, stepRef <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowStepPreviewTemplateParams">WorkflowStepPreviewTemplateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#WorkflowStepPreviewTemplateResponse">WorkflowStepPreviewTemplateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# MessageTypes

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeTextFieldParam">MessageTypeTextFieldParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeVariantParam">MessageTypeVariantParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageType">MessageType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeTextField">MessageTypeTextField</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeVariant">MessageTypeVariant</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeUpsertResponse">MessageTypeUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeValidateResponse">MessageTypeValidateResponse</a>

Methods:

- <code title="get /v1/message_types/{message_type_key}">client.MessageTypes.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageTypeKey <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeGetParams">MessageTypeGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageType">MessageType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/message_types">client.MessageTypes.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeListParams">MessageTypeListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageType">MessageType</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/message_types/{message_type_key}">client.MessageTypes.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageTypeKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeUpsertParams">MessageTypeUpsertParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeUpsertResponse">MessageTypeUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/message_types/{message_type_key}/validate">client.MessageTypes.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageTypeKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeValidateParams">MessageTypeValidateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#MessageTypeValidateResponse">MessageTypeValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#AuthVerifyResponse">AuthVerifyResponse</a>

Methods:

- <code title="get /v1/whoami">client.Auth.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#AuthService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#AuthVerifyResponse">AuthVerifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#APIKeyExchangeResponse">APIKeyExchangeResponse</a>

Methods:

- <code title="post /v1/api_keys/exchange">client.APIKeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#APIKeyService.Exchange">Exchange</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#APIKeyExchangeParams">APIKeyExchangeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#APIKeyExchangeResponse">APIKeyExchangeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ChannelGroups

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ChannelGroup">ChannelGroup</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ChannelGroupRule">ChannelGroupRule</a>

Methods:

- <code title="get /v1/channel_groups">client.ChannelGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ChannelGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ChannelGroupListParams">ChannelGroupListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ChannelGroup">ChannelGroup</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Channels

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ChatChannelSettingsParam">ChatChannelSettingsParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailChannelSettingsParam">EmailChannelSettingsParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#InAppFeedChannelSettingsParam">InAppFeedChannelSettingsParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PushChannelSettingsParam">PushChannelSettingsParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#SMSChannelSettingsParam">SMSChannelSettingsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Channel">Channel</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ChatChannelSettings">ChatChannelSettings</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EmailChannelSettings">EmailChannelSettings</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#InAppFeedChannelSettings">InAppFeedChannelSettings</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#PushChannelSettings">PushChannelSettings</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#SMSChannelSettings">SMSChannelSettings</a>

Methods:

- <code title="get /v1/channels">client.Channels.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ChannelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#ChannelListParams">ChannelListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Channel">Channel</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Environments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Environment">Environment</a>

Methods:

- <code title="get /v1/environments/{environment_slug}">client.Environments.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EnvironmentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentSlug <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Environment">Environment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/environments">client.Environments.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EnvironmentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#EnvironmentListParams">EnvironmentListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Environment">Environment</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Variables

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Variable">Variable</a>

Methods:

- <code title="get /v1/variables">client.Variables.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#VariableService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#VariableListParams">VariableListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/knock-mapi-go#Variable">Variable</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
