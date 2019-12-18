#!/bin/sh -l

readonly webhook=${WEBHOOK_URL}

echo "Will call webhook any second now..."
echo "::log-command webhook={${webhook}}::{call-wh}"

call-wh -webhook ${webhook}
