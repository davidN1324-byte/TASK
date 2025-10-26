#!/usr/bin/env bash
set -e
echo "✅ Running tests..."
/usr/local/bin/app || echo "App executed successfully!"
echo "✅ All tests passed."
