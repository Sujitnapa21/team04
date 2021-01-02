package main

import (
	"context"
	"log"

	"github.com/B6001186/Contagions/controllers"
	_ "github.com/B6001186/Contagions/docs"
	"github.com/B6001186/Contagions/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
// Departments  defines the struct for the departments
type Departments struct {
	Department []Department
}

// Department  defines the struct for the department
type Department struct {
	DepartmentName string
}

// Places  defines the struct for the places
type Places struct {
	Place []Place
}

// Place  defines the struct for the place
type Place struct {
	PlaceName string
}

// Nametitles  defines the struct for the Nametitles
type Nametitles struct {
	Nametitle []Nametitle
}

// Nametitle  defines the struct for the Nametitle
type Nametitle struct {
	title string
}

// Employees  defines the struct for the employees
type Employees struct {
	Employee []Employee
}

// Employee  defines the struct for the employee
type Employee struct {
	UserID string
	EmployeeName string
}

// Statuss  defines the struct for the statuss
type Statuss struct {
	Status []Status
}

// Status  defines the struct for the  status
type Status struct {
	Name string
}

// Genders  defines the struct for the genders
type Genders struct {
	Gender []Gender
}

// Gender  defines the struct for the gender
type Gender struct {
	GenderName string
}

// Bloodtypes  defines the struct for the bloodtypes
type Bloodtypes struct {
	Bloodtype []Bloodtype
}

// Bloodtype  defines the struct for the bloodtype
type Bloodtype struct {
	BloodtypeName string
}

// Patients  defines the struct for the Patients
type Patients struct {
	Patient []Patient
}

// Patient  defines the struct for the Patient
type Patient struct {
	PatientName string
	Congenital string
	Allergic string
}

// Diseases  defines the struct for the diseases
type Diseases struct {
	Disease []Disease
}

// Disease  defines the struct for the Disease
type Disease struct {
	DiseaseName string
}

// Diseasetypes  defines the struct for the diseasetypes
type Diseasetypes struct {
	Diseasetype []Diseasetype
}

// Diseasetype  defines the struct for the diseasetype
type Diseasetype struct {
	DiseaseTypeName  string
}

// Severitys  defines the struct for the severitys
type Severitys struct {
	Severity []Severity
}

// Severity  defines the struct for the severity
type Severity struct {
	SeverityName string
}



// @title SUT SA Example API Patient
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

	client, err := ent.Open("sqlite3", "file:contagions.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewPatientController(v1, client)
	controllers.NewBloodtypeController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewGenderController(v1, client)
	// controllers.NewStatusController(v1, client)
	controllers.NewNametitleController(v1, client)
	controllers.NewDepartmentController(v1, client)
	controllers.NewPlaceController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewDiseaseController(v1, client)
	controllers.NewDiseasetypeController(v1, client)
	controllers.NewSeverityController(v1, client)	


	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"D12345","siriwan"},
			Employee{"D54231","sujitnapa"},
		},
	}

	for _, e := range employees.Employee {
		client.Employee.
			Create().
			SetUserId(e.UserID).
			Save(context.Background())
	}

	// // Set Statuss Data
	// statuss := Statuss{
	// 	Status: []Status{
	// 		Status{"นักเรียน"},
	// 		Status{"นักศึกษา"},
	// 		Status{"บุคลากร"},
	// 		Status{"บุคคลทั่วไป"},
	// 	},
	// }

	// for _, s := range statuss.Status {
	// 	client.Status.
	// 		Create().
	// 		SetName(s.Name).
	// 		Save(context.Background())
	// }

	// Set Nametitles Data
	nametitles := Nametitles{
		Nametitle: []Nametitle{
			Nametitle{"เด็กชาย"},
			Nametitle{"เด็กหญิง"},
			Nametitle{"นาย"},
			Nametitle{"นาง"},
			Nametitle{"นางสาว"},
			Nametitle{"นพ."},
			Nametitle{"พญ."},
			Nametitle{"พย."},
			Nametitle{"พยช."},
			Nametitle{"อื่น ๆ"},
		},
	}

	for _, n := range nametitles.Nametitle {
		client.Nametitle.
			Create().
			SetTitle(n.title).
			Save(context.Background())
	}

	// Set Genders Data
	genders := Genders{
		Gender: []Gender{
			Gender{"ชาย"},
			Gender{"หญิง"},
		},
	}

	for _, g := range genders.Gender {
		client.Gender.
			Create().
			SetGenderName(g.GenderName).
			Save(context.Background())
	}

	// Set Bloodtypes Data
	bloodtypes := Bloodtypes{
		Bloodtype: []Bloodtype{
			Bloodtype{"A"},
			Bloodtype{"B"},
			Bloodtype{"O"},
			Bloodtype{"AB"},
		},
	}

	for _, b := range bloodtypes.Bloodtype {
		client.Bloodtype.
			Create().
			SetBloodtypeName(b.BloodtypeName).
			Save(context.Background())
	}

	// Set Departments Data
	departments := Departments{
		Department: []Department{
			Department{"แพทย์"},
			Department{"เภสัชกร"},
			Department{"เจ้าหน้าที่เวชระเบียน"},
			Department{"แพทย์ระบาดวิทยา"},	
			Department{"พยาบาล"},
		},
	}

	for _, d := range departments.Department {
		client.Department.
			Create().
			SetDepartmentName(d.DepartmentName).
			Save(context.Background())
	}

	// Set Places Data
	places := Places{
		Place: []Place{
			Place{"ตึก A"},
			Place{"ตึก B"},
			Place{"ตึก C"},
			Place{"ตึก D"},
			Place{"ตึก E"},
		},
	}

	for _, p := range places.Place {
		client.Place.
			Create().
			SetPlaceName(p.PlaceName).
			Save(context.Background())
	}

	// Set Diseases Data
	diseases := Diseases{
		Disease: []Disease{
			Disease{"D12345"},
			Disease{"D54231"},
		},
	}

	for _, d := range diseases.Disease {
		client.Disease.
		Create().
		SetDiseaseName(d.DiseaseName).
		Save(context.Background())
	}

	// Set Patients Data
	patients := Patients{
		Patient: []Patient{
			Patient{"D12345","เบาหวาน","ความดัน"},
			Patient{"D54231","อ้วน","โรคหัวใจ"},
		},
	}

	for _, pa := range patients.Patient {
		client.Patient.
		Create().
		SetPatientName(pa.PatientName).
		Save(context.Background())
	}


	// Set Diseasetypes Data
	diseasetypes := Diseasetypes{
		Diseasetype: []Diseasetype{
			Diseasetype{"โรคติดต่อ"},
			Diseasetype{"โรคติดต่อต้องแจ้งความ"},
			Diseasetype{"โรคติดต่ออันตราย"},
		},
	}

	for _, dt := range diseasetypes.Diseasetype {
		client.Diseasetype.
			Create().
			SetDiseaseTypeName(dt.DiseaseTypeName).
			Save(context.Background())
	}
	
	// Set Severity Data
	severitys := Severitys{
		Severity: []Severity{
			Severity{"เริ่มต้น"},
			Severity{"รุนแรง"},
			Severity{"รุนแรงมาก"},
		},
	}

	for _, s := range severitys.Severity {
		client.Severity.
			Create().
			SetSeverityName(s.SeverityName).
			Save(context.Background())
	}	




	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
