#!/usr/bin/env bash

set -e

if [[ $# -eq 0 ]] ; then
    echo 'Tweeter pick winner script.'
    echo 'Usage: ./tweeter-pick-winner.sh [tweet id] [waves block height] [proving private key]'
    echo " Output 'participants_with_block_signature.txt' file will contains a json array of participants in the first line and the waves block signature at given height."
    echo ' The block height for the competition and the public key must be available publicly before the competition begins.'
    echo ' The list of participants is fixed at the time of the drawing and together with the created proof and is a confirmation of the randomness of the lottery.'
    exit 0
fi

./retweets-parser $1 > participants_with_block_signature.txt
PARTICIPANTS_TOTAL=$(jq length participants_with_block_signature.txt)
curl -sS -X GET --header 'Accept: application/json' "https://nodes.wavesnodes.com/blocks/headers/at/$2" | jq '.["signature"]' >> participants_with_block_signature.txt
echo "Retweeters and block signature was saved to 'participants_with_block_signature.txt'"
PROOFS=$(./vrf-prove participants_with_block_signature.txt "$3" "$PARTICIPANTS_TOTAL")
echo "$PROOFS"
WINNER_INDEX=$(echo "$PROOFS" | tail -n 1 | awk 'BEGIN {FS=": ";} { print $2 }')
WINNER=$(head -n 1 participants_with_block_signature.txt | jq ".[$WINNER_INDEX]")

echo "Total participants: $PARTICIPANTS_TOTAL"
echo "Winner index: $WINNER_INDEX"
echo "Winner: $WINNER"