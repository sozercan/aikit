#!/usr/bin/env bats

load helpers

WAIT_TIME=120
SLEEP_TIME=1

@test "send request to llama-2-7b-chat" {
    run curl --retry 20 --retry-all-errors http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
     "model": "llama-2-7b-chat",
     "messages": [{"role": "user", "content": "explain kubernetes in a sentence"}],
    }'
    assert_success
}
