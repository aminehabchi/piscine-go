#!/bin/bash
curl -s  https://learn.zone01oujda.ma/assets/superhero/all.json | jq '.[119].name'
curl -s  https://learn.zone01oujda.ma/assets/superhero/all.json | jq '.[119].powerstats.power'
curl -s  https://learn.zone01oujda.ma/assets/superhero/all.json | jq '.[11].appearance.gender'
