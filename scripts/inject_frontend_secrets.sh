#!/bin/bash

# This script injects the necessary frontend secrets defined in the .env file into the frontend application.

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

RELATIVE_ENV_PATH = "../

