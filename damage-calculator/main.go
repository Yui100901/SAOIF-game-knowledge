package main

import "math"

type Parameters struct {
	BaseAttack       float64 //基础攻擊力
	AttackRate       float64 //攻擊加成區間
	TotalAttack      float64 //縂攻擊力
	SkillDamageRate  float64 //技能倍率
	TargetDefense    float64 //目標防禦
	AdditionalDamage float64 //追加傷害
	DamageRate       float64 //增傷區
	MarkDamageRate   float64 //刻印增傷區
	AwakenRate       float64 //覺醒率
	ChaosRate        float64 //混沌追憶
}

// CalculateTotalAttack 縂攻擊計算
func (p *Parameters) CalculateTotalAttack() {
	p.TotalAttack = math.Floor(p.BaseAttack * p.AttackRate)
}

// CalculateNormalDamage 普通傷害計算
func (p *Parameters) CalculateNormalDamage() float64 {
	if p.MarkDamageRate <= 0 {
		return 1
	}
	//技能傷害
	skillDamage := p.TotalAttack*p.DamageRate - p.TargetDefense
	if skillDamage <= 0 {
		skillDamage = 1
	}
	//基礎傷害
	baseDamage := math.Floor(skillDamage + p.AdditionalDamage)
	//增傷影響
	damageRateEffected := math.Floor(baseDamage * p.DamageRate)
	//刻印影響
	markDamageRateEffected := math.Floor(damageRateEffected * +p.MarkDamageRate)
	//混沌追憶影響係數
	chaosEffectRate := 1 - p.ChaosRate + p.AwakenRate
	if chaosEffectRate > 1 {
		chaosEffectRate = 1
	}
	//混沌追憶影響
	chaosEffected := math.Floor(markDamageRateEffected * chaosEffectRate)
	return chaosEffected
}

// CalculateAwakenDamage 覺醒傷害計算
func (p *Parameters) CalculateAwakenDamage() float64 {
	if p.AwakenRate == 0 {
		return 0
	}
	//基礎傷害
	baseDamage := math.Floor(p.TotalAttack*p.DamageRate + p.AdditionalDamage)
	//增傷影響
	damageRateEffected := math.Floor(baseDamage * p.DamageRate)
	//覺醒影響
	awakenRateEffected := math.Floor(damageRateEffected * p.AwakenRate)
	return awakenRateEffected
}

// CalculateTotal 傷害計算
func (p *Parameters) CalculateTotal() float64 {
	p.CalculateTotalAttack()
	if p.DamageRate <= 0 {
		return 1
	}
	total := p.CalculateNormalDamage() + p.CalculateAwakenDamage()
	return total
}
