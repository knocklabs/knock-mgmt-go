# Shared Response Types

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/shared#PageInfo">PageInfo</a>

# Templates

Params Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChatTemplateParam">ChatTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailTemplateParam">EmailTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#InAppFeedTemplateParam">InAppFeedTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PushTemplateParam">PushTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#RequestTemplateParam">RequestTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#SMSTemplateParam">SMSTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WebhookTemplateParam">WebhookTemplateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChatTemplate">ChatTemplate</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailTemplate">EmailTemplate</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#InAppFeedTemplate">InAppFeedTemplate</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PushTemplate">PushTemplate</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#RequestTemplate">RequestTemplate</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#SMSTemplate">SMSTemplate</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WebhookTemplate">WebhookTemplate</a>

# EmailLayouts

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayout">EmailLayout</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutUpsertResponse">EmailLayoutUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutValidateResponse">EmailLayoutValidateResponse</a>

Methods:

- <code title="get /v1/email_layouts/{email_layout_key}">client.EmailLayouts.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, emailLayoutKey <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutGetParams">EmailLayoutGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayout">EmailLayout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/email_layouts">client.EmailLayouts.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutListParams">EmailLayoutListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayout">EmailLayout</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/email_layouts/{email_layout_key}">client.EmailLayouts.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, emailLayoutKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutUpsertParams">EmailLayoutUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutUpsertResponse">EmailLayoutUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/email_layouts/{email_layout_key}/validate">client.EmailLayouts.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, emailLayoutKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutValidateParams">EmailLayoutValidateParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailLayoutValidateResponse">EmailLayoutValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Commits

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Commit">Commit</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitCommitAllResponse">CommitCommitAllResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitPromoteAllResponse">CommitPromoteAllResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitPromoteOneResponse">CommitPromoteOneResponse</a>

Methods:

- <code title="get /v1/commits/{id}">client.Commits.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Commit">Commit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/commits">client.Commits.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitListParams">CommitListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Commit">Commit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/commits">client.Commits.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitService.CommitAll">CommitAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitCommitAllParams">CommitCommitAllParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitCommitAllResponse">CommitCommitAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/commits/promote">client.Commits.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitService.PromoteAll">PromoteAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitPromoteAllParams">CommitPromoteAllParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitPromoteAllResponse">CommitPromoteAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/commits/{id}/promote">client.Commits.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitService.PromoteOne">PromoteOne</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#CommitPromoteOneResponse">CommitPromoteOneResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Partials

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Partial">Partial</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialUpsertResponse">PartialUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialValidateResponse">PartialValidateResponse</a>

Methods:

- <code title="get /v1/partials/{partial_key}">client.Partials.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, partialKey <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialGetParams">PartialGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Partial">Partial</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/partials">client.Partials.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialListParams">PartialListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Partial">Partial</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/partials/{partial_key}">client.Partials.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, partialKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialUpsertParams">PartialUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialUpsertResponse">PartialUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/partials/{partial_key}/validate">client.Partials.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, partialKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialValidateParams">PartialValidateParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PartialValidateResponse">PartialValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Translations

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Translation">Translation</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationGetResponse">TranslationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationUpsertResponse">TranslationUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationValidateResponse">TranslationValidateResponse</a>

Methods:

- <code title="get /v1/translations/{locale_code}">client.Translations.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, localeCode <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationGetParams">TranslationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationGetResponse">TranslationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/translations">client.Translations.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationListParams">TranslationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Translation">Translation</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/translations/{locale_code}">client.Translations.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, localeCode <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationUpsertParams">TranslationUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationUpsertResponse">TranslationUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/translations/{locale_code}/validate">client.Translations.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, localeCode <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationValidateParams">TranslationValidateParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#TranslationValidateResponse">TranslationValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workflows

Params Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ConditionParam">ConditionParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ConditionGroupUnionParam">ConditionGroupUnionParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#DurationParam">DurationParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#SendWindowParam">SendWindowParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowAIAgentStepParam">WorkflowAIAgentStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowBatchStepParam">WorkflowBatchStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowBranchStepParam">WorkflowBranchStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowChatStepParam">WorkflowChatStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowDelayStepParam">WorkflowDelayStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowEmailStepParam">WorkflowEmailStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowFetchStepParam">WorkflowFetchStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowInAppFeedStepParam">WorkflowInAppFeedStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowPushStepParam">WorkflowPushStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowRandomCohortStepParam">WorkflowRandomCohortStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowSMSStepParam">WorkflowSMSStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowStepUnionParam">WorkflowStepUnionParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowThrottleStepParam">WorkflowThrottleStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowTriggerWorkflowStepParam">WorkflowTriggerWorkflowStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowUpdateDataStepParam">WorkflowUpdateDataStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowUpdateObjectStepParam">WorkflowUpdateObjectStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowUpdateTenantStepParam">WorkflowUpdateTenantStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowUpdateUserStepParam">WorkflowUpdateUserStepParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowWebhookStepParam">WorkflowWebhookStepParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Condition">Condition</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ConditionGroupUnion">ConditionGroupUnion</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Duration">Duration</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#SendWindow">SendWindow</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Workflow">Workflow</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowAIAgentStep">WorkflowAIAgentStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowBatchStep">WorkflowBatchStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowBranchStep">WorkflowBranchStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowChatStep">WorkflowChatStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowDelayStep">WorkflowDelayStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowEmailStep">WorkflowEmailStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowFetchStep">WorkflowFetchStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowInAppFeedStep">WorkflowInAppFeedStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowPushStep">WorkflowPushStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowRandomCohortStep">WorkflowRandomCohortStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowSMSStep">WorkflowSMSStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowStepUnion">WorkflowStepUnion</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowThrottleStep">WorkflowThrottleStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowTriggerWorkflowStep">WorkflowTriggerWorkflowStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowUpdateDataStep">WorkflowUpdateDataStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowUpdateObjectStep">WorkflowUpdateObjectStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowUpdateTenantStep">WorkflowUpdateTenantStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowUpdateUserStep">WorkflowUpdateUserStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowWebhookStep">WorkflowWebhookStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowGetResponse">WorkflowGetResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowActivateResponse">WorkflowActivateResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowRunResponse">WorkflowRunResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowUpsertResponse">WorkflowUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowValidateResponse">WorkflowValidateResponse</a>

Methods:

- <code title="get /v1/workflows/{workflow_key}">client.Workflows.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowKey <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowGetParams">WorkflowGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowGetResponse">WorkflowGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/workflows">client.Workflows.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowListParams">WorkflowListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Workflow">Workflow</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/workflows/{workflow_key}/activate">client.Workflows.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowService.Activate">Activate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowActivateParams">WorkflowActivateParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowActivateResponse">WorkflowActivateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/workflows/{workflow_key}/run">client.Workflows.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowRunParams">WorkflowRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowRunResponse">WorkflowRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/workflows/{workflow_key}">client.Workflows.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowUpsertParams">WorkflowUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowUpsertResponse">WorkflowUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/workflows/{workflow_key}/validate">client.Workflows.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workflowKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowValidateParams">WorkflowValidateParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowValidateResponse">WorkflowValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Steps

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowStepPreviewTemplateResponse">WorkflowStepPreviewTemplateResponse</a>

Methods:

- <code title="post /v1/workflows/{workflow_key}/steps/{step_ref}/preview_template">client.Workflows.Steps.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowStepService.PreviewTemplate">PreviewTemplate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stepRef <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowStepPreviewTemplateParams">WorkflowStepPreviewTemplateParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#WorkflowStepPreviewTemplateResponse">WorkflowStepPreviewTemplateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# MessageTypes

Params Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeTextFieldParam">MessageTypeTextFieldParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeVariantParam">MessageTypeVariantParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageType">MessageType</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeTextField">MessageTypeTextField</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeVariant">MessageTypeVariant</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeUpsertResponse">MessageTypeUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeValidateResponse">MessageTypeValidateResponse</a>

Methods:

- <code title="get /v1/message_types/{message_type_key}">client.MessageTypes.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageTypeKey <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeGetParams">MessageTypeGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageType">MessageType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/message_types">client.MessageTypes.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeListParams">MessageTypeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageType">MessageType</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/message_types/{message_type_key}">client.MessageTypes.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageTypeKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeUpsertParams">MessageTypeUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeUpsertResponse">MessageTypeUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/message_types/{message_type_key}/validate">client.MessageTypes.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, messageTypeKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeValidateParams">MessageTypeValidateParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MessageTypeValidateResponse">MessageTypeValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#AuthVerifyResponse">AuthVerifyResponse</a>

Methods:

- <code title="get /v1/whoami">client.Auth.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#AuthService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#AuthVerifyResponse">AuthVerifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#APIKeyExchangeResponse">APIKeyExchangeResponse</a>

Methods:

- <code title="post /v1/api_keys/exchange">client.APIKeys.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#APIKeyService.Exchange">Exchange</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#APIKeyExchangeParams">APIKeyExchangeParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#APIKeyExchangeResponse">APIKeyExchangeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ChannelGroups

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroup">ChannelGroup</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroupRule">ChannelGroupRule</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroupUpsertResponse">ChannelGroupUpsertResponse</a>

Methods:

- <code title="get /v1/channel_groups/{channel_group_key}">client.ChannelGroups.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, channelGroupKey <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroup">ChannelGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/channel_groups">client.ChannelGroups.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroupListParams">ChannelGroupListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroup">ChannelGroup</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/channel_groups/{channel_group_key}">client.ChannelGroups.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, channelGroupKey <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /v1/channel_groups/{channel_group_key}">client.ChannelGroups.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroupService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, channelGroupKey <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroupUpsertParams">ChannelGroupUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelGroupUpsertResponse">ChannelGroupUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Channels

Params Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChatChannelSettingsParam">ChatChannelSettingsParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailChannelSettingsParam">EmailChannelSettingsParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#InAppFeedChannelSettingsParam">InAppFeedChannelSettingsParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PushChannelSettingsParam">PushChannelSettingsParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#SMSChannelSettingsParam">SMSChannelSettingsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Channel">Channel</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChatChannelSettings">ChatChannelSettings</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EmailChannelSettings">EmailChannelSettings</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#InAppFeedChannelSettings">InAppFeedChannelSettings</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#PushChannelSettings">PushChannelSettings</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#SMSChannelSettings">SMSChannelSettings</a>

Methods:

- <code title="get /v1/channels">client.Channels.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#ChannelListParams">ChannelListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Channel">Channel</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Members

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Member">Member</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MemberUser">MemberUser</a>

Methods:

- <code title="get /v1/members/{id}">client.Members.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MemberService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Member">Member</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/members">client.Members.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MemberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MemberListParams">MemberListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Member">Member</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/members/{id}">client.Members.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#MemberService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Environments

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Environment">Environment</a>

Methods:

- <code title="get /v1/environments/{environment_slug}">client.Environments.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EnvironmentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentSlug <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Environment">Environment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/environments">client.Environments.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EnvironmentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#EnvironmentListParams">EnvironmentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Environment">Environment</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Variables

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Variable">Variable</a>

Methods:

- <code title="get /v1/variables">client.Variables.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#VariableService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#VariableListParams">VariableListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Variable">Variable</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Guides

Params Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideActivationURLPatternParam">GuideActivationURLPatternParam</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideStepParam">GuideStepParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Guide">Guide</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideActivationURLPattern">GuideActivationURLPattern</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideStep">GuideStep</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideActivateResponse">GuideActivateResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideArchiveResponse">GuideArchiveResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideUpsertResponse">GuideUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideValidateResponse">GuideValidateResponse</a>

Methods:

- <code title="get /v1/guides/{guide_key}">client.Guides.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, guideKey <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideGetParams">GuideGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Guide">Guide</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/guides">client.Guides.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideListParams">GuideListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Guide">Guide</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/guides/{guide_key}/activate">client.Guides.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideService.Activate">Activate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, guideKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideActivateParams">GuideActivateParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideActivateResponse">GuideActivateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/guides/{guide_key}">client.Guides.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, guideKey <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideArchiveResponse">GuideArchiveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/guides/{guide_key}">client.Guides.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, guideKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideUpsertParams">GuideUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideUpsertResponse">GuideUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/guides/{guide_key}/validate">client.Guides.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, guideKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideValidateParams">GuideValidateParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#GuideValidateResponse">GuideValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Branches

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Branch">Branch</a>

Methods:

- <code title="post /v1/branches/{branch_slug}">client.Branches.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BranchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, branchSlug <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BranchNewParams">BranchNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Branch">Branch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/branches/{branch_slug}">client.Branches.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BranchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, branchSlug <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BranchGetParams">BranchGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Branch">Branch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/branches">client.Branches.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BranchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BranchListParams">BranchListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Branch">Branch</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/branches/{branch_slug}">client.Branches.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BranchService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, branchSlug <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BranchDeleteParams">BranchDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Broadcasts

Params Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastRequestParam">BroadcastRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Broadcast">Broadcast</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastCancelResponse">BroadcastCancelResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastSendResponse">BroadcastSendResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastUpsertResponse">BroadcastUpsertResponse</a>
- <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastValidateResponse">BroadcastValidateResponse</a>

Methods:

- <code title="get /v1/broadcasts/{broadcast_key}">client.Broadcasts.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, broadcastKey <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastGetParams">BroadcastGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Broadcast">Broadcast</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/broadcasts">client.Broadcasts.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastListParams">BroadcastListParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go/packages/pagination#EntriesCursor">EntriesCursor</a>[<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#Broadcast">Broadcast</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/broadcasts/{broadcast_key}/cancel">client.Broadcasts.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, broadcastKey <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastCancelParams">BroadcastCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastCancelResponse">BroadcastCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/broadcasts/{broadcast_key}/send">client.Broadcasts.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, broadcastKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastSendParams">BroadcastSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastSendResponse">BroadcastSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/broadcasts/{broadcast_key}">client.Broadcasts.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, broadcastKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastUpsertParams">BroadcastUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastUpsertResponse">BroadcastUpsertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/broadcasts/{broadcast_key}/validate">client.Broadcasts.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, broadcastKey <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastValidateParams">BroadcastValidateParams</a>) (\*<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go">knockmapi</a>.<a href="https://pkg.go.dev/github.com/knocklabs/knock-mgmt-go#BroadcastValidateResponse">BroadcastValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
