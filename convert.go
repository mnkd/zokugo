package main

// var zokugo = []string{
// 	"亜", "慰", "宇", "餌", "悪", // あいうえお
// 	"華", "鬼", "苦", "怪", "狐", // かきくけこ
// 	"沙", "死", "酢", "勢", "訴", // さしすせそ
// 	"駄", "血", "津", "帝", "斗", // たちつてと
// 	"那", "尼", "奴", "根", "乃", // なにぬねの
// 	"覇", "卑", "父", "屁", "帆", // はひふへほ
// 	"魔", "魅", "夢", "女", "喪", // まみむめも
// 	"夜", "遊", "世", // やゆよ
// 	"羅", "璃", "流", "零", "露", // らりるれろ
// 	"我", "魚", "云", // わをん
// 	"蛾", "戯", "愚", "外", "誤", // がぎぐげご
// 	"座", "慈", "頭", "是", "憎", // ざじずぜぞ
// 	"堕", "痔", "尽", "泥", "怒", // だぢづでど
// 	"婆", "微", "撫", "塀", "墓", // ばびぶべぼ
// 	"波", "碑", "腐", "餅", "歩"} // ぱぴぷぺぽ

var zokugo = []string{
	"亜", "亜", "慰", "慰", "宇", "宇", "餌", "餌", "悪", "悪", // ぁあぃいぅうぇえぉお
	"華", "蛾", "鬼", "戯", "苦", "愚", "怪", "外", "狐", "誤", // かがきぎくぐけげこご
	"沙", "座", "死", "慈", "酢", "頭", "勢", "是", "訴", "憎", // さざしじすずせぜそぞ
	"駄", "堕", "血", "痔", "津", "津", "尽", "帝", "泥", "斗", "怒", // ただちぢっつづてでとど
	"那", "尼", "奴", "根", "乃", // なにぬねの
	"覇", "婆", "波", "卑", "微", "碑", "父", "撫", "腐", "屁", "塀", "餅", "帆", "墓", "歩", // はばぱひびぴふぶぷへべぺほぼぽ
	"魔", "魅", "夢", "女", "喪", // まみむめも
	"夜", "夜", "遊", "遊", "世", "世", // ゃやゅゆょよ
	"羅", "璃", "流", "零", "露", // らりるれろ
	"我", "我", "ゐ", "ゑ", "魚", "云"} // わゐゑをん

func Convert(value string) string {
	var s string
	for _, v := range value {
		if v >= 0x3041 && v <= 0x3093 {
			s = s + zokugo[v-0x3041]
		} else {
			s = s + string(v)
		}
	}
	return s
}