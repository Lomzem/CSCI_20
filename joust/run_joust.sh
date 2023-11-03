#!/bin/bash

Knight1Name="King Arthur"
Knight1Stamina=100
Knight1WeaponType="Excalibur"
Knight1HitChance=0
Knight1StaminaCost=80
Knight2Name="Black Knight"
Knight2Stamina=100
Knight2WeaponType="Longsword"
Knight2HitChance=0
Knight2StaminaCost=80

go run ./joust.go -Knight1Name "$Knight1Name" -Knight1Stamina $Knight1Stamina -Knight1WeaponType "$Knight1WeaponType" -Knight1HitChance $Knight1HitChance -Knight1StaminaCost $Knight1StaminaCost -Knight2Name "$Knight2Name" -Knight2Stamina $Knight2Stamina -Knight2WeaponType "$Knight2WeaponType" -Knight2HitChance $Knight2HitChance -Knight2StaminaCost $Knight2StaminaCost
