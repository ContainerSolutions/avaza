### WORK STACK

-[x] Make embedded HTTP server work.
-[x] Use short term token to register long term token
  --> problem; wrong request. Maybe need to request different kind of input in 1st step?
-[~] Persist long term token + refresh funcionality.
-[x] -> Datetime parsing issue. go run ./cmd/avaza p l --> time issue.
      ---> Write down  given return format of date; then fork go-opensapi/strfmt
           (already in client/custom_dataetime_strfmt)
- [ ] List tasks

### DESIRED FEATURES
- Task list / management
- View / edit hours per week with TUI
- Register expenses from CLI

## Data model
