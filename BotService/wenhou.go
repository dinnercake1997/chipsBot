package BotService

import "chipsBot/utils"

func DrinkTeaContent()(string){

	randDrink:=utils.GetRand(1,11)
	randTea:=utils.GetRand(1,14)
	dirnkMap:=map[int]string{}
	dirnkMap[1]="可好喝了，欧尼酱也要记得补充水分噢"
	dirnkMap[2]="，不去试试吗？正好也是喝茶时间了。"
	dirnkMap[3]="，不买一杯试试吗？努力之余也要适当放松！"
	dirnkMap[4]="，独角兽今天推荐这个。"
	dirnkMap[5]="，优酱喝了都说好！"
	dirnkMap[6]="，独角兽最近喜欢喝这个？"
	dirnkMap[7]="，独角兽想喝这个，要陪独角兽去一趟奶茶店吗？"
	dirnkMap[8]="，我猜欧尼酱一定会喜欢。"
	dirnkMap[9]="听说很不错，优酱也流口水了，正好独角兽也口渴了。"
	dirnkMap[10]="一定要试试，独角兽，喜欢。"


	teaMap:=map[int]string{}
	teaMap[1]="一点点的百香奶绿去冰3分➕波霸➕茶冻"
	teaMap[2]="一点点的茶冻芒果青少冰无糖➕波霸"
	teaMap[3]="一点点的红茶玛奇朵无糖去冰➕波霸➕冰淇淋"
	teaMap[4]="一点点的鲜奶➕红豆➕波霸➕冰淇淋"
	teaMap[5]="CoCo家的双球冰淇淋红茶+芋头+珍珠、去冰、无糖"
	teaMap[6]="CoCo家的鲜百香果双响炮少冰、半糖"
	teaMap[7]="厝内小眷村的龙涎鲜奶茶无糖去冰➕芝麻芋圆➕芋泥2"
	teaMap[8]="厝内小眷村的花生牛奶沐芋乳微糖去冰➕红宝石"
	teaMap[9]="茶百道的豆乳玉麒麟➕血糯米分装➕芋泥分装➕豆花布丁分装➕冻冻分装"
	teaMap[10]="茶百道的杨枝甘露➕双倍芒果底➕脆波波➕冻冻"
	teaMap[11]="喜茶的芝芝桃桃少冰➕红柚粒➕脆波波➕芝士换冰淇淋"
	teaMap[12]="喜茶的纯绿妍少冰少糖➕脆啵啵➕红柚果粒"
	teaMap[13]="喜茶的嫣红奶茶少糖➕芝士换冰淇淋➕黑波波"
	if randDrink> len(dirnkMap)||randTea> len(teaMap){
		return "欧尼酱，独角兽来提醒你喝水了！"
	}
	return  "欧尼酱！"+teaMap[randTea]+dirnkMap[randDrink]
}
func WaterContent()( string){
	randDrink:=utils.GetRand(1,11)
	dirnkMap:=map[int]string{}
	dirnkMap[0]="欧尼酱，喝水水"
	dirnkMap[1]="欧尼酱，喝水水"
	dirnkMap[2]="喝水辣，欧尼酱。"
	dirnkMap[3]="欧尼酱知道独角兽为什么水这么多吗？因为独角兽爱喝水"
	dirnkMap[4]="欧尼酱去趟厕所，然后喝水吧"
	dirnkMap[5]="独角兽，刚好买多了一瓶水，欧尼酱要喝吗？"
	dirnkMap[6]="补充水分，多运动，身体好"
	dirnkMap[7]="欧尼酱今天喝了几杯水了？"
	dirnkMap[8]="水是生命之源，这是独角兽水多的真正原因。"
	dirnkMap[9]="知道吗，如果欧尼酱7天不喝水，会死掉的喔！"
	dirnkMap[10]="欧尼酱，站起来活动一下筋骨，然后喝水"
	if randDrink> len(dirnkMap){
		return "欧尼酱，独角兽来提醒你喝水了！"
	}
	return  dirnkMap[randDrink]
}
