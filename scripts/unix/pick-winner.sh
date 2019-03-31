#!/usr/bin/env bash

set -e

./retweets-parser $1 > participants_with_block_signature.txt
PARTICIPANTS_TOTAL=$(jq length participants_with_block_signature.txt)
curl -sS -X GET --header 'Accept: application/json' "https://nodes.wavesnodes.com/blocks/headers/at/$2" | jq '.["signature"]' >> participants_with_block_signature.txt
echo "Retweeters and block signature was saved to 'participants_with_block_signature.txt'"
PROOFS=$(./vrf-proof participants_with_block_signature.txt "$3" "$PARTICIPANTS_TOTAL")
echo "$PROOFS"
WINNER_INDEX=$(echo "$PROOFS" | tail -n 1 | awk 'BEGIN {FS=": ";} { print $2 }')
WINNER=$(head -n 1 participants_with_block_signature.txt | jq ".[$WINNER_INDEX]")

echo "Total participants: $PARTICIPANTS_TOTAL"
echo "Winner index: $WINNER_INDEX"
echo "Winner: $WINNER"