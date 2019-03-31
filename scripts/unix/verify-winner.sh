#!/usr/bin/env bash

set -e

if [[ $# -eq 0 ]] ; then
    echo 'Winner verify script.'
    echo 'Usage: ./verify-winner.sh [provable file] [proof] [proving public key file]'
    echo ' Provable file should contains a json array of participants in the first line'
    echo ' Other lines may contain data unknown at the time of the beginning of the lottery, for example, the block signature from the future height.'
    exit 0
fi

PARTICIPANTS_TOTAL=$(head -n 1 "$1" | jq length)
PROOFS=$(./vrf-verify participants_with_block_signature.txt "$2" "$3" "$PARTICIPANTS_TOTAL")
echo "$PROOFS"
WINNER_INDEX=$(echo "$PROOFS" | tail -n 1 | awk 'BEGIN {FS=": ";} { print $2 }')
WINNER=$(head -n 1 participants_with_block_signature.txt | jq ".[$WINNER_INDEX]")

echo "Total participants: $PARTICIPANTS_TOTAL"
echo "Winner index: $WINNER_INDEX"
echo "Winner: $WINNER"