package main

import (
	"context"
	"log"

	"github.com/User/app/controllers"
	_ "github.com/User/app/docs"
	"github.com/User/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Units struct{
	Unit []Unit
}
type Unit struct{
	UnitType	string
}

type Forms struct{
	Form []Form
}
type Form struct{
	FormType	string
}

type Volumes struct{
	Volume []Volume
}
type Volume struct{
	VolumeType	string
}

type Users struct{
	User []User
}
type User struct{
	Username	string
	Email	string
	Password	string
}

// type Drugs struct{
// 	Drug []Drug
// }
// type Drug struct{
// 	DrugType	string
// 	Strength	int
// 	Information		string	
// }

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewDrugController(v1, client)
	controllers.NewUserController(v1, client)
	controllers.NewUnitController(v1, client)
	controllers.NewFormController(v1, client)
	controllers.NewVolumeController(v1, client)

	// // Set User
	// users := Users{
	// 	User: []User{
	// 		User{"สมหมาย ใจเกินร้อย", "sommai@gmail.com", "1212312121"},
	// 	},
	// }

	// for _, u := range users.User {
	// 	client.User.
	// 		Create().
	// 		SetName(u.Name).
	// 		SetEmail(u.Email).
	// 		SetPassword(u.Password).
	// 		Save(context.Background())
	// }

	// Set User
	users := Users{
		User: []User{
			User{"นายคนนั้น  นั่นเอง","khonnan@gmail.com", "11111111"},
			User{"นายลำเทียน จ้องผสมพันธุ์","lamtean@gmail.com", "12121212"},
			User{"นายชัยยศ พรหมจารีย์พินาศ","chaiyot@gmail.com", "13131313"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetEmail(u.Email).
			SetUsername(u.Username).
			SetPassword(u.Password).
			Save(context.Background())
	}

	// // Set Drug
	// drugs := Drugs{
	// 	Drug: []Drug{
	// 		Drug{"aspirin"},
	// 	},
	// }

	// for _, dg := range drugs.Drug {
	// 	client.Drug.
	// 		Create().
	// 		SetDrugType(dg.DrugType).
	// 		Save(context.Background())
	// }

	// Set Unit
	units := Units{
		Unit: []Unit{
			// Unit{"tablet"},
			// Unit{"bottle"},
			// Unit{"pack"},
			Unit{"เม็ด"},
			Unit{"ขวด"},
			Unit{"ซอง"},
		},
	}

	for _, un := range units.Unit {
		client.Unit.
			Create().
			SetUnitType(un.UnitType).
			Save(context.Background())
	}

	// Set Volume
	volumes := Volumes{
		Volume: []Volume{
			// Volume{"mg"},
			// Volume{"mcg"},
			Volume{"มิลลิกรัม"},
			Volume{"ไมโครกรัม"},
		},
	}

	for _, v := range volumes.Volume {
		client.Volume.
			Create().
			SetVolumeType(v.VolumeType).
			Save(context.Background())
	}

	// Set Form
	forms := Forms{ 
		Form: []Form{
			// Form{"capsul"},
			// Form{"tablet"},
			// Form{"powder"},
			// Form{"liquid"},
			Form{"แคปซูล"},
			Form{"เม็ด"},
			Form{"ผง"},
			Form{"น้ำ"},
		},
	}

	for _, f := range forms.Form {
		client.Form.
			Create().
			SetFormType(f.FormType).
			Save(context.Background())
	}

	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
