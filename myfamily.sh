#!/bin/bash
curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r .[$HERO_ID-1].connections.relatives
