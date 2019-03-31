#!/usr/bin/env bash

set -e

PARTICIPANTS_TOTAL=$(head -n 1 "$1" | jq length)
PROOFS=$(./vrf-verify participants_with_block_signature.txt "$2" "$3" "$PARTICIPANTS_TOTAL")
echo "$PROOFS"
WINNER_INDEX=$(echo "$PROOFS" | tail -n 1 | awk 'BEGIN {FS=": ";} { print $2 }')
WINNER=$(head -n 1 participants_with_block_signature.txt | jq ".[$WINNER_INDEX]")

echo "Total participants: $PARTICIPANTS_TOTAL"
echo "Winner index: $WINNER_INDEX"
echo "Winner: $WINNER"