#!/usr/bin/env bats

load helpers

WAIT_TIME=120
SLEEP_TIME=1

teardown() {
}

teardown_file() {
}

@test "send request to localhost:8080/v1/chat/completions" {
    run curl --retry 10 --retry-all-errors http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
     "model": "mistral-7b-openorca.Q6_K.gguf",
     "messages": [{"role": "user", "content": "hello world"}],
     "temperature": 0
    }'
    assert_success
    # TODO check response
}
