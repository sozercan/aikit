#!/bin/bash

# Define the functions to extract each part
extract_model_name() {
    # Capture the base name, handling the special case for llama-3.1
    echo "$1" | sed -E 's/^llama-(3\.1)-([0-9]+\.?[0-9]*b)-.*/llama\1/;t; s/^([a-z]+)-([0-9]+x[0-9]+b|[0-9]+\.?[0-9]*b)-.*/\1/; s/^([a-z]+)-([0-9]+)-.*/\1\2/; s/^([a-z]+)-([0-9]+\.?[0-9]*b)$/\1/'
}

extract_model_size() {
    # Capture the size part, handling the special case for llama-3.1
    echo "$1" | sed -E 's/^llama-(3\.1)-([0-9]+\.?[0-9]*b)-.*/\2/;t; s/^[a-z]+-([0-9]+x[0-9]+b|[0-9]+\.?[0-9]*b)-.*/\1/; s/^[a-z]+-[0-9]+-([0-9]+\.?[0-9]*b).*/\1/; s/^[a-z]+-([0-9]+\.?[0-9]*b)$/\1/'
}

extract_model_type() {
    # Capture the type part if present, otherwise return an empty string
    echo "$1" | sed -n -e 's/.*\(chat\).*/\1/p' -e 's/.*\(instruct\).*/\1/p'
}

# Run and display results for each example
for MODEL in "llama-2-7b-chat" "llama-2-13b-chat" "llama-3-8b-instruct" "llama-3.1-8b-instruct" "phi-3-3.8b" "gemma-2b-instruct" "gemma-2-2b-instruct" "codestral-22b" "llama-3-70b-instruct" "llama-3.1-70b-instruct" "mixtral-8x7b-instruct"; do
    echo "Model: $MODEL"
    echo "  Name: $(extract_model_name $MODEL)"
    echo "  Size: $(extract_model_size $MODEL)"
    echo "  Type: $(extract_model_type $MODEL)"
done
