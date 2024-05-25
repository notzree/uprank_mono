#!/bin/bash

# This script checks if the queue stack is deployed, and if not, deploys it.
# It then injects the queue urls into the .env file so applications can use them.

# Default environment
ENV="dev"

# Parse command line arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        --env) ENV="$2"; shift ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

# Define the Pulumi stack name and paths based on the environment
FULLY_QUALIFIED_STACK_NAME="notzree/queues/$ENV"
RELATIVE_QUEUE_PATH="../infrastructure/$ENV/queues" # Path to the Pulumi stack from this file
RELATIVE_ENV_PATH="../.env" # Path to the .env file from this file

# Flag to track if the stack was deployed by this script
STACK_DEPLOYED=0

# Function to check if a Pulumi stack is deployed
is_queue_deployed() {
    pushd $RELATIVE_QUEUE_PATH > /dev/null
    pulumi stack select "$FULLY_QUALIFIED_STACK_NAME" > /dev/null 2>&1
    if [ $? -ne 0 ]; then
        echo "Stack $FULLY_QUALIFIED_STACK_NAME does not exist."
        popd > /dev/null
        return 1
    fi
    
    OUTPUT=$(pulumi stack output 2>&1)
    if echo "$OUTPUT" | grep -q "No output values currently in this stack"; then
        echo "Stack $FULLY_QUALIFIED_STACK_NAME is not deployed."
        popd > /dev/null
        return 1
    fi
    
    echo "Stack $FULLY_QUALIFIED_STACK_NAME is deployed."
    popd > /dev/null
    return 0
}

# Function to deploy the Pulumi stack if it is not deployed
deploy_stack() {
    pushd $RELATIVE_QUEUE_PATH > /dev/null
    pulumi up --yes
    if [ $? -ne 0 ]; then
        echo "Failed to deploy the stack $FULLY_QUALIFIED_STACK_NAME."
        popd > /dev/null
        exit 1
    fi
    STACK_DEPLOYED=1
    popd > /dev/null
}

# Function to append or overwrite stack outputs to the .env file
append_outputs_to_env() {
    pushd $RELATIVE_QUEUE_PATH > /dev/null
    OUTPUTS=$(pulumi stack output --json)
    popd > /dev/null
    
    if [ $STACK_DEPLOYED -eq 1 ]; then
        # Remove existing values for keys that will be overwritten
        for key in $(echo "$OUTPUTS" | jq -r 'keys[]'); do
            UPPER_KEY=$(echo "$key" | tr '[:lower:]' '[:upper:]')
            sed -i "/^$UPPER_KEY=/d" $RELATIVE_ENV_PATH
        done
    fi
    
    printf "\n" >> $RELATIVE_ENV_PATH
    for key in $(echo "$OUTPUTS" | jq -r 'keys[]'); do
        value=$(echo "$OUTPUTS" | jq -r --arg key "$key" '.[$key]')
        UPPER_KEY=$(echo "$key" | tr '[:lower:]' '[:upper:]')
        # Check if the key already exists in the .env file
        if ! grep -q "^$UPPER_KEY=" $RELATIVE_ENV_PATH; then
            echo "$UPPER_KEY=$value" >> $RELATIVE_ENV_PATH
        fi
    done
}


# Check if the Pulumi stack is deployed
if ! is_queue_deployed; then
    echo "Deploying the stack $FULLY_QUALIFIED_STACK_NAME..."
    deploy_stack
else
    echo "Stack $FULLY_QUALIFIED_STACK_NAME is already deployed."
fi

# Append stack outputs to the .env file
append_outputs_to_env

echo "Done."
