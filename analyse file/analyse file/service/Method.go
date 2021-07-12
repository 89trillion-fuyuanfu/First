package service

// 定义一个map
		type Vmap map[string]struct{
			Id string			// id
			Rarity string		//稀有度
			UnlockArena string	//解锁阶段
			CombatPoints string //战力
			Name string //姓名
			Cvc string //稀有度
		}


		// 输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
		func GetSolider(rarity string,unkockArena string,cvc string,vmap2 Vmap) string{
			var a string
			for _,value := range vmap2{
				if rarity == value.Rarity && unkockArena == value.UnlockArena && cvc == value.Cvc {
					print(value.Name)
					a = value.Name
				}
			}
			return a
		}
		// 输入士兵id获取稀有度
		func GetRarity(id string,vmap2 Vmap)string{
			var b string
			for _,value := range vmap2{
				if id == value.Id{
					println(value.Rarity)
					b = value.Rarity
				}

			}
			return b
		}
		// 输入士兵id获取战力
		func GetCombatPoints(id string,vmap2 Vmap) string {
			var c string
			for _,value := range vmap2{
				if id == value.Id{
					println(value.CombatPoints)
					c = value.CombatPoints
				}
			}
			return c
		}

		// 输入cvc获取所有合法的士兵
		func GetLegalsoldier(cvc string,vmap2 Vmap) string {
			var d string
			for _,value := range vmap2{
				if value.Cvc == cvc{
					print(value.Name)
					d = value.Name
				}
			}
			return d
		}
		// 获取每个阶段解锁相应士兵的json数据
		func GetSoldierjson(unlockArena string,vmap2 Vmap) string{
			var e string
			for _,value := range vmap2{
				if value.UnlockArena == unlockArena{
					print(value.Cvc,value.Id,value.Rarity,value.CombatPoints)
					e = value.Cvc+" "+value.Id+" "+value.Rarity+" "+value.CombatPoints
				}

			}
			return e
		}





