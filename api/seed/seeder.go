package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/maksimmysak/cook/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

var recipests = []models.Recipes{
	models.Recipes{
		Title:          "Овсяное печенье с кофейной глазурью - пошаговый рецепт с фото",
		Description:    "Рецепт приготовления печенья из овсяных хлопьев и шоколада с глазурью из кофе и сливочного масла. По этому рецепту у вас получится около 2 десятка печенья.",
		Name:           "Овсяное печенье с кофейной глазурью",
		MainImg:        "https://img.povar.ru/main/34/82/b6/99/ovsyanoe_pechene_s_kofeinoi_glazuriu-23399.jpg",
		Detailed:       "Рецепт приготовления печенья из овсяных хлопьев и шоколада с глазурью из кофе и сливочного масла. По этому рецепту у вас получится около 2 десятка печенья.",
		DetailedFull:   "",
		Ingredients:    "сода  — 1/2 Чайных ложки корица  — 1/2 Чайных ложки соль  — 1/4 Чайных ложки сахар  — 1/4 Стакана яйца  — 1 Штука ванильный экстракт  — 1/2 Чайных ложки молоко  — 2 Ст. ложки (глазурь) растворимый кофе или эспрессо  — 1/2 Чайных ложки (глазурь) ванильный экстракт  — 1/2 Чайных ложки (глазурь)",
		Appointment:    "На завтрак / На обед / Для детей",
		MainIngredient: "Крупы / Тесто / Овсянка / Шоколад",
		Dish:           "",
		Time:           "30 мин",
		Diet:           "",
		CountServing:   "6-8",
		Content: `<!-- step_id = 16348 -->                                                    <div class="detailed_step_photo_big"><a
		title="1. Сделать печенье. Разогреть духовку до 175 градусов. Выстелить противень пергаментной бумагой или силиконовым ковриком. В средней миске смешать муку, соду, корицу и соль, отложить в сторону. В большой миске миксером взбить сливочное масло. Добавить оба вида сахара и взбить. Взбить с яйцом и перемешать с ванильным экстрактом. " rel="stepphotos"
		class="stepphotos"
		href="https://img.povar.ru/uploads/5f/0a/97/58/ovsyanoe_pechene_s_kofeinoi_glazuriu-23395.jpg"><img src="https://img.povar.ru/steps/5f/0a/97/58/ovsyanoe_pechene_s_kofeinoi_glazuriu-23395.jpg" width="298" height="198" itemprop="image" border="0" alt="Овсяное печенье с кофейной глазурью - фото шаг 1"/></a>
		</div>
		<div class="detailed_step_description_big">1. Сделать печенье. Разогреть духовку до 175 градусов. Выстелить противень пергаментной бумагой или силиконовым ковриком. В средней миске смешать муку, соду, корицу и соль, отложить в сторону. В большой миске миксером взбить сливочное масло. Добавить оба вида сахара и взбить. Взбить с яйцом и перемешать с ванильным экстрактом. </div>
		<div class="cleaner"></div>
		<!-- step_id = 16349 -->                                                    <div class="detailed_step_photo_big"><a
		title="2.Добавить смесь муки в два захода и перемешать до однородности. Добавить овсяные хлопья и перемешать. Перемешать с шоколадными чипсами, пока они не будут равномерно распределены по всему тесту. " rel="stepphotos"
		class="stepphotos"
		href="https://img.povar.ru/uploads/73/21/db/e8/ovsyanoe_pechene_s_kofeinoi_glazuriu-23396.jpg"><img src="https://img.povar.ru/steps/73/21/db/e8/ovsyanoe_pechene_s_kofeinoi_glazuriu-23396.jpg" width="298" height="198" itemprop="image" border="0" alt="Овсяное печенье с кофейной глазурью - фото шаг 2"/></a>
		</div>
		<div class="detailed_step_description_big">2.Добавить смесь муки в два захода и перемешать до однородности. Добавить овсяные хлопья и перемешать. Перемешать с шоколадными чипсами, пока они не будут равномерно распределены по всему тесту. </div>
		<div class="cleaner"></div>
		<!-- step_id = 16350 -->                                                    <div class="detailed_step_photo_big"><a
		title="3. Выложить по одной столовой ложке теста на подготовленный противень. Выпекать 10-12 минут, до краев золотисто-коричневого цвета. Дать остыть на противне в течение 2 минут, затем дать полностью остыть на стойке перед нанесением глазури." rel="stepphotos"
		class="stepphotos"
		href="https://img.povar.ru/uploads/b4/e5/10/44/ovsyanoe_pechene_s_kofeinoi_glazuriu-23397.jpg"><img src="https://img.povar.ru/steps/b4/e5/10/44/ovsyanoe_pechene_s_kofeinoi_glazuriu-23397.jpg" width="298" height="198" itemprop="image" border="0" alt="Овсяное печенье с кофейной глазурью - фото шаг 3"/></a>
		</div>
		<div class="detailed_step_description_big">3. Выложить по одной столовой ложке теста на подготовленный противень. Выпекать 10-12 минут, до краев золотисто-коричневого цвета. Дать остыть на противне в течение 2 минут, затем дать полностью остыть на стойке перед нанесением глазури.</div>
		<div class="cleaner"></div>
		<!-- step_id = 16351 -->                                                    <div class="detailed_step_photo_big"><a
		title="4. Чтобы сделать глазурь, нагреть молоко, добавить кофе и дать настояться в течение 5-10 минут. Процедить смесь. Дать остыть до комнатной температуры около 15 минут. В средней миске взбить сливочное масло до кремовой консистенции. Добавить половину сахара и взбить. Перемешать с ванильным экстрактом и половиной кофейной смеси. Добавить оставшийся сахар и взбить. Размешать с оставшейся кофейной смесью и тщательно взбить. Залить остывшее печенье глазурью. При желании украсить сверху шоколадными чипсами." rel="stepphotos"
		class="stepphotos"
		href="https://img.povar.ru/uploads/29/4b/c8/fe/ovsyanoe_pechene_s_kofeinoi_glazuriu-23398.jpg"><img src="https://img.povar.ru/steps/29/4b/c8/fe/ovsyanoe_pechene_s_kofeinoi_glazuriu-23398.jpg" width="298" height="198" itemprop="image" border="0" alt="Овсяное печенье с кофейной глазурью - фото шаг 4"/></a>
		</div>
		<div class="detailed_step_description_big">4. Чтобы сделать глазурь, нагреть молоко, добавить кофе и дать настояться в течение 5-10 минут. Процедить смесь. Дать остыть до комнатной температуры около 15 минут. В средней миске взбить сливочное масло до кремовой консистенции. Добавить половину сахара и взбить. Перемешать с ванильным экстрактом и половиной кофейной смеси. Добавить оставшийся сахар и взбить. Размешать с оставшейся кофейной смесью и тщательно взбить. Залить остывшее печенье глазурью. При желании украсить сверху шоколадными чипсами.</div>
		`,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}, &models.Recipes{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Recipes{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*
		err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
	*/

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}

		recipests[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Recipes{}).Create(&recipests[i]).Error
		if err != nil {
			log.Fatalf("cannot seed Recipest table: %v", err)
		}
	}
}
