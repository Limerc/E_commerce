package mysql

import (
	"fmt"
	"os"

	"github.com/Limerc/E_commerce/gomall/app/product/conf"
	"github.com/Limerc/E_commerce/gomall/app/product/biz/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// 从环境变量中获取配置
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !DB.Migrator().HasTable(&model.Product{})
		DB.AutoMigrate( //nolint:errcheck
			&model.Product{},
			&model.Category{},
		)
		// if needDemoData {
		// 	DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06','T-Shirt','T-Shirt'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06','Sticker','Sticker')")
		// 	DB.Exec("INSERT INTO `product`.`product` VALUES ( 1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Notebook', 'The cloudwego notebook is a highly efficient and feature-rich notebook designed to meet all your note-taking needs. ', '/static/image/notebook.jpeg', 9.90 ), ( 2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Mouse-Pad', 'The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ', '/static/image/mouse-pad.jpeg', 8.80 ), ( 3, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt.jpeg', 6.60 ), ( 4, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-1.jpeg', 2.20 ), ( 5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', 'Sweatshirt', 'The cloudwego Sweatshirt is a cozy and fashionable garment that provides warmth and style during colder weather.', '/static/image/sweatshirt.jpeg', 1.10 ), ( 6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'T-Shirt', 'The cloudwego t-shirt is a stylish and comfortable clothing item that allows you to showcase your fashion sense while enjoying maximum comfort.', '/static/image/t-shirt-2.jpeg', 1.80 ), ( 7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'mascot', 'The cloudwego mascot is a charming and captivating representation of the brand, designed to bring joy and a playful spirit to any environment.', '/static/image/logo.jpg', 4.80 )")
		// 	DB.Exec("INSERT INTO `product`.`product_category` (product_id,category_id) VALUES ( 1, 2 ), ( 2, 2 ), ( 3, 1 ), ( 4, 1 ), ( 5, 1 ), ( 6, 1 ),( 7, 2 )")
		// }
		if needDemoData {
			DB.Exec("INSERT INTO `product`.`category` VALUES (1, '2025-02-17 00:00:00', '2025-02-17 00:00:00', 'Garage_Kits', 'Garage Kits'),(2, '2025-02-17 00:00:00', '2025-02-17 00:00:00', 'Pendant', 'Pendant'),(3, '2025-02-17 00:00:00', '2025-02-17 00:00:00', 'Phone_Shell', 'Phone Shell')")
			DB.Exec("INSERT INTO `product`.`product` VALUES (1, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '藕粉哪吒', 'A garage kit of cute Nezha.', '/static/image/product/Product1.png', 59),(2, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '藕粉敖丙', 'A garage kit of cute A Bing.', '/static/image/product/Product2.png', 59),(3, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '敖丙版哪吒', 'A special garage kit of Nezha with A Bing style.', '/static/image/product/Product3.png', 59),(4, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '乖巧敖丙', 'A garage kit of well - behaved A Bing.', '/static/image/product/Product4.png', 69),(5, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '捣蛋哪吒', 'A garage kit of naughty Nezha.', '/static/image/product/Product5.png', 69),(6, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '牵手敖丙', 'A garage kit of A Bing holding hands.', '/static/image/product/Product6.png', 79),(7, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '牵手哪吒', 'A garage kit of Nezha holding hands.', '/static/image/product/Product7.png', 79),(8, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '哪吒飞猪', 'A pendant of Nezha on a flying pig.', '/static/image/product/Product2_1.png', 19),(9, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '哪吒挂件', 'A pendant of Nezha.', '/static/image/product/Product2_2.png', 12),(10, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '敖丙挂件', 'A pendant of A Bing.', '/static/image/product/Product2_3.png', 12),(11, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '太乙挂件', 'A pendant of Taiyi.', '/static/image/product/Product2_4.png', 12),(12, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '哪吒手机壳', 'A phone shell with Nezha pattern.', '/static/image/product/Product3_1.png', 18),(13, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '敖光手机壳', 'A phone shell with Ao Guang pattern.', '/static/image/product/Product3_2.png', 18),(14, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '敖润手机壳', 'A phone shell with Ao Run pattern.', '/static/image/product/Product3_3.png', 18),(15, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '敖丙手机壳', 'A phone shell with A Bing pattern.', '/static/image/product/Product3_4.png', 18),(16, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '敖顺手机壳', 'A phone shell with Ao Shun pattern.', '/static/image/product/Product3_5.png', 18),(17, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '敖钦手机壳', 'A phone shell with Ao Qin pattern.', '/static/image/product/Product3_6.png', 18),(18, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '申公豹手机壳', 'A phone shell with Shen Gongbao pattern.', '/static/image/product/Product3_7.png', 18),(19, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '哪吒', 'A phone shell with Nezha theme.', '/static/image/product/Product3_8.png', 23),(20, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '少年哪吒', 'A phone shell with young Nezha pattern.', '/static/image/product/Product3_9.png', 23),(21, '2025-02-17 00:00:00', '2025-02-17 00:00:00', '三头六臂哪吒', 'A phone shell with three - headed and six - armed Nezha pattern.', '/static/image/product/Product3_10.png', 29)")
			DB.Exec("INSERT INTO `product`.`product_category` (product_id, category_id) VALUES (1, 1), (2, 1), (3, 1), (4, 1), (5, 1), (6, 1), (7, 1),(8, 2), (9, 2), (10, 2), (11, 2),(12, 3), (13, 3), (14, 3), (15, 3), (16, 3), (17, 3), (18, 3), (19, 3), (20, 3), (21, 3)")
		}
	}
}