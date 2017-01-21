package main

import (
	"testing"
)

//
func isEqualToString(v1, v2 string, t *testing.T) {
	if v1 != v2 {
		t.Error(v1, "not equal to", v2)
	}
}

func TestHiragana(t *testing.T) {
	isEqualToString(Convert("あいうえお"), "亜慰宇餌悪", t)
	isEqualToString(Convert("ぁぃぅぇぉ"), "亜慰宇餌悪", t)

	isEqualToString(Convert("かきくけこ"), "華鬼苦怪狐", t)
	isEqualToString(Convert("がぎぐげご"), "蛾戯愚外誤", t)

	isEqualToString(Convert("さしすせそ"), "沙死酢勢訴", t)
	isEqualToString(Convert("ざじずぜぞ"), "座慈頭是憎", t)

	isEqualToString(Convert("たちっつてと"), "駄血津津帝斗", t)
	isEqualToString(Convert("だぢづでど"), "堕痔尽泥怒", t)

	isEqualToString(Convert("なにぬねの"), "那尼奴根乃", t)

	isEqualToString(Convert("はひふへほ"), "覇卑父屁帆", t)
	isEqualToString(Convert("ばびぶべぼ"), "婆微撫塀墓", t)
	isEqualToString(Convert("ぱぴぷぺぽ"), "波碑腐餅歩", t)

	isEqualToString(Convert("まみむめも"), "魔魅夢女喪", t)
	isEqualToString(Convert("やゆよ"), "夜遊世", t)
	isEqualToString(Convert("ゃゅょ"), "夜遊世", t)

	isEqualToString(Convert("らりるれろ"), "羅璃流零露", t)
	isEqualToString(Convert("わをん"), "我魚云", t)
}

func TestAlphaNumeric(t *testing.T) {
	isEqualToString(Convert("abcあ.い.う.え.お😄"), "abc亜.慰.宇.餌.悪😄", t)
}
